package handlers

import (
	"ejercicio3/internal/domain"
	"ejercicio3/internal/product"
	"ejercicio3/pkg"
	"net/http"
	"strconv"

	"github.com/gin-gonic/gin"
)

// controlador de product

type Product struct {
	sv product.Service
}

func NewProduct(sv product.Service) *Product {
	return &Product{sv: sv}
}

func (p *Product) GetById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}

		prod, err := p.sv.GetById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}
		ctx.JSON(http.StatusOK, pkg.Response{Message: "ok", Data: prod})

	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.sv.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}
		ctx.JSON(http.StatusOK, pkg.Response{Message: "ok", Data: products})

	}
}

func (p *Product) PingPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := p.sv.PingPong()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}
		ctx.String(http.StatusOK, response)
	}
}

func (p *Product) GetProductsByPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}

		products, err := p.sv.GetProductsByPrice(price)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}

		ctx.JSON(http.StatusOK, pkg.Response{Message: "ok", Data: products})
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product
		if err := ctx.ShouldBind(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}
		p, err := p.sv.Create(product)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, pkg.Response{Message: err.Error(), Data: nil})
			return
		}

		ctx.JSON(http.StatusOK, pkg.Response{Message: "ok", Data: p})

	}
}
