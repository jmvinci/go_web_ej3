package routes

import (
	handlers "ejercicio3/cmd/server/handler"
	"ejercicio3/cmd/server/middlewares"
	"ejercicio3/internal/product"
	"ejercicio3/pkg/store"

	"github.com/gin-gonic/gin"
)

type Router struct {
	st store.Storage

	en *gin.Engine
}

func NewRouter(en *gin.Engine, storage *store.Storage) *Router {
	return &Router{en: en, st: *storage}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// website
func (r *Router) SetProduct() {
	// instances
	rp := product.NewRepository(&r.st)
	sv := product.NewService(&rp)
	h := handlers.NewProduct(sv)

	p := r.en.Group("/product")

	r.en.GET("/ping", h.PingPong())

	p.GET("/", h.GetAll())
	p.GET("/:id", h.GetById())
	p.GET("/search", h.GetProductsByPrice())

	p.Use(middlewares.AuthMiddleware())
	{
		p.POST("/", h.Create())
		p.DELETE("/:id", h.DeleteById())
		p.PUT("/:id", h.UpdateProduct())
		p.PATCH("/:id", h.PartialUpdate())
	}

}
