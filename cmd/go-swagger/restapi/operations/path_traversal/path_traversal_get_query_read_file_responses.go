// Code generated by go-swagger; DO NOT EDIT.

package path_traversal

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"

	"github.com/go-openapi/runtime"
)

// PathTraversalGetQueryReadFileOKCode is the HTTP code returned for type PathTraversalGetQueryReadFileOK
const PathTraversalGetQueryReadFileOKCode int = 200

/*
PathTraversalGetQueryReadFileOK returns the rendered response as a string

swagger:response pathTraversalGetQueryReadFileOK
*/
type PathTraversalGetQueryReadFileOK struct {

	/*The response when succesful query happens
	  In: Body
	*/
	Payload string `json:"body,omitempty"`
}

// NewPathTraversalGetQueryReadFileOK creates PathTraversalGetQueryReadFileOK with default headers values
func NewPathTraversalGetQueryReadFileOK() *PathTraversalGetQueryReadFileOK {

	return &PathTraversalGetQueryReadFileOK{}
}

// WithPayload adds the payload to the path traversal get query read file o k response
func (o *PathTraversalGetQueryReadFileOK) WithPayload(payload string) *PathTraversalGetQueryReadFileOK {
	o.Payload = payload
	return o
}

// SetPayload sets the payload to the path traversal get query read file o k response
func (o *PathTraversalGetQueryReadFileOK) SetPayload(payload string) {
	o.Payload = payload
}

// WriteResponse to the client
func (o *PathTraversalGetQueryReadFileOK) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.WriteHeader(200)
	payload := o.Payload
	if err := producer.Produce(rw, payload); err != nil {
		panic(err) // let the recovery middleware deal with this
	}
}

/*
PathTraversalGetQueryReadFileDefault Error occured

swagger:response pathTraversalGetQueryReadFileDefault
*/
type PathTraversalGetQueryReadFileDefault struct {
	_statusCode int
}

// NewPathTraversalGetQueryReadFileDefault creates PathTraversalGetQueryReadFileDefault with default headers values
func NewPathTraversalGetQueryReadFileDefault(code int) *PathTraversalGetQueryReadFileDefault {
	if code <= 0 {
		code = 500
	}

	return &PathTraversalGetQueryReadFileDefault{
		_statusCode: code,
	}
}

// WithStatusCode adds the status to the path traversal get query read file default response
func (o *PathTraversalGetQueryReadFileDefault) WithStatusCode(code int) *PathTraversalGetQueryReadFileDefault {
	o._statusCode = code
	return o
}

// SetStatusCode sets the status to the path traversal get query read file default response
func (o *PathTraversalGetQueryReadFileDefault) SetStatusCode(code int) {
	o._statusCode = code
}

// WriteResponse to the client
func (o *PathTraversalGetQueryReadFileDefault) WriteResponse(rw http.ResponseWriter, producer runtime.Producer) {

	rw.Header().Del(runtime.HeaderContentType) //Remove Content-Type on empty responses

	rw.WriteHeader(o._statusCode)
}
