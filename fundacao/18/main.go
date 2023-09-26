package main

import (
	"fmt"

	"github.com/vyctor/curso-go/matematica"
)

func main() {
	soma := matematica.Soma(10, 10)
	fmt.Println("Resultado: ", soma)
}
