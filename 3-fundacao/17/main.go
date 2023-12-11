package main

func SomaInteiro(m map[string]int) int {
	var soma int
	for _, v := range m {
		soma += v
	}
	return soma
}

func SomaFloat(m map[string]float64) float64 {
	var soma float64
	for _, v := range m {
		soma += v
	}
	return soma
}

type Number interface {
	int | float64
}

func Soma[T Number](m map[string]T) T {
	var soma T
	for _, v := range m {
		soma += v
	}
	return soma
}

func Compara[T comparable](a T, b T) bool {
	if a == b {
		return true
	}
	return false
}

func main() {
	m := map[string]int{"Wesley": 1000, "Vyctor": 10000, "Maria": 100}
	m2 := map[string]float64{"Wesley": 1000.13, "Vyctor": 10000.10, "Maria": 100.25}

	println(Soma(m))
	println(Soma(m2))

	println(Compara(10, 10))
	println(Compara(false, true))
}
