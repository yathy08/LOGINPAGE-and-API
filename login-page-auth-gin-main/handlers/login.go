package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/cache"
)

func GetLogin(g *gin.Context) {
	cache.ClearCache(g)
	fmt.Println("login page loaded")
	g.HTML(http.StatusOK, "login.html", nil)
}

func PostLogin(g *gin.Context) {
	cache.ClearCache(g)

	userSession := GetSessionValue(g, "sessionToken", "username")
	if userSession != nil {
		fmt.Println("redirected to home")
		g.HTML(http.StatusBadRequest, "index.html", userSession)
		return
	}

	g.Request.ParseForm()
	pwd := g.Request.FormValue("password")
	email := g.Request.FormValue("email")

	user, exist := GetUser(email)
	if exist != nil {
		g.HTML(http.StatusBadRequest, "login.html", user)
		return
	}
	if user.Email != email {
		g.HTML(http.StatusBadRequest, "login.html", "invalid email ")
		return
	}
	if user.Password != pwd {
		g.HTML(http.StatusBadRequest, "login.html", "invalid password")
		return
	}

	err := CreateSession(g, "sessionToken", "username", user)
	if err != nil {
		g.Error(err)
		return
	}
	g.HTML(http.StatusOK, "index.html", user)
	g.Redirect(http.StatusOK, "/")

}
