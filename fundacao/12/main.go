package main

import "fmt"

type Cliente struct {
	nome string
}

type Conta struct {
	saldo int
}

func NewConta() *Conta {
	return &Conta{saldo: 0}
}

func (c *Cliente) andou() {
	c.nome = "Vyctor Guimarães"
	fmt.Printf("O cliente %v andou\n", c.nome)
}

func (c Conta) simular(valor int) int {
	c.saldo += valor
	return c.saldo
}

func (c *Conta) depositar(valor int) int {
	c.saldo += valor
	return c.saldo
}

func main() {
	vyctor := Cliente{
		nome: "Vyctor",
	}
	vyctor.andou()
	fmt.Printf("O valor da struct com nome é %v\n", vyctor.nome)

	conta := NewConta()
	conta.simular(200)
	println(conta.saldo)
	conta.depositar(1000)
	println(conta.saldo)
}
