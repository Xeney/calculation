package main

import (
	"app/pkg/logic"
	"encoding/gob"
	"fmt"

	"github.com/gin-contrib/sessions"
	"github.com/gin-contrib/sessions/cookie"
	"github.com/gin-gonic/gin"
)

var router *gin.Engine

var IP_ADDR string = "localhost:5000"

func main() {
	router = gin.Default()
	router.LoadHTMLGlob("ui/html/*")
	router.StaticFile("style.css", "./ui/static/styles/style.css")
	router.StaticFile("icon.svg", "./ui/static/images/icon.svg")
	router.StaticFile("icon-light.svg", "./ui/static/images/icon-light.svg")
	store := cookie.NewStore([]byte("secret"))
	router.Use(sessions.Sessions("DATA", store))
	initializeRoutes()
	gob.Register(logic.User{})
	fmt.Println("\nЗАПУСК СЕРВЕРА http://" + IP_ADDR + "\n ")
	err := router.Run(IP_ADDR)
	if err != nil {
		router.Run("localhost:8080")
	}
}
