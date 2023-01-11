package product

import (
	"ejercicio3/internal/domain"
	"ejercicio3/pkg/store"
	"errors"
	"fmt"
)

type repository struct {
	st *store.Storage
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

func NewRepository(st *store.Storage) Repository {
	return &repository{st: st}
}

func (r *repository) GetById(id int) (product domain.Product, err error) {
	aux, _ := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	db, err := (*r.st).Get()

	for _, v := range db {
		if v.Id == id {
			product = v
			break
		}
	}
	return
}

func (r *repository) GetAll() (products []domain.Product, err error) {
	products, err = (*r.st).Get()
	fmt.Println(products)
	if err != nil {
		return
	}
	return
}

func (r *repository) GetProducstByPrice(price float64) (products []domain.Product, err error) {
	db, err := (*r.st).Get()
	if err != nil {
		return
	}
	for _, v := range db {
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
	db, err := (*r.st).Get()
	if err != nil {
		return
	}
	len := len(db)
	p.Id = db[len-1].Id + 1
	product = p
	db = append(db, p)
	if err = (*r.st).Set(db); err != nil {
		return
	}
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
	db, err := (*r.st).Get()

	if err != nil {
		return
	}

	if option && db[ind].CodeValue != p.CodeValue {
		err = ErrCodeValue
		return
	}

	p.Id = id
	db[index] = p
	if err = (*r.st).Set(db); err != nil {
		return
	}
	prod = p
	return
}

func (r *repository) Delete(id int) (msg string, err error) {

	aux, index := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	db, err := (*r.st).Get()
	if err != nil {
		return
	}
	db = append(db[:index], db[index+1:]...)
	if err = (*r.st).Set(db); err != nil {
		return
	}
	msg = fmt.Sprintf("Producto con id: %d borrado correctamente", id)
	return
}

func (r *repository) PartialUpdate(id int, product domain.Product) (prod domain.Product, err error) {
	aux, index := r.checkId(id)
	if !aux {
		err = ErrIdNotFound
		return
	}
	db, err := (*r.st).Get()
	if err != nil {
		return
	}
	db[index] = product
	if err = (*r.st).Set(db); err != nil {
		return
	}
	prod = db[index]
	return
}

func (r *repository) checkCodeValue(codeValue string) (option bool, index int) {
	db, _ := (*r.st).Get()
	for i, v := range db {
		if v.CodeValue == codeValue {
			option = true
			index = i
			break
		}
	}
	return
}

func (r *repository) checkId(id int) (aux bool, index int) {
	db, _ := (*r.st).Get()
	for i, v := range db {
		if v.Id == id {
			aux = true
			index = i
			break
		}
	}
	return
}
