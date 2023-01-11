package main

import (
	"ejercicio3/cmd/server/routes"
	"ejercicio3/pkg"
	"fmt"

	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
)

func main() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println(err)
		}
		fmt.Println("Ejecución finalizada")
	}()
	err := godotenv.Load()
	if err != nil {
		panic(err)
	}
	db, err := pkg.GetProductsStruct()

	if err != nil {
		panic(err)
	}

	en := gin.Default()
	rt := routes.NewRouter(en, &db)
	rt.SetProduct()

	if err := en.Run(); err != nil {
		panic(err)
	}
}
