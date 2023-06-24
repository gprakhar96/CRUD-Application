package app

import (
	"github.com/gin-gonic/gin"
)

var (
	router = gin.Default()
)

func Start() {
	mapURLs()
	router.Run("localhost:1111")
}
