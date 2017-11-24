// Code generated by go-swagger; DO NOT EDIT.

package tasks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/roscopecoltran/go-swagger/examples/task-tracker/models"
)

// DeleteTaskReader is a Reader for the DeleteTask structure.
type DeleteTaskReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteTaskReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 204:
		result := NewDeleteTaskNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		result := NewDeleteTaskDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewDeleteTaskNoContent creates a DeleteTaskNoContent with default headers values
func NewDeleteTaskNoContent() *DeleteTaskNoContent {
	return &DeleteTaskNoContent{}
}

/*DeleteTaskNoContent handles this case with default header values.

Task deleted
*/
type DeleteTaskNoContent struct {
}

func (o *DeleteTaskNoContent) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTaskNoContent ", 204)
}

func (o *DeleteTaskNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteTaskDefault creates a DeleteTaskDefault with default headers values
func NewDeleteTaskDefault(code int) *DeleteTaskDefault {
	return &DeleteTaskDefault{
		_statusCode: code,
	}
}

/*DeleteTaskDefault handles this case with default header values.

Error response
*/
type DeleteTaskDefault struct {
	_statusCode int

	XErrorCode string

	Payload *models.Error
}

// Code gets the status code for the delete task default response
func (o *DeleteTaskDefault) Code() int {
	return o._statusCode
}

func (o *DeleteTaskDefault) Error() string {
	return fmt.Sprintf("[DELETE /tasks/{id}][%d] deleteTask default  %+v", o._statusCode, o.Payload)
}

func (o *DeleteTaskDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	// response header X-Error-Code
	o.XErrorCode = response.GetHeader("X-Error-Code")

	o.Payload = new(models.Error)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
