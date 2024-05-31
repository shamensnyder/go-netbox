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

// ExtrasCustomLinksListButtonClassParameter the model 'ExtrasCustomLinksListButtonClassParameter'
type ExtrasCustomLinksListButtonClassParameter string

// List of extras_custom_links_list_button_class_parameter
const (
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_BLACK        ExtrasCustomLinksListButtonClassParameter = "black"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_BLUE         ExtrasCustomLinksListButtonClassParameter = "blue"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_CYAN         ExtrasCustomLinksListButtonClassParameter = "cyan"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_GHOST_DARK   ExtrasCustomLinksListButtonClassParameter = "ghost-dark"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_GRAY         ExtrasCustomLinksListButtonClassParameter = "gray"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_GREEN        ExtrasCustomLinksListButtonClassParameter = "green"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_INDIGO       ExtrasCustomLinksListButtonClassParameter = "indigo"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_ORANGE       ExtrasCustomLinksListButtonClassParameter = "orange"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_OUTLINE_DARK ExtrasCustomLinksListButtonClassParameter = "outline-dark"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_PINK         ExtrasCustomLinksListButtonClassParameter = "pink"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_PURPLE       ExtrasCustomLinksListButtonClassParameter = "purple"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_RED          ExtrasCustomLinksListButtonClassParameter = "red"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_TEAL         ExtrasCustomLinksListButtonClassParameter = "teal"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_WHITE        ExtrasCustomLinksListButtonClassParameter = "white"
	EXTRASCUSTOMLINKSLISTBUTTONCLASSPARAMETER_YELLOW       ExtrasCustomLinksListButtonClassParameter = "yellow"
)

// All allowed values of ExtrasCustomLinksListButtonClassParameter enum
var AllowedExtrasCustomLinksListButtonClassParameterEnumValues = []ExtrasCustomLinksListButtonClassParameter{
	"black",
	"blue",
	"cyan",
	"ghost-dark",
	"gray",
	"green",
	"indigo",
	"orange",
	"outline-dark",
	"pink",
	"purple",
	"red",
	"teal",
	"white",
	"yellow",
}

func (v *ExtrasCustomLinksListButtonClassParameter) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := ExtrasCustomLinksListButtonClassParameter(value)
	for _, existing := range AllowedExtrasCustomLinksListButtonClassParameterEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid ExtrasCustomLinksListButtonClassParameter", value)
}

// NewExtrasCustomLinksListButtonClassParameterFromValue returns a pointer to a valid ExtrasCustomLinksListButtonClassParameter
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewExtrasCustomLinksListButtonClassParameterFromValue(v string) (*ExtrasCustomLinksListButtonClassParameter, error) {
	ev := ExtrasCustomLinksListButtonClassParameter(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for ExtrasCustomLinksListButtonClassParameter: valid values are %v", v, AllowedExtrasCustomLinksListButtonClassParameterEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v ExtrasCustomLinksListButtonClassParameter) IsValid() bool {
	for _, existing := range AllowedExtrasCustomLinksListButtonClassParameterEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to extras_custom_links_list_button_class_parameter value
func (v ExtrasCustomLinksListButtonClassParameter) Ptr() *ExtrasCustomLinksListButtonClassParameter {
	return &v
}

type NullableExtrasCustomLinksListButtonClassParameter struct {
	value *ExtrasCustomLinksListButtonClassParameter
	isSet bool
}

func (v NullableExtrasCustomLinksListButtonClassParameter) Get() *ExtrasCustomLinksListButtonClassParameter {
	return v.value
}

func (v *NullableExtrasCustomLinksListButtonClassParameter) Set(val *ExtrasCustomLinksListButtonClassParameter) {
	v.value = val
	v.isSet = true
}

func (v NullableExtrasCustomLinksListButtonClassParameter) IsSet() bool {
	return v.isSet
}

func (v *NullableExtrasCustomLinksListButtonClassParameter) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableExtrasCustomLinksListButtonClassParameter(val *ExtrasCustomLinksListButtonClassParameter) *NullableExtrasCustomLinksListButtonClassParameter {
	return &NullableExtrasCustomLinksListButtonClassParameter{value: val, isSet: true}
}

func (v NullableExtrasCustomLinksListButtonClassParameter) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableExtrasCustomLinksListButtonClassParameter) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
