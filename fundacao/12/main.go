package main

func soma(a, b *int) int {
	*a = 50
	return *a + *b
}

func main() {
	minhaVar1 := 10
	minhaVar2 := 20

	println(soma(&minhaVar1, &minhaVar2))
	println(minhaVar1)
}
