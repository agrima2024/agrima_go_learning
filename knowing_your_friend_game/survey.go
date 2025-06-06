package main

import (
	"fmt"
	"html/template"
	"net/http"
)

func serveSurvey(w http.ResponseWriter, r *http.Request) {
	tmpl, err := template.ParseFiles("survey.html")
	if err != nil {
		fmt.Fprintf(w, "Error parsing template: %v", err)
		return
	}
	tmpl.Execute(w, nil)
}

func submitSurvey(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		fmt.Fprintf(w, "Unsupported request method")
		return
	}
	fmt.Println("Favorite Color:", r.FormValue("question1"))
	fmt.Println("Favorite Food:", r.FormValue("question2"))
	fmt.Println("Favorite Book:", r.FormValue("question3"))
	fmt.Fprintf(w, "Survey submitted successfully!")
}

func main() {
	http.HandleFunc("/", serveSurvey)
	http.HandleFunc("/submit", submitSurvey)
	fmt.Println("Server listening on port 8080")
	http.ListenAndServe(":8080", nil)
}
