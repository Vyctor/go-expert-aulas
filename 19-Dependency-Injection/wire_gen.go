// Code generated by Wire. DO NOT EDIT.

//go:generate go run -mod=mod github.com/google/wire/cmd/wire
//go:build !wireinject
// +build !wireinject

package main

import (
	"database/sql"
	"dependency-injection/product"
	"github.com/google/wire"
)

import (
	_ "github.com/mattn/go-sqlite3"
)

// Injectors from wire.go:

func NewUseCase(db *sql.DB) *product.ProductUsecase {
	productRepository := product.NewProductRepository(db)
	productUsecase := product.NewProductUseCase(productRepository)
	return productUsecase
}

// wire.go:

var setRepositoryDependency = wire.NewSet(product.NewProductRepository, wire.Bind(new(product.ProductRepositoryInterface), new(*product.ProductRepository)))
