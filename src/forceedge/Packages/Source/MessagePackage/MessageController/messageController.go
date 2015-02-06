/*
Package MessageController deals with any requests coming in relation to messages

The methods provided by this package are:
SaveMessage()
RetrieveMessage()
*/
package MessageController

import (
	"fmt"
	"forceedge/Packages/Source/LogPackage/Controller"
	"forceedge/Packages/Source/MessagePackage/MessageModel"
	"forceedge/Packages/Utilities/Response"
	"net/http"
)

// SaveMessage will store the message in postgres
func SaveMessage(w http.ResponseWriter, r *http.Request) {
	resp := responseUtil.New(w)
	// logPackage.Print("Request received, save message")
	MessageModel.RegisterMessageAndNotify(resp, "some message goes here")
	// logPackage.Print(fmt.Sprintf("Request completed, message saved with id: %v", resp.GetIntResult()))

	resp.Handle()
}

// RetrieveMessage will retrieve a message for a given id in the url
func RetrieveMessage(w http.ResponseWriter, r *http.Request) {
	logPackage.Print("Request received, get message")
	fmt.Fprint(w, MessageModel.GetMessage(5))
}
