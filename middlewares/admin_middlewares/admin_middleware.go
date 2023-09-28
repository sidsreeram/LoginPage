package adminmiddlewares

import (
	"fmt"
	"net/http"
	"project1/auth"
	"project1/database"
	"project1/models"
	"github.com/gin-gonic/gin"
)

func SecureadminHome() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenjwt, err := context.Cookie("adminjwt") // Updated to use context.Cookie
		if err != nil || tokenjwt == "" {
			context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
			context.Abort()
			return
		}

		if token.ValidateToken(tokenjwt) {
			context.Next()
		} else {
			context.Set("message", "Session invalid")
			context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
		}
	}
}

func Validateadmin() gin.HandlerFunc {
	return func(context *gin.Context) {
		var admin models.Admins
		username := context.PostForm("username")
		password := context.PostForm("password")
		database.DB.Where("username = ? AND password = ?", username, password).First(&admin)

		if admin.Username == "" {
			context.HTML(http.StatusOK, "loginadmin.html", gin.H{"message": "not an admin"})
			context.Abort()
			return
		}

		jwttoken, err := token.Generatejwt(username, password) // Updated to handle potential error
		if err != nil {
			fmt.Println(err)
			context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
			context.Abort()
			return
		}

		// Set the JWT token as a cookie
		context.SetCookie("adminjwt", jwttoken, 3600, "/", "localhost", false, true)

		if token.ValidateToken(jwttoken) {
			context.Next()
		} else {
			context.Set("message", "Session invalid")
			context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
		}
	}
}
