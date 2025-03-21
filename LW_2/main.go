package main

import (
	"html/template"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

var templates = template.Must(template.ParseGlob("templates/*.html"))

func main() {
	r := mux.NewRouter()

	// Головна сторінка
	r.HandleFunc("/", homeHandler).Methods("GET", "POST")

	// Статичні файли (CSS, JS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущено на порту 8080...")
	http.ListenAndServe(":8080", r)
}

// Головна сторінка (містить і форму, і обробку POST)
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result := calculateTask(r.Form)
		templates.ExecuteTemplate(w, "index.html", result)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}
