package main

import (
	"fmt"

	"github.com/google/uuid"
	"github.com/vyctor/curso-go/matematica"
)

func main() {
	soma := matematica.Soma(10, 10)
	fmt.Println("Resultado: ", soma)
	fmt.Println("UUID: ", uuid.New())
}
