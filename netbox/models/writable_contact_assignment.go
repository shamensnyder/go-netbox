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
	"math"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// WritableContactAssignment writable contact assignment
//
// swagger:model WritableContactAssignment
type WritableContactAssignment struct {

	// Contact
	// Required: true
	Contact *int64 `json:"contact"`

	// Content type
	// Required: true
	ContentType *string `json:"content_type"`

	// Created
	// Read Only: true
	// Format: date-time
	Created *strfmt.DateTime `json:"created,omitempty"`

	// Display
	// Read Only: true
	Display string `json:"display,omitempty"`

	// ID
	// Read Only: true
	ID int64 `json:"id,omitempty"`

	// Last updated
	// Read Only: true
	// Format: date-time
	LastUpdated *strfmt.DateTime `json:"last_updated,omitempty"`

	// Object
	// Read Only: true
	Object interface{} `json:"object,omitempty"`

	// Object id
	// Required: true
	// Maximum: math.MaxInt64
	// Minimum: 0
	ObjectID *int64 `json:"object_id"`

	// Priority
	// Enum: [primary secondary tertiary inactive]
	Priority string `json:"priority,omitempty"`

	// Role
	// Required: true
	Role *int64 `json:"role"`

	// Url
	// Read Only: true
	// Format: uri
	URL strfmt.URI `json:"url,omitempty"`
}

// Validate validates this writable contact assignment
func (m *WritableContactAssignment) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateContact(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateContentType(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateCreated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateLastUpdated(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateObjectID(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validatePriority(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateRole(formats); err != nil {
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

func (m *WritableContactAssignment) validateContact(formats strfmt.Registry) error {

	if err := validate.Required("contact", "body", m.Contact); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) validateContentType(formats strfmt.Registry) error {

	if err := validate.Required("content_type", "body", m.ContentType); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) validateCreated(formats strfmt.Registry) error {
	if swag.IsZero(m.Created) { // not required
		return nil
	}

	if err := validate.FormatOf("created", "body", "date-time", m.Created.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) validateLastUpdated(formats strfmt.Registry) error {
	if swag.IsZero(m.LastUpdated) { // not required
		return nil
	}

	if err := validate.FormatOf("last_updated", "body", "date-time", m.LastUpdated.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) validateObjectID(formats strfmt.Registry) error {

	if err := validate.Required("object_id", "body", m.ObjectID); err != nil {
		return err
	}

	if err := validate.MinimumInt("object_id", "body", *m.ObjectID, 0, false); err != nil {
		return err
	}

	if err := validate.MaximumInt("object_id", "body", *m.ObjectID, math.MaxInt64, false); err != nil {
		return err
	}

	return nil
}

var writableContactAssignmentTypePriorityPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["primary","secondary","tertiary","inactive"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		writableContactAssignmentTypePriorityPropEnum = append(writableContactAssignmentTypePriorityPropEnum, v)
	}
}

const (

	// WritableContactAssignmentPriorityPrimary captures enum value "primary"
	WritableContactAssignmentPriorityPrimary string = "primary"

	// WritableContactAssignmentPrioritySecondary captures enum value "secondary"
	WritableContactAssignmentPrioritySecondary string = "secondary"

	// WritableContactAssignmentPriorityTertiary captures enum value "tertiary"
	WritableContactAssignmentPriorityTertiary string = "tertiary"

	// WritableContactAssignmentPriorityInactive captures enum value "inactive"
	WritableContactAssignmentPriorityInactive string = "inactive"
)

// prop value enum
func (m *WritableContactAssignment) validatePriorityEnum(path, location string, value string) error {
	if err := validate.EnumCase(path, location, value, writableContactAssignmentTypePriorityPropEnum, true); err != nil {
		return err
	}
	return nil
}

func (m *WritableContactAssignment) validatePriority(formats strfmt.Registry) error {
	if swag.IsZero(m.Priority) { // not required
		return nil
	}

	// value enum
	if err := m.validatePriorityEnum("priority", "body", m.Priority); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) validateRole(formats strfmt.Registry) error {

	if err := validate.Required("role", "body", m.Role); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) validateURL(formats strfmt.Registry) error {
	if swag.IsZero(m.URL) { // not required
		return nil
	}

	if err := validate.FormatOf("url", "body", "uri", m.URL.String(), formats); err != nil {
		return err
	}

	return nil
}

// ContextValidate validate this writable contact assignment based on the context it is used
func (m *WritableContactAssignment) ContextValidate(ctx context.Context, formats strfmt.Registry) error {
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

	if err := m.contextValidateLastUpdated(ctx, formats); err != nil {
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

func (m *WritableContactAssignment) contextValidateCreated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "created", "body", m.Created); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) contextValidateDisplay(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "display", "body", string(m.Display)); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) contextValidateID(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "id", "body", int64(m.ID)); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) contextValidateLastUpdated(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "last_updated", "body", m.LastUpdated); err != nil {
		return err
	}

	return nil
}

func (m *WritableContactAssignment) contextValidateURL(ctx context.Context, formats strfmt.Registry) error {

	if err := validate.ReadOnly(ctx, "url", "body", strfmt.URI(m.URL)); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *WritableContactAssignment) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *WritableContactAssignment) UnmarshalBinary(b []byte) error {
	var res WritableContactAssignment
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
