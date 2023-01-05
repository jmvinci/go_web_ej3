package handlers

import (
	"ejercicio3/models"
	"encoding/json"
	"io"
	"net/http"
	"os"
	"strconv"

	"github.com/gin-gonic/gin"
)

// Variable global con los productos cargados para mantener la estructura en memoria durante la ejecución
var products, _ = getProductsStruct()

func openJsonFile() (jsonFile *os.File, err error) {
	jsonFile, err = os.Open("products.json")

	if err != nil {
		return
	}
	return
}

func getProductsStruct() (products models.Products, err error) {

	jsonFile, err := openJsonFile()

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

func GetProducts(ctx *gin.Context) {

	// response
	ctx.JSON(http.StatusOK, models.Response{Message: "ok", Data: products})
}

func GetProductById(ctx *gin.Context) {

	//request
	id, err := strconv.Atoi(ctx.Param("id"))

	if err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})

		return
	}
	// process

	var prod models.Product

	for _, v := range products.Products {
		if v.Id == id {
			prod = v
			break
		}
	}

	// response

	ctx.JSON(http.StatusOK, models.Response{Message: "ok", Data: prod})
}

func GetProductByPrice(ctx *gin.Context) {
	// request
	price, err := strconv.ParseFloat(ctx.Query("priceGt"), 64)

	// process

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
}

func AddProduct(ctx *gin.Context) {
	var product models.Product

	if err := ctx.BindJSON(&product); err != nil {
		ctx.JSON(http.StatusBadRequest, models.Response{Message: err.Error(), Data: nil})
	}

	product.Id = len(products.Products) + 1

	products.Products = append(products.Products, product)

	ctx.JSON(http.StatusOK, models.Response{Message: "Agregado correctamente", Data: products})
}

func PingPong(ctx *gin.Context) {
	ctx.String(http.StatusOK, "pong")
}
