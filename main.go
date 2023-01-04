package main

import (
	"ejercicio3/models"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

func main() {

	r := gin.Default()
	p := r.Group("/products")

	r.GET("/ping", func(ctx *gin.Context) {
		ctx.String(http.StatusOK, "pong")
	})

	p.GET("/", func(ctx *gin.Context) {

		// process
		products, err := GetProductsStruct()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})
			return
		}

		// response
		ctx.JSON(http.StatusOK, models.Response{Message: "ok", Data: products})

	})
	p.GET("/:id", func(ctx *gin.Context) {

		//request
		id, err := strconv.Atoi(ctx.Param("id"))

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})

			return
		}
		// process
		products, err := GetProductsStruct()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})

			return
		}

		var prod models.Product

		for _, v := range products.Products {
			if v.Id == id {
				prod = v
				break
			}
		}

		// response

		ctx.JSON(http.StatusOK, models.Response{Message: "ok", Data: prod})
	})
	p.GET("/search", func(ctx *gin.Context) {

		// request
		price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)

		// process

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})

			return
		}

		products, err := GetProductsStruct()

		if err != nil {
			ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})

			return
		}

		var p models.Products
		for _, v := range products.Products {
			if v.Price > price {
				p.Products = append(p.Products, v)
			}
		}

		// response

		ctx.JSON(http.StatusOK, models.Response{Message: "ok", Data: p})

	})

	r.Run()
}

func OpenJsonFile() (jsonFile *os.File, err error) {
	jsonFile, err = os.Open("products.json")

	if err != nil {
		return
	}
	return
}

func GetProductsStruct() (products models.Products, err error) {

	jsonFile, err := OpenJsonFile()

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &products.Products); err != nil {
		return
	}

	defer jsonFile.Close()
	return
}
