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

// checks if the WebhookRequest type satisfies the MappedNullable interface at compile time
var _ MappedNullable = &WebhookRequest{}

// WebhookRequest Adds support for custom fields and tags.
type WebhookRequest struct {
	Name        string  `json:"name"`
	Description *string `json:"description,omitempty"`
	// This URL will be called using the HTTP method defined when the webhook is called. Jinja2 template processing is supported with the same context as the request body.
	PayloadUrl string                           `json:"payload_url"`
	HttpMethod *PatchedWebhookRequestHttpMethod `json:"http_method,omitempty"`
	// The complete list of official content types is available <a href=\"https://www.iana.org/assignments/media-types/media-types.xhtml\">here</a>.
	HttpContentType *string `json:"http_content_type,omitempty"`
	// User-supplied HTTP headers to be sent with the request in addition to the HTTP content type. Headers should be defined in the format <code>Name: Value</code>. Jinja2 template processing is supported with the same context as the request body (below).
	AdditionalHeaders *string `json:"additional_headers,omitempty"`
	// Jinja2 template for a custom request body. If blank, a JSON object representing the change will be included. Available context data includes: <code>event</code>, <code>model</code>, <code>timestamp</code>, <code>username</code>, <code>request_id</code>, and <code>data</code>.
	BodyTemplate *string `json:"body_template,omitempty"`
	// When provided, the request will include a <code>X-Hook-Signature</code> header containing a HMAC hex digest of the payload body using the secret as the key. The secret is not transmitted in the request.
	Secret *string `json:"secret,omitempty"`
	// Enable SSL certificate verification. Disable with caution!
	SslVerification *bool `json:"ssl_verification,omitempty"`
	// The specific CA certificate file to use for SSL verification. Leave blank to use the system defaults.
	CaFilePath           NullableString         `json:"ca_file_path,omitempty"`
	CustomFields         map[string]interface{} `json:"custom_fields,omitempty"`
	Tags                 []NestedTagRequest     `json:"tags,omitempty"`
	AdditionalProperties map[string]interface{}
}

type _WebhookRequest WebhookRequest

// NewWebhookRequest instantiates a new WebhookRequest object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewWebhookRequest(name string, payloadUrl string) *WebhookRequest {
	this := WebhookRequest{}
	this.Name = name
	this.PayloadUrl = payloadUrl
	return &this
}

// NewWebhookRequestWithDefaults instantiates a new WebhookRequest object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewWebhookRequestWithDefaults() *WebhookRequest {
	this := WebhookRequest{}
	return &this
}

// GetName returns the Name field value
func (o *WebhookRequest) GetName() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.Name
}

// GetNameOk returns a tuple with the Name field value
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.Name, true
}

// SetName sets field value
func (o *WebhookRequest) SetName(v string) {
	o.Name = v
}

// GetDescription returns the Description field value if set, zero value otherwise.
func (o *WebhookRequest) GetDescription() string {
	if o == nil || IsNil(o.Description) {
		var ret string
		return ret
	}
	return *o.Description
}

// GetDescriptionOk returns a tuple with the Description field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetDescriptionOk() (*string, bool) {
	if o == nil || IsNil(o.Description) {
		return nil, false
	}
	return o.Description, true
}

// HasDescription returns a boolean if a field has been set.
func (o *WebhookRequest) HasDescription() bool {
	if o != nil && !IsNil(o.Description) {
		return true
	}

	return false
}

// SetDescription gets a reference to the given string and assigns it to the Description field.
func (o *WebhookRequest) SetDescription(v string) {
	o.Description = &v
}

// GetPayloadUrl returns the PayloadUrl field value
func (o *WebhookRequest) GetPayloadUrl() string {
	if o == nil {
		var ret string
		return ret
	}

	return o.PayloadUrl
}

// GetPayloadUrlOk returns a tuple with the PayloadUrl field value
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetPayloadUrlOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return &o.PayloadUrl, true
}

// SetPayloadUrl sets field value
func (o *WebhookRequest) SetPayloadUrl(v string) {
	o.PayloadUrl = v
}

