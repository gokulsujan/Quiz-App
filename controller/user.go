package controller

import (
	"mhb_heart_day_quiz_2023/config"
	models "mhb_heart_day_quiz_2023/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func UserSignUp(c *gin.Context) {
	var user models.User
	if err := c.ShouldBind(&user); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}

	//checking the email id already exists or not
	checkEmail := config.DB.Where("email = ?", user.Email).First(&models.User{})
	if checkEmail.RowsAffected != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Email Id already registered with us.", "status": "false"})
		return
	}

	//checking employee id exist or not
	checkEmp := config.DB.Where("emp_id = ?", user.EmpId).First(&models.User{})
	if checkEmp.RowsAffected != 0 {
		c.JSON(http.StatusUnauthorized, gin.H{"message": "Employee id already registered with us", "status": "false"})
		return
	}

	//inserting data into database
	result := config.DB.Create(&user)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.JSON(http.StatusSeeOther, gin.H{"message": "Registration SUccessfull!!", "status": "true"})
}

func UserSignUpForm(c *gin.Context) {
	var department []models.Department
	result := config.DB.Find(&department)
	if result.Error != nil {
		if result.Error != nil {
			c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
			return
		}
	}
	c.HTML(http.StatusOK, "registration.html", gin.H{"departments": department})
}
