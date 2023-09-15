package adminmiddlewares

import (
	"fmt"
	"net/http"
	"project1/auth"
	"project1/database"
	"project1/models"

	"github.com/gin-contrib/sessions"
	"github.com/gin-gonic/gin"
)

func SecureadminHome() gin.HandlerFunc {
	return func(context *gin.Context) {
		tokenjwt, err := context.Request.Cookie("adminjwt")
		if err != nil {
			context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
			fmt.Println(err)
			context.Abort()
			return
		}
		if tokenjwt == nil || tokenjwt.Value == "" {
			context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
			context.Abort()
			return
		}
		if token.ValidateToken(tokenjwt.Value) {
			context.Next()
			return
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
        jwttoken := token.Generatejwt(username, password)
        fmt.Println(jwttoken)
        
        // Store the JWT token in the user's session
        session := sessions.Default(context)
        session.Set("adminjwt", jwttoken)
        session.Save()
        
        context.SetCookie("adminjwt", jwttoken, 3600, "/", "localhost", true, true)
        if jwttoken == "" {
            context.Redirect(http.StatusPermanentRedirect, "/loginadmin")
            context.Abort()
            return
        }
        context.Next()
    }
}