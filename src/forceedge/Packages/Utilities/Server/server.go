package server

import (
	"log"
	"net/http"
)

func Start() {
	log.Printf("About to listen on 8888. Go to https://127.0.0.1:8888/")
	err := http.ListenAndServe(":8888", nil)
	if err != nil {
		log.Fatal(err)
	}	
}