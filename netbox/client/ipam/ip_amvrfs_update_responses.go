// Code generated by go-swagger; DO NOT EDIT.

package ipam

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// IPAMVrfsUpdateReader is a Reader for the IPAMVrfsUpdate structure.
type IPAMVrfsUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *IPAMVrfsUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewIPAMVrfsUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewIPAMVrfsUpdateOK creates a IPAMVrfsUpdateOK with default headers values
func NewIPAMVrfsUpdateOK() *IPAMVrfsUpdateOK {
	return &IPAMVrfsUpdateOK{}
}

/*IPAMVrfsUpdateOK handles this case with default header values.

IPAMVrfsUpdateOK ipam vrfs update o k
*/
type IPAMVrfsUpdateOK struct {
	Payload *models.WritableVRF
}

func (o *IPAMVrfsUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /ipam/vrfs/{id}/][%d] ipamVrfsUpdateOK  %+v", 200, o.Payload)
}

func (o *IPAMVrfsUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableVRF)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
