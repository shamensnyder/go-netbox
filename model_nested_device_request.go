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

// checks if the NestedDeviceRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &NestedDeviceRequest{}

// NestedDeviceRequest Represents an object related through a ForeignKey field. On write, it accepts a primary key (PK) value or a dictionary of attributes which can be used to uniquely identify the related object. This class should be subclassed to return a full representation of the related object on read.
type NestedDeviceRequest struct {
	Name                 NullableString `json:"name,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _NestedDeviceRequest NestedDeviceRequest

// NewNestedDeviceRequest instantiates a new NestedDeviceRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewNestedDeviceRequest() *NestedDeviceRequest {
	this := NestedDeviceRequest{}
	return &this
}

// NewNestedDeviceRequestWithDefaults instantiates a new NestedDeviceRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewNestedDeviceRequestWithDefaults() *NestedDeviceRequest {
	this := NestedDeviceRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *NestedDeviceRequest) GetName() string {
	if o == nil || IsNil(o.Name.Get()) {
		var ret string
		return ret
	}
	return *o.Name.Get()
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *NestedDeviceRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.Name.Get(), o.Name.IsSet()
}

// HasName returns a boolean if a field has been set.
func (o *NestedDeviceRequest) HasName() bool {
	if o != nil && o.Name.IsSet() {
		return true
	}

	return false
}

// SetName gets a reference to the given NullableString and assigns it to the Name field.
func (o *NestedDeviceRequest) SetName(v string) {
	o.Name.Set(&v)
}

// SetNameNil sets the value for Name to be an explicit nil
func (o *NestedDeviceRequest) SetNameNil() {
	o.Name.Set(nil)
}

// UnsetName ensures that no value is present for Name, not even an explicit nil
func (o *NestedDeviceRequest) UnsetName() {
	o.Name.Unset()
}

func (o NestedDeviceRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o NestedDeviceRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if o.Name.IsSet() {
		toSerialize["name"] = o.Name.Get()
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *NestedDeviceRequest) UnmarshalJSON(data []byte) (err error) {
	varNestedDeviceRequest := _NestedDeviceRequest{}

	err = json.Unmarshal(data, &varNestedDeviceRequest)

	if err != nil {
		return err
	}

	*o = NestedDeviceRequest(varNestedDeviceRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableNestedDeviceRequest struct {
	value *NestedDeviceRequest
	isSet bool
}

func (v NullableNestedDeviceRequest) Get() *NestedDeviceRequest {
	return v.value
}

func (v *NullableNestedDeviceRequest) Set(val *NestedDeviceRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableNestedDeviceRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableNestedDeviceRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableNestedDeviceRequest(val *NestedDeviceRequest) *NullableNestedDeviceRequest {
	return &NullableNestedDeviceRequest{value: val, isSet: true}
}

func (v NullableNestedDeviceRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableNestedDeviceRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
