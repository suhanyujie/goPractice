package main

import (
	"io"
	"log"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	defer r.Body.Close()
	io.WriteString(w, "Hello world...\n")
}
func main() {
	http.HandleFunc("/hello", helloHandler)
	err := http.ListenAndServe(":8001", nil)
	if err != nil {
		log.Println(err)
		return
	}
}
