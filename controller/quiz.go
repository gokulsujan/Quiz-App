package controller

import (
	"fmt"
	"mhb_heart_day_quiz_2023/config"
	models "mhb_heart_day_quiz_2023/model"
	"net/http"
	"strconv"

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
func StopQuiz(c *gin.Context) {
	id := c.Param("id")
	result := config.DB.Model(&models.Quiz{}).Where("id = ?", id).Updates(&models.Quiz{Status: "ended"})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz?edit=true")
}

func QuizSetup(c *gin.Context) {
	id := c.Param("id")
	var quiz models.Quiz
	var guideline []models.Guidelines
	var questions []models.Question

	//get quiz data
	getQuiz := config.DB.Find(&quiz, id)
	if getQuiz.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": getQuiz.Error.Error()})
		return
	}

	//getting questions
	getQuestions := config.DB.Where("quiz_id = ?", id).Find(&questions)
	if getQuestions.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "false", "error": getQuestions.Error.Error()})
	}

	//getting guidelines
	getGuidelines := config.DB.Where("quiz_id = ?", id).Find(&guideline)
	if getGuidelines.Error != nil {
		c.JSON(http.StatusOK, gin.H{"status": "false", "error": getGuidelines.Error.Error()})
	}

	c.HTML(http.StatusOK, "quizSetup.html", gin.H{"status": "true", "quiz": quiz, "guidlines": guideline, "questions": questions})
}

func AddGuidelines(c *gin.Context) {
	var guideline models.Guidelines
	if err := c.ShouldBind(&guideline); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": err.Error()})
		return
	}
	result := config.DB.Create(&guideline)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz/setup/"+strconv.FormatUint(uint64(guideline.QuizId), 10)+"?guideadd=true")
}

func DeleteGuidelines(c *gin.Context) {
	id := c.Param("id")
	var guideline models.Guidelines
	getGuideline := config.DB.First(&guideline, id)
	if getGuideline.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": getGuideline.Error.Error()})
		return
	}
	fmt.Println(guideline)
	result := config.DB.Delete(&models.Guidelines{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz/setup/"+strconv.FormatUint(uint64(guideline.QuizId), 10)+"?guidedelete=true")
}

func AddQuestion(c *gin.Context) {
	var question models.Question
	if err := c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": err.Error()})
		return
	}
	if question.CorrectAnswer > 4 || question.CorrectAnswer < 1 {
		c.JSON(http.StatusUnauthorized, gin.H{"status": "false", "Message": "Answer is not a valid option"})
		return
	}
	result := config.DB.Create(&question)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz/setup/"+strconv.FormatUint(uint64(question.QuizID), 10)+"?questionAdd=true")
}

func DeleteQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	getQuestion := config.DB.Preload("Quiz").First(&question, id)
	if getQuestion.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": getQuestion.Error.Error()})
		return
	}
	result := config.DB.Delete(&models.Question{}, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz/setup/"+strconv.FormatUint(uint64(question.QuizID), 10)+"?questionDelete=true")
}

func EditQuestionForm(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	result := config.DB.First(&question, id)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.HTML(http.StatusOK, "questionEdit.html", gin.H{"status": "true", "question": question, "id": id})
}

func EditQuestion(c *gin.Context) {
	id := c.Param("id")
	var question models.Question
	if err := c.ShouldBind(&question); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": err.Error()})
		return
	}
	result := config.DB.Where("id = ?", id).Updates(&models.Question{Text: question.Text, OptionA: question.OptionA, OptionB: question.OptionB, OptionC: question.OptionC, OptionD: question.OptionD, CorrectAnswer: question.CorrectAnswer})
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": result.Error.Error()})
		return
	}
	c.Redirect(http.StatusSeeOther, "/admin/quiz/setup/"+strconv.FormatUint(uint64(question.QuizID), 10)+"?questionEdit=true")
}
