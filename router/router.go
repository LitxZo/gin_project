package router

import (
	"github.com/gin-gonic/gin"
)

func WebStart() *gin.Engine {
	router := gin.Default()

	return router
}
