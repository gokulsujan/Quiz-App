package config

import (
	models "mhb_heart_day_quiz_2023/model"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

var DB *gorm.DB

func ConnectToDB() {
	var err error
	dsn := "host=localhost user=postgres password=superadmin dbname=mhb_heart_day_quiz port=5432 sslmode=disable"
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}

	DB.AutoMigrate(&models.User{})
	DB.AutoMigrate(&models.Quiz{})
	DB.AutoMigrate(&models.Question{})
	DB.AutoMigrate(&models.UserQuiz{})
	DB.AutoMigrate(&models.Department{})
	DB.AutoMigrate(&models.Guidelines{})
	DB.AutoMigrate(&models.UserQuestion{})
}
