package asciiart

import (
	"fmt"
	"html/template"
	"net/http"
	"os"

	"ascii-art-web/artgen"
)

type FormData struct {
	InputText string
	Banner    string
	ASCIIArt  string
	Error     string
}

var tmpl *template.Template

func init() {
	var err error

	tmpl, err = template.ParseFiles("templates/form.html")
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

// FormHandler handles form submissions.
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}

	if r.Method != http.MethodGet && r.Method != http.MethodPost {
		http.Error(w, "405 Method not allowed", http.StatusMethodNotAllowed)
		return
	}

	if r.Method == http.MethodGet {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
	}
	if r.Method == http.MethodPost {

		inputText := r.FormValue("inputText")
		banner := r.FormValue("banner")

		inputText, err := artgen.HandleInput(inputText)
		if err != nil {
			fmt.Println("Error:", err)
		}

		if inputText == "" || banner == "" {
			http.Error(w, "400 bad request", http.StatusBadRequest)
			return
		}

		asciiArt, err := artgen.PrintingAscii(inputText, banner)
		if err != nil {
			http.Error(w, "404 Not Found", http.StatusNotFound)
			return
		}

		formData := FormData{
			InputText: inputText,
			Banner:    banner,
			ASCIIArt:  asciiArt,
		}

		err = tmpl.Execute(w, formData)
		if err != nil {
			http.Error(w, "500 Internal Server Error", http.StatusInternalServerError)
		}
	}
}
