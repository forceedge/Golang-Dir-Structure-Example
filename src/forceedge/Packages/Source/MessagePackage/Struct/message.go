package messagePackage

import "fmt"

type Message struct {
	id   int
	text string
}

func NewMessage() *Message {
	return new(Message)
}

func (msg *Message) Save() int {
	// Save message here
	msg.id = 56

	// Return response
	return msg.id
}

func (msg *Message) Get(id int) string {
	// Get message here
	msg.text = fmt.Sprintf("This is a message for id no :%d", id)

	// return it
	return msg.text
}
