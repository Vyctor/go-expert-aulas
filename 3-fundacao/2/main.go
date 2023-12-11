package main

const a = "Hello, World!"

type User struct {
	id   int
	nome string
}

func main() {
	c := User{1, "João"}

	println(c.id, c.nome)
}
