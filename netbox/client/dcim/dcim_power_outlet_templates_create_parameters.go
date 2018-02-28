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

	strfmt "github.com/go-openapi/strfmt"

	"github.com/digitalocean/go-netbox/netbox/models"
)

// NewDcimPowerOutletTemplatesCreateParams creates a new DcimPowerOutletTemplatesCreateParams object
// with the default values initialized.
func NewDcimPowerOutletTemplatesCreateParams() *DcimPowerOutletTemplatesCreateParams {
	var ()
	return &DcimPowerOutletTemplatesCreateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewDcimPowerOutletTemplatesCreateParamsWithTimeout creates a new DcimPowerOutletTemplatesCreateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewDcimPowerOutletTemplatesCreateParamsWithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesCreateParams {
	var ()
	return &DcimPowerOutletTemplatesCreateParams{

		timeout: timeout,
	}
}

// NewDcimPowerOutletTemplatesCreateParamsWithContext creates a new DcimPowerOutletTemplatesCreateParams object
// with the default values initialized, and the ability to set a context for a request
func NewDcimPowerOutletTemplatesCreateParamsWithContext(ctx context.Context) *DcimPowerOutletTemplatesCreateParams {
	var ()
	return &DcimPowerOutletTemplatesCreateParams{

		Context: ctx,
	}
}

// NewDcimPowerOutletTemplatesCreateParamsWithHTTPClient creates a new DcimPowerOutletTemplatesCreateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewDcimPowerOutletTemplatesCreateParamsWithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesCreateParams {
	var ()
	return &DcimPowerOutletTemplatesCreateParams{
		HTTPClient: client,
	}
}

/*DcimPowerOutletTemplatesCreateParams contains all the parameters to send to the API endpoint
for the dcim power outlet templates create operation typically these are written to a http.Request
*/
type DcimPowerOutletTemplatesCreateParams struct {

	/*Data*/
	Data *models.WritablePowerOutletTemplate

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) WithTimeout(timeout time.Duration) *DcimPowerOutletTemplatesCreateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) WithContext(ctx context.Context) *DcimPowerOutletTemplatesCreateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) WithHTTPClient(client *http.Client) *DcimPowerOutletTemplatesCreateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) WithData(data *models.WritablePowerOutletTemplate) *DcimPowerOutletTemplatesCreateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the dcim power outlet templates create params
func (o *DcimPowerOutletTemplatesCreateParams) SetData(data *models.WritablePowerOutletTemplate) {
	o.Data = data
}

// WriteToRequest writes these params to a swagger request
func (o *DcimPowerOutletTemplatesCreateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
