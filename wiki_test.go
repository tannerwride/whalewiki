package main

import (
	"testing"
)

func TestPage(t *testing.T) {
	page := &Page{Title: "Test Title", Body: "Here we go!"}
	got := page.Body
	want := "Here we go!"

	if got != want {

	}
}