package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
)

func CreateRouter() *gin.Engine {
	r := gin.Default()
	r.ForwardedByClientIP = true
	if err := r.SetTrustedProxies([]string{"192.168.1.2", "10.0.0.0/8", "87.228.13.94"}); err != nil {
		return nil
	}

	r.Static("/assets/", "frontend/")
	r.LoadHTMLGlob("frontend/templates/*.html")

	r.GET("/", indexHandler)
	r.GET("/api/messages", handlerInfo)
	r.POST("/messages", messagesRegistrationHandler)

	return r
}

func Run(router *gin.Engine) {
	if err := router.Run(":8080"); err != nil {
		log.Fatalf("Error - starting HTTP server: %v", err)
		return
	}
}
