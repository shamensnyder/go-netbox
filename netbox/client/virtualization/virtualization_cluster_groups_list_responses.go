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

// VirtualizationClusterGroupsListReader is a Reader for the VirtualizationClusterGroupsList structure.
type VirtualizationClusterGroupsListReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *VirtualizationClusterGroupsListReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 200:
		result := NewVirtualizationClusterGroupsListOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewVirtualizationClusterGroupsListOK creates a VirtualizationClusterGroupsListOK with default headers values
func NewVirtualizationClusterGroupsListOK() *VirtualizationClusterGroupsListOK {
	return &VirtualizationClusterGroupsListOK{}
}

/*VirtualizationClusterGroupsListOK handles this case with default header values.

VirtualizationClusterGroupsListOK virtualization cluster groups list o k
*/
type VirtualizationClusterGroupsListOK struct {
	Payload *models.VirtualizationClusterGroupsListOKBody
}

func (o *VirtualizationClusterGroupsListOK) Error() string {
	return fmt.Sprintf("[GET /virtualization/cluster-groups/][%d] virtualizationClusterGroupsListOK  %+v", 200, o.Payload)
}

func (o *VirtualizationClusterGroupsListOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.VirtualizationClusterGroupsListOKBody)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
