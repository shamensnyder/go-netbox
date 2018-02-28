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

// DcimPlatformsListReader is a Reader for the DcimPlatformsList structure.
type DcimPlatformsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPlatformsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPlatformsListOK creates a DcimPlatformsListOK with default headers values
func NewDcimPlatformsListOK() *DcimPlatformsListOK {
	return &DcimPlatformsListOK{}
}

/*DcimPlatformsListOK handles this case with default header values.

DcimPlatformsListOK dcim platforms list o k
*/
type DcimPlatformsListOK struct {
	Payload *models.DcimPlatformsListOKBody
}

func (o *DcimPlatformsListOK) Error() string {
	return fmt.Sprintf("[GET /dcim/platforms/][%d] dcimPlatformsListOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.DcimPlatformsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
