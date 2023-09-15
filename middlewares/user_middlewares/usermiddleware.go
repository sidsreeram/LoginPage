package usermiddlewares

import (
	"net/http"
	"project1/database"
	"project1/models"
	"project1/auth"
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
)

func SecureHome() gin.HandlerFunc {
	return func(context *gin.Context) {
		
		tokenjwt, err := context.Request.Cookie("jwt")
		if err != nil || tokenjwt == nil || tokenjwt.Value == "" {
			
			context.Redirect(http.StatusPermanentRedirect, "/login")
			context.Abort()
			return
		}

		// Validating the token using its credentials and checking if the expiration time has expired
		if token.ValidateToken(tokenjwt.Value) {
			context.Next()
			return
		}

		// Token is invalid, redirect to login
		context.Set("message", "Session invalid")
		context.Redirect(http.StatusPermanentRedirect, "/login")
		context.Abort()
	}
}

func ValidateUser() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var user models.Users
		username := ctx.PostForm("username")
		password := ctx.PostForm("password")

		// Query the database to find the user
		database.DB.Where("username = ?", username).First(&user)

		// Compare hashed password
		if user.Username == "" || bcrypt.CompareHashAndPassword([]byte(user.Password), []byte(password)) != nil {
			// Invalid user or password, return an error response
			ctx.HTML(http.StatusOK, "loginuser.html", gin.H{"message": "User not found or invalid password"})
			ctx.Abort()
			return
		}

		// Generate and set JWT token as a cookie
		jwttoken := token.Generatejwt(username, password)
ctx.SetCookie("jwt", jwttoken, 3600, "/", "localhost", true, true)
		ctx.SetCookie("username", username, 3600, "/", "localhost", true, true)

		// Continue processing the request
		ctx.Next()
	}
}

func ClearCache() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		
		ctx.Header("Cache-Control", "no-cache, no-store, must-revalidate")
		ctx.Header("Pragma", "no-cache")
		ctx.Header("Expires", "0")
		ctx.Header("Vary", "Cookie")
		ctx.Next()
	}
}
