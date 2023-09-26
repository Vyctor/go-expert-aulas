package main

import "fmt"

func main() {
	var minhaVar interface{} = "Vyctor Guimarães"
	println(minhaVar.(string))

	resultado, ok := minhaVar.(int)

	fmt.Printf("O valor de res é %v e o resultado de ok é %v", resultado, ok)

	res2 := minhaVar.(int)
	fmt.Printf("O valor de res2 é %v", res2)
}