// GetHttpMethod returns the HttpMethod field value if set, zero value otherwise.
func (o *WebhookRequest) GetHttpMethod() PatchedWebhookRequestHttpMethod {
	if o == nil || IsNil(o.HttpMethod) {
		var ret PatchedWebhookRequestHttpMethod
		return ret
	}
	return *o.HttpMethod
}

// GetHttpMethodOk returns a tuple with the HttpMethod field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetHttpMethodOk() (*PatchedWebhookRequestHttpMethod, bool) {
	if o == nil || IsNil(o.HttpMethod) {
		return nil, false
	}
	return o.HttpMethod, true
}

// HasHttpMethod returns a boolean if a field has been set.
func (o *WebhookRequest) HasHttpMethod() bool {
	if o != nil && !IsNil(o.HttpMethod) {
		return true
	}

	return false
}

// SetHttpMethod gets a reference to the given PatchedWebhookRequestHttpMethod and assigns it to the HttpMethod field.
func (o *WebhookRequest) SetHttpMethod(v PatchedWebhookRequestHttpMethod) {
	o.HttpMethod = &v
}

// GetHttpContentType returns the HttpContentType field value if set, zero value otherwise.
func (o *WebhookRequest) GetHttpContentType() string {
	if o == nil || IsNil(o.HttpContentType) {
		var ret string
		return ret
	}
	return *o.HttpContentType
}

// GetHttpContentTypeOk returns a tuple with the HttpContentType field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetHttpContentTypeOk() (*string, bool) {
	if o == nil || IsNil(o.HttpContentType) {
		return nil, false
	}
	return o.HttpContentType, true
}

// HasHttpContentType returns a boolean if a field has been set.
func (o *WebhookRequest) HasHttpContentType() bool {
	if o != nil && !IsNil(o.HttpContentType) {
		return true
	}

	return false
}

// SetHttpContentType gets a reference to the given string and assigns it to the HttpContentType field.
func (o *WebhookRequest) SetHttpContentType(v string) {
	o.HttpContentType = &v
}

// GetAdditionalHeaders returns the AdditionalHeaders field value if set, zero value otherwise.
func (o *WebhookRequest) GetAdditionalHeaders() string {
	if o == nil || IsNil(o.AdditionalHeaders) {
		var ret string
		return ret
	}
	return *o.AdditionalHeaders
}

// GetAdditionalHeadersOk returns a tuple with the AdditionalHeaders field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetAdditionalHeadersOk() (*string, bool) {
	if o == nil || IsNil(o.AdditionalHeaders) {
		return nil, false
	}
	return o.AdditionalHeaders, true
}

// HasAdditionalHeaders returns a boolean if a field has been set.
func (o *WebhookRequest) HasAdditionalHeaders() bool {
	if o != nil && !IsNil(o.AdditionalHeaders) {
		return true
	}

	return false
}

// SetAdditionalHeaders gets a reference to the given string and assigns it to the AdditionalHeaders field.
func (o *WebhookRequest) SetAdditionalHeaders(v string) {
	o.AdditionalHeaders = &v
}

// GetBodyTemplate returns the BodyTemplate field value if set, zero value otherwise.
func (o *WebhookRequest) GetBodyTemplate() string {
	if o == nil || IsNil(o.BodyTemplate) {
		var ret string
		return ret
	}
	return *o.BodyTemplate
}

// GetBodyTemplateOk returns a tuple with the BodyTemplate field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetBodyTemplateOk() (*string, bool) {
	if o == nil || IsNil(o.BodyTemplate) {
		return nil, false
	}
	return o.BodyTemplate, true
}

// HasBodyTemplate returns a boolean if a field has been set.
func (o *WebhookRequest) HasBodyTemplate() bool {
	if o != nil && !IsNil(o.BodyTemplate) {
		return true
	}

	return false
}

// SetBodyTemplate gets a reference to the given string and assigns it to the BodyTemplate field.
func (o *WebhookRequest) SetBodyTemplate(v string) {
	o.BodyTemplate = &v
}

// GetSecret returns the Secret field value if set, zero value otherwise.
func (o *WebhookRequest) GetSecret() string {
	if o == nil || IsNil(o.Secret) {
		var ret string
		return ret
	}
	return *o.Secret
}

