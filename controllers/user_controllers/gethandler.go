package usercontrollers

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
)

func LoginGetHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "loginuser.html", gin.H{})
}
func SignUpGetHandler(ctx *gin.Context) {
	ctx.HTML(http.StatusOK, "signup.html", gin.H{})
}
func LogoutGetHandler(ctx *gin.Context) {
	ctx.SetCookie("jwt", "", -1, "/", "localhost", true, true)
	ctx.SetCookie("username", "", -1, "/", "localhost", true, true)
	fmt.Println("saodgtdysufgsfiuasssfjksghjgfjuyefgjdtyufvwejumg")
	ctx.Redirect(http.StatusPermanentRedirect, "/login")
}
func HomeGetHandler(ctx *gin.Context) {
	log.Println("Homelander1")
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("An error occurred:", r)
			ctx.SetCookie("jwt", "", -1, "/", "localhost", true, true)
			ctx.SetCookie("username", "", -1, "/", "localhost", true, true)
			ctx.Redirect(http.StatusPermanentRedirect, "/login")
		}
	}()
	fmt.Println("Home called")
	username:=ctx.PostForm("username")
	ctx.HTML(http.StatusOK, "homeuser.html", gin.H{"name":username})
}
