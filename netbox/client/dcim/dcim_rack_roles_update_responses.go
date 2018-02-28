// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// DcimRackRolesUpdateReader is a Reader for the DcimRackRolesUpdate structure.
type DcimRackRolesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRackRolesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRackRolesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRackRolesUpdateOK creates a DcimRackRolesUpdateOK with default headers values
func NewDcimRackRolesUpdateOK() *DcimRackRolesUpdateOK {
	return &DcimRackRolesUpdateOK{}
}

/*DcimRackRolesUpdateOK handles this case with default header values.

DcimRackRolesUpdateOK dcim rack roles update o k
*/
type DcimRackRolesUpdateOK struct {
	Payload *models.RackRole
}

func (o *DcimRackRolesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/rack-roles/{id}/][%d] dcimRackRolesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRackRolesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RackRole)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
