package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/cache"
)

func GetHome(g *gin.Context) {
	cache.ClearCache(g)

	userSession := GetSessionValue(g, "sessionToken", "username")
	if userSession != nil {
		fmt.Println("home loaded", userSession)
		g.HTML(http.StatusSeeOther, "index.html", userSession)
		return
	}
	fmt.Println("loginpage")
	g.HTML(http.StatusOK, "login.html", nil)
}
