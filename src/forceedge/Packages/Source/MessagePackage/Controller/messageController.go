/*
Package messageService deals with any requests coming in relation to messages

The methods provided by this package are:
SaveMessage()
RetrieveMessage()
*/
package messagePackage

import (
	"fmt"
	"forceedge/Packages/Source/LogPackage/Controller"
	"forceedge/Packages/Source/MessagePackage/Struct"
	"net/http"
)

// SaveMessage will store the message in postgres
func SaveMessage(w http.ResponseWriter, r *http.Request) {
	msg := messagePackage.NewMessage()
	logPackage.Print("Request received, save message")
	fmt.Fprint(w, msg.Save())
}

// RetrieveMessage will retrieve a message for a given id in the url
func RetrieveMessage(w http.ResponseWriter, r *http.Request) {
	logPackage.Print("Request received, get message")
	fmt.Fprint(w, messagePackage.NewMessage().Get(5))
}
