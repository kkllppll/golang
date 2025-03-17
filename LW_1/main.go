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
	r.HandleFunc("/", homeHandler).Methods("GET")
	r.HandleFunc("/task1", task1Handler).Methods("GET", "POST")
	r.HandleFunc("/task2", task2Handler).Methods("GET", "POST")

	// статичні файли (CSS)
	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущено на порту 8080...")
	http.ListenAndServe(":8080", r)
}

func homeHandler(w http.ResponseWriter, r *http.Request) {
	templates.ExecuteTemplate(w, "index.html", nil)
}

func task1Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result := calculateTask1(r.Form)
		templates.ExecuteTemplate(w, "task1.html", result)
		return
	}
	templates.ExecuteTemplate(w, "task1.html", nil)
}

func task2Handler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		result := calculateTask2(r.Form)
		templates.ExecuteTemplate(w, "task2.html", result)
		return
	}
	templates.ExecuteTemplate(w, "task2.html", nil)
}
