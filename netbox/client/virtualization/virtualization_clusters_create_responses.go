// Code generated by go-swagger; DO NOT EDIT.

package virtualization

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// VirtualizationClustersCreateReader is a Reader for the VirtualizationClustersCreate structure.
type VirtualizationClustersCreateReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClustersCreateReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 201:
		result := NewVirtualizationClustersCreateCreated()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClustersCreateCreated creates a VirtualizationClustersCreateCreated with default headers values
func NewVirtualizationClustersCreateCreated() *VirtualizationClustersCreateCreated {
	return &VirtualizationClustersCreateCreated{}
}

/*VirtualizationClustersCreateCreated handles this case with default header values.

VirtualizationClustersCreateCreated virtualization clusters create created
*/
type VirtualizationClustersCreateCreated struct {
	Payload *models.WritableCluster
}

func (o *VirtualizationClustersCreateCreated) Error() string {
	return fmt.Sprintf("[POST /virtualization/clusters/][%d] virtualizationClustersCreateCreated  %+v", 201, o.Payload)
}

func (o *VirtualizationClustersCreateCreated) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.WritableCluster)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
