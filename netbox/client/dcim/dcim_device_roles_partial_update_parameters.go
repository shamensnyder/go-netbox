// Code generated by go-swagger; DO NOT EDIT.

package dcim

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewDcimDeviceRolesPartialUpdateParams creates a new DcimDeviceRolesPartialUpdateParams object
// with the default values initialized.
func NewDcimDeviceRolesPartialUpdateParams() *DcimDeviceRolesPartialUpdateParams {
	var ()
	return &DcimDeviceRolesPartialUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimDeviceRolesPartialUpdateParamsWithTimeout creates a new DcimDeviceRolesPartialUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimDeviceRolesPartialUpdateParamsWithTimeout(timeout time.Duration) *DcimDeviceRolesPartialUpdateParams {
	var ()
	return &DcimDeviceRolesPartialUpdateParams{

		timeout: timeout,
	}
}

// NewDcimDeviceRolesPartialUpdateParamsWithContext creates a new DcimDeviceRolesPartialUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimDeviceRolesPartialUpdateParamsWithContext(ctx context.Context) *DcimDeviceRolesPartialUpdateParams {
	var ()
	return &DcimDeviceRolesPartialUpdateParams{

		Context: ctx,
	}
}

// NewDcimDeviceRolesPartialUpdateParamsWithHTTPClient creates a new DcimDeviceRolesPartialUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimDeviceRolesPartialUpdateParamsWithHTTPClient(client *http.Client) *DcimDeviceRolesPartialUpdateParams {
	var ()
	return &DcimDeviceRolesPartialUpdateParams{
		HTTPClient: client,
	}
}

/*DcimDeviceRolesPartialUpdateParams contains all the parameters to send to the API endpoint
for the dcim device roles partial update operation typically these are written to a http.Request
*/
type DcimDeviceRolesPartialUpdateParams struct {

	/*Data*/
	Data *models.DeviceRole
	/*ID
	  A unique integer value identifying this device role.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) WithTimeout(timeout time.Duration) *DcimDeviceRolesPartialUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) WithContext(ctx context.Context) *DcimDeviceRolesPartialUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) WithHTTPClient(client *http.Client) *DcimDeviceRolesPartialUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) WithData(data *models.DeviceRole) *DcimDeviceRolesPartialUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) SetData(data *models.DeviceRole) {
	o.Data = data
}

// WithID adds the id to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) WithID(id int64) *DcimDeviceRolesPartialUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the dcim device roles partial update params
func (o *DcimDeviceRolesPartialUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *DcimDeviceRolesPartialUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	// path param id
	if err := r.SetPathParam("id", swag.FormatInt64(o.ID)); err != nil {
		return err
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
