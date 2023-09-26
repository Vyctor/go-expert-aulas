package main

func main() {
	a := 1
	b := 2
	c := 3

	if a == 2 {
		println(a)
	} else {
		println(b)
	}

	switch c {
	case 1:
		print(a)
	case 2:
		print(b)
	default:
		print("default")
	}
}
