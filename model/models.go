package models

import (
	"time"

	"gorm.io/gorm"
)

// Login Credentials
type LoginCredentials struct {
	Username string `json:"username" form:"username"`
	Password string `json:"password" form:"password"`
}

// User represents a user in the system.
type User struct {
	gorm.Model
	Name       string     `json:"name" form:"name" gorm:"not null"`
	EmpId      uint       `json:"emp_id" form:"emp_id" gorm:"not null"`
	Department uint       `json:"dept_id" form:"dept_id" gorm:"not null"`
	Dept       Department `gorm:"ForeignKey:Department"`
	Email      string     `json:"email" form:"email" gorm:"not null"`
	Password   string     `json:"password" form:"password" gorm:"not null"`
	Role       string     `json:"role" form:"role" gorm:"default:user"`
	Status     string     `json:"status" form:"status" gorm:"default:pending"`
}

type Department struct {
	gorm.Model
	Name string `json:"dept_name" form:"dept_name" gorm:"not null"`
}

// Quiz represents a quiz that can be created by an admin.
type Quiz struct {
	gorm.Model
	Title       string `json:"title" form:"title" gorm:"not null"`
	Description string `json:"description" form:"description" gorm:"not null"`
	Questions   int    `json:"questions" form:"questions" gorm:"not null"`
	Duration    int    `json:"duration" form:"duration" gorm:"not null"` // Duration in minutes
	Status      string `json:"status" form:"status" gorm:"default:pending"`
}

// Question represents a quiz question.
type Question struct {
	gorm.Model
	QuizID        uint   `json:"quiz_id" form:"quiz_id" gorm:"not null"`
	Quiz          Quiz   `gorm:"ForeignKey:QuizID"`
	Text          string `json:"text" form:"text" gorm:"not null"`
	OptionA       string `json:"option_a" form:"option_a" gorm:"not null"`
	OptionB       string `json:"option_b" form:"option_b" gorm:"not null"`
	OptionC       string `json:"option_c" form:"option_c" gorm:"not null"`
	OptionD       string `json:"option_d" form:"option_d" gorm:"not null"`
	CorrectAnswer int    `json:"correct_answer" form:"correct_answer" gorm:"not null"`
}

// UserQuiz represents the participation of a user in a quiz.
type UserQuiz struct {
	gorm.Model
	UserID      uint       `json:"user_id" form:"user_id" gorm:"not null"`
	User        User       `gorm:"ForeignKey:UserID"`
	QuizID      uint       `json:"quiz_id" form:"quiz_id" gorm:"not null"`
	Quiz        Quiz       `gorm:"ForeignKey:QuizID"`
	CompletedAt *time.Time `json:"completed_at" form:"completed_at" gorm:"default:null"`
}

// UserQuestion
type UserQuestion struct {
	gorm.Model
	UserQuizId uint     `json:"user_quiz_id" form:"user_quiz_id" gorm:"not null"`
	UserQuiz   UserQuiz `gorm:"ForeignKey:UserQuizId"`
	QuestionId uint     `json:"question_id" form:"questin_id" gorm:"not null"`
	Question   Question `gorm:"ForeignKey:"`
	Answer     int      `json:"answer" form:"answer" gorm:"not null"`
}

// Guidelines represents the guidelines for the quiz

type Guidelines struct {
	gorm.Model
	Guideline string `json:"guideline" form:"guideline" gorm:"not null"`
	QuizId    uint   `json:"quiz_id" form:"quiz_id" gorm:"not null"`
	Quiz      Quiz   `gorm:"ForeignKey:QuizId"`
}
