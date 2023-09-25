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
	Title    string `json:"title" gorm:"not null"`
	Duration int    `json:"duration" gorm:"not null"` // Duration in minutes
	Status   string `json:"status" gorm:"not null"`
}

// Question represents a quiz question.
type Question struct {
	gorm.Model
	QuizID        uint   `json:"quiz_id" gorm:"not null"`
	Text          string `json:"text" gorm:"not null"`
	OptionA       string `json:"option_a" gorm:"not null"`
	OptionB       string `json:"option_b" gorm:"not null"`
	OptionC       string `json:"option_c" gorm:"not null"`
	OptionD       string `json:"option_d" gorm:"not null"`
	CorrectAnswer int    `json:"correct_answer" gorm:"not null"`
}

// UserQuiz represents the participation of a user in a quiz.
type UserQuiz struct {
	gorm.Model
	UserID      uint       `json:"user_id"  gorm:"not null"`
	QuizID      uint       `json:"quiz_id" gorm:"not null"`
	Score       int        `json:"score" gorm:"not null"`
	StartedAt   *time.Time `json:"started_at" gorm:"not null"`
	CompletedAt *time.Time `json:"completed_at" gorm:"not null"`
}
