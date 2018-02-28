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

// DcimPlatformsPartialUpdateReader is a Reader for the DcimPlatformsPartialUpdate structure.
type DcimPlatformsPartialUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimPlatformsPartialUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimPlatformsPartialUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimPlatformsPartialUpdateOK creates a DcimPlatformsPartialUpdateOK with default headers values
func NewDcimPlatformsPartialUpdateOK() *DcimPlatformsPartialUpdateOK {
	return &DcimPlatformsPartialUpdateOK{}
}

/*DcimPlatformsPartialUpdateOK handles this case with default header values.

DcimPlatformsPartialUpdateOK dcim platforms partial update o k
*/
type DcimPlatformsPartialUpdateOK struct {
	Payload *models.WritablePlatform
}

func (o *DcimPlatformsPartialUpdateOK) Error() string {
	return fmt.Sprintf("[PATCH /dcim/platforms/{id}/][%d] dcimPlatformsPartialUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimPlatformsPartialUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritablePlatform)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
