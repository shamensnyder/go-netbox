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

// ExtrasExportTemplatesCreateReader is a Reader for the ExtrasExportTemplatesCreate structure.
type ExtrasExportTemplatesCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *ExtrasExportTemplatesCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewExtrasExportTemplatesCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewExtrasExportTemplatesCreateCreated creates a ExtrasExportTemplatesCreateCreated with default headers values
func NewExtrasExportTemplatesCreateCreated() *ExtrasExportTemplatesCreateCreated {
	return &ExtrasExportTemplatesCreateCreated{}
}

/*ExtrasExportTemplatesCreateCreated handles this case with default header values.

ExtrasExportTemplatesCreateCreated extras export templates create created
*/
type ExtrasExportTemplatesCreateCreated struct {
	Payload *models.ExportTemplate
}

func (o *ExtrasExportTemplatesCreateCreated) Error() string {
	return fmt.Sprintf("[POST /extras/export-templates/][%d] extrasExportTemplatesCreateCreated  %+v", 201, o.Payload)
}

func (o *ExtrasExportTemplatesCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ExportTemplate)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
