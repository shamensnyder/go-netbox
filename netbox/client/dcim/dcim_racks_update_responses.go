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

// DcimRacksUpdateReader is a Reader for the DcimRacksUpdate structure.
type DcimRacksUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimRacksUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimRacksUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimRacksUpdateOK creates a DcimRacksUpdateOK with default headers values
func NewDcimRacksUpdateOK() *DcimRacksUpdateOK {
	return &DcimRacksUpdateOK{}
}

/*DcimRacksUpdateOK handles this case with default header values.

DcimRacksUpdateOK dcim racks update o k
*/
type DcimRacksUpdateOK struct {
	Payload *models.WritableRack
}

func (o *DcimRacksUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/racks/{id}/][%d] dcimRacksUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimRacksUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableRack)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
