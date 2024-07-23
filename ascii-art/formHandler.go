package asciiart

import (
	"ascii-art-web/artgen"
	"fmt"
	"html/template"
	"net/http"
	"os"
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

	tmpl, err = template.ParseFiles("../templates/index.html")
	if err != nil {
		fmt.Println("Error parsing template:", err)
		os.Exit(1)
	}
}


// FormHandler handles form submissions.
func FormHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodPost {
		http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
		return
	}

	inputText := r.FormValue("inputText")
	banner := r.FormValue("banner")

	if inputText == "" || banner == "" {
		http.Error(w, "Bad Request: Missing input text or banner", http.StatusBadRequest)
		return
	}

	asciiArt, err := artgen.PrintingAscii(inputText, banner)
	if err != nil {
		http.Error(w, "Internal Server Error:"+err.Error(), http.StatusInternalServerError)
		return
	}

	formData := FormData{
		InputText: inputText,
		Banner:    banner,
		ASCIIArt:  asciiArt,
	}

	err = tmpl.Execute(w, formData)
	if err != nil {
		http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
	}
}
