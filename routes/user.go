package routes

import (
	"mhb_heart_day_quiz_2023/auth"
	"mhb_heart_day_quiz_2023/controller"

	"github.com/gin-gonic/gin"
)

func RouteSetup(r *gin.Engine) {
	User := r.Group("/")
	{
		User.GET("/registration", controller.UserSignUpForm)
		User.POST("/registration", controller.UserSignUp)

		User.GET("/", controller.LoginGetHTML)
		User.POST("login", controller.Login)
	}
	Admin := r.Group("/admin")
	{
		Admin.GET("", auth.AdminAuth, controller.UserRoleSelect)
		Admin.GET("/home", auth.AdminAuth, controller.AdminHome)
		Admin.GET("/department", auth.AdminAuth, controller.DeptManagement)
		Admin.POST("/department", auth.AdminAuth, controller.AddDepartment)
		Admin.POST("/delete/:id", auth.AdminAuth, controller.DeleteDept)
		Admin.GET("/user", auth.AdminAuth, controller.UserManagerment)
		Admin.GET("/user/edit/:id", auth.AdminAuth, controller.EditUserForm)
		Admin.POST("/user/edit/:id", auth.AdminAuth, controller.EditUser)
		Admin.GET("/user/delete/:id", auth.AdminAuth, controller.DeleteUser)
		Admin.GET("/quiz", auth.AdminAuth, controller.QuizHome)
		Admin.POST("/quiz/add", auth.AdminAuth, controller.AddQuiz)
		Admin.GET("/quiz/delete/:id", auth.AdminAuth, controller.DeleteQuiz)
		Admin.GET("/quiz/edit/:id", auth.AdminAuth, controller.EditQuizForm)
		Admin.POST("/quiz/edit/:id", auth.AdminAuth, controller.EditQuiz)
		Admin.GET("/quiz/start/:id", auth.AdminAuth, controller.StartQuiz)
	}
}
