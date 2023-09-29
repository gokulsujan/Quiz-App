package controller

import (
	"math/rand"
	"mhb_heart_day_quiz_2023/config"
	models "mhb_heart_day_quiz_2023/model"
	"net/http"
	"strconv"
	"time"

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
	//checking quiz already started
	quizStat := config.DB.Where("user_id = ? and quiz_id = ?", user.ID, id).First(&models.UserQuiz{})
	if quizStat.RowsAffected != 0 {
		c.Redirect(http.StatusSeeOther, "/quiz/questions/"+id)
		return
	}

	c.HTML(http.StatusOK, "userQuizHome.html", gin.H{"guidelines": guidelines, "user": user})
}

func UserQuizStart(c *gin.Context) {
	var userQuiz models.UserQuiz
	if err := c.ShouldBind(&userQuiz); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": err.Error(), "status": "false"})
		return
	}

	//getQUizDetails
	var quiz models.Quiz
	getQuiz := config.DB.First(&quiz, userQuiz.QuizID)
	if getQuiz.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getQuiz.Error.Error(), "status": "false"})
		return
	}

	//checking quiz already started
	quizStat := config.DB.Where("user_id = ? and quiz_id = ?", userQuiz.UserID, quiz.ID).First(&models.UserQuiz{})
	if quizStat.RowsAffected != 0 {
		c.Redirect(http.StatusSeeOther, "/quiz/questions/"+strconv.Itoa(int(userQuiz.QuizID)))
		return
	}

	result := config.DB.Create(&userQuiz)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}

	//get all the questions in the quiz
	var questions []models.Question
	getQuestions := config.DB.Where("quiz_id = ?", userQuiz.QuizID).Find(&questions)
	if getQuestions.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": getQuestions.Error.Error(), "status": "false"})
		return
	}

	userQuestions := make([]int, 0, quiz.Questions)
	if quiz.Questions > len(questions)+1 {
		c.JSON(http.StatusBadGateway, gin.H{"status": "false", "message": "There is not much question available in the database. Contact the host"})
	}
	rand.Seed(time.Now().UnixNano())
	usedNumbers := make(map[string]bool)
	for len(userQuestions) < quiz.Questions {
		num := rand.Intn(len(questions))
		if !usedNumbers[strconv.Itoa(int(userQuiz.UserID))+strconv.Itoa(int(num))] {
			usedNumbers[strconv.Itoa(int(userQuiz.UserID))+strconv.Itoa(int(num))] = true
			userQuestions = append(userQuestions, num)
			result := config.DB.Create(&models.UserQuestion{UserQuizId: userQuiz.ID, QuestionId: questions[num].ID})
			if result.Error != nil {
				c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
				return
			}
		}
		// fmt.Println(num)
	}
	c.Redirect(http.StatusSeeOther, "/quiz/questions/"+strconv.Itoa(int(userQuiz.QuizID)))
}

func QuizQuestion(c *gin.Context) {
	id := c.Param("id")
	var questions []models.Question
	var userDetails models.User
	user_id, _ := c.Get("empID")

	//getting the user
	getUser := config.DB.Where("emp_id = ?", user_id).First(&userDetails)
	if getUser.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"status": "false", "error": getUser.Error.Error()})
		return
	}

	result := config.DB.Table("questions").Select("questions.*").Joins("INNER JOIN user_questions ON questions.id = user_questions.question_id").Joins("INNER JOIN user_quizzes ON user_questions.user_quiz_id = user_quizzes.id").Where("user_quizzes.user_id = ? AND user_quizzes.quiz_id = ?", userDetails.ID, id).Find(&questions)
	if result.Error != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": result.Error.Error(), "status": "false"})
		return
	}
	c.HTML(http.StatusAccepted, "quizQuestionsUsers.html", gin.H{"status": "true", "questions": questions})
}
