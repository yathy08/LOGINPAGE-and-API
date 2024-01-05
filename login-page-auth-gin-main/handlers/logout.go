package handlers

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
	"main.go/cache"
)

func Logout(g *gin.Context) {
	err := SessionDelete(g, "sessionToken")
	if err != nil {
		fmt.Println("Error Ocurred:", err)
		return
	}
	cache.ClearCache(g)
	g.HTML(http.StatusOK, "login.html", "successfully logged out")
}
