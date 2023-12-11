package main

import (
	"io/ioutil"
	"net/http"
	"time"
)

func main() {
	c := http.Client{Timeout: time.Duration(300 * time.Millisecond)}
	res, err := c.Get("http://www.google.com.br")

	if err != nil {
		panic(err)
	}

	defer res.Body.Close()
	body, err := ioutil.ReadAll(res.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
