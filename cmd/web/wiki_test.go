package main

import (
	"net/http"
	"net/http/httptest"
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

func TestLoad(t *testing.T) {
	got, err := loadPage("tweep")
	want := "EOF"

	if err == nil {
		t.Errorf("got %q want %q", got, want)

	}
}

func TestPing(t *testing.T) {
	rr := httptest.NewRecorder()
	r, err := http.NewRequest(http.MethodGet, "/", nil)
	if err != nil {
		t.Fatal(err)
	}

	ping(rr, r)

	rs := rr.Result()

	if rs.StatusCode != http.StatusOK {
		t.Errorf("want %d, got %d", http.StatusOK, rs.StatusCode)
	}
}