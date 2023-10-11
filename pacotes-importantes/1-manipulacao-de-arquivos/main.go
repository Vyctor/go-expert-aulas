package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	// Criação e escrita
	f, error := os.Create("arquivo.txt")

	if error != nil {
		panic(error)
	}

	//tamanho, error := f.WriteString("Olá, mundo!")
	tamanho, error := f.Write([]byte("Hello, world!"))

	if error != nil {
		panic(error)
	}

	fmt.Printf("Tamanho: %d bytes\n", tamanho)

	defer f.Close()

	// leitura
	arquivo, error := os.ReadFile("arquivo.txt")
	if error != nil {
		panic(error)
	}

	fmt.Println(string(arquivo))

	// leitura de pouco em pouco abrindo o arquivo
	arquivo2, error := os.Open("arquivo.txt")

	if error != nil {
		panic(error)
	}

	reader := bufio.NewReader(arquivo2)
	buffer := make([]byte, 3)

	for {
		n, err := reader.Read(buffer)
		if err != nil {
			break
		}

		fmt.Println(string(buffer[:n]))
	}

	err := os.Remove("arquivo.txt")

	if err != nil {
		panic(err)
	}

}
