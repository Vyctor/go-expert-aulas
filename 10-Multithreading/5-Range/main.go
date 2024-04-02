package main

import (
	"fmt"
	"sync"
)

func main() {
	channel := make(chan int)
	wg := sync.WaitGroup{}
	wg.Add(10)

	go publisher(channel)
	reader(channel, &wg)

	wg.Wait()
}

func reader(ch chan int, wg *sync.WaitGroup) {
	for i := range ch {
		fmt.Printf("Received %d\n", (i+1)*10)
		wg.Done()
	}
}

func publisher(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
	}
	close(ch)
}
