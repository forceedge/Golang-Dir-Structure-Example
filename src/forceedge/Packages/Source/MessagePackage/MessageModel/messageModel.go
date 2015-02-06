package MessageModel

import (
	"forceedge/Packages/Source/MessagePackage/Message"
	"forceedge/Packages/Utilities/Response"
)

// RegisterMessageAndNotify function will save the message in the db and notify the user who created it
func RegisterMessageAndNotify(message string) string {
	msg := Message.New()
	msg.SetText(message)
	msg.Save()

	resp := responseUtil.New()
	resp.SetFlashMessage("Your message has been saved")

	return resp.Handle()
}

func GetMessage(id int) *Message.Instance {
	return Message.Get(id)
}
