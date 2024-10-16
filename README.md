# Go API client for ionoscloud

An enterprise-grade Database is provided as a Service (DBaaS) solution that
can be managed through a browser-based \"Data Center Designer\" (DCD) tool or
via an easy to use API.

The API allows you to create additional MariaDB database clusters or modify existing
ones. It is designed to allow users to leverage the same power and
flexibility found within the DCD visual tool. Both tools are consistent with
their concepts and lend well to making the experience smooth and intuitive.


## Overview
This API client was generated by the [OpenAPI Generator](https://openapi-generator.tech) project.  By using the [OpenAPI-spec](https://www.openapis.org/) from a remote server, you can easily generate an API client.

- API version: 0.1.0
- Package version: 1.1.0
- Build package: org.openapitools.codegen.languages.GoClientCodegen

## Installation

Install the following dependencies:

```shell
go get github.com/stretchr/testify/assert
go get golang.org/x/net/context
```

Put the package under your project folder and add the following in import:

```golang
import ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mariadb"
```

To use a proxy, set the environment variable `HTTP_PROXY`:

```golang
os.Setenv("HTTP_PROXY", "http://proxy_name:proxy_port")
```

## Configuration of Server URL

Default configuration comes with `Servers` field that contains server objects as defined in the OpenAPI specification.

### Select Server Configuration

For using other server than the one defined on index 0 set context value `sw.ContextServerIndex` of type `int`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerIndex, 1)
```

### Templated Server URL

Templated server URL is formatted using default variables from configuration or from context value `sw.ContextServerVariables` of type `map[string]string`.

```golang
ctx := context.WithValue(context.Background(), ionoscloud.ContextServerVariables, map[string]string{
	"basePath": "v2",
})
```

Note, enum values are always validated and all unused variables are silently ignored.

## Documentation for API Endpoints

All URIs are relative to *https://mariadb.de-txl.ionos.com*

Class | Method | HTTP request | Description
------------ | ------------- | ------------- | -------------
*BackupsApi* | [**BackupsFindById**](docs/api/BackupsApi.md#backupsfindbyid) | **Get** /backups/{backupId} | Fetch backups
*BackupsApi* | [**BackupsGet**](docs/api/BackupsApi.md#backupsget) | **Get** /backups | List of backups.
*BackupsApi* | [**ClusterBackupsGet**](docs/api/BackupsApi.md#clusterbackupsget) | **Get** /clusters/{clusterId}/backups | List backups of cluster
*ClustersApi* | [**ClustersDelete**](docs/api/ClustersApi.md#clustersdelete) | **Delete** /clusters/{clusterId} | Delete a cluster
*ClustersApi* | [**ClustersFindById**](docs/api/ClustersApi.md#clustersfindbyid) | **Get** /clusters/{clusterId} | Fetch a cluster
*ClustersApi* | [**ClustersGet**](docs/api/ClustersApi.md#clustersget) | **Get** /clusters | List clusters
*ClustersApi* | [**ClustersPost**](docs/api/ClustersApi.md#clusterspost) | **Post** /clusters | Create a cluster


## Documentation For Models

 - [Backup](docs/models/Backup.md)
 - [BackupList](docs/models/BackupList.md)
 - [BackupListAllOf](docs/models/BackupListAllOf.md)
 - [BackupResponse](docs/models/BackupResponse.md)
 - [BaseBackup](docs/models/BaseBackup.md)
 - [ClusterList](docs/models/ClusterList.md)
 - [ClusterListAllOf](docs/models/ClusterListAllOf.md)
 - [ClusterMetadata](docs/models/ClusterMetadata.md)
 - [ClusterProperties](docs/models/ClusterProperties.md)
 - [ClusterResponse](docs/models/ClusterResponse.md)
 - [ClustersGet400Response](docs/models/ClustersGet400Response.md)
 - [ClustersGet401Response](docs/models/ClustersGet401Response.md)
 - [ClustersGet403Response](docs/models/ClustersGet403Response.md)
 - [ClustersGet404Response](docs/models/ClustersGet404Response.md)
 - [ClustersGet405Response](docs/models/ClustersGet405Response.md)
 - [ClustersGet415Response](docs/models/ClustersGet415Response.md)
 - [ClustersGet422Response](docs/models/ClustersGet422Response.md)
 - [ClustersGet429Response](docs/models/ClustersGet429Response.md)
 - [ClustersGet500Response](docs/models/ClustersGet500Response.md)
 - [ClustersGet503Response](docs/models/ClustersGet503Response.md)
 - [Connection](docs/models/Connection.md)
 - [CreateClusterProperties](docs/models/CreateClusterProperties.md)
 - [CreateClusterRequest](docs/models/CreateClusterRequest.md)
 - [DBUser](docs/models/DBUser.md)
 - [DayOfTheWeek](docs/models/DayOfTheWeek.md)
 - [ErrorMessage](docs/models/ErrorMessage.md)
 - [MaintenanceWindow](docs/models/MaintenanceWindow.md)
 - [MariadbVersion](docs/models/MariadbVersion.md)
 - [Pagination](docs/models/Pagination.md)
 - [PaginationLinks](docs/models/PaginationLinks.md)
 - [State](docs/models/State.md)


## Documentation For Authorization


Authentication schemes defined for the API:
### tokenAuth

- **Type**: API key
- **API key parameter name**: Authorization
- **Location**: HTTP header

Note, each API key must be added to a map of `map[string]APIKey` where the key is: Authorization and passed in as the auth context for each request.

Example

```golang
auth := context.WithValue(
		context.Background(),
		sw.ContextAPIKeys,
		map[string]sw.APIKey{
			"Authorization": {Key: "API_KEY_STRING"},
		},
	)
r, err := client.Service.Operation(auth, args)
```


## Documentation for Utility Methods

Due to the fact that model structure members are all pointers, this package contains
a number of utility functions to easily obtain pointers to values of basic types.
Each of these functions takes a value of the given basic type and returns a pointer to it:

* `PtrBool`
* `PtrInt`
* `PtrInt32`
* `PtrInt64`
* `PtrFloat`
* `PtrFloat32`
* `PtrFloat64`
* `PtrString`
* `PtrTime`

## Author



