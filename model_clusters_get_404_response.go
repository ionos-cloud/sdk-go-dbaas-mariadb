/*
 * IONOS DBaaS MariaDB REST API
 *
 * An enterprise-grade Database is provided as a Service (DBaaS) solution that can be managed through a browser-based \"Data Center Designer\" (DCD) tool or via an easy to use API.  The API allows you to create additional MariaDB database clusters or modify existing ones. It is designed to allow users to leverage the same power and flexibility found within the DCD visual tool. Both tools are consistent with their concepts and lend well to making the experience smooth and intuitive.
 *
 * API version: 0.1.0
 */

// Code generated by OpenAPI Generator (https://openapi-generator.tech); DO NOT EDIT.

package ionoscloud

import (
	"encoding/json"
)

// ClustersGet404Response struct for ClustersGet404Response
type ClustersGet404Response struct {
	// The HTTP status code of the operation.
	HttpStatus *int32          `json:"httpStatus"`
	Messages   *[]ErrorMessage `json:"messages"`
}

// NewClustersGet404Response instantiates a new ClustersGet404Response object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClustersGet404Response(httpStatus int32, messages []ErrorMessage) *ClustersGet404Response {
	this := ClustersGet404Response{}

	this.HttpStatus = &httpStatus
	this.Messages = &messages

	return &this
}

// NewClustersGet404ResponseWithDefaults instantiates a new ClustersGet404Response object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClustersGet404ResponseWithDefaults() *ClustersGet404Response {
	this := ClustersGet404Response{}
	return &this
}

// GetHttpStatus returns the HttpStatus field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ClustersGet404Response) GetHttpStatus() *int32 {
	if o == nil {
		return nil
	}

	return o.HttpStatus

}

// GetHttpStatusOk returns a tuple with the HttpStatus field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClustersGet404Response) GetHttpStatusOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.HttpStatus, true
}

// SetHttpStatus sets field value
func (o *ClustersGet404Response) SetHttpStatus(v int32) {

	o.HttpStatus = &v

}

// HasHttpStatus returns a boolean if a field has been set.
func (o *ClustersGet404Response) HasHttpStatus() bool {
	if o != nil && o.HttpStatus != nil {
		return true
	}

	return false
}

// GetMessages returns the Messages field value
// If the value is explicit nil, the zero value for []ErrorMessage will be returned
func (o *ClustersGet404Response) GetMessages() *[]ErrorMessage {
	if o == nil {
		return nil
	}

	return o.Messages

}

// GetMessagesOk returns a tuple with the Messages field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClustersGet404Response) GetMessagesOk() (*[]ErrorMessage, bool) {
	if o == nil {
		return nil, false
	}

	return o.Messages, true
}

// SetMessages sets field value
func (o *ClustersGet404Response) SetMessages(v []ErrorMessage) {

	o.Messages = &v

}

// HasMessages returns a boolean if a field has been set.
func (o *ClustersGet404Response) HasMessages() bool {
	if o != nil && o.Messages != nil {
		return true
	}

	return false
}

func (o ClustersGet404Response) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.HttpStatus != nil {
		toSerialize["httpStatus"] = o.HttpStatus
	}

	if o.Messages != nil {
		toSerialize["messages"] = o.Messages
	}

	return json.Marshal(toSerialize)
}

type NullableClustersGet404Response struct {
	value *ClustersGet404Response
	isSet bool
}

func (v NullableClustersGet404Response) Get() *ClustersGet404Response {
	return v.value
}

func (v *NullableClustersGet404Response) Set(val *ClustersGet404Response) {
	v.value = val
	v.isSet = true
}

func (v NullableClustersGet404Response) IsSet() bool {
	return v.isSet
}

func (v *NullableClustersGet404Response) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClustersGet404Response(val *ClustersGet404Response) *NullableClustersGet404Response {
	return &NullableClustersGet404Response{value: val, isSet: true}
}

func (v NullableClustersGet404Response) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClustersGet404Response) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
