package main

func main() {
	a := 10
	var ponteiro *int = &a
	println(a)
	*ponteiro = 20
	b := &a
	println(a)
	println(*b)
	*b = 30
	println(a)
}
