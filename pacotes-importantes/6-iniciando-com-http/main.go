package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCEP)
	http.ListenAndServe(":8000", nil)
}

func BuscaCEP(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Olá mundo!"))
}
