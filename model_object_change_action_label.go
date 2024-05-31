/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.3 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
	"fmt"
)

// ObjectChangeActionLabel the model 'ObjectChangeActionLabel'
type ObjectChangeActionLabel string

// List of ObjectChange_action_label
const (
	OBJECTCHANGEACTIONLABEL_CREATED ObjectChangeActionLabel = "Created"
	OBJECTCHANGEACTIONLABEL_UPDATED ObjectChangeActionLabel = "Updated"
	OBJECTCHANGEACTIONLABEL_DELETED ObjectChangeActionLabel = "Deleted"
)

// All allowed values of ObjectChangeActionLabel enum
var AllowedObjectChangeActionLabelEnumValues = []ObjectChangeActionLabel{
	"Created",
	"Updated",
	"Deleted",
}

func (v *ObjectChangeActionLabel) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ObjectChangeActionLabel(value)
	for _, existing := range AllowedObjectChangeActionLabelEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ObjectChangeActionLabel", value)
}

// NewObjectChangeActionLabelFromValue returns a pointer to a valid ObjectChangeActionLabel
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewObjectChangeActionLabelFromValue(v string) (*ObjectChangeActionLabel, error) {
	ev := ObjectChangeActionLabel(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ObjectChangeActionLabel: valid values are %v", v, AllowedObjectChangeActionLabelEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ObjectChangeActionLabel) IsValid() bool {
	for _, existing := range AllowedObjectChangeActionLabelEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to ObjectChange_action_label value
func (v ObjectChangeActionLabel) Ptr() *ObjectChangeActionLabel {
	return &v
}

type NullableObjectChangeActionLabel struct {
	value *ObjectChangeActionLabel
	isSet bool
}

func (v NullableObjectChangeActionLabel) Get() *ObjectChangeActionLabel {
	return v.value
}

func (v *NullableObjectChangeActionLabel) Set(val *ObjectChangeActionLabel) {
	v.value = val
	v.isSet = true
}

func (v NullableObjectChangeActionLabel) IsSet() bool {
	return v.isSet
}

func (v *NullableObjectChangeActionLabel) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableObjectChangeActionLabel(val *ObjectChangeActionLabel) *NullableObjectChangeActionLabel {
	return &NullableObjectChangeActionLabel{value: val, isSet: true}
}

func (v NullableObjectChangeActionLabel) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableObjectChangeActionLabel) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
