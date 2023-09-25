package main

import "fmt"

type Endereco struct {
	Logradouro string
	Numero     int
	Bairro     string
	Cidade     string
	Estado     string
}

func (client *Cliente) Desativar() {
	client.Ativo = false
	fmt.Printf("O Cliente %s foi desativado!", client.Nome)
}

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
	Endereco
}

func main() {
	vyctor := Cliente{
		Nome:  "Vyctor",
		Idade: 28,
		Ativo: true,
		Endereco: Endereco{
			Logradouro: "Rua dos Bobos",
			Numero:     0,
			Bairro:     "Centro",
			Cidade:     "São Paulo",
			Estado:     "SP",
		},
	}

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", vyctor.Nome, vyctor.Idade, vyctor.Ativo)

	vyctor.Desativar()
}
