package handlers

import (
	"encoding/gob"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/gorilla/sessions"
	"main.go/models"
)

var Cookie = sessions.NewCookieStore([]byte("1234567890"))

func CreateSession(g *gin.Context, sessionName string, Key string, Value models.User) error {
	gob.Register(models.User{})
	session, err := Cookie.Get(g.Request, sessionName)
	if err != nil {
		return err
	}
	session.Values[Key] = Value
	return session.Save(g.Request, g.Writer)

}

func GetSessionValue(g *gin.Context, sessionName string, Key string) interface{} {
	gob.Register(models.User{})
	session, _ := Cookie.Get(g.Request, sessionName)
	fmt.Println(session)
	return session.Values[Key]
}

func SessionDelete(g *gin.Context, sessionName string) error {
	gob.Register(models.User{})
	session, err := Cookie.Get(g.Request, sessionName)
	if err != nil {
		return err
	}
	session.Options.MaxAge = -1
	return session.Save(g.Request, g.Writer)
}
