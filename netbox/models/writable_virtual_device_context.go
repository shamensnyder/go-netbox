// Code generated by go-swagger; DO NOT EDIT.

// Copyright 2020 The go-netbox Authors.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//   http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.
//

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"encoding/json"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableVirtualDeviceContext writable virtual device context
//
// swagger:model WritableVirtualDeviceContext
type WritableVirtualDeviceContext struct {

	// Comments
	Comments string `json:"comments,omitempty"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Custom fields
	CustomFields interface{} `json:"custom_fields,omitempty"`

	// Device
	Device *int64 `json:"device,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Identifier
	//
	// Numeric identifier unique to the parent device
	// Maximum: 32767
	// Minimum: 0
	Identifier *int64 `json:"identifier,omitempty"`

	// Interface count
	// Read Only: true
	InterfaceCount int64 `json:"interface_count,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Name
	// Required: true
	// Max Length: 64
	// Min Length: 1
	Name *string `json:"name"`

	// Primary ip
	// Read Only: true
	PrimaryIP string `json:"primary_ip,omitempty"`

	// Primary IPv4
	PrimaryIp4 *int64 `json:"primary_ip4,omitempty"`

	// Primary IPv6
	PrimaryIp6 *int64 `json:"primary_ip6,omitempty"`

	// Status
	// Required: true
	// Enum: [active planned offline]
	Status *string `json:"status"`

	// tags
	Tags []*NestedTag `json:"tags,omitempty"`

	// Tenant
	Tenant *int64 `json:"tenant,omitempty"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable virtual device context
func (m *WritableVirtualDeviceContext) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIdentifier(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateTags(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateURL(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVirtualDeviceContext) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) validateIdentifier(formats strfmt.Registry) error {
	if swag.IsZero(m.Identifier) { // not required
		return nil
	}

	if err := validate.MinimumInt("identifier", "body", *m.Identifier, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("identifier", "body", *m.Identifier, 32767, false); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	if err := validate.MinLength("name", "body", *m.Name, 1); err != nil {
		return err
	}

	if err := validate.MaxLength("name", "body", *m.Name, 64); err != nil {
		return err
	}

	return nil
}

var writableVirtualDeviceContextTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["active","planned","offline"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableVirtualDeviceContextTypeStatusPropEnum = append(writableVirtualDeviceContextTypeStatusPropEnum, v)
	}
}

const (

	// WritableVirtualDeviceContextStatusActive captures enum value "active"
	WritableVirtualDeviceContextStatusActive string = "active"

	// WritableVirtualDeviceContextStatusPlanned captures enum value "planned"
	WritableVirtualDeviceContextStatusPlanned string = "planned"

	// WritableVirtualDeviceContextStatusOffline captures enum value "offline"
	WritableVirtualDeviceContextStatusOffline string = "offline"
)

// prop value enum
func (m *WritableVirtualDeviceContext) validateStatusEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableVirtualDeviceContextTypeStatusPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableVirtualDeviceContext) validateStatus(formats strfmt.Registry) error {

	if err := validate.Required("status", "body", m.Status); err != nil {
		return err
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) validateTags(formats strfmt.Registry) error {
	if swag.IsZero(m.Tags) { // not required
		return nil
	}

	for i := 0; i < len(m.Tags); i++ {
		if swag.IsZero(m.Tags[i]) { // not required
			continue
		}

		if m.Tags[i] != nil {
			if err := m.Tags[i].Validate(formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableVirtualDeviceContext) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable virtual device context based on the context it is used
func (m *WritableVirtualDeviceContext) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
	var res []error

	if err := m.contextValidateCreated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateDisplay(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateID(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateInterfaceCount(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidatePrimaryIP(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateTags(ctx, formats); err != nil {
		res = append(res, err)
	}

	if err := m.contextValidateURL(ctx, formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateInterfaceCount(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "interface_count", "body", int64(m.InterfaceCount)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidatePrimaryIP(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "primary_ip", "body", string(m.PrimaryIP)); err != nil {
		return err
	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateTags(ctx context.Context, formats strfmt.Registry) error {

	for i := 0; i < len(m.Tags); i++ {

		if m.Tags[i] != nil {
			if err := m.Tags[i].ContextValidate(ctx, formats); err != nil {
				if ve, ok := err.(*errors.Validation); ok {
					return ve.ValidateName("tags" + "." + strconv.Itoa(i))
				} else if ce, ok := err.(*errors.CompositeError); ok {
					return ce.ValidateName("tags" + "." + strconv.Itoa(i))
				}
				return err
			}
		}

	}

	return nil
}

func (m *WritableVirtualDeviceContext) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableVirtualDeviceContext) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableVirtualDeviceContext) UnmarshalBinary(b []byte) error {
	var res WritableVirtualDeviceContext
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
