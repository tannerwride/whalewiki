package main

import (
	"testing"
)

func TestPage(t *testing.T) {
	page := &Page{Title: "Test Title", Body: []byte("Here we go!")}
	got := page.Title
	want := "Test Title"

	if got != want {
		t.Errorf("got %q want %q", got, want)

	}
}