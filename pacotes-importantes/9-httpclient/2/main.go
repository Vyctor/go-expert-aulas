package main

import (
	"bytes"
	"io"
	"net/http"
	"os"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(300 * time.Millisecond)}
	jsonVar := bytes.NewBuffer([]byte(`{"nome":"Curso de Go", "carga_horaria":40}`))
	res, err := c.Post("https://jsonplaceholder.typicode.com/posts", "application/json", jsonVar)

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	io.CopyBuffer(os.Stdout, res.Body, nil)
}
