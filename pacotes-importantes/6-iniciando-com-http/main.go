package main

import "net/http"

func main() {
	http.HandleFunc("/", BuscaCepHandler)
	http.ListenAndServe(":8000", nil)
}

func BuscaCepHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		w.WriteHeader(http.StatusNotFound)
		return
	}

	cepParam := r.URL.Query().Get("cep")

	if cepParam == "" {
		w.WriteHeader(http.StatusBadRequest)
		return
	}

	w.Header().Set("Content-Type", "application/json charset=utf-8")
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("Olá mundo!"))
}
