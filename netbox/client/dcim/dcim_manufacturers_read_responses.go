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

// DcimManufacturersReadReader is a Reader for the DcimManufacturersRead structure.
type DcimManufacturersReadReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimManufacturersReadReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimManufacturersReadOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimManufacturersReadOK creates a DcimManufacturersReadOK with default headers values
func NewDcimManufacturersReadOK() *DcimManufacturersReadOK {
	return &DcimManufacturersReadOK{}
}

/*DcimManufacturersReadOK handles this case with default header values.

DcimManufacturersReadOK dcim manufacturers read o k
*/
type DcimManufacturersReadOK struct {
	Payload *models.Manufacturer
}

func (o *DcimManufacturersReadOK) Error() string {
	return fmt.Sprintf("[GET /dcim/manufacturers/{id}/][%d] dcimManufacturersReadOK  %+v", 200, o.Payload)
}

func (o *DcimManufacturersReadOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Manufacturer)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
