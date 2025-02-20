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

// ClusterProperties Properties of a database cluster.
type ClusterProperties struct {
	// The friendly name of your cluster.
	DisplayName    *string         `json:"displayName,omitempty"`
	MariadbVersion *MariadbVersion `json:"mariadbVersion,omitempty"`
	// The DNS name pointing to your cluster.
	DnsName *string `json:"dnsName,omitempty"`
	// The total number of instances in the cluster (one primary and n-1 secondary).
	Instances *int32 `json:"instances,omitempty"`
	// The amount of memory per instance in gigabytes (GB).
	Ram *int32 `json:"ram,omitempty"`
	// The number of CPU cores per instance.
	Cores *int32 `json:"cores,omitempty"`
	// The amount of storage per instance in gigabytes (GB).
	StorageSize       *int32             `json:"storageSize,omitempty"`
	Connections       *[]Connection      `json:"connections,omitempty"`
	MaintenanceWindow *MaintenanceWindow `json:"maintenanceWindow,omitempty"`
	Backup            *BackupProperties  `json:"backup,omitempty"`
}

// NewClusterProperties instantiates a new ClusterProperties object
// This constructor will assign default values to properties that have it defined,
// and makes sure properties required by API are set, but the set of arguments
// will change when the set of required properties is changed
func NewClusterProperties() *ClusterProperties {
	this := ClusterProperties{}

	return &this
}

// NewClusterPropertiesWithDefaults instantiates a new ClusterProperties object
// This constructor will only assign default values to properties that have it defined,
// but it doesn't guarantee that properties required by API are set
func NewClusterPropertiesWithDefaults() *ClusterProperties {
	this := ClusterProperties{}
	return &this
}

// GetDisplayName returns the DisplayName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClusterProperties) GetDisplayName() *string {
	if o == nil {
		return nil
	}

	return o.DisplayName

}

// GetDisplayNameOk returns a tuple with the DisplayName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetDisplayNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DisplayName, true
}

// SetDisplayName sets field value
func (o *ClusterProperties) SetDisplayName(v string) {

	o.DisplayName = &v

}

// HasDisplayName returns a boolean if a field has been set.
func (o *ClusterProperties) HasDisplayName() bool {
	if o != nil && o.DisplayName != nil {
		return true
	}

	return false
}

// GetMariadbVersion returns the MariadbVersion field value
// If the value is explicit nil, the zero value for MariadbVersion will be returned
func (o *ClusterProperties) GetMariadbVersion() *MariadbVersion {
	if o == nil {
		return nil
	}

	return o.MariadbVersion

}

// GetMariadbVersionOk returns a tuple with the MariadbVersion field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetMariadbVersionOk() (*MariadbVersion, bool) {
	if o == nil {
		return nil, false
	}

	return o.MariadbVersion, true
}

// SetMariadbVersion sets field value
func (o *ClusterProperties) SetMariadbVersion(v MariadbVersion) {

	o.MariadbVersion = &v

}

// HasMariadbVersion returns a boolean if a field has been set.
func (o *ClusterProperties) HasMariadbVersion() bool {
	if o != nil && o.MariadbVersion != nil {
		return true
	}

	return false
}

// GetDnsName returns the DnsName field value
// If the value is explicit nil, the zero value for string will be returned
func (o *ClusterProperties) GetDnsName() *string {
	if o == nil {
		return nil
	}

	return o.DnsName

}

// GetDnsNameOk returns a tuple with the DnsName field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetDnsNameOk() (*string, bool) {
	if o == nil {
		return nil, false
	}

	return o.DnsName, true
}

// SetDnsName sets field value
func (o *ClusterProperties) SetDnsName(v string) {

	o.DnsName = &v

}

// HasDnsName returns a boolean if a field has been set.
func (o *ClusterProperties) HasDnsName() bool {
	if o != nil && o.DnsName != nil {
		return true
	}

	return false
}

// GetInstances returns the Instances field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ClusterProperties) GetInstances() *int32 {
	if o == nil {
		return nil
	}

	return o.Instances

}

// GetInstancesOk returns a tuple with the Instances field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetInstancesOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Instances, true
}

// SetInstances sets field value
func (o *ClusterProperties) SetInstances(v int32) {

	o.Instances = &v

}

// HasInstances returns a boolean if a field has been set.
func (o *ClusterProperties) HasInstances() bool {
	if o != nil && o.Instances != nil {
		return true
	}

	return false
}

// GetRam returns the Ram field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ClusterProperties) GetRam() *int32 {
	if o == nil {
		return nil
	}

	return o.Ram

}

// GetRamOk returns a tuple with the Ram field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetRamOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Ram, true
}

// SetRam sets field value
func (o *ClusterProperties) SetRam(v int32) {

	o.Ram = &v

}

// HasRam returns a boolean if a field has been set.
func (o *ClusterProperties) HasRam() bool {
	if o != nil && o.Ram != nil {
		return true
	}

	return false
}

// GetCores returns the Cores field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ClusterProperties) GetCores() *int32 {
	if o == nil {
		return nil
	}

	return o.Cores

}

// GetCoresOk returns a tuple with the Cores field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetCoresOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.Cores, true
}

// SetCores sets field value
func (o *ClusterProperties) SetCores(v int32) {

	o.Cores = &v

}

// HasCores returns a boolean if a field has been set.
func (o *ClusterProperties) HasCores() bool {
	if o != nil && o.Cores != nil {
		return true
	}

	return false
}

