package main 

import (
	"fmt"
	"os"
	"net/http"
	"log"
)

func main() {
	fmt.Println("Starting server")
	port := os.Getenv("PORT")
	http.HandleFunc("/", homePage)
	http.HandleFunc("/view/", viewHandler)
	http.HandleFunc("/edit/", editHandler)
	http.HandleFunc("/save/", saveHandler)
	log.Fatal(http.ListenAndServe(":" + port, nil))
}