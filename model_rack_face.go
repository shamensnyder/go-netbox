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

// RackFace the model 'RackFace'
type RackFace string

// List of Rack_face
const (
	RACKFACE_FRONT RackFace = "front"
	RACKFACE_REAR  RackFace = "rear"
)

// All allowed values of RackFace enum
var AllowedRackFaceEnumValues = []RackFace{
	"front",
	"rear",
}

func (v *RackFace) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := RackFace(value)
	for _, existing := range AllowedRackFaceEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid RackFace", value)
}

// NewRackFaceFromValue returns a pointer to a valid RackFace
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewRackFaceFromValue(v string) (*RackFace, error) {
	ev := RackFace(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for RackFace: valid values are %v", v, AllowedRackFaceEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v RackFace) IsValid() bool {
	for _, existing := range AllowedRackFaceEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to Rack_face value
func (v RackFace) Ptr() *RackFace {
	return &v
}

type NullableRackFace struct {
	value *RackFace
	isSet bool
}

func (v NullableRackFace) Get() *RackFace {
	return v.value
}

func (v *NullableRackFace) Set(val *RackFace) {
	v.value = val
	v.isSet = true
}

func (v NullableRackFace) IsSet() bool {
	return v.isSet
}

func (v *NullableRackFace) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackFace(val *RackFace) *NullableRackFace {
	return &NullableRackFace{value: val, isSet: true}
}

func (v NullableRackFace) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackFace) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
