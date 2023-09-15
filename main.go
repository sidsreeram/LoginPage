package main

import (
	"fmt"
	"os"
	admincontrollers "project1/controllers/admin_controllers"
	usercontrollers "project1/controllers/user_controllers"
	"project1/database"
	adminmiddlewares "project1/middlewares/admin_middlewares"
	usermiddlewares "project1/middlewares/user_middlewares"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	// Loading environment variables from "globals.env"
	err := godotenv.Load("globals.env")
	if err != nil {
		fmt.Println(err)
		panic("Cannot load environment variables")
	}
	
	// Getting the database address from the environment variables
	env := os.Getenv("DATABASE_ADDRESS")

	// Connecting to the database and migrating the models to tables
	database.Connect(env)
	database.Migrator()
	// database.Migrateadmin()

	// Initializing the Gin router
	router := InitGin()
	router.Run("localhost:1011")
}

// Initializing the router and setting up routes
func InitGin() *gin.Engine {
	router := gin.Default()
	router.LoadHTMLGlob("templates/*.html") // To render HTML files

	// Routes for user-related functionality
	router.Use(func(ctx *gin.Context) {
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Next()
	})
	router.GET("/", usermiddlewares.SecureHome(), usercontrollers.LoginGetHandler)          // Home for users
	router.GET("/login", usermiddlewares.ClearCache(), usercontrollers.LoginGetHandler)     // User login page
	router.GET("/signup", usercontrollers.SignUpGetHandler)                             // User signup page
	router.POST("/submitsignup", usercontrollers.SignUpPostHandler)                      // Create a user account
	router.POST("/gethome", usermiddlewares.ValidateUser(), usercontrollers.HomeGetHandler) // User login
	router.GET("/loginadmin", admincontrollers.Getadminlogin)//calling admin login page
	router.POST("/getadminhome", adminmiddlewares.Validateadmin(), admincontrollers.Getadminhome)//called when the admin presses the login button
	router.GET("/adminhome", adminmiddlewares.SecureadminHome(), admincontrollers.Getadminhome)//home admin
	router.POST("/search", admincontrollers.Getsearches)//to get the search results about users in admin home
	router.GET("/delete/:username", admincontrollers.Deleteuser)//called for deleting a user
	router.GET("/edit/:username", admincontrollers.Edituser)//calles the edit page of the users
	router.POST("/editusers/:edituser", admincontrollers.Postedit)//posting the edit on user
	router.GET("/logout", usercontrollers.LogoutGetHandler)
	router.GET("/logoutadmin", admincontrollers.Logoutadmin)



	return router
}
