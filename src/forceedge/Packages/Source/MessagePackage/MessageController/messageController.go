/*
Package messageService deals with any requests coming in relation to messages

The methods provided by this package are:
SaveMessage()
RetrieveMessage()
*/
package MessageController

import (
	"fmt"
	"forceedge/Packages/Source/LogPackage/Controller"
	"forceedge/Packages/Source/MessagePackage/MessageModel"
	"net/http"
)

// SaveMessage will store the message in postgres
func SaveMessage(w http.ResponseWriter, r *http.Request) {
	response := MessageModel.RegisterMessageAndNotify("some message goes here")
	logPackage.Print("Request received, save message")
	fmt.Fprint(w, response)
}

// RetrieveMessage will retrieve a message for a given id in the url
func RetrieveMessage(w http.ResponseWriter, r *http.Request) {
	logPackage.Print("Request received, get message")
	fmt.Fprint(w, MessageModel.GetMessage(5))
}
