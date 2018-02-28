// Code generated by go-swagger; DO NOT EDIT.

package circuits

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

// NewCircuitsCircuitTypesUpdateParams creates a new CircuitsCircuitTypesUpdateParams object
// with the default values initialized.
func NewCircuitsCircuitTypesUpdateParams() *CircuitsCircuitTypesUpdateParams {
	var ()
	return &CircuitsCircuitTypesUpdateParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewCircuitsCircuitTypesUpdateParamsWithTimeout creates a new CircuitsCircuitTypesUpdateParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewCircuitsCircuitTypesUpdateParamsWithTimeout(timeout time.Duration) *CircuitsCircuitTypesUpdateParams {
	var ()
	return &CircuitsCircuitTypesUpdateParams{

		timeout: timeout,
	}
}

// NewCircuitsCircuitTypesUpdateParamsWithContext creates a new CircuitsCircuitTypesUpdateParams object
// with the default values initialized, and the ability to set a context for a request
func NewCircuitsCircuitTypesUpdateParamsWithContext(ctx context.Context) *CircuitsCircuitTypesUpdateParams {
	var ()
	return &CircuitsCircuitTypesUpdateParams{

		Context: ctx,
	}
}

// NewCircuitsCircuitTypesUpdateParamsWithHTTPClient creates a new CircuitsCircuitTypesUpdateParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewCircuitsCircuitTypesUpdateParamsWithHTTPClient(client *http.Client) *CircuitsCircuitTypesUpdateParams {
	var ()
	return &CircuitsCircuitTypesUpdateParams{
		HTTPClient: client,
	}
}

/*CircuitsCircuitTypesUpdateParams contains all the parameters to send to the API endpoint
for the circuits circuit types update operation typically these are written to a http.Request
*/
type CircuitsCircuitTypesUpdateParams struct {

	/*Data*/
	Data *models.CircuitType
	/*ID
	  A unique integer value identifying this circuit type.

	*/
	ID int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithTimeout(timeout time.Duration) *CircuitsCircuitTypesUpdateParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithContext(ctx context.Context) *CircuitsCircuitTypesUpdateParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithHTTPClient(client *http.Client) *CircuitsCircuitTypesUpdateParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithData adds the data to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithData(data *models.CircuitType) *CircuitsCircuitTypesUpdateParams {
	o.SetData(data)
	return o
}

// SetData adds the data to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetData(data *models.CircuitType) {
	o.Data = data
}

// WithID adds the id to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) WithID(id int64) *CircuitsCircuitTypesUpdateParams {
	o.SetID(id)
	return o
}

// SetID adds the id to the circuits circuit types update params
func (o *CircuitsCircuitTypesUpdateParams) SetID(id int64) {
	o.ID = id
}

// WriteToRequest writes these params to a swagger request
func (o *CircuitsCircuitTypesUpdateParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

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
