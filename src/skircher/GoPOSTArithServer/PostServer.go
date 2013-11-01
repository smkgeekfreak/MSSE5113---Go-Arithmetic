package main

import (
	"fmt"
	"html"
	//"io/ioutil"
	"log"
	"net/http"
	//"skircher/WikiArith"
	//"time"
	"html/template"
)

func someHandler(w http.ResponseWriter, r *http.Request) {
	// read form value
	//value := r.FormValue("value")
	//if r.Method == "POST" {
	//	// receive posted data
	//	body, err := ioutil.ReadAll(r.Body)
	//	fmt.Fprintf(w, "POST Answer %v! %v", body, err)
	//} else {
	//	fmt.Fprintf(w, "Non-POST %v!", value)
	//}
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

func editHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/edit/"):]
	p, err := loadPage(title)
	if err != nil {
		p = &Page{Title: title}
	}
	renderTemplate(w, "edit", p)
	//fmt.Fprintf(w, "<h1>Editing %s</h1>"+
	//    "<form action=\"/save/%s\" method=\"POST\">"+
	//    "<textarea name=\"body\">%s</textarea><br>"+
	//    "<input type=\"submit\" value=\"Save\">"+
	//    "</form>",
	//    p.Title, p.Title, p.Body)
}

func openHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/open/"):]
	p, err := loadPage(title)
	if err != nil {
		http.Redirect(w, r, "/edit/"+title, http.StatusFound)
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
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	http.HandleFunc("/", indexHandler)
	//http.Handle("/", http.FileServer(http.Dir("./static/")))
	http.HandleFunc("/arithmetic", doJSONArithmeticPOST)
	//http.Handle("/static", )
	log.Fatal(http.ListenAndServe(":8081", nil))
	//http.HandleFunc("/", handler)
	//http.HandleFunc("/summing", someHandler)
	////http.ListenAndServe(":8081", nil)
	//s := &http.Server{
	//	Addr:           ":8081",
	//	Handler:        nil,
	//	ReadTimeout:    10 * time.Second,
	//	WriteTimeout:   10 * time.Second,
	//	MaxHeaderBytes: 1 << 20,
	//}
	//init()
	//log.Fatal(s.ListenAndServe())
}
