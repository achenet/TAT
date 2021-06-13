package main

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"net/http"
)

func main() {
	fmt.Println("Test server")
	r := gin.Default()
	r.GET("/", GetHandler)

	r.Run()
}

func GetHandler(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{"success": "Request received"})
}
