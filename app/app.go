package app

import (
	"fmt"
	"net/http"
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func StartApp() {

	fmt.Println("app starting ... ")


	router.GET("/test", testFunc)
	router.Run()
}

func testFunc(c *gin.Context) {
	c.JSON(http.StatusAccepted, gin.H{
		"message": "great success",
	})
}