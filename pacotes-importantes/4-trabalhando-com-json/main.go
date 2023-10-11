package main

import (
	"encoding/json"
	"os"
)

type Conta struct {
	Numero int `json:"n"`
	Saldo  int `json:"s"`
}

func main() {
	conta := Conta{Numero: 1, Saldo: 1000}

	res, err := toJson(conta)

	if err != nil {
		println(err)
	}

	println(string(res))

	encoder := json.NewEncoder(os.Stdout)
	encoder.Encode(conta)

	var contaX Conta

	json.Unmarshal(res, &contaX)

	encoder.Encode(contaX)

}

func toJson(v any) ([]byte, error) {
	return json.Marshal(v)
}
