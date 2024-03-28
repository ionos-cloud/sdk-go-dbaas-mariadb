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
	"fmt"
)

// MariadbVersion The MariaDB version of your cluster.
type MariadbVersion string

// List of MariadbVersion
const (
	MARIADBVERSION__6  MariadbVersion = "10.6"
	MARIADBVERSION__11 MariadbVersion = "10.11"
)

func (v *MariadbVersion) UnmarshalJSON(src []byte) error {
	var value string
	err := json.Unmarshal(src, &value)
	if err != nil {
		return err
	}
	enumTypeValue := MariadbVersion(value)
	for _, existing := range []MariadbVersion{"10.6", "10.11"} {
		if existing == enumTypeValue {
			*v = enumTypeValue
			return nil
		}
	}

	return fmt.Errorf("%+v is not a valid MariadbVersion", value)
}

// Ptr returns reference to MariadbVersion value
func (v MariadbVersion) Ptr() *MariadbVersion {
	return &v
}

type NullableMariadbVersion struct {
	value *MariadbVersion
	isSet bool
}

func (v NullableMariadbVersion) Get() *MariadbVersion {
	return v.value
}

func (v *NullableMariadbVersion) Set(val *MariadbVersion) {
	v.value = val
	v.isSet = true
}

func (v NullableMariadbVersion) IsSet() bool {
	return v.isSet
}

func (v *NullableMariadbVersion) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableMariadbVersion(val *MariadbVersion) *NullableMariadbVersion {
	return &NullableMariadbVersion{value: val, isSet: true}
}

func (v NullableMariadbVersion) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableMariadbVersion) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}