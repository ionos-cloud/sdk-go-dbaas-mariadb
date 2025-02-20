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

// BackupProperties Properties configuring the backup of the cluster.
type BackupProperties struct {
	// The S3 location where the backups will be stored.
	Location *string `json:"location,omitempty"`
}

// NewBackupProperties instantiates a new BackupProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewBackupProperties() *BackupProperties {
	this := BackupProperties{}

	return &this
}

// NewBackupPropertiesWithDefaults instantiates a new BackupProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewBackupPropertiesWithDefaults() *BackupProperties {
	this := BackupProperties{}
	return &this
}

// GetLocation returns the Location field value
// If the value is explicit nil, the zero value for string will be returned
func (o *BackupProperties) GetLocation() *string {
	if o == nil {
		return nil
	}

	return o.Location

}

// GetLocationOk returns a tuple with the Location field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *BackupProperties) GetLocationOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.Location, true
}

// SetLocation sets field value
func (o *BackupProperties) SetLocation(v string) {

	o.Location = &v

}

// HasLocation returns a boolean if a field has been set.
func (o *BackupProperties) HasLocation() bool {
	if o != nil && o.Location != nil {
		return true
	}

	return false
}

func (o BackupProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.Location != nil {
		toSerialize["location"] = o.Location
	}

	return json.Marshal(toSerialize)
}

type NullableBackupProperties struct {
	value *BackupProperties
	isSet bool
}

func (v NullableBackupProperties) Get() *BackupProperties {
	return v.value
}

func (v *NullableBackupProperties) Set(val *BackupProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableBackupProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableBackupProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableBackupProperties(val *BackupProperties) *NullableBackupProperties {
	return &NullableBackupProperties{value: val, isSet: true}
}

func (v NullableBackupProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableBackupProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
