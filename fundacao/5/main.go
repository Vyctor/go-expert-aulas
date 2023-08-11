package main

import "fmt"

const a = "Hello, World!"

type ID int

var (
	b bool
	c int
	d string
	e float64
	f ID = 1
)

func main() {
	println(a)
	println(b)
	b = true
	println(b)
	println(c)
	println(d)
	println(e)
	x()
}

func x() {
	fmt.Printf("O tipo de f é %T", f)
}
