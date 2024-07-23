package asciiart

import "net/http"

// ServeIndexPage serves the main HTML page.
func ServeIndexPage(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/" {
		http.NotFound(w, r)
		return
	}
	if r.Method == http.MethodGet {
		err := tmpl.Execute(w, nil)
		if err != nil {
			http.Error(w, "Internal Server Error: "+err.Error(), http.StatusInternalServerError)
		}
	}
}
