package admincontrollers

import (
	"project1/database"
	"project1/models"
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Getadminlogin(context *gin.Context) {
	context.HTML(http.StatusOK, "loginadmin.html", gin.H{})
}

var users []models.Users

func Getadminhome(ctx *gin.Context) {
	getusers := getusers()
	fmt.Println("starting...")
	for _, v := range users {
		fmt.Println(v.Username)
	}
	ctx.HTML(http.StatusOK, "adminhome.html", gin.H{"users": getusers})
}

var searchusers []models.Users

func Getsearches(ctx *gin.Context) {
	getusers := getusers()
	search := ctx.PostForm("query")
	if search == "" {
		ctx.HTML(http.StatusOK, "adminhome.html", gin.H{"users": users, "search": searchusers})
		return
	}
	query := "SELECT * FROM users WHERE username LIKE ?"
	database.DB.Raw(query, "%"+search+"%").Scan(&searchusers)
	ctx.HTML(http.StatusOK, "adminhome.html", gin.H{"users": getusers, "search": searchusers})
}

func Deleteuser(ctx *gin.Context) {
	username := ctx.Param("username")
	fmt.Println(username)
	query := "DELETE FROM users WHERE username = ?"
	database.DB.Exec(query, username)
	getuser := getusers()
	ctx.HTML(http.StatusOK, "adminhome.html", gin.H{"users": getuser})
}

func Edituser(ctx *gin.Context) {
	var sigleuser models.Users
	username := ctx.Param("username")
	query := "select * from users where username = ?"
	database.DB.Raw(query, username).Scan(&sigleuser)
	ctx.HTML(http.StatusOK, "editusers.html", gin.H{"user": sigleuser})
}

func Postedit(ctx *gin.Context) {
	editusers := ctx.Param("edituser")
	username := ctx.PostForm("username")
	email := ctx.PostForm("email")
	query := "UPDATE users SET username = ?, email = ? WHERE username = ?"
	res := database.DB.Exec(query, username, email, editusers)
	if res.Error != nil {
		fmt.Print(res.Error)
	}
	getuser := getusers() 
	ctx.HTML(http.StatusOK, "adminhome.html", gin.H{"users": getuser})
}

func getusers() []models.Users {
	var getusers []models.Users
	getquery := "select * from users"
	database.DB.Raw(getquery).Scan(&getusers)
	return getusers
}

func Logoutadmin(ctx *gin.Context)  {
	ctx.SetCookie("adminjwt", "", -1, "/", "localhost", true, true)
	ctx.Redirect(http.StatusPermanentRedirect,"/login")
}