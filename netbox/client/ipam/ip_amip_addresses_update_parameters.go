// Code generated by go-swagger; DO NOT EDIT.

package ipam

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

// NewIPAMIPAddressesUpdateParams creates a new IPAMIPAddressesUpdateParams object
// with the default values initialized.
func NewIPAMIPAddressesUpdateParams() *IPAMIPAddressesUpdateParams {
	var ()
	return &IPAMIPAddressesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMIPAddressesUpdateParamsWithTimeout creates a new IPAMIPAddressesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMIPAddressesUpdateParamsWithTimeout(timeout time.Duration) *IPAMIPAddressesUpdateParams {
	var ()
	return &IPAMIPAddressesUpdateParams{

		timeout: timeout,
	}
}

// NewIPAMIPAddressesUpdateParamsWithContext creates a new IPAMIPAddressesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMIPAddressesUpdateParamsWithContext(ctx context.Context) *IPAMIPAddressesUpdateParams {
	var ()
	return &IPAMIPAddressesUpdateParams{

		Context: ctx,
	}
}

// NewIPAMIPAddressesUpdateParamsWithHTTPClient creates a new IPAMIPAddressesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMIPAddressesUpdateParamsWithHTTPClient(client *http.Client) *IPAMIPAddressesUpdateParams {
	var ()
	return &IPAMIPAddressesUpdateParams{
		HTTPClient: client,
	}
}

/*IPAMIPAddressesUpdateParams contains all the parameters to send to the API endpoint
for the ipam ip addresses update operation typically these are written to a http.Request
*/
type IPAMIPAddressesUpdateParams struct {

	/*Data*/
	Data *models.WritableIPAddress
	/*ID
	  A unique integer value identifying this IP address.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) WithTimeout(timeout time.Duration) *IPAMIPAddressesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) WithContext(ctx context.Context) *IPAMIPAddressesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) WithHTTPClient(client *http.Client) *IPAMIPAddressesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) WithData(data *models.WritableIPAddress) *IPAMIPAddressesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) SetData(data *models.WritableIPAddress) {
	o.Data = data
}

// WithID adds the id to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) WithID(id int64) *IPAMIPAddressesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the ipam ip addresses update params
func (o *IPAMIPAddressesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMIPAddressesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
