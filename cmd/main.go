package main

import (
	"ejercicio3/cmd/server/routes"
	"ejercicio3/pkg"
	"log"

	"github.com/gin-gonic/gin"
)

func main() {
	db, err := pkg.GetProductsStruct()

	if err != nil {
		panic(err)
	}

	en := gin.Default()
	rt := routes.NewRouter(en, &db)
	rt.SetProduct()

	if err := en.Run(); err != nil {
		log.Fatal(err)
	}
}
