package main

import "fmt"

func main() {
	channel := make(chan int)
	go publisher(channel)
	reader(channel)
}

func reader(ch chan int) {
	for i := range ch {
		fmt.Printf("Received %d\n", (i+1)*10)
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
