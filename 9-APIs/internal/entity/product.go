package entity

import (
	"apis/pkg/entity"
	"errors"
	"time"
)

var (
	ErrorIdIsRequired    = errors.New("id is required")
	ErrorNameIsRequired  = errors.New("name is required")
	ErrorPriceIsRequired = errors.New("price is required")
	ErrorInvalidIDI      = errors.New("invalid id")
	ErrorInvalidPrice    = errors.New("invalid price")
)

type Product struct {
	ID        entity.ID `json:"id"`
	Name      string    `json:"name"`
	Price     float64   `json:"price"`
	CreatedAt time.Time `json:"created_at"`
}

func NewProduct(name string, price float64) (*Product, error) {
	id := entity.NewID()

	product := &Product{
		ID:        id,
		Name:      name,
		Price:     price,
		CreatedAt: time.Now(),
	}

	if err := product.Validate(); err != nil {
		return nil, err
	}

	return product, nil
}

func (p *Product) Validate() error {
	if p.ID.String() == "" {
		return ErrorIdIsRequired
	}
	if _, err := entity.ParseID(p.ID.String()); err != nil {
		return ErrorInvalidIDI
	}

	if p.Name == "" {
		return ErrorNameIsRequired
	}

	if p.Price == 0 {
		return ErrorPriceIsRequired
	}

	if p.Price < 0 {
		return ErrorInvalidPrice
	}

	return nil
}
