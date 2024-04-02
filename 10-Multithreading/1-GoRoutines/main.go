package main

import (
	"fmt"
	"sync"
	"time"
)

func task(name string, wg *sync.WaitGroup) {
	for i := 0; i < 10; i++ {
		fmt.Printf("%d: Task %s is running\n", i, name)
		time.Sleep(1 * time.Second)
		wg.Done()
	}
}

func main() {
	wg := sync.WaitGroup{}
	wg.Add(20)
	go task("A", &wg)
	go task("B", &wg)

	fmt.Println("Main function is done")
	wg.Wait()
}
