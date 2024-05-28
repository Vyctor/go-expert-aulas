package main

import (
	"fmt"
	"os"
)

func main() {
	i := 10000
	for {
		f, err := os.Create(fmt.Sprintf("./tmp/file%d.txt", i))
		if err != nil {
			fmt.Printf("Error creating file: %s\n", err)
			panic(err)
		}
		defer f.Close()
		f.WriteString("Hello, World!")
		i++
		if i == 20000 {
			break
		}
	}
}
