package main

import "fmt"

func main() {
	var meuArray [3]int

	meuArray[0] = 1
	meuArray[1] = 2
	meuArray[2] = 3

	fmt.Println(meuArray[len(meuArray)-1])
	fmt.Println(meuArray[0])

	for index, value := range meuArray {
		fmt.Printf("O valor do índice é %d e o valor é %d\n", index, value)
	}

}
