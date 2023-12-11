package main

import "fmt"

type Cliente struct {
	Nome  string
	Idade int
	Ativo bool
}

func main() {
	vyctor := Cliente{
		Nome:  "Vyctor",
		Idade: 28,
		Ativo: true,
	}

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", vyctor.Nome, vyctor.Idade, vyctor.Ativo)

	vyctor.Ativo = false

	fmt.Printf("Nome: %s\nIdade: %d\nAtivo: %t\n", vyctor.Nome, vyctor.Idade, vyctor.Ativo)

}
