package store

import (
	"ejercicio3/internal/domain"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"os"
)

var (
	ErrCodeValue  error = errors.New("error: código repetido")
	ErrIdNotFound error = errors.New("error: id no hallado")
)

type Products interface {
	GetById(id int) (prod domain.Product, err error)
	Delete(id int) (msg string, err error)
}

type products struct {
	Products *[]domain.Product
}

func NewStore() Products {
	var prod products
	db, err := prod.getProductsStruct()
	if err != nil {
		panic(err)
	}
	return &products{Products: db.Products}
}

func (p *products) openJsonFile() (jsonFile *os.File, err error) {
	jsonFile, err = os.Open("products.json")

	if err != nil {
		return
	}
	return
}

func (p *products) getProductsStruct() (products products, err error) {

	jsonFile, err := p.openJsonFile()

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

func (p *products) GetById(id int) (prod domain.Product, err error) {
	for _, v := range *p.Products {
		if v.Id == id {
			prod = v
			break
		}
	}
	return
}

func (p *products) Delete(id int) (msg string, err error) {

	aux, index := p.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	db := *p.Products
	*p.Products = append(db[:index], db[index+1:]...)
	msg = fmt.Sprintf("Producto con id: %d borrado correctamente", id)
	return
}

func (p *products) Update(id int, pr domain.Product) (prod domain.Product, err error) {
	aux, index := p.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}

	option, ind := p.checkCodeValue(pr.CodeValue)
	db := *p.Products

	if option && db[ind].CodeValue != pr.CodeValue {
		err = ErrCodeValue
		return
	}

	pr.Id = id
	db[index] = pr
	prod = pr
	return
}

func (p *products) checkId(id int) (aux bool, index int) {
	for i, v := range *p.Products {
		if v.Id == id {
			aux = true
			index = i
			break
		}
	}
	return
}

func (p *products) checkCodeValue(codeValue string) (option bool, index int) {

	for i, v := range *p.Products {
		if v.CodeValue == codeValue {
			option = true
			index = i
			break
		}
	}
	return
}
