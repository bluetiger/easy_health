package main

import (
	"log"
	"net/http"
)

func main() {
	handler := func(_ http.ResponseWriter,_ *http.Request){
		return
	}
	http.HandleFunc("/health/", handler)
	log.Fatal(http.ListenAndServe(":8080",nil))
}
