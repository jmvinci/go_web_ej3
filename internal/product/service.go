package product

import (
	"ejercicio3/internal/domain"
	"encoding/json"
	"io"
)

type Service interface {
	GetById(id int) (product domain.Product, err error)
	GetAll() (products []domain.Product, err error)
	GetProductsByPrice(price float64) (product []domain.Product, err error)
	PingPong() (pong string, err error)
	Create(p domain.Product) (product domain.Product, err error)
	Delete(id int) (msg string, err error)
	Update(id int, p domain.Product) (prod domain.Product, err error)
	PartialUpdate(id int, body io.ReadCloser) (prod domain.Product, err error)
}

type service struct {
	rp Repository
}

func NewService(rp *Repository) Service {
	return &service{rp: *rp}
}

func (s *service) GetById(id int) (product domain.Product, err error) {
	return s.rp.GetById(id)
}

func (s *service) GetAll() (product []domain.Product, err error) {
	return s.rp.GetAll()
}

func (s *service) GetProductsByPrice(price float64) (product []domain.Product, err error) {
	return s.rp.GetProducstByPrice(price)
}
func (s *service) PingPong() (pong string, err error) {
	return s.rp.PingPong()
}
func (s *service) Create(p domain.Product) (product domain.Product, err error) {
	return s.rp.Create(p)
}

func (s *service) Delete(id int) (msg string, err error) {
	return s.rp.Delete(id)
}

func (s *service) Update(id int, p domain.Product) (prod domain.Product, err error) {
	return s.rp.Update(id, p)
}

func (s *service) PartialUpdate(id int, body io.ReadCloser) (prod domain.Product, err error) {
	p, err := s.GetById(id)
	if err != nil {
		return
	}

	if err = json.NewDecoder(body).Decode(&p); err != nil {
		return
	}

	return s.rp.PartialUpdate(id, p)
}
