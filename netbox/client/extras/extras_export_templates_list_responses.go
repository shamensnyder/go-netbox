// Code generated by go-swagger; DO NOT EDIT.

package extras

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// ExtrasExportTemplatesListReader is a Reader for the ExtrasExportTemplatesList structure.
type ExtrasExportTemplatesListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewExtrasExportTemplatesListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasExportTemplatesListOK creates a ExtrasExportTemplatesListOK with default headers values
func NewExtrasExportTemplatesListOK() *ExtrasExportTemplatesListOK {
	return &ExtrasExportTemplatesListOK{}
}

/*ExtrasExportTemplatesListOK handles this case with default header values.

ExtrasExportTemplatesListOK extras export templates list o k
*/
type ExtrasExportTemplatesListOK struct {
	Payload *models.ExtrasExportTemplatesListOKBody
}

func (o *ExtrasExportTemplatesListOK) Error() string {
	return fmt.Sprintf("[GET /extras/export-templates/][%d] extrasExportTemplatesListOK  %+v", 200, o.Payload)
}

func (o *ExtrasExportTemplatesListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExtrasExportTemplatesListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
