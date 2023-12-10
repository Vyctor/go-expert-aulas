package main

import (
	"io/ioutil"
	"net/http"
)

func main() {
	httpClient := http.Client{}

	request, err := http.NewRequest("POST", "http://google.com", nil)

	if err != nil {
		panic(err)
	}
	request.Header.Set("Accept", "Application/json")
	response, err := httpClient.Do(request)

	if err != nil {
		panic(err)
	}

	defer response.Body.Close()

	body, err := ioutil.ReadAll(response.Body)

	if err != nil {
		panic(err)
	}

	println(string(body))
}
