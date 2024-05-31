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

// checks if the FrontPortRearPortRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &FrontPortRearPortRequest{}

// FrontPortRearPortRequest NestedRearPortSerializer but with parent device omitted (since front and rear ports must belong to same device)
type FrontPortRearPortRequest struct {
	Name string `json:"name"`
	// Physical label
	Label                *string `json:"label,omitempty"`
	Description          *string `json:"description,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _FrontPortRearPortRequest FrontPortRearPortRequest

// NewFrontPortRearPortRequest instantiates a new FrontPortRearPortRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewFrontPortRearPortRequest(name string) *FrontPortRearPortRequest {
	this := FrontPortRearPortRequest{}
	this.Name = name
	return &this
}

// NewFrontPortRearPortRequestWithDefaults instantiates a new FrontPortRearPortRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewFrontPortRearPortRequestWithDefaults() *FrontPortRearPortRequest {
	this := FrontPortRearPortRequest{}
	return &this
}

// GetName returns the Name field value
func (o *FrontPortRearPortRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *FrontPortRearPortRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *FrontPortRearPortRequest) SetName(v string) {
	o.Name = v
}

// GetLabel returns the Label field value if set, zero value otherwise.
func (o *FrontPortRearPortRequest) GetLabel() string {
	if o == nil || IsNil(o.Label) {
		var ret string
		return ret
	}
	return *o.Label
}

// GetLabelOk returns a tuple with the Label field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortRearPortRequest) GetLabelOk() (*string, bool) {
	if o == nil || IsNil(o.Label) {
		return nil, false
	}
	return o.Label, true
}

// HasLabel returns a boolean if a field has been set.
func (o *FrontPortRearPortRequest) HasLabel() bool {
	if o != nil && !IsNil(o.Label) {
		return true
	}

	return false
}

// SetLabel gets a reference to the given string and assigns it to the Label field.
func (o *FrontPortRearPortRequest) SetLabel(v string) {
	o.Label = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *FrontPortRearPortRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *FrontPortRearPortRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *FrontPortRearPortRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *FrontPortRearPortRequest) SetDescription(v string) {
	o.Description = &v
}

func (o FrontPortRearPortRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o FrontPortRearPortRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Label) {
		toSerialize["label"] = o.Label
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *FrontPortRearPortRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
	}

	allProperties := make(map[string]interface{})

	err = json.Unmarshal(data, &allProperties)

	if err != nil {
		return err
	}

	for _, requiredProperty := range requiredProperties {
		if _, exists := allProperties[requiredProperty]; !exists {
			return fmt.Errorf("no value given for required property %v", requiredProperty)
		}
	}

	varFrontPortRearPortRequest := _FrontPortRearPortRequest{}

	err = json.Unmarshal(data, &varFrontPortRearPortRequest)

	if err != nil {
		return err
	}

	*o = FrontPortRearPortRequest(varFrontPortRearPortRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "label")
		delete(additionalProperties, "description")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableFrontPortRearPortRequest struct {
	value *FrontPortRearPortRequest
	isSet bool
}

func (v NullableFrontPortRearPortRequest) Get() *FrontPortRearPortRequest {
	return v.value
}

func (v *NullableFrontPortRearPortRequest) Set(val *FrontPortRearPortRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableFrontPortRearPortRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableFrontPortRearPortRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableFrontPortRearPortRequest(val *FrontPortRearPortRequest) *NullableFrontPortRearPortRequest {
	return &NullableFrontPortRearPortRequest{value: val, isSet: true}
}

func (v NullableFrontPortRearPortRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableFrontPortRearPortRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
