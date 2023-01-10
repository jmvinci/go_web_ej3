package product

import "ejercicio3/internal/domain"

type Service interface {
	GetById(id int) (product domain.Product, err error)
	GetAll() (products []domain.Product, err error)
	GetProductsByPrice(price float64) (product []domain.Product, err error)
	PingPong() (pong string, err error)
	Create(p domain.Product) (product domain.Product, err error)
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