// GetStorageSize returns the StorageSize field value
// If the value is explicit nil, the zero value for int32 will be returned
func (o *ClusterProperties) GetStorageSize() *int32 {
	if o == nil {
		return nil
	}

	return o.StorageSize

}

// GetStorageSizeOk returns a tuple with the StorageSize field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetStorageSizeOk() (*int32, bool) {
	if o == nil {
		return nil, false
	}

	return o.StorageSize, true
}

// SetStorageSize sets field value
func (o *ClusterProperties) SetStorageSize(v int32) {

	o.StorageSize = &v

}

// HasStorageSize returns a boolean if a field has been set.
func (o *ClusterProperties) HasStorageSize() bool {
	if o != nil && o.StorageSize != nil {
		return true
	}

	return false
}

// GetConnections returns the Connections field value
// If the value is explicit nil, the zero value for []Connection will be returned
func (o *ClusterProperties) GetConnections() *[]Connection {
	if o == nil {
		return nil
	}

	return o.Connections

}

// GetConnectionsOk returns a tuple with the Connections field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetConnectionsOk() (*[]Connection, bool) {
	if o == nil {
		return nil, false
	}

	return o.Connections, true
}

// SetConnections sets field value
func (o *ClusterProperties) SetConnections(v []Connection) {

	o.Connections = &v

}

// HasConnections returns a boolean if a field has been set.
func (o *ClusterProperties) HasConnections() bool {
	if o != nil && o.Connections != nil {
		return true
	}

	return false
}

// GetMaintenanceWindow returns the MaintenanceWindow field value
// If the value is explicit nil, the zero value for MaintenanceWindow will be returned
func (o *ClusterProperties) GetMaintenanceWindow() *MaintenanceWindow {
	if o == nil {
		return nil
	}

	return o.MaintenanceWindow

}

// GetMaintenanceWindowOk returns a tuple with the MaintenanceWindow field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetMaintenanceWindowOk() (*MaintenanceWindow, bool) {
	if o == nil {
		return nil, false
	}

	return o.MaintenanceWindow, true
}

// SetMaintenanceWindow sets field value
func (o *ClusterProperties) SetMaintenanceWindow(v MaintenanceWindow) {

	o.MaintenanceWindow = &v

}

// HasMaintenanceWindow returns a boolean if a field has been set.
func (o *ClusterProperties) HasMaintenanceWindow() bool {
	if o != nil && o.MaintenanceWindow != nil {
		return true
	}

	return false
}

// GetBackup returns the Backup field value
// If the value is explicit nil, the zero value for BackupProperties will be returned
func (o *ClusterProperties) GetBackup() *BackupProperties {
	if o == nil {
		return nil
	}

	return o.Backup

}

// GetBackupOk returns a tuple with the Backup field value
// and a boolean to check if the value has been set.
// NOTE: If the value is an explicit nil, `nil, true` will be returned
func (o *ClusterProperties) GetBackupOk() (*BackupProperties, bool) {
	if o == nil {
		return nil, false
	}

	return o.Backup, true
}

// SetBackup sets field value
func (o *ClusterProperties) SetBackup(v BackupProperties) {

	o.Backup = &v

}

// HasBackup returns a boolean if a field has been set.
func (o *ClusterProperties) HasBackup() bool {
	if o != nil && o.Backup != nil {
		return true
	}

	return false
}

func (o ClusterProperties) MarshalJSON() ([]byte, error) {
	toSerialize := map[string]interface{}{}
	if o.DisplayName != nil {
		toSerialize["displayName"] = o.DisplayName
	}

	if o.MariadbVersion != nil {
		toSerialize["mariadbVersion"] = o.MariadbVersion
	}

	if o.DnsName != nil {
		toSerialize["dnsName"] = o.DnsName
	}

	if o.Instances != nil {
		toSerialize["instances"] = o.Instances
	}

	if o.Ram != nil {
		toSerialize["ram"] = o.Ram
	}

	if o.Cores != nil {
		toSerialize["cores"] = o.Cores
	}

	if o.StorageSize != nil {
		toSerialize["storageSize"] = o.StorageSize
	}

	if o.Connections != nil {
		toSerialize["connections"] = o.Connections
	}

	if o.MaintenanceWindow != nil {
		toSerialize["maintenanceWindow"] = o.MaintenanceWindow
	}

	if o.Backup != nil {
		toSerialize["backup"] = o.Backup
	}

	return json.Marshal(toSerialize)
}

type NullableClusterProperties struct {
	value *ClusterProperties
	isSet bool
}

func (v NullableClusterProperties) Get() *ClusterProperties {
	return v.value
}

func (v *NullableClusterProperties) Set(val *ClusterProperties) {
	v.value = val
	v.isSet = true
}

func (v NullableClusterProperties) IsSet() bool {
	return v.isSet
}

func (v *NullableClusterProperties) Unset() {
	v.value = nil
	v.isSet = false
}

func NewNullableClusterProperties(val *ClusterProperties) *NullableClusterProperties {
	return &NullableClusterProperties{value: val, isSet: true}
}

func (v NullableClusterProperties) MarshalJSON() ([]byte, error) {
	return json.Marshal(v.value)
}

func (v *NullableClusterProperties) UnmarshalJSON(src []byte) error {
	v.isSet = true
	return json.Unmarshal(src, &v.value)
}
