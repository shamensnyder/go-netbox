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

// checks if the RackReservationRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &RackReservationRequest{}

// RackReservationRequest Adds support for custom fields and tags.
type RackReservationRequest struct {
	Rack                 RackRequest            `json:"rack"`
	Units                []int32                `json:"units"`
	User                 UserRequest            `json:"user"`
	Tenant               NullableTenantRequest  `json:"tenant,omitempty"`
	Description          string                 `json:"description"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _RackReservationRequest RackReservationRequest

// NewRackReservationRequest instantiates a new RackReservationRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewRackReservationRequest(rack RackRequest, units []int32, user UserRequest, description string) *RackReservationRequest {
	this := RackReservationRequest{}
	this.Rack = rack
	this.Units = units
	this.User = user
	this.Description = description
	return &this
}

// NewRackReservationRequestWithDefaults instantiates a new RackReservationRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewRackReservationRequestWithDefaults() *RackReservationRequest {
	this := RackReservationRequest{}
	return &this
}

// GetRack returns the Rack field value
func (o *RackReservationRequest) GetRack() RackRequest {
	if o == nil {
		var ret RackRequest
		return ret
	}

	return o.Rack
}

// GetRackOk returns a tuple with the Rack field value
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetRackOk() (*RackRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Rack, true
}

// SetRack sets field value
func (o *RackReservationRequest) SetRack(v RackRequest) {
	o.Rack = v
}

// GetUnits returns the Units field value
func (o *RackReservationRequest) GetUnits() []int32 {
	if o == nil {
		var ret []int32
		return ret
	}

	return o.Units
}

// GetUnitsOk returns a tuple with the Units field value
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetUnitsOk() ([]int32, bool) {
	if o == nil {
		return nil, false
	}
	return o.Units, true
}

// SetUnits sets field value
func (o *RackReservationRequest) SetUnits(v []int32) {
	o.Units = v
}

// GetUser returns the User field value
func (o *RackReservationRequest) GetUser() UserRequest {
	if o == nil {
		var ret UserRequest
		return ret
	}

	return o.User
}

// GetUserOk returns a tuple with the User field value
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetUserOk() (*UserRequest, bool) {
	if o == nil {
		return nil, false
	}
	return &o.User, true
}

// SetUser sets field value
func (o *RackReservationRequest) SetUser(v UserRequest) {
	o.User = v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *RackReservationRequest) GetTenant() TenantRequest {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret TenantRequest
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *RackReservationRequest) GetTenantOk() (*TenantRequest, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *RackReservationRequest) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableTenantRequest and assigns it to the Tenant field.
func (o *RackReservationRequest) SetTenant(v TenantRequest) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *RackReservationRequest) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *RackReservationRequest) UnsetTenant() {
	o.Tenant.Unset()
}

// GetDescription returns the Description field value
func (o *RackReservationRequest) GetDescription() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Description
}

// GetDescriptionOk returns a tuple with the Description field value
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetDescriptionOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Description, true
}

// SetDescription sets field value
func (o *RackReservationRequest) SetDescription(v string) {
	o.Description = v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *RackReservationRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *RackReservationRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *RackReservationRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *RackReservationRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *RackReservationRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *RackReservationRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *RackReservationRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *RackReservationRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *RackReservationRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *RackReservationRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o RackReservationRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o RackReservationRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["rack"] = o.Rack
	toSerialize["units"] = o.Units
	toSerialize["user"] = o.User
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	toSerialize["description"] = o.Description
	if !IsNil(o.Comments) {
		toSerialize["comments"] = o.Comments
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *RackReservationRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"rack",
		"units",
		"user",
		"description",
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

	varRackReservationRequest := _RackReservationRequest{}

	err = json.Unmarshal(data, &varRackReservationRequest)

	if err != nil {
		return err
	}

	*o = RackReservationRequest(varRackReservationRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "rack")
		delete(additionalProperties, "units")
		delete(additionalProperties, "user")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableRackReservationRequest struct {
	value *RackReservationRequest
	isSet bool
}

func (v NullableRackReservationRequest) Get() *RackReservationRequest {
	return v.value
}

func (v *NullableRackReservationRequest) Set(val *RackReservationRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableRackReservationRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableRackReservationRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableRackReservationRequest(val *RackReservationRequest) *NullableRackReservationRequest {
	return &NullableRackReservationRequest{value: val, isSet: true}
}

func (v NullableRackReservationRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableRackReservationRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
