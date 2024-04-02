package main

import (
	"fmt"
	"time"
)

func main() {
	data := make(chan int)
	workerCount := 200000

	// inicializa os workers
	for i := 0; i < workerCount; i++ {
		go worker(i, data)
	}

	// envia dados para os workers
	for i := 0; i < 10000000; i++ {
		data <- i
	}
}

func worker(workerId int, data chan int) {
	for x := range data {
		fmt.Printf("Worker %d received %d\n", workerId, x)
		time.Sleep(time.Second)
	}
}
