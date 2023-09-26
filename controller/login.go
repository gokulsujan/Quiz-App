package controller

import (
	"mhb_heart_day_quiz_2023/auth"
	"mhb_heart_day_quiz_2023/config"
	models "mhb_heart_day_quiz_2023/model"
	"net/http"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func LoginGetHTML(c *gin.Context) {
	c.HTML(http.StatusOK, "login.html", nil)
}

func Login(c *gin.Context) {
	var user models.User
	var loginCred models.LoginCredentials

	if err := c.ShouldBind(&loginCred); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}

	//Getting the userdetails
	userData := config.DB.Where("email = ? OR emp_id = ?", loginCred.Username, loginCred.Username).First(&user)
	if userData.RowsAffected == 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials", "status": "false"})
		return
	}

	if user.Password != loginCred.Password {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Invalid Credentials", "status": "false"})
		return
	}

	if user.Status != "active" {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "User is not activated", "status": "false"})
		return
	}

	token, err := auth.CreateToken(user.EmpId, user.Role)
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}
	session := sessions.Default(c)
	session.Set("token", token)
	session.Save()
	c.Redirect(http.StatusSeeOther, "/admin")
}
