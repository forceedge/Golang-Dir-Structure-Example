package main

import (
	"forceedge/Packages/Source/MessagePackage/Controller"
	"forceedge/Packages/Utilities/Server"
	"net/http"
)

////////////////////

func main() {
	//Register handlers
	http.HandleFunc("/save/:message", messagePackage.SaveMessage)
	http.HandleFunc("/retrieve/:id", messagePackage.RetrieveMessage)

	server.Start()
}
