package main

import (
	"net/http"
	"github.com/gin-gonic/gin"
)

func main() {
	router := gin.Default()

	// GET /
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "hallo world")
	})

	// GET /get-json
	router.GET("/get-json", func(c *gin.Context) {
		c.JSON(http.StatusOK, gin.H{
			"msg": "hallo world",
		})
	})
	
	// GET /get/:id
	router.GET("/get/", func(c *gin.Context) {
		id := c.Param("id")
		c.String(http.StatusOK, "hello %s", id)
	})

	log.Printf("Go gin webserver running. Access it at: http://localhost:8080")
	router.Run(":8080")
}