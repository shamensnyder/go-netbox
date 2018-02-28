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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewIPAMVlansCreateParams creates a new IPAMVlansCreateParams object
// with the default values initialized.
func NewIPAMVlansCreateParams() *IPAMVlansCreateParams {
	var ()
	return &IPAMVlansCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewIPAMVlansCreateParamsWithTimeout creates a new IPAMVlansCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewIPAMVlansCreateParamsWithTimeout(timeout time.Duration) *IPAMVlansCreateParams {
	var ()
	return &IPAMVlansCreateParams{

		timeout: timeout,
	}
}

// NewIPAMVlansCreateParamsWithContext creates a new IPAMVlansCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewIPAMVlansCreateParamsWithContext(ctx context.Context) *IPAMVlansCreateParams {
	var ()
	return &IPAMVlansCreateParams{

		Context: ctx,
	}
}

// NewIPAMVlansCreateParamsWithHTTPClient creates a new IPAMVlansCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewIPAMVlansCreateParamsWithHTTPClient(client *http.Client) *IPAMVlansCreateParams {
	var ()
	return &IPAMVlansCreateParams{
		HTTPClient: client,
	}
}

/*IPAMVlansCreateParams contains all the parameters to send to the API endpoint
for the ipam vlans create operation typically these are written to a http.Request
*/
type IPAMVlansCreateParams struct {

	/*Data*/
	Data *models.WritableVLAN

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the ipam vlans create params
func (o *IPAMVlansCreateParams) WithTimeout(timeout time.Duration) *IPAMVlansCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the ipam vlans create params
func (o *IPAMVlansCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the ipam vlans create params
func (o *IPAMVlansCreateParams) WithContext(ctx context.Context) *IPAMVlansCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the ipam vlans create params
func (o *IPAMVlansCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the ipam vlans create params
func (o *IPAMVlansCreateParams) WithHTTPClient(client *http.Client) *IPAMVlansCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the ipam vlans create params
func (o *IPAMVlansCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the ipam vlans create params
func (o *IPAMVlansCreateParams) WithData(data *models.WritableVLAN) *IPAMVlansCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the ipam vlans create params
func (o *IPAMVlansCreateParams) SetData(data *models.WritableVLAN) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *IPAMVlansCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Data != nil {
		if err := r.SetBodyParam(o.Data); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
