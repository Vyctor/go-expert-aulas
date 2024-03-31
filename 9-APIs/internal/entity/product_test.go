package entity

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestNewProduct(t *testing.T) {
	product, err := NewProduct("Product 1", 10)
	assert.Nil(t, err)
	assert.NotNil(t, product)
	assert.NotEmpty(t, product.ID)
	assert.Equal(t, product.Name, "Product 1")
	assert.Equal(t, product.Price, 10.0)
}

func TestProductWhenNameIsRequired(t *testing.T) {
	product, err := NewProduct("", 10)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrorNameIsRequired)
}

func TestProductWhenPriceIsRequired(t *testing.T) {
	product, err := NewProduct("Product 1", 0)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrorPriceIsRequired)
}

func TestProductWhenPriceIsInvalid(t *testing.T) {
	product, err := NewProduct("Product 1", -10)
	assert.NotNil(t, err)
	assert.Nil(t, product)
	assert.Equal(t, err, ErrorInvalidPrice)
}
