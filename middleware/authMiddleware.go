package middleware

import (
	"ejercicio3/pkg"
	"log"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func AuthMiddleware() gin.HandlerFunc {
	requiredToken := os.Getenv("TOKEN")
	if requiredToken == "" {
		log.Fatal("Please set API_TOKEN environment variable")
	}
	return func(ctx *gin.Context) {
		token := ctx.GetHeader("TOKEN")
		if token == "" {
			ctx.JSON(http.StatusUnauthorized, pkg.Response{Message: "API token required", Data: nil})
			ctx.Abort()
			return
		}
		if token != requiredToken {
			ctx.JSON(http.StatusUnauthorized, pkg.Response{Message: "Not token provided", Data: nil})
			ctx.Abort()
			return
		}
		ctx.Next()
	}
}
