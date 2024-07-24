package handlers

import (
	"github.com/gin-gonic/gin"
	"log"
	"messagio-gin-postrgresql-kafka/internal/db"
	"messagio-gin-postrgresql-kafka/internal/models"
	"net/http"
)

func indexHandler(c *gin.Context) {
	c.HTML(http.StatusOK, "index.html", gin.H{
		"Error": nil,
	})
}

func messagesRegistrationHandler(c *gin.Context) {
	var id int
	var err error
	var message models.Message

	if err = c.BindJSON(&message); err != nil {
		log.Printf("Error - parsing the message: %s", err)
		c.JSON(http.StatusBadRequest, gin.H{
			"Error": err.Error(),
		})
		return
	}

	if id, err = message.Create(); err != nil {
		log.Printf("Error - creating message: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
		"id":    id,
	})
}

func handlerInfo(c *gin.Context) {
	msg, err := db.GetAllMessages()
	if err != nil {
		log.Printf("Error - getting messages: %s", err)
		c.JSON(http.StatusInternalServerError, gin.H{
			"Error": err.Error(),
		})
	}
	log.Println(msg)
	c.JSON(http.StatusOK, gin.H{
		"Error": nil,
		"data":  msg,
	})
}
