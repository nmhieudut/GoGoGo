package middlewares

import (
	"log"

	"github.com/gin-gonic/gin"
)

func CustomMiddleware(context *gin.Context) {
	log.Print(("I'm in a global middleware"))
	context.Next()
}