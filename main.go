package main

import (
	"mhb_heart_day_quiz_2023/config"
	"mhb_heart_day_quiz_2023/routes"
	"text/template"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

func init() {
	config.ConnectToDB()
}

func main() {
	// Define a custom "add" function for serial numbers
	funcMap := template.FuncMap{
		"add": func(a, b int) int {
			return a + b
		},
		"eq": func(a, b interface{}) bool {
			return a == b
		},
	}

	r := gin.Default()

	// Register the custom function with the template engine
	r.SetFuncMap(funcMap)

	// Serve HTML templates from the "templates" directory
	r.LoadHTMLGlob("view/*.html")

	// Serve static files from the "static" directory
	r.Static("/static", "./view")

	// Use cookies middleware
	store := cookie.NewStore([]byte("your-secret-key"))
	r.Use(sessions.Sessions("mysession", store))

	//routes setup
	routes.RouteSetup(r)

	r.Run()
}
