/*
Package Message will hold the basic entity data, interact with the db and provide basic CRUD functions
*/
package Message

import "fmt"

// Instance of a message entity
type Instance struct {
	id   int
	text string
}

// New is the constructor method
func New() *Instance {
	return new(Instance)
}

// SetText is a setter for the text property of Instance struct
func (msg *Instance) SetText(text string) {
	msg.text = text
}

// Create will save the the Instance data to the db and return the id
func (msg *Instance) Create() int {
	// Create message here
	msg.id = 56

	// Return response
	return msg.id
}

// Get will get a message.instance back from the db
func Get(id int) *Instance {
	// Get message here
	msg := New()
	msg.text = fmt.Sprintf("This is a message for id no :%d", id)

	// return it
	return msg
}
