package main

import (
	"fmt"
	"log"
	"net/http"
)

func hellowHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello World")
}

func main() {
	fs := http.FileServer(http.Dir("./"))
	http.Handle("/", fs)
	http.HandleFunc("/hello", hellowHandler)
	log.Println("Listening...")
	http.ListenAndServe(":8080", nil)

}
