package main

import (
	"net/http"
	"os"
	"text/template"
)

type Curso struct {
	Nome         string
	CargaHoraria int
}

type Cursos []Curso

func main() {
	templates := []string{
		"header.html",
		"content.html",
		"footer.html",
	}

	t := template.Must(template.New("content.html").ParseFiles(templates...))

	err := t.Execute(os.Stdout, Cursos{
		Curso{"Curso de Go", 40},
		Curso{"Curso de Python", 35},
		Curso{"Curso de Java", 60},
	})

	if err != nil {
		panic(err)
	}

	http.ListenAndServe(":8000", nil)
}
