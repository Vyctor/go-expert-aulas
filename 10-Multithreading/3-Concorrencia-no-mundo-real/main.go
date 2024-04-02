package main

import (
	"fmt"
	"net/http"
	"sync"
	"sync/atomic"
)

var number uint64 = 0

func main() {
	//somaMutex()
	somaAtomica(number)
	http.ListenAndServe(":3000", nil)
}

func somaMutex() {
	m := sync.Mutex{}
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		m.Lock()
		number++
		m.Unlock()
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d!\n", number)))
	})
}

func somaAtomica(number uint64) {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		atomic.AddUint64(&number, 1)
		w.Write([]byte(fmt.Sprintf("Você é o visitante número %d!\n", number)))
	})
}
