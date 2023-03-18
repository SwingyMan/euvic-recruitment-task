package main

import (
	"fmt"
	"log"
	"net/http"
)

func hello(response http.ResponseWriter, request *http.Request) {
	log.Println("Requested / from " + request.RemoteAddr)
	fmt.Fprint(response, "Welcome from backend1!")
}
func main() {
	http.HandleFunc("/", hello)
	log.Println("Loaded!")
	log.Fatal(http.ListenAndServe(":80", nil))
}
