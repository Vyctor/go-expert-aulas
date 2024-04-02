package main

import "fmt"

func main() {
	hello := make(chan string)
	go receive("vvg", hello)
	read(hello)
}

func receive(name string, hello chan<- string) {
	hello <- name
}

func read(data <-chan string) {
	fmt.Println(<-data)
}
