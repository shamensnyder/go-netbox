/*
NetBox REST API

No description provided (generated by Openapi Generator https://github.com/openapitools/openapi-generator)

API version: 4.0.3 (4.0)
*/

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package netbox

import (
	"encoding/json"
)

// checks if the IPSecProfileMode type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &IPSecProfileMode{}

// IPSecProfileMode struct for IPSecProfileMode
type IPSecProfileMode struct {
	Value                *IPSecProfileModeValue `json:"value,omitempty"`
	Label                *IPSecProfileModeLabel `json:"label,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _IPSecProfileMode IPSecProfileMode

// NewIPSecProfileMode instantiates a new IPSecProfileMode object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewIPSecProfileMode() *IPSecProfileMode {
	this := IPSecProfileMode{}
	return &this
}

// NewIPSecProfileModeWithDefaults instantiates a new IPSecProfileMode object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewIPSecProfileModeWithDefaults() *IPSecProfileMode {
	this := IPSecProfileMode{}
	return &this
}

// GetValue returns the Value field value if set, zero value otherwise.
func (o *IPSecProfileMode) GetValue() IPSecProfileModeValue {
	if o == nil || IsNil(o.Value) {
		var ret IPSecProfileModeValue
		return ret
	}
	return *o.Value
}

// GetValueOk returns a tuple with the Value field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfileMode) GetValueOk() (*IPSecProfileModeValue, bool) {
	if o == nil || IsNil(o.Value) {
		return nil, false
	}
	return o.Value, true
}

// HasValue returns a boolean if a field has been set.
func (o *IPSecProfileMode) HasValue() bool {
	if o != nil && !IsNil(o.Value) {
		return true
	}

	return false
}

// SetValue gets a reference to the given IPSecProfileModeValue and assigns it to the Value field.
func (o *IPSecProfileMode) SetValue(v IPSecProfileModeValue) {
	o.Value = &v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *IPSecProfileMode) GetLabel() IPSecProfileModeLabel {
	if o == nil || IsNil(o.Label) {
		var ret IPSecProfileModeLabel
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *IPSecProfileMode) GetLabelOk() (*IPSecProfileModeLabel, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *IPSecProfileMode) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given IPSecProfileModeLabel and assigns it to the Label field.
func (o *IPSecProfileMode) SetLabel(v IPSecProfileModeLabel) {
	o.Label = &v
}

func (o IPSecProfileMode) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o IPSecProfileMode) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Value) {
		toSerialize["value"] = o.Value
	}
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *IPSecProfileMode) UnmarshalJSON(data []byte) (err error) {
	varIPSecProfileMode := _IPSecProfileMode{}

	err = json.Unmarshal(data, &varIPSecProfileMode)

	if err != nil {
		return err
	}

	*o = IPSecProfileMode(varIPSecProfileMode)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "value")
		delete(additionalProperties, "label")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableIPSecProfileMode struct {
	value *IPSecProfileMode
	isSet bool
}

func (v NullableIPSecProfileMode) Get() *IPSecProfileMode {
	return v.value
}

func (v *NullableIPSecProfileMode) Set(val *IPSecProfileMode) {
	v.value = val
	v.isSet = true
}

func (v NullableIPSecProfileMode) IsSet() bool {
	return v.isSet
}

func (v *NullableIPSecProfileMode) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableIPSecProfileMode(val *IPSecProfileMode) *NullableIPSecProfileMode {
	return &NullableIPSecProfileMode{value: val, isSet: true}
}

func (v NullableIPSecProfileMode) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableIPSecProfileMode) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
