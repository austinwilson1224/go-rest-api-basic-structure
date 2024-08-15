package controller

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

func TestController(c *gin.Context) {
	fmt.Println("test controller called ")
	c.JSON(http.StatusOK, gin.H{
		"message": "test success",
	})
}


func TestController2(c *gin.Context) {
	fmt.Println("test controller2 called ")
	c.JSON(http.StatusOK, gin.H{
		"message222": "test success????",
	})
}