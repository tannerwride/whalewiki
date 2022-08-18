package main 

import (
	"fmt"
	"os"
	"net/http"
	"log"
)
// Page is a structure for holding a wiki page, describes how it will be stored.
type Page struct {
	Title string
	Body []byte
}

func (p *Page) save() error {
	filename := p.Title + ".txt"
	return os.WriteFile(filename, p.Body, 0600)
}

func loadPage(title string) (*Page, error) {
	filename := title + ".txt"
	body, err := os.ReadFile(filename)
	if err != nil {
		return nil, err
	}
	return &Page{Title: title, Body: body}, nil
}

func viewHandler(w http.ResponseWriter, r *http.Request) {
	title := r.URL.Path[len("/view/"):]
	p, _ := loadPage(title)
	fmt.Fprintf(w, "<h1>%s</h1><div>%s</div>", p.Title, p.Body)
}

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w,"Hi")
}

func main() {
	fmt.Println("Starting server on port 4000")
	http.HandleFunc("/", viewHandler)
	log.Fatal(http.ListenAndServe(":4000", nil))
}