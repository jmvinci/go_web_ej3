package store

import (
	"ejercicio3/internal/domain"
	"encoding/json"
	"errors"
	"io"
	"os"
)

var (
	ErrCodeValue  error = errors.New("error: código repetido")
	ErrIdNotFound error = errors.New("error: id no hallado")
)

type Storage interface {
	Get() (products []domain.Product, err error)
	Set(newJson []domain.Product) (err error)
}

type storage struct {
	Path string
}

func NewStore(path string) Storage {
	return &storage{Path: path}
}

func (s *storage) Get() (products []domain.Product, err error) {
	jsonFile, err := os.Open(s.Path)

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
func (s *storage) Set(newJson []domain.Product) (err error) {
	byteValue, err := json.Marshal(newJson)
	if err != nil {
		return
	}
	if err = os.WriteFile(s.Path, byteValue, 0644); err != nil {
		return
	}
	return
}
