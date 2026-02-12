package main

import (
	"fmt"
	"log"
	"net/http"
)

func mh(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hellow World Service")
}

func main() {
	http.HandleFunc("/", mh)
	log.Println("listenin on port 8585")
	http.ListenAndServe(":8585", nil)
}
