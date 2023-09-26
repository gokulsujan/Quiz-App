package auth

import (
	"fmt"
	"net/http"
	"strconv"
	"time"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v4"
)

var jwtKey = []byte("your_secret_key")

type CustomClaims struct {
	EmpID string `json:"emp_id"`
	Role  string `json:"role"`
	jwt.StandardClaims
}

// CreateToken generates a JWT token with the specified emp_id, role, and expiry time.
func CreateToken(empID uint, role string) (string, error) {
	empIDStr := strconv.FormatUint(uint64(empID), 10)
	// Create a new token with custom claims
	token := jwt.New(jwt.SigningMethodHS256)

	// Set the claims (payload)
	claims := &CustomClaims{
		EmpID: empIDStr,
		Role:  role,
		StandardClaims: jwt.StandardClaims{
			ExpiresAt: time.Now().Add(2 * time.Hour).Unix(), // Token expires in 2 hours
		},
	}

	token.Claims = claims

	// Sign the token with the secret key
	tokenString, err := token.SignedString(jwtKey)
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func AdminAuth(c *gin.Context) {
	session := sessions.Default(c)
	tokenString := session.Get("token")
	if tokenString == nil {
		c.Redirect(http.StatusSeeOther, "")
		return
	}

	token, err := jwt.Parse(tokenString.(string), func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method and return the secret key
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if token.Valid {
		// Access the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Access individual claims
			role := claims["role"].(string)
			if role != "admin" {
				c.Redirect(http.StatusSeeOther, "/home")
				c.AbortWithStatus(http.StatusSeeOther)
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid claims format", "status": "false"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed token", "status": "false"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token is either expired or not valid yet", "status": "false"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid token", "status": "false"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Next()
}

func UserAuth(c *gin.Context) {
	session := sessions.Default(c)
	tokenString := session.Get("token")
	if tokenString == nil {
		c.Redirect(http.StatusSeeOther, "")
		return
	}

	token, err := jwt.Parse(tokenString.(string), func(token *jwt.Token) (interface{}, error) {
		// Verify the signing method and return the secret key
		if token.Method.Alg() != jwt.SigningMethodHS256.Alg() {
			return nil, fmt.Errorf("invalid signing method")
		}
		return jwtKey, nil
	})
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	if token.Valid {
		// Access the claims
		if claims, ok := token.Claims.(jwt.MapClaims); ok {
			// Access individual claims
			role := claims["role"].(string)
			emp := claims["emp_id"].(string)
			c.Set("empID", emp)
			if (role != "user") && (role != "admin") {
				c.Redirect(http.StatusSeeOther, "")
				c.AbortWithStatus(http.StatusSeeOther)
				return
			}
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid claims format", "status": "false"})
			c.AbortWithStatus(http.StatusUnauthorized)
			return
		}
	} else if ve, ok := err.(*jwt.ValidationError); ok {
		if ve.Errors&jwt.ValidationErrorMalformed != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Malformed token", "status": "false"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else if ve.Errors&(jwt.ValidationErrorExpired|jwt.ValidationErrorNotValidYet) != 0 {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Token is either expired or not valid yet", "status": "false"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		} else {
			c.JSON(http.StatusBadRequest, gin.H{"message": "Invalid token", "status": "false"})
			c.AbortWithStatus(http.StatusBadRequest)
			return
		}
	} else {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		c.AbortWithStatus(http.StatusBadRequest)
		return
	}

	c.Next()
}
