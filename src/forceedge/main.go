package main

import (
	"forceedge/Packages/Source/MessagePackage/MessageController"
	"forceedge/Packages/Utilities/Server"
	"net/http"
)

////////////////////

func main() {
	http.HandleFunc("/save/:message", MessageController.SaveMessage)
	http.HandleFunc("/retrieve/:id", MessageController.RetrieveMessage)

	server.Start()
}
