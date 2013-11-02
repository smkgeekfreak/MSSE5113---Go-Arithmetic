package main

import (
	"fmt"
	"html"
	"html/template"
	"log"
	"net/http"
)

func someHandler(w http.ResponseWriter, r *http.Request) {

	switch r.Method {
	case "GET":
		// Serve the resource.
	case "POST":
		fmt.Fprintf(w, "POST, %q", html.EscapeString(r.URL.Path))
	case "PUT":
		// Update an existing record.
	case "DELETE":
		// Remove the record.
	default:
		http.Error(w, "Invalid request method.", 405)
	}

}

func constructHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/construct/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "construct", p)
}

func openHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/open/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/construct/"+title, http.StatusFound)
		return
	}
	renderTemplate(w, "open", p)
	//fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func saveHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/save/"):]
	body := r.FormValue("body")
	p := &Page{Title: title, Body: []byte(body)}
	p.save()
	http.Redirect(w, r, "/open/"+title, http.StatusFound)
}

func renderTemplate(w http.ResponseWriter, tmpl string, p *Page) {
	t, _ := template.ParseFiles(tmpl + ".html")
	t.Execute(w, p)
}

func indexHandler(w http.ResponseWriter, r *http.Request) {
	http.Redirect(w, r, "/open", http.StatusFound)
}
func main() {
	http.HandleFunc("/open/", openHandler)
	http.HandleFunc("/construct/", constructHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/", indexHandler)
	http.HandleFunc("/arithmetic", doJSONArithmeticPOST)
	log.Fatal(http.ListenAndServe(":8081", nil))
}
