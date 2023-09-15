package usercontrollers

import (
	"fmt"
	"net/http"
	"project1/database"
	"project1/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	// "golang.org/x/net/context/ctxhttp"
)

func SignUpPostHandler(context *gin.Context) {
	var user models.Users
	username := context.PostForm("username")
	email := context.PostForm("email")
	password := context.PostForm("password")
	confirmpassword := context.PostForm("confirmPassword")
	fmt.Println(password)
	fmt.Println(confirmpassword)
	if password != confirmpassword {
		context.HTML(http.StatusOK, "signup.html", gin.H{"message": "both the passwords should be the same"})
		return
	}
	database.DB.Where("username = ?", username).First(&user)
	if user.Username != "" {
		context.HTML(http.StatusOK, "signup.html", gin.H{"message": "user already exists"})
		return
	}
	//hashing the password
	pass := []byte(password)
	hashedpass, err := bcrypt.GenerateFromPassword(pass, bcrypt.DefaultCost)
	if err != nil {
		context.HTML(http.StatusOK, "signup.html", gin.H{"message": "error... try again later"})
		return
	}

	//creating the specific row to be inserted to the data base
	user = models.Users{
		Username: username,
		Password: string(hashedpass),
		Email:    email,
	}
	//inserting the user to the data base && checks for any errors occured during the transaction
	record := database.DB.Create(&user)
	if record.Error != nil {
		fmt.Println("Error occured while inseerting ton the data base")
		context.JSON(http.StatusInternalServerError, gin.H{"error": record.Error.Error()})
		context.Abort()
		return
	}
	context.HTML(http.StatusOK, "loginuser.html", gin.H{})
}
func LoginPostHandler(ctx *gin.Context) {
	username := ctx.PostForm("username")
	password := ctx.PostForm("password")

	var user models.Users

	// Fetch the user by username from the database
	if err := database.DB.Where("username = ?", username).First(&user).Error; err != nil {
		// User not found
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid username or password"})
		return
	}

	// Compare the provided password with the hashed password in the database
	if err := bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)); err != nil {
		// Passwords do not match
		ctx.HTML(http.StatusUnauthorized, "login.html", gin.H{"message": "Invalid username or password"})
		return
	}

	// Passwords match, user is authenticated
	// You can set a session or token to keep the user logged in here if needed

	ctx.HTML(http.StatusOK, "homeuser.html", gin.H{"message": "Logged in successfully"})
}
