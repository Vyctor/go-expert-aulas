package main

import (
	"log"
	"net/http"
	"time"
)

func main() {
	http.HandleFunc("/", handler)
	http.ListenAndServe(":8080", nil)
}

func handler(w http.ResponseWriter, r *http.Request) {
	ctx := r.Context()
	log.Println("Request iniciated")
	defer log.Println("Request ended")
	select {
	case <-time.After(5 * time.Second):
		log.Println("Request proccessed with success")
		w.Write([]byte("Request proccessed with success"))
	case <-ctx.Done():
		log.Println("Request cancelled by client")
		http.Error(w, "Request cancelled by client", http.StatusGone)
	}
}
