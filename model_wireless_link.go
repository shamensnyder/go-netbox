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
	"time"
)

// checks if the WirelessLink type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WirelessLink{}

// WirelessLink Adds support for custom fields and tags.
type WirelessLink struct {
	Id                   int32                  `json:"id"`
	Url                  string                 `json:"url"`
	Display              string                 `json:"display"`
	InterfaceA           Interface              `json:"interface_a"`
	InterfaceB           Interface              `json:"interface_b"`
	Ssid                 *string                `json:"ssid,omitempty"`
	Status               *WirelessLinkStatus    `json:"status,omitempty"`
	Tenant               NullableTenant         `json:"tenant,omitempty"`
	AuthType             *WirelessLANAuthType   `json:"auth_type,omitempty"`
	AuthCipher           *WirelessLANAuthCipher `json:"auth_cipher,omitempty"`
	AuthPsk              *string                `json:"auth_psk,omitempty"`
	Description          *string                `json:"description,omitempty"`
	Comments             *string                `json:"comments,omitempty"`
	Tags                 []NestedTag            `json:"tags,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Created              NullableTime           `json:"created"`
	LastUpdated          NullableTime           `json:"last_updated"`
	AdditionalProperties map[string]interface{}
}

type _WirelessLink WirelessLink

// NewWirelessLink instantiates a new WirelessLink object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWirelessLink(id int32, url string, display string, interfaceA Interface, interfaceB Interface, created NullableTime, lastUpdated NullableTime) *WirelessLink {
	this := WirelessLink{}
	this.Id = id
	this.Url = url
	this.Display = display
	this.InterfaceA = interfaceA
	this.InterfaceB = interfaceB
	this.Created = created
	this.LastUpdated = lastUpdated
	return &this
}

// NewWirelessLinkWithDefaults instantiates a new WirelessLink object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWirelessLinkWithDefaults() *WirelessLink {
	this := WirelessLink{}
	return &this
}

// GetId returns the Id field value
func (o *WirelessLink) GetId() int32 {
	if o == nil {
		var ret int32
		return ret
	}

	return o.Id
}

// GetIdOk returns a tuple with the Id field value
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetIdOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Id, true
}

// SetId sets field value
func (o *WirelessLink) SetId(v int32) {
	o.Id = v
}

// GetUrl returns the Url field value
func (o *WirelessLink) GetUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Url
}

// GetUrlOk returns a tuple with the Url field value
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Url, true
}

// SetUrl sets field value
func (o *WirelessLink) SetUrl(v string) {
	o.Url = v
}

// GetDisplay returns the Display field value
func (o *WirelessLink) GetDisplay() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Display
}

// GetDisplayOk returns a tuple with the Display field value
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetDisplayOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Display, true
}

// SetDisplay sets field value
func (o *WirelessLink) SetDisplay(v string) {
	o.Display = v
}

// GetInterfaceA returns the InterfaceA field value
func (o *WirelessLink) GetInterfaceA() Interface {
	if o == nil {
		var ret Interface
		return ret
	}

	return o.InterfaceA
}

// GetInterfaceAOk returns a tuple with the InterfaceA field value
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetInterfaceAOk() (*Interface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceA, true
}

// SetInterfaceA sets field value
func (o *WirelessLink) SetInterfaceA(v Interface) {
	o.InterfaceA = v
}

// GetInterfaceB returns the InterfaceB field value
func (o *WirelessLink) GetInterfaceB() Interface {
	if o == nil {
		var ret Interface
		return ret
	}

	return o.InterfaceB
}

// GetInterfaceBOk returns a tuple with the InterfaceB field value
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetInterfaceBOk() (*Interface, bool) {
	if o == nil {
		return nil, false
	}
	return &o.InterfaceB, true
}

// SetInterfaceB sets field value
func (o *WirelessLink) SetInterfaceB(v Interface) {
	o.InterfaceB = v
}

// GetSsid returns the Ssid field value if set, zero value otherwise.
func (o *WirelessLink) GetSsid() string {
	if o == nil || IsNil(o.Ssid) {
		var ret string
		return ret
	}
	return *o.Ssid
}

// GetSsidOk returns a tuple with the Ssid field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetSsidOk() (*string, bool) {
	if o == nil || IsNil(o.Ssid) {
		return nil, false
	}
	return o.Ssid, true
}

// HasSsid returns a boolean if a field has been set.
func (o *WirelessLink) HasSsid() bool {
	if o != nil && !IsNil(o.Ssid) {
		return true
	}

	return false
}

// SetSsid gets a reference to the given string and assigns it to the Ssid field.
func (o *WirelessLink) SetSsid(v string) {
	o.Ssid = &v
}

// GetStatus returns the Status field value if set, zero value otherwise.
func (o *WirelessLink) GetStatus() WirelessLinkStatus {
	if o == nil || IsNil(o.Status) {
		var ret WirelessLinkStatus
		return ret
	}
	return *o.Status
}

// GetStatusOk returns a tuple with the Status field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetStatusOk() (*WirelessLinkStatus, bool) {
	if o == nil || IsNil(o.Status) {
		return nil, false
	}
	return o.Status, true
}

// HasStatus returns a boolean if a field has been set.
func (o *WirelessLink) HasStatus() bool {
	if o != nil && !IsNil(o.Status) {
		return true
	}

	return false
}

// SetStatus gets a reference to the given WirelessLinkStatus and assigns it to the Status field.
func (o *WirelessLink) SetStatus(v WirelessLinkStatus) {
	o.Status = &v
}

// GetTenant returns the Tenant field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WirelessLink) GetTenant() Tenant {
	if o == nil || IsNil(o.Tenant.Get()) {
		var ret Tenant
		return ret
	}
	return *o.Tenant.Get()
}

// GetTenantOk returns a tuple with the Tenant field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLink) GetTenantOk() (*Tenant, bool) {
	if o == nil {
		return nil, false
	}
	return o.Tenant.Get(), o.Tenant.IsSet()
}

// HasTenant returns a boolean if a field has been set.
func (o *WirelessLink) HasTenant() bool {
	if o != nil && o.Tenant.IsSet() {
		return true
	}

	return false
}

// SetTenant gets a reference to the given NullableTenant and assigns it to the Tenant field.
func (o *WirelessLink) SetTenant(v Tenant) {
	o.Tenant.Set(&v)
}

// SetTenantNil sets the value for Tenant to be an explicit nil
func (o *WirelessLink) SetTenantNil() {
	o.Tenant.Set(nil)
}

// UnsetTenant ensures that no value is present for Tenant, not even an explicit nil
func (o *WirelessLink) UnsetTenant() {
	o.Tenant.Unset()
}

// GetAuthType returns the AuthType field value if set, zero value otherwise.
func (o *WirelessLink) GetAuthType() WirelessLANAuthType {
	if o == nil || IsNil(o.AuthType) {
		var ret WirelessLANAuthType
		return ret
	}
	return *o.AuthType
}

// GetAuthTypeOk returns a tuple with the AuthType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetAuthTypeOk() (*WirelessLANAuthType, bool) {
	if o == nil || IsNil(o.AuthType) {
		return nil, false
	}
	return o.AuthType, true
}

// HasAuthType returns a boolean if a field has been set.
func (o *WirelessLink) HasAuthType() bool {
	if o != nil && !IsNil(o.AuthType) {
		return true
	}

	return false
}

// SetAuthType gets a reference to the given WirelessLANAuthType and assigns it to the AuthType field.
func (o *WirelessLink) SetAuthType(v WirelessLANAuthType) {
	o.AuthType = &v
}

// GetAuthCipher returns the AuthCipher field value if set, zero value otherwise.
func (o *WirelessLink) GetAuthCipher() WirelessLANAuthCipher {
	if o == nil || IsNil(o.AuthCipher) {
		var ret WirelessLANAuthCipher
		return ret
	}
	return *o.AuthCipher
}

// GetAuthCipherOk returns a tuple with the AuthCipher field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetAuthCipherOk() (*WirelessLANAuthCipher, bool) {
	if o == nil || IsNil(o.AuthCipher) {
		return nil, false
	}
	return o.AuthCipher, true
}

// HasAuthCipher returns a boolean if a field has been set.
func (o *WirelessLink) HasAuthCipher() bool {
	if o != nil && !IsNil(o.AuthCipher) {
		return true
	}

	return false
}

// SetAuthCipher gets a reference to the given WirelessLANAuthCipher and assigns it to the AuthCipher field.
func (o *WirelessLink) SetAuthCipher(v WirelessLANAuthCipher) {
	o.AuthCipher = &v
}

// GetAuthPsk returns the AuthPsk field value if set, zero value otherwise.
func (o *WirelessLink) GetAuthPsk() string {
	if o == nil || IsNil(o.AuthPsk) {
		var ret string
		return ret
	}
	return *o.AuthPsk
}

// GetAuthPskOk returns a tuple with the AuthPsk field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetAuthPskOk() (*string, bool) {
	if o == nil || IsNil(o.AuthPsk) {
		return nil, false
	}
	return o.AuthPsk, true
}

// HasAuthPsk returns a boolean if a field has been set.
func (o *WirelessLink) HasAuthPsk() bool {
	if o != nil && !IsNil(o.AuthPsk) {
		return true
	}

	return false
}

// SetAuthPsk gets a reference to the given string and assigns it to the AuthPsk field.
func (o *WirelessLink) SetAuthPsk(v string) {
	o.AuthPsk = &v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WirelessLink) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WirelessLink) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WirelessLink) SetDescription(v string) {
	o.Description = &v
}

// GetComments returns the Comments field value if set, zero value otherwise.
func (o *WirelessLink) GetComments() string {
	if o == nil || IsNil(o.Comments) {
		var ret string
		return ret
	}
	return *o.Comments
}

// GetCommentsOk returns a tuple with the Comments field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetCommentsOk() (*string, bool) {
	if o == nil || IsNil(o.Comments) {
		return nil, false
	}
	return o.Comments, true
}

// HasComments returns a boolean if a field has been set.
func (o *WirelessLink) HasComments() bool {
	if o != nil && !IsNil(o.Comments) {
		return true
	}

	return false
}

// SetComments gets a reference to the given string and assigns it to the Comments field.
func (o *WirelessLink) SetComments(v string) {
	o.Comments = &v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WirelessLink) GetTags() []NestedTag {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTag
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetTagsOk() ([]NestedTag, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WirelessLink) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTag and assigns it to the Tags field.
func (o *WirelessLink) SetTags(v []NestedTag) {
	o.Tags = v
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WirelessLink) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WirelessLink) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WirelessLink) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WirelessLink) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetCreated returns the Created field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WirelessLink) GetCreated() time.Time {
	if o == nil || o.Created.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.Created.Get()
}

// GetCreatedOk returns a tuple with the Created field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLink) GetCreatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.Created.Get(), o.Created.IsSet()
}

// SetCreated sets field value
func (o *WirelessLink) SetCreated(v time.Time) {
	o.Created.Set(&v)
}

// GetLastUpdated returns the LastUpdated field value
// If the value is explicit nil, the zero value for time.Time will be returned
func (o *WirelessLink) GetLastUpdated() time.Time {
	if o == nil || o.LastUpdated.Get() == nil {
		var ret time.Time
		return ret
	}

	return *o.LastUpdated.Get()
}

// GetLastUpdatedOk returns a tuple with the LastUpdated field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WirelessLink) GetLastUpdatedOk() (*time.Time, bool) {
	if o == nil {
		return nil, false
	}
	return o.LastUpdated.Get(), o.LastUpdated.IsSet()
}

// SetLastUpdated sets field value
func (o *WirelessLink) SetLastUpdated(v time.Time) {
	o.LastUpdated.Set(&v)
}

func (o WirelessLink) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WirelessLink) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["id"] = o.Id
	toSerialize["url"] = o.Url
	toSerialize["display"] = o.Display
	toSerialize["interface_a"] = o.InterfaceA
	toSerialize["interface_b"] = o.InterfaceB
	if !IsNil(o.Ssid) {
		toSerialize["ssid"] = o.Ssid
	}
	if !IsNil(o.Status) {
		toSerialize["status"] = o.Status
	}
	if o.Tenant.IsSet() {
		toSerialize["tenant"] = o.Tenant.Get()
	}
	if !IsNil(o.AuthType) {
		toSerialize["auth_type"] = o.AuthType
	}
	if !IsNil(o.AuthCipher) {
		toSerialize["auth_cipher"] = o.AuthCipher
	}
	if !IsNil(o.AuthPsk) {
		toSerialize["auth_psk"] = o.AuthPsk
	}
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
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
	toSerialize["created"] = o.Created.Get()
	toSerialize["last_updated"] = o.LastUpdated.Get()

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WirelessLink) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"id",
		"url",
		"display",
		"interface_a",
		"interface_b",
		"created",
		"last_updated",
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

	varWirelessLink := _WirelessLink{}

	err = json.Unmarshal(data, &varWirelessLink)

	if err != nil {
		return err
	}

	*o = WirelessLink(varWirelessLink)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "id")
		delete(additionalProperties, "url")
		delete(additionalProperties, "display")
		delete(additionalProperties, "interface_a")
		delete(additionalProperties, "interface_b")
		delete(additionalProperties, "ssid")
		delete(additionalProperties, "status")
		delete(additionalProperties, "tenant")
		delete(additionalProperties, "auth_type")
		delete(additionalProperties, "auth_cipher")
		delete(additionalProperties, "auth_psk")
		delete(additionalProperties, "description")
		delete(additionalProperties, "comments")
		delete(additionalProperties, "tags")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "created")
		delete(additionalProperties, "last_updated")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWirelessLink struct {
	value *WirelessLink
	isSet bool
}

func (v NullableWirelessLink) Get() *WirelessLink {
	return v.value
}

func (v *NullableWirelessLink) Set(val *WirelessLink) {
	v.value = val
	v.isSet = true
}

func (v NullableWirelessLink) IsSet() bool {
	return v.isSet
}

func (v *NullableWirelessLink) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWirelessLink(val *WirelessLink) *NullableWirelessLink {
	return &NullableWirelessLink{value: val, isSet: true}
}

func (v NullableWirelessLink) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWirelessLink) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
