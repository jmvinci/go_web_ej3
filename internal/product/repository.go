package product

import (
	"ejercicio3/internal/domain"
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
}

type ErrCodeValue struct{}

func (e *ErrCodeValue) Error() string {
	return "Código repetido"
}

func NewRepository(db *[]domain.Product) Repository {
	return &repository{db: db}
}

func (r *repository) GetById(id int) (product domain.Product, err error) {
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

	if r.checkCodeValue(p.CodeValue) {
		err = &ErrCodeValue{}
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

func (r *repository) checkCodeValue(codeValue string) (option bool) {

	for _, v := range *r.db {
		if v.CodeValue == codeValue {
			option = true
			break
		}
	}
	return
}
