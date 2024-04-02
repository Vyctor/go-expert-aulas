package main

func main() {
	ch := make(chan string, 3)
	ch <- "Hello"
	ch <- "World"
	ch <- "World"

	println(<-ch)
	println(<-ch)
}
