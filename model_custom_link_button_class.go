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

// CustomLinkButtonClass The class of the first link in a group will be used for the dropdown button  * `outline-dark` - Default * `blue` - Blue * `indigo` - Indigo * `purple` - Purple * `pink` - Pink * `red` - Red * `orange` - Orange * `yellow` - Yellow * `green` - Green * `teal` - Teal * `cyan` - Cyan * `gray` - Gray * `black` - Black * `white` - White * `ghost-dark` - Link
type CustomLinkButtonClass string

// List of CustomLink_button_class
const (
	CUSTOMLINKBUTTONCLASS_OUTLINE_DARK CustomLinkButtonClass = "outline-dark"
	CUSTOMLINKBUTTONCLASS_BLUE         CustomLinkButtonClass = "blue"
	CUSTOMLINKBUTTONCLASS_INDIGO       CustomLinkButtonClass = "indigo"
	CUSTOMLINKBUTTONCLASS_PURPLE       CustomLinkButtonClass = "purple"
	CUSTOMLINKBUTTONCLASS_PINK         CustomLinkButtonClass = "pink"
	CUSTOMLINKBUTTONCLASS_RED          CustomLinkButtonClass = "red"
	CUSTOMLINKBUTTONCLASS_ORANGE       CustomLinkButtonClass = "orange"
	CUSTOMLINKBUTTONCLASS_YELLOW       CustomLinkButtonClass = "yellow"
	CUSTOMLINKBUTTONCLASS_GREEN        CustomLinkButtonClass = "green"
	CUSTOMLINKBUTTONCLASS_TEAL         CustomLinkButtonClass = "teal"
	CUSTOMLINKBUTTONCLASS_CYAN         CustomLinkButtonClass = "cyan"
	CUSTOMLINKBUTTONCLASS_GRAY         CustomLinkButtonClass = "gray"
	CUSTOMLINKBUTTONCLASS_BLACK        CustomLinkButtonClass = "black"
	CUSTOMLINKBUTTONCLASS_WHITE        CustomLinkButtonClass = "white"
	CUSTOMLINKBUTTONCLASS_GHOST_DARK   CustomLinkButtonClass = "ghost-dark"
)

// All allowed values of CustomLinkButtonClass enum
var AllowedCustomLinkButtonClassEnumValues = []CustomLinkButtonClass{
	"outline-dark",
	"blue",
	"indigo",
	"purple",
	"pink",
	"red",
	"orange",
	"yellow",
	"green",
	"teal",
	"cyan",
	"gray",
	"black",
	"white",
	"ghost-dark",
}

func (v *CustomLinkButtonClass) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := CustomLinkButtonClass(value)
	for _, existing := range AllowedCustomLinkButtonClassEnumValues {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid CustomLinkButtonClass", value)
}

// NewCustomLinkButtonClassFromValue returns a pointer to a valid CustomLinkButtonClass
// for the value passed as argument, or an error if the value passed is not allowed by the enum
func NewCustomLinkButtonClassFromValue(v string) (*CustomLinkButtonClass, error) {
	ev := CustomLinkButtonClass(v)
	if ev.IsValid() {
		return &ev, nil
	} else {
		return nil, fmt.Errorf("invalid value '%v' for CustomLinkButtonClass: valid values are %v", v, AllowedCustomLinkButtonClassEnumValues)
	}
}

// IsValid return true if the value is valid for the enum, false otherwise
func (v CustomLinkButtonClass) IsValid() bool {
	for _, existing := range AllowedCustomLinkButtonClassEnumValues {
		if existing == v {
			return true
		}
	}
	return false
}

// Ptr returns reference to CustomLink_button_class value
func (v CustomLinkButtonClass) Ptr() *CustomLinkButtonClass {
	return &v
}

type NullableCustomLinkButtonClass struct {
	value *CustomLinkButtonClass
	isSet bool
}

func (v NullableCustomLinkButtonClass) Get() *CustomLinkButtonClass {
	return v.value
}

func (v *NullableCustomLinkButtonClass) Set(val *CustomLinkButtonClass) {
	v.value = val
	v.isSet = true
}

func (v NullableCustomLinkButtonClass) IsSet() bool {
	return v.isSet
}

func (v *NullableCustomLinkButtonClass) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableCustomLinkButtonClass(val *CustomLinkButtonClass) *NullableCustomLinkButtonClass {
	return &NullableCustomLinkButtonClass{value: val, isSet: true}
}

func (v NullableCustomLinkButtonClass) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableCustomLinkButtonClass) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
