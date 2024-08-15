package app

import (
	"github.com/austinwilson1224/go-rest-api-basic-structure/controller"
)

func MapUrls() {
	router.GET("/test-package", controller.TestController2)
}