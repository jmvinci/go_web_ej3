package main

import (
	"ejercicio3/handlers"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	p := r.Group("/products")

	r.GET("/ping", handlers.PingPong)
	p.GET("/", handlers.GetProducts)
	p.GET("/:id", handlers.GetProductById)
	p.GET("/search", handlers.GetProductByPrice)
	p.POST("/", handlers.AddProduct)

	r.Run()
}
