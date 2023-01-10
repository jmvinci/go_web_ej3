package routes

import (
	handlers "ejercicio3/cmd/server/handler"
	"ejercicio3/internal/domain"
	"ejercicio3/internal/product"

	"github.com/gin-gonic/gin"
)

type Router struct {
	db *[]domain.Product
	en *gin.Engine
}

func NewRouter(en *gin.Engine, db *[]domain.Product) *Router {
	return &Router{en: en, db: db}
}

func (r *Router) SetRoutes() {
	r.SetProduct()
}

// website
func (r *Router) SetProduct() {
	// instances
	rp := product.NewRepository(r.db)
	sv := product.NewService(&rp)
	h := handlers.NewProduct(sv)

	p := r.en.Group("/product")

	r.en.GET("/ping", h.PingPong())
	p.GET("/", h.GetAll())
	p.GET("/:id", h.GetById())
	p.GET("/search", h.GetProductsByPrice())
	p.POST("/", h.Create())

}
