/*
Package responseUtil will hold any response to be returned by the controller
*/
package responseUtil

import "log"

// Response type describes itself
type Response struct {
	flash     string
	response  string
	error     string
	exception string
}

// New Response Constructor
func New() *Response {
	return new(Response)
}

// SetFlashMessage sets a flash for further processing
func (resp *Response) SetFlashMessage(text string) {
	resp.flash = text
}

// GetFlashMessages will get the flash message stored
func (resp *Response) GetFlashMessages() string {
	return resp.flash
}

// SetJSONResponse will set a json as the response to be returned back from the controller
func (resp *Response) SetJSONResponse(json string) {
	resp.response = json
}

// GetResponse gets the response stored in this object
func (resp *Response) GetResponse() string {
	return resp.response
}

// SetError Sets an error message to be logged
func (resp *Response) SetError(err string) {
	resp.error = err
}

// SetException Sets an exception message
func (resp *Response) SetException(exc string, code int) {
	resp.exception = exc
}

// Handle the response
func (resp *Response) Handle() string {
	if resp.exception != "" {
		log.Fatal(resp.exception)
	}

	if resp.error != "" {
		panic(resp.error)
	}

	return resp.GetResponse()
}
