package main

const a = "Hello, World!"

var (
	b bool
	c int
	d string
	e float64
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
	println(a + " Again!")
}
