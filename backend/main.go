package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func helloWorld(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Hello World",
	})
}

func main() {
	r := gin.Default()

	r.GET("/", helloWorld)

	r.Run()
}
