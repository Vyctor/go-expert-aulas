package main

import "fmt"

func main() {
	salarios := map[string]int{"Wesley": 1000, "Amanda": 2000, "João": 3000}
	// fmt.Printf("O salário de Wesley é %d\n", salarios["Wesley"])
	// fmt.Printf("O salário de Amanda é %d\n", salarios["Amanda"])
	// fmt.Printf("O salário de João é %d\n", salarios["João"])
	delete(salarios, "João")
	// fmt.Printf("O salário de João é %d\n", salarios["João"])
	salarios["João"] = 4000
	// fmt.Printf("O salário de João é %d\n", salarios["João"])

	sal := make(map[string]int)
	sal["Wesley"] = 1000

	sal1 := map[string]int{}
	sal1["Wesley"] = 1000

	for key, value := range salarios {
		fmt.Printf("O salário de %s é %d\n", key, value)
	}

	totalSalarios := 0

	for _, value := range salarios {
		totalSalarios += value
	}

	fmt.Printf("O total de salários é %d\n", totalSalarios)
}
