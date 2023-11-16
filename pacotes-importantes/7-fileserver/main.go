package main

import "net/http"

func main() {
	fileserver := http.FileServer(http.Dir("public"))
	mux := http.NewServeMux()
	mux.Handle("/", fileserver)
	mux.HandleFunc("/blog", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello from blog!"))
	})
	http.ListenAndServe(":8000", mux)
}
