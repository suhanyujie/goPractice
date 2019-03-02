package main

import (
	"log"
	"net/http"
)

func main() {
	h := http.FileServer(http.Dir("."))
	err := http.ListenAndServeTLS(":8001", "/home/www/common/keys/other2/erp1.com.crt", "/home/www/common/keys/other2/erp1.com.key", h)
	if err != nil {
		log.Println(err)
	}
}
