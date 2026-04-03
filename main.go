package main

import (
	"html/template"
	"net/http"
	"strings"

	"stringtransformer/stringtransformer"
)

// PageData holds data passed to the template
type PageData struct {
	Result string
}

// handler serves the form and processes user input
func handler(w http.ResponseWriter, r *http.Request) {
	// Parse template
	tmpl := template.Must(template.ParseFiles("templates/index.html"))

	data := PageData{}

	if r.Method == "POST" {
		// Get form values
		text := r.FormValue("text")
		action := r.FormValue("action")

		// Split input into words
		words := strings.Fields(text)

		// Transform using stringtransformer
		data.Result = stringtransformer.Transform(action, words)
	}

	// Execute template with data
	tmpl.Execute(w, data)
}

func main() {
	// Handle root path
	http.HandleFunc("/", handler)

	// Start server
	println("Server running at http://localhost:8080")
	http.ListenAndServe(":8080", nil)
}
