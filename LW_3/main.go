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

	// головна сторінка
	r.HandleFunc("/", homeHandler).Methods("GET", "POST")

	r.PathPrefix("/static/").Handler(http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Println("Сервер запущено на порту 8081...")
	http.ListenAndServe(":8081", r)
}

// Головна сторінка
func homeHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == "POST" {
		r.ParseForm()
		calculator := Calculator{}         // створюємо екземпляр структури
		calculator.InputParameters(r.Form) // передаємо вхідні дані
		result := calculator.Calculate()   // викликаємо метод
		templates.ExecuteTemplate(w, "index.html", result)
		return
	}
	templates.ExecuteTemplate(w, "index.html", nil)
}
