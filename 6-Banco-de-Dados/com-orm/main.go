package main

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type Product struct {
	ID           string `gorm:"primaryKey"`
	Name         string
	Price        float64
	CategoryID   string
	Category     Category
	SerialNumber SerialNumber
	gorm.Model
}

type Category struct {
	ID       string `gorm:"primaryKey"`
	Name     string
	Products []Product
}

type SerialNumber struct {
	ID        int `gorm:"primaryKey"`
	Number    string
	ProductID int
}

func main() {
	dsn := "root:root@tcp(localhost:3306)/goexpert?parseTime=true"
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
	db.AutoMigrate(&Product{}, &Category{}, &SerialNumber{})

}

func findWithPreload(db *gorm.DB) (*[]Product, error) {
	var products []Product
	err := db.Preload("Category").Find(&products)
	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return nil, err.Error
	}
	return &products, nil
}

func createWithCategory(db *gorm.DB) {
	category := Category{
		ID:   "1",
		Name: "Category 1",
	}
	product := Product{
		ID:         "1",
		Name:       "ORM PRODUCT 1",
		Price:      19.99,
		CategoryID: "1",
		Category:   category,
	}

	db.Create(&category)
	db.Create(&product)
}

func delete(db *gorm.DB, product *Product) error {
	err := db.Delete(&product)
	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return err.Error
	}
	return nil
}

func update(db *gorm.DB, product *Product) error {
	err := db.Save(&product)
	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return err.Error
	}
	return nil
}

func selectWithLike(db *gorm.DB) (*[]Product, error) {
	var products []Product
	err := db.Where("name LIKE ?", "%ORM%").Order("name").Find(&products)
	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return nil, err.Error
	}
	return &products, nil
}

func selectWithWhere(db *gorm.DB) (*[]Product, error) {
	var products []Product
	err := db.Where("price < 100").Find(&products)
	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return nil, err.Error
	}
	return &products, nil
}

func selectAll(db *gorm.DB, limit int) (*[]Product, error) {
	var products []Product
	err := db.Limit(limit).Offset(1).Find(&products)
	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return nil, err.Error
	}
	return &products, nil
}

func selectOne(db *gorm.DB) (*Product, error) {
	var p Product
	err := db.First(&p)

	if err.Error != nil {
		fmt.Printf("Error: %v\n", err.Error)
		return nil, err.Error
	}
	return &p, nil
}

func autoMigrate(db *gorm.DB) {
	err := db.AutoMigrate(&Product{})
	if err != nil {
		panic(err)
	}
}

func create(db *gorm.DB) *Product {
	var product = Product{
		ID:    "5",
		Name:  "ORM PRODUCT 1",
		Price: 19.99,
	}

	db.Create(&product)
	return &product
}