// GetSecretOk returns a tuple with the Secret field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetSecretOk() (*string, bool) {
	if o == nil || IsNil(o.Secret) {
		return nil, false
	}
	return o.Secret, true
}

// HasSecret returns a boolean if a field has been set.
func (o *WebhookRequest) HasSecret() bool {
	if o != nil && !IsNil(o.Secret) {
		return true
	}

	return false
}

// SetSecret gets a reference to the given string and assigns it to the Secret field.
func (o *WebhookRequest) SetSecret(v string) {
	o.Secret = &v
}

// GetSslVerification returns the SslVerification field value if set, zero value otherwise.
func (o *WebhookRequest) GetSslVerification() bool {
	if o == nil || IsNil(o.SslVerification) {
		var ret bool
		return ret
	}
	return *o.SslVerification
}

// GetSslVerificationOk returns a tuple with the SslVerification field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetSslVerificationOk() (*bool, bool) {
	if o == nil || IsNil(o.SslVerification) {
		return nil, false
	}
	return o.SslVerification, true
}

// HasSslVerification returns a boolean if a field has been set.
func (o *WebhookRequest) HasSslVerification() bool {
	if o != nil && !IsNil(o.SslVerification) {
		return true
	}

	return false
}

// SetSslVerification gets a reference to the given bool and assigns it to the SslVerification field.
func (o *WebhookRequest) SetSslVerification(v bool) {
	o.SslVerification = &v
}

// GetCaFilePath returns the CaFilePath field value if set, zero value otherwise (both if not set or set to explicit null).
func (o *WebhookRequest) GetCaFilePath() string {
	if o == nil || IsNil(o.CaFilePath.Get()) {
		var ret string
		return ret
	}
	return *o.CaFilePath.Get()
}

// GetCaFilePathOk returns a tuple with the CaFilePath field value if set, nil otherwise
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *WebhookRequest) GetCaFilePathOk() (*string, bool) {
	if o == nil {
		return nil, false
	}
	return o.CaFilePath.Get(), o.CaFilePath.IsSet()
}

// HasCaFilePath returns a boolean if a field has been set.
func (o *WebhookRequest) HasCaFilePath() bool {
	if o != nil && o.CaFilePath.IsSet() {
		return true
	}

	return false
}

// SetCaFilePath gets a reference to the given NullableString and assigns it to the CaFilePath field.
func (o *WebhookRequest) SetCaFilePath(v string) {
	o.CaFilePath.Set(&v)
}

// SetCaFilePathNil sets the value for CaFilePath to be an explicit nil
func (o *WebhookRequest) SetCaFilePathNil() {
	o.CaFilePath.Set(nil)
}

// UnsetCaFilePath ensures that no value is present for CaFilePath, not even an explicit nil
func (o *WebhookRequest) UnsetCaFilePath() {
	o.CaFilePath.Unset()
}

// GetCustomFields returns the CustomFields field value if set, zero value otherwise.
func (o *WebhookRequest) GetCustomFields() map[string]interface{} {
	if o == nil || IsNil(o.CustomFields) {
		var ret map[string]interface{}
		return ret
	}
	return o.CustomFields
}

// GetCustomFieldsOk returns a tuple with the CustomFields field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetCustomFieldsOk() (map[string]interface{}, bool) {
	if o == nil || IsNil(o.CustomFields) {
		return map[string]interface{}{}, false
	}
	return o.CustomFields, true
}

// HasCustomFields returns a boolean if a field has been set.
func (o *WebhookRequest) HasCustomFields() bool {
	if o != nil && !IsNil(o.CustomFields) {
		return true
	}

	return false
}

// SetCustomFields gets a reference to the given map[string]interface{} and assigns it to the CustomFields field.
func (o *WebhookRequest) SetCustomFields(v map[string]interface{}) {
	o.CustomFields = v
}

// GetTags returns the Tags field value if set, zero value otherwise.
func (o *WebhookRequest) GetTags() []NestedTagRequest {
	if o == nil || IsNil(o.Tags) {
		var ret []NestedTagRequest
		return ret
	}
	return o.Tags
}

