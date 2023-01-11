package product

import (
	"ejercicio3/internal/domain"
	"errors"
	"fmt"
)

type repository struct {
	db *[]domain.Product
}

type Repository interface {
	GetById(id int) (product domain.Product, err error)
	GetAll() (products []domain.Product, err error)
	Create(p domain.Product) (product domain.Product, err error)
	PingPong() (pong string, err error)
	GetProducstByPrice(price float64) (products []domain.Product, err error)
	Delete(id int) (msg string, err error)
	Update(id int, p domain.Product) (prod domain.Product, err error)
	PartialUpdate(id int, product domain.Product) (prod domain.Product, err error)
}

var (
	ErrCodeValue  error = errors.New("error: código repetido")
	ErrIdNotFound error = errors.New("error: id no hallado")
)

func NewRepository(db *[]domain.Product) Repository {
	return &repository{db: db}
}

func (r *repository) GetById(id int) (product domain.Product, err error) {
	aux, _ := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	for _, v := range *r.db {
		if v.Id == id {
			product = v
			break
		}
	}
	return
}

func (r *repository) GetAll() (products []domain.Product, err error) {
	products = *r.db
	return
}

func (r *repository) GetProducstByPrice(price float64) (products []domain.Product, err error) {
	for _, v := range *r.db {
		if v.Price > price {
			products = append(products, v)
		}
	}
	return
}

func (r *repository) Create(p domain.Product) (product domain.Product, err error) {

	option, _ := r.checkCodeValue(p.CodeValue)
	if option {
		err = ErrCodeValue
		return
	}

	len := len(*r.db)
	db := *r.db
	p.Id = db[len-1].Id + 1
	product = p
	*r.db = append(*r.db, p)
	return
}

func (r *repository) PingPong() (pong string, err error) {
	pong = "Pong"
	return
}

func (r *repository) Update(id int, p domain.Product) (prod domain.Product, err error) {
	aux, index := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}

	option, ind := r.checkCodeValue(p.CodeValue)
	db := *r.db

	if option && db[ind].CodeValue != p.CodeValue {
		err = ErrCodeValue
		return
	}

	p.Id = id
	db[index] = p
	prod = p
	return
}

func (r *repository) Delete(id int) (msg string, err error) {

	aux, index := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	db := *r.db
	*r.db = append(db[:index], db[index+1:]...)
	msg = fmt.Sprintf("Producto con id: %d borrado correctamente", id)
	return
}

func (r *repository) PartialUpdate(id int, product domain.Product) (prod domain.Product, err error) {
	aux, index := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	db := *r.db
	db[index] = product
	prod = db[index]
	return
}

func (r *repository) checkCodeValue(codeValue string) (option bool, index int) {

	for i, v := range *r.db {
		if v.CodeValue == codeValue {
			option = true
			index = i
			break
		}
	}
	return
}

func (r *repository) checkId(id int) (aux bool, index int) {
	for i, v := range *r.db {
		if v.Id == id {
			aux = true
			index = i
			break
		}
	}
	return
}
