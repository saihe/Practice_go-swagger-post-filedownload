// Code generated by go-swagger; DO NOT EDIT.

package operations

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PostFileOKCode is the HTTP code returned for type PostFileOK
const PostFileOKCode int = 200

/*PostFileOK ファイル

swagger:response postFileOK
*/
type PostFileOK struct {

	/*
	  In: Body
	*/
	Payload *PostFileOKBody `json:"body,omitempty"`
}

// NewPostFileOK creates PostFileOK with default headers values
func NewPostFileOK() *PostFileOK {

	return &PostFileOK{}
}

// WithPayload adds the payload to the post file o k response
func (o *PostFileOK) WithPayload(payload *PostFileOKBody) *PostFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the post file o k response
func (o *PostFileOK) SetPayload(payload *PostFileOKBody) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PostFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	if o.Payload != nil {
		payload := o.Payload
		if err := producer.Produce(rw, payload); err != nil {
			panic(err) // let the recovery middleware deal with this
		}
	}
}