// GetTagsOk returns a tuple with the Tags field value if set, nil otherwise
// and a boolean to check if the value has been set.
func (o *WebhookRequest) GetTagsOk() ([]NestedTagRequest, bool) {
	if o == nil || IsNil(o.Tags) {
		return nil, false
	}
	return o.Tags, true
}

// HasTags returns a boolean if a field has been set.
func (o *WebhookRequest) HasTags() bool {
	if o != nil && !IsNil(o.Tags) {
		return true
	}

	return false
}

// SetTags gets a reference to the given []NestedTagRequest and assigns it to the Tags field.
func (o *WebhookRequest) SetTags(v []NestedTagRequest) {
	o.Tags = v
}

func (o WebhookRequest) MarshalJSON() ([]byte, error) {
	toSerialize, err := o.ToMap()
	if err != nil {
		return []byte{}, err
	}
	return json.Marshal(toSerialize)
}

func (o WebhookRequest) ToMap() (map[string]interface{}, error) {
	toSerialize := map[string]interface{}{}
	toSerialize["name"] = o.Name
	if !IsNil(o.Description) {
		toSerialize["description"] = o.Description
	}
	toSerialize["payload_url"] = o.PayloadUrl
	if !IsNil(o.HttpMethod) {
		toSerialize["http_method"] = o.HttpMethod
	}
	if !IsNil(o.HttpContentType) {
		toSerialize["http_content_type"] = o.HttpContentType
	}
	if !IsNil(o.AdditionalHeaders) {
		toSerialize["additional_headers"] = o.AdditionalHeaders
	}
	if !IsNil(o.BodyTemplate) {
		toSerialize["body_template"] = o.BodyTemplate
	}
	if !IsNil(o.Secret) {
		toSerialize["secret"] = o.Secret
	}
	if !IsNil(o.SslVerification) {
		toSerialize["ssl_verification"] = o.SslVerification
	}
	if o.CaFilePath.IsSet() {
		toSerialize["ca_file_path"] = o.CaFilePath.Get()
	}
	if !IsNil(o.CustomFields) {
		toSerialize["custom_fields"] = o.CustomFields
	}
	if !IsNil(o.Tags) {
		toSerialize["tags"] = o.Tags
	}

	for key, value := range o.AdditionalProperties {
		toSerialize[key] = value
	}

	return toSerialize, nil
}

func (o *WebhookRequest) UnmarshalJSON(data []byte) (err error) {
	// This validates that all required properties are included in the JSON object
	// by unmarshalling the object into a generic map with string keys and checking
	// that every required field exists as a key in the generic map.
	requiredProperties := []string{
		"name",
		"payload_url",
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

	varWebhookRequest := _WebhookRequest{}

	err = json.Unmarshal(data, &varWebhookRequest)

	if err != nil {
		return err
	}

	*o = WebhookRequest(varWebhookRequest)

	additionalProperties := make(map[string]interface{})

	if err = json.Unmarshal(data, &additionalProperties); err == nil {
		delete(additionalProperties, "name")
		delete(additionalProperties, "description")
		delete(additionalProperties, "payload_url")
		delete(additionalProperties, "http_method")
		delete(additionalProperties, "http_content_type")
		delete(additionalProperties, "additional_headers")
		delete(additionalProperties, "body_template")
		delete(additionalProperties, "secret")
		delete(additionalProperties, "ssl_verification")
		delete(additionalProperties, "ca_file_path")
		delete(additionalProperties, "custom_fields")
		delete(additionalProperties, "tags")
		o.AdditionalProperties = additionalProperties
	}

	return err
}

type NullableWebhookRequest struct {
	value *WebhookRequest
	isSet bool
}

func (v NullableWebhookRequest) Get() *WebhookRequest {
	return v.value
}

func (v *NullableWebhookRequest) Set(val *WebhookRequest) {
	v.value = val
	v.isSet = true
}

func (v NullableWebhookRequest) IsSet() bool {
	return v.isSet
}

func (v *NullableWebhookRequest) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableWebhookRequest(val *WebhookRequest) *NullableWebhookRequest {
	return &NullableWebhookRequest{value: val, isSet: true}
}

func (v NullableWebhookRequest) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableWebhookRequest) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
