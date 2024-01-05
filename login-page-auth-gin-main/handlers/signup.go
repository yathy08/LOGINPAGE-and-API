package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/cache"
)

func GetSignup(g *gin.Context) {
	cache.ClearCache(g)
	userSession := GetSessionValue(g, "sessionToken", "username")
	if userSession != nil {
		fmt.Println("Redirected to home ")
		g.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}
	fmt.Println("loaded signup page")
	g.HTML(http.StatusOK, "signup.html", nil)
}

func PostSignup(g *gin.Context) {
	cache.ClearCache(g)

	userSession := GetSessionValue(g, "sessionToken", "username")
	if userSession != nil {
		fmt.Println("Redirected to home ")
		g.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}

	g.Request.ParseForm()
	username := g.Request.FormValue("name")
	email := g.Request.FormValue("email")
	password := g.Request.Form.Get("password")
	cfmpassword := g.Request.Form.Get("cfm-password")

	if password != cfmpassword {
		g.HTML(http.StatusBadRequest, "signup.html", "password does not match")
		return
	}

	exist := MakeUser(username, email, password)
	if exist != nil {
		g.HTML(http.StatusBadRequest, "signup.html", exist)
		return
	}

	user, _ := GetUser(email)
	err := CreateSession(g, "sessionToken", "username", user)
	if err != nil {
		g.Error(err)
	}
	g.HTML(http.StatusOK, "index.html", user)
	g.Redirect(http.StatusOK, "/")

}
