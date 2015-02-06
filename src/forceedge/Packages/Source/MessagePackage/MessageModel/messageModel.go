/*
Package MessageModel will communicate between any controller and the Message struct
*/
package MessageModel

import (
	"forceedge/Packages/Source/MessagePackage/Message"
	"forceedge/Packages/Utilities/Response"
)

// RegisterMessageAndNotify function will save the message in the db and notify the user who created it
func RegisterMessageAndNotify(resp *responseUtil.Response, message string) {
	msg := Message.New()
	msg.SetText(message)
	resp.SetIntResult(msg.Create()) // Or have Message.Create(msg)
	resp.SetFlashMessage("Your message has been saved")
}

// GetMessage gets a message instance back
func GetMessage(id int) *Message.Instance {
	return Message.Get(id)
}
