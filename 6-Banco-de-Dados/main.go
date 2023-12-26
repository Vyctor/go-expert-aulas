package main

import (
	"database/sql"
	"fmt"

	_ "github.com/go-sql-driver/mysql"
	"github.com/google/uuid"
)

type Product struct {
	ID    string
	Name  string
	Price float64
}

func main() {
	db, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/goexpert")
	if err != nil {
		panic(err)
	}
	defer db.Close()
	product := NewProduct("Product 1", 19.99)
	err = insertProduct(db, product)
	if err != nil {
		panic(err)
	}
	product.Price = 100000.00
	err = updateProduct(db, product)
	if err != nil {
		panic(err)
	}

	p, err := selectOneProduct(db, product.ID)

	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v, com id %v possui o preço de %.2f\n", p.Name, p.ID, p.Price)

	err = deleteProduct(db, product)
	if err != nil {
		panic(err)
	}

	fmt.Printf("Product: %v, com id %v foi deletado\n", p.Name, p.ID)

	products, err := selectAllProducs(db)

	if err != nil {
		panic(err)
	}

	for _, product := range products {
		fmt.Printf("ID: %v. Name: %v Price: %.2f\n", product.ID, product.Name, product.Price)
	}

}

func NewProduct(name string, price float64) *Product {
	return &Product{
		ID:    uuid.New().String(),
		Name:  name,
		Price: price,
	}
}

func insertProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("insert into products(id, name, price) values(?, ?, ?)")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID, product.Name, product.Price)
	if err != nil {
		return err
	}
	return nil
}

func updateProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("update products set name = ?, price = ? where id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.Name, product.Price, product.ID)
	if err != nil {
		return err
	}
	return nil
}

func selectOneProduct(db *sql.DB, id string) (*Product, error) {
	var product Product
	stmt, err := db.Prepare("select id, name, price from products where id = ?")
	if err != nil {
		return &product, err
	}
	defer stmt.Close()
	err = stmt.QueryRow(id).Scan(&product.ID, &product.Name, &product.Price)
	if err != nil {
		return &product, err
	}
	return &product, nil
}

func selectAllProducs(db *sql.DB) ([]Product, error) {
	rows, err := db.Query("SELECT * from products")
	if err != nil {
		return nil, err
	}
	defer rows.Close()
	var products []Product
	for rows.Next() {
		var product Product
		err := rows.Scan(&product.ID, &product.Name, &product.Price)
		if err != nil {
			return nil, err
		}
		products = append(products, product)
	}
	return products, nil
}

func deleteProduct(db *sql.DB, product *Product) error {
	stmt, err := db.Prepare("DELETE FROM products WHERE id = ?")
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(product.ID)
	if err != nil {
		return err
	}
	return nil
}
