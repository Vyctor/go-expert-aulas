package main

import (
	"errors"
	"fmt"
)

func main() {
	result, condition, error := sum(1, 2)
	if error != nil {
		fmt.Println(error.Error())
		return
	}
	fmt.Println(result, condition)
}

func sum(a, b int) (int, bool, error) {
	if a+b > 50 {
		return a + b, true, errors.New("O valor é maior que 50")
	}

	return a + b, false, nil
}
