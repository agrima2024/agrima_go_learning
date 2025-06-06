package main

import (
	"encoding/json"
	"html/template"
	"io"
	"net/http"
	"os"
)

type Question struct {
	Question string   `json:"question"`
	Type     string   `json:"type"`
	Options  []string `json:"options,omitempty"`
}

func buildQuestionHTML(question Question) string {
	html := ""
	html += "<div class='question'>"
	html += "<p>" + question.Question + "</p>"
	switch question.Type {
	case "text":
		html += "<input type='text' name='question_" + question.Question + "'>"
	case "textarea":
		html += "<textarea name='question_" + question.Question + "'></textarea>"
	case "radio":
		for _, option := range question.Options {
			html += "<label><input type='radio' name='question_" + question.Question + "' value='" + option + "'>" + option + "</label><br>"
		}
	case "checkbox":
		for _, option := range question.Options {
			html += "<label><input type='checkbox' name='question_" +
				question.Question + "[]" value= "+option + ">"
			+option + "</label><br>"
		}
	}
	html += "</div>"
	return html
}

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		// Open questions.json file with error handling
		file, err := os.Open("questions.json")
		if err != nil {
			http.Error(w, "Error opening questions.json file", http.StatusInternalServerError)
			return
		}
		defer file.Close() // Close the file after use

		// Read file content
		data, err := io.ReadAll(file)
		if err != nil {
			http.Error(w, "Error reading questions.json file", http.StatusInternalServerError)
			return
		}

		var questions []Question
		err = json.Unmarshal(data, &questions)
		if err != nil {
			http.Error(w, "Error parsing questions.json file", http.StatusInternalServerError)
			return
		}

		var questionHTML string
		for _, question := range questions {
			questionHTML += buildQuestionHTML(question)
		}

		// Parse survey template with error handling
		tmpl, err := template.ParseFiles("survey.html.tmpl")
		if err != nil {
			http.Error(w, "Error parsing survey template", http.StatusInternalServerError)
			return
		}

		data = struct {
			Questions string
		}{Questions: questionHTML}
		err = tmpl.Execute(w, data)
		if err != nil {
			http.Error(w, "Error executing survey template", http.StatusInternalServerError)
			return
		}
	})

	http.ListenAndServe(":8080", nil)
}
