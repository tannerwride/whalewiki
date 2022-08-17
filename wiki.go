package main 

import (
	"fmt"
	"os"
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

func main() {
	f1 := &Page{Title: "Test", Body: []byte("Sample page, ok?")}
	f1.save()
	f2, _ := loadPage("Test")
	fmt.Println(string(f2.Body))
}