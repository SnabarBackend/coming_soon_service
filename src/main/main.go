package main

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	// Getting default API Engine
	// of Gin
	r := gin.Default()

	// GET handler to test
	// health of this service
	r.GET("/healthz", func(c *gin.Context) {
		c.JSON(http.StatusOK,gin.H{
			"msg": "OK",
		})
	})

	// Starting server on 8080 port
	r.Run(":8080")
}