package main

import (
	"io"
	"net/http"
)

func main() {
	chamada, err := http.Get("https://www.google.com.br")

	if err != nil {
		panic(err)
	}

	res, err := io.ReadAll(chamada.Body)

	if err != nil {
		panic(err)
	}

	println(string(res))

	defer chamada.Body.Close()
}
