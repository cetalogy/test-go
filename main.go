package main

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func main() {
	// Initialize the Gin router
	router := gin.Default()

	// Define a GET route to handle "/"
	router.GET("/", func(c *gin.Context) {
		c.String(http.StatusOK, "Hello, World! test pushsas")
	})

	// Run the server on port 8080
	router.Run(":8080")
}
