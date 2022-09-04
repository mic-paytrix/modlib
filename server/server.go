package server

import (
	"log"

	"github.com/gin-gonic/gin"
)

func GinServer() {
	r := gin.Default()
	r.GET("/ping", func(c *gin.Context) {
		log.Print("received a ping, responding")
		c.JSON(200, gin.H{
			"reply": "pong",
		})
	})
	r.Run() // listen and serve on 0.0.0.0:8080
}
