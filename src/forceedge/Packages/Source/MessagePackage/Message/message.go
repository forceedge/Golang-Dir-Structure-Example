package Message

import "fmt"

type Instance struct {
	id   int
	text string
}

func New() *Instance {
	return new(Instance)
}

func (msg *Instance) SetText(text string) {
	msg.text = text
}

func (msg *Instance) Save() int {
	// Save message here
	msg.id = 56

	// Return response
	return msg.id
}

func Get(id int) *Instance {
	// Get message here
	msg := New()
	msg.text = fmt.Sprintf("This is a message for id no :%d", id)

	// return it
	return msg
}
