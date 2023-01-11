package handlers

import (
	"ejercicio3/internal/domain"
	"ejercicio3/internal/product"
	"ejercicio3/pkg/web"
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
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		prod, err := p.sv.GetById(id)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		ctx.JSON(http.StatusOK, web.Response{Message: "ok", Data: prod})

	}
}

func (p *Product) GetAll() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		products, err := p.sv.GetAll()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		ctx.JSON(http.StatusOK, web.Response{Message: "ok", Data: products})

	}
}

func (p *Product) PingPong() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		response, err := p.sv.PingPong()
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		ctx.String(http.StatusOK, response)
	}
}

func (p *Product) GetProductsByPrice() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		products, err := p.sv.GetProductsByPrice(price)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		ctx.JSON(http.StatusOK, web.Response{Message: "ok", Data: products})
	}
}

func (p *Product) Create() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		var product domain.Product
		if err := ctx.ShouldBindJSON(&product); err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		p, err := p.sv.Create(product)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		ctx.JSON(http.StatusOK, web.Response{Message: "ok", Data: p})

	}
}

func (p *Product) DeleteById() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		msg, err := p.sv.Delete(id)
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		ctx.JSON(http.StatusOK, web.Response{Message: msg, Data: nil})
	}
}

func (p *Product) UpdateProduct() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}
		var prod domain.Product
		if err := ctx.ShouldBind(&prod); err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		pr, err := p.sv.Update(id, prod)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		ctx.JSON(http.StatusOK, web.Response{Message: "Producto actualizado", Data: pr})

	}
}

func (p *Product) PartialUpdate() gin.HandlerFunc {
	return func(ctx *gin.Context) {
		id, err := strconv.Atoi(ctx.Param("id"))
		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		prod, err := p.sv.PartialUpdate(id, ctx.Request.Body)

		if err != nil {
			ctx.JSON(http.StatusBadRequest, web.ErrorResponse{Message: err.Error(), Status: http.StatusBadRequest, Code: "400"})
			return
		}

		ctx.JSON(http.StatusOK, web.Response{Message: "Producto actualizado", Data: prod})

	}
}
