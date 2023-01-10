package pkg

import (
	"ejercicio3/internal/domain"
	"encoding/json"
	"io"
	"os"
)

func openJsonFile() (jsonFile *os.File, err error) {
	jsonFile, err = os.Open("products.json")

	if err != nil {
		return
	}
	return
}

func GetProductsStruct() (products []domain.Product, err error) {

	jsonFile, err := openJsonFile()

	if err != nil {
		return
	}

	byteValue, err := io.ReadAll(jsonFile)

	if err != nil {
		return
	}

	if err = json.Unmarshal(byteValue, &products); err != nil {
		return
	}

	defer jsonFile.Close()

	return
}
