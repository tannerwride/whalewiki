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