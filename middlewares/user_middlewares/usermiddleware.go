package usermiddlewares

import (
	"fmt"
	"net/http"
	"project1/auth"
	"project1/database"
	"project1/models"

	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SecureHome() gin.HandlerFunc {
	return func(context *gin.Context) {
		// Retrieve the JWT token from the cookie
		tokenjwt, err := context.Cookie("jwt")
		if err != nil || tokenjwt == "" {
			// If the token is missing or there's an error, redirect to the login page
			context.Redirect(http.StatusPermanentRedirect, "/login")
			context.Abort()
			return
		}

		// Validate the token using its credentials and check if it has expired
		if !token.ValidateToken(tokenjwt) {
			// If the token is invalid, redirect to the login page
			context.Set("message", "Session invalid")
			context.Redirect(http.StatusPermanentRedirect, "/login")
			context.Abort()
			return
		}

		// Token is valid, continue processing the request
		context.Next()
	}
}

func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Retrieve user credentials from the login form
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		// Query the database to find the user by username
		var user models.Users
		database.DB.Where("username = ?", username).First(&user)

		// Compare hashed password
		if user.Username == "" || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			// If the user is not found or the password is invalid, return an error response
			ctx.HTML(http.StatusOK, "loginuser.html", gin.H{"message": "User not found or invalid password"})
			ctx.Abort()
			return
		}

		// Generate JWT token
		jwttoken, err := token.Generatejwt(username, password)
		if err != nil {
			fmt.Println(err)
			ctx.Redirect(http.StatusPermanentRedirect, "/loginadmin")
			ctx.Abort()
			return
		}

		// Set JWT token as a cookie
		ctx.SetCookie("jwt", jwttoken, 3600, "/", "localhost", true, true)
		ctx.SetCookie("username", username, 3600, "/", "localhost", true, true)

		// Continue processing the request
		ctx.Next()
	}
}


func ClearCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		// Clear cache headers to prevent caching of sensitive data
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Header("Vary", "Cookie")
		ctx.Next()
	}
}

