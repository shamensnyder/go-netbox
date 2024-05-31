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

// checks if the PatchedWritableIPSecProfileRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &PatchedWritableIPSecProfileRequest{}

// PatchedWritableIPSecProfileRequest Adds support for custom fields and tags.
type PatchedWritableIPSecProfileRequest struct {
	Name                 *string                `json:"name,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Mode                 *IPSecProfileModeValue `json:"mode,omitempty"`
	IkePolicy            *IKEPolicyRequest      `json:"ike_policy,omitempty"`
	IpsecPolicy          *IPSecPolicyRequest    `json:"ipsec_policy,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _PatchedWritableIPSecProfileRequest PatchedWritableIPSecProfileRequest

// NewPatchedWritableIPSecProfileRequest instantiates a new PatchedWritableIPSecProfileRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewPatchedWritableIPSecProfileRequest() *PatchedWritableIPSecProfileRequest {
	this := PatchedWritableIPSecProfileRequest{}
	return &this
}

// NewPatchedWritableIPSecProfileRequestWithDefaults instantiates a new PatchedWritableIPSecProfileRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewPatchedWritableIPSecProfileRequestWithDefaults() *PatchedWritableIPSecProfileRequest {
	this := PatchedWritableIPSecProfileRequest{}
	return &this
}

// GetName returns the Name field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetName() string {
	if o == nil || IsNil(o.Name) {
		var ret string
		return ret
	}
	return *o.Name
}

// GetNameOk returns a tuple with the Name field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetNameOk() (*string, bool) {
	if o == nil || IsNil(o.Name) {
		return nil, false
	}
	return o.Name, true
}

// HasName returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasName() bool {
	if o != nil && !IsNil(o.Name) {
		return true
	}

	return false
}

// SetName gets a reference to the given string and assigns it to the Name field.
func (o *PatchedWritableIPSecProfileRequest) SetName(v string) {
	o.Name = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *PatchedWritableIPSecProfileRequest) SetDescription(v string) {
	o.Description = &v
}

// GetMode returns the Mode field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetMode() IPSecProfileModeValue {
	if o == nil || IsNil(o.Mode) {
		var ret IPSecProfileModeValue
		return ret
	}
	return *o.Mode
}

// GetModeOk returns a tuple with the Mode field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetModeOk() (*IPSecProfileModeValue, bool) {
	if o == nil || IsNil(o.Mode) {
		return nil, false
	}
	return o.Mode, true
}

// HasMode returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasMode() bool {
	if o != nil && !IsNil(o.Mode) {
		return true
	}

	return false
}

// SetMode gets a reference to the given IPSecProfileModeValue and assigns it to the Mode field.
func (o *PatchedWritableIPSecProfileRequest) SetMode(v IPSecProfileModeValue) {
	o.Mode = &v
}

// GetIkePolicy returns the IkePolicy field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetIkePolicy() IKEPolicyRequest {
	if o == nil || IsNil(o.IkePolicy) {
		var ret IKEPolicyRequest
		return ret
	}
	return *o.IkePolicy
}

// GetIkePolicyOk returns a tuple with the IkePolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetIkePolicyOk() (*IKEPolicyRequest, bool) {
	if o == nil || IsNil(o.IkePolicy) {
		return nil, false
	}
	return o.IkePolicy, true
}

// HasIkePolicy returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasIkePolicy() bool {
	if o != nil && !IsNil(o.IkePolicy) {
		return true
	}

	return false
}

// SetIkePolicy gets a reference to the given IKEPolicyRequest and assigns it to the IkePolicy field.
func (o *PatchedWritableIPSecProfileRequest) SetIkePolicy(v IKEPolicyRequest) {
	o.IkePolicy = &v
}

// GetIpsecPolicy returns the IpsecPolicy field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetIpsecPolicy() IPSecPolicyRequest {
	if o == nil || IsNil(o.IpsecPolicy) {
		var ret IPSecPolicyRequest
		return ret
	}
	return *o.IpsecPolicy
}

// GetIpsecPolicyOk returns a tuple with the IpsecPolicy field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetIpsecPolicyOk() (*IPSecPolicyRequest, bool) {
	if o == nil || IsNil(o.IpsecPolicy) {
		return nil, false
	}
	return o.IpsecPolicy, true
}

// HasIpsecPolicy returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasIpsecPolicy() bool {
	if o != nil && !IsNil(o.IpsecPolicy) {
		return true
	}

	return false
}

// SetIpsecPolicy gets a reference to the given IPSecPolicyRequest and assigns it to the IpsecPolicy field.
func (o *PatchedWritableIPSecProfileRequest) SetIpsecPolicy(v IPSecPolicyRequest) {
	o.IpsecPolicy = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *PatchedWritableIPSecProfileRequest) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *PatchedWritableIPSecProfileRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *PatchedWritableIPSecProfileRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *PatchedWritableIPSecProfileRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *PatchedWritableIPSecProfileRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *PatchedWritableIPSecProfileRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

func (o PatchedWritableIPSecProfileRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o PatchedWritableIPSecProfileRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	if !IsNil(o.Name) {
		toSerialize["name"] = o.Name
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	if !IsNil(o.Mode) {
		toSerialize["mode"] = o.Mode
	}
	if !IsNil(o.IkePolicy) {
		toSerialize["ike_policy"] = o.IkePolicy
	}
	if !IsNil(o.IpsecPolicy) {
		toSerialize["ipsec_policy"] = o.IpsecPolicy
	}
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

func (o *PatchedWritableIPSecProfileRequest) UnmarshalJSON(data []byte) (err error) {
	varPatchedWritableIPSecProfileRequest := _PatchedWritableIPSecProfileRequest{}

	err = json.Unmarshal(data, &varPatchedWritableIPSecProfileRequest)

	if err != nil {
		return err
	}

	*o = PatchedWritableIPSecProfileRequest(varPatchedWritableIPSecProfileRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "mode")
		delete(additionalProperties, "ike_policy")
		delete(additionalProperties, "ipsec_policy")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullablePatchedWritableIPSecProfileRequest struct {
	value *PatchedWritableIPSecProfileRequest
	isSet bool
}

func (v NullablePatchedWritableIPSecProfileRequest) Get() *PatchedWritableIPSecProfileRequest {
	return v.value
}

func (v *NullablePatchedWritableIPSecProfileRequest) Set(val *PatchedWritableIPSecProfileRequest) {
	v.value = val
	v.isSet = true
}

func (v NullablePatchedWritableIPSecProfileRequest) IsSet() bool {
	return v.isSet
}

func (v *NullablePatchedWritableIPSecProfileRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullablePatchedWritableIPSecProfileRequest(val *PatchedWritableIPSecProfileRequest) *NullablePatchedWritableIPSecProfileRequest {
	return &NullablePatchedWritableIPSecProfileRequest{value: val, isSet: true}
}

func (v NullablePatchedWritableIPSecProfileRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullablePatchedWritableIPSecProfileRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
