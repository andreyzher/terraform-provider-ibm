// Code generated by go-swagger; DO NOT EDIT.

package compute

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.ibm.com/Bluemix/riaas-go-client/riaas/models"
)

// GetInstancesInstanceIDActionsIDReader is a Reader for the GetInstancesInstanceIDActionsID structure.
type GetInstancesInstanceIDActionsIDReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetInstancesInstanceIDActionsIDReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewGetInstancesInstanceIDActionsIDOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 404:
		result := NewGetInstancesInstanceIDActionsIDNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetInstancesInstanceIDActionsIDOK creates a GetInstancesInstanceIDActionsIDOK with default headers values
func NewGetInstancesInstanceIDActionsIDOK() *GetInstancesInstanceIDActionsIDOK {
	return &GetInstancesInstanceIDActionsIDOK{}
}

/*GetInstancesInstanceIDActionsIDOK handles this case with default header values.

dummy
*/
type GetInstancesInstanceIDActionsIDOK struct {
	Payload *models.InstanceAction
}

func (o *GetInstancesInstanceIDActionsIDOK) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/actions/{id}][%d] getInstancesInstanceIdActionsIdOK  %+v", 200, o.Payload)
}

func (o *GetInstancesInstanceIDActionsIDOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.InstanceAction)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetInstancesInstanceIDActionsIDNotFound creates a GetInstancesInstanceIDActionsIDNotFound with default headers values
func NewGetInstancesInstanceIDActionsIDNotFound() *GetInstancesInstanceIDActionsIDNotFound {
	return &GetInstancesInstanceIDActionsIDNotFound{}
}

/*GetInstancesInstanceIDActionsIDNotFound handles this case with default header values.

error
*/
type GetInstancesInstanceIDActionsIDNotFound struct {
	Payload *models.Riaaserror
}

func (o *GetInstancesInstanceIDActionsIDNotFound) Error() string {
	return fmt.Sprintf("[GET /instances/{instance_id}/actions/{id}][%d] getInstancesInstanceIdActionsIdNotFound  %+v", 404, o.Payload)
}

func (o *GetInstancesInstanceIDActionsIDNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Riaaserror)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}