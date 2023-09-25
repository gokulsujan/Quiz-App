package controller

import (
	"mhb_heart_day_quiz_2023/config"
	models "mhb_heart_day_quiz_2023/model"
	"net/http"

	"github.com/gin-gonic/gin"
)

func QuizHome(c *gin.Context) {
	var quizes []models.Quiz
	result := config.DB.Find(&quizes)
	if result.Error != nil {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "false", "error": result.Error.Error()})
	}
	c.HTML(http.StatusOK, "quizHome.html", gin.H{"status": "true", "quizes": quizes})
}

func AddQuiz(c *gin.Context) {
	var quiz models.Quiz
	if err := c.ShouldBind(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}
	result := config.DB.Create(&quiz)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz?add=true")
}

func EditQuizForm(c *gin.Context) {
	var quiz models.Quiz
	id := c.Param("id")
	result := config.DB.First(&quiz, id)
	if result.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.HTML(http.StatusOK, "editQuiz.html", gin.H{"quiz": quiz})
}

func EditQuiz(c *gin.Context) {
	id := c.Param("id")
	var quiz models.Quiz
	if err := c.ShouldBind(&quiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}
	result := config.DB.Model(&quiz).Where("id = ?", id).Updates(&models.Quiz{Title: quiz.Title, Description: quiz.Description, Duration: quiz.Duration})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz?edit=true")
}

func DeleteQuiz(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Delete(&models.Quiz{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz?delete=true")
}

func StartQuiz(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.Quiz{}).Where("id = ?", id).Updates(&models.Quiz{Status: "active"})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz?edit=true")
}
