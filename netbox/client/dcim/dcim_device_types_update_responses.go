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

// DcimDeviceTypesUpdateReader is a Reader for the DcimDeviceTypesUpdate structure.
type DcimDeviceTypesUpdateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DcimDeviceTypesUpdateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewDcimDeviceTypesUpdateOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDcimDeviceTypesUpdateOK creates a DcimDeviceTypesUpdateOK with default headers values
func NewDcimDeviceTypesUpdateOK() *DcimDeviceTypesUpdateOK {
	return &DcimDeviceTypesUpdateOK{}
}

/*DcimDeviceTypesUpdateOK handles this case with default header values.

DcimDeviceTypesUpdateOK dcim device types update o k
*/
type DcimDeviceTypesUpdateOK struct {
	Payload *models.WritableDeviceType
}

func (o *DcimDeviceTypesUpdateOK) Error() string {
	return fmt.Sprintf("[PUT /dcim/device-types/{id}/][%d] dcimDeviceTypesUpdateOK  %+v", 200, o.Payload)
}

func (o *DcimDeviceTypesUpdateOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableDeviceType)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
