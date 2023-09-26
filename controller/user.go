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
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.HTML(http.StatusOK, "registration.html", gin.H{"departments": department})
}

func UserHome(c *gin.Context) {
	var quiz []models.Quiz
	var user models.User
	userId, _ := c.Get("empID")
	getUser := config.DB.Where("emp_id = ?", userId).First(&user)
	if getUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getUser.Error.Error(), "status": "false"})
		return
	}
	result := config.DB.Where("status = ? OR status = ?", "active", "ended").Find(&quiz)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.HTML(http.StatusOK, "userHome.html", gin.H{"quiz": quiz, "user": user})
}

func QuizGuidelines(c *gin.Context) {
	id := c.Param("id")
	var guidelines []models.Guidelines
	var user models.User
	result := config.DB.Where("quiz_id = ?", id).Find(&guidelines)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	userId, _ := c.Get("empID")
	getUser := config.DB.Where("emp_id = ?", userId).First(&user)
	if getUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getUser.Error.Error(), "status": "false"})
		return
	}
	c.HTML(http.StatusOK, "userQuizHome.html", gin.H{"guidelines": guidelines, "user": user})
}
