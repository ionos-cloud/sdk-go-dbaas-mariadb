# \BackupsApi

All URIs are relative to *https://mariadb.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**BackupsFindById**](BackupsApi.md#BackupsFindById) | **Get** /backups/{backupId} | Fetch backups|
|[**BackupsGet**](BackupsApi.md#BackupsGet) | **Get** /backups | List of backups.|
|[**ClusterBackupsGet**](BackupsApi.md#ClusterBackupsGet) | **Get** /clusters/{clusterId}/backups | List backups of cluster|



## BackupsFindById

```go
var result BackupResponse = BackupsFindById(ctx, backupId)
                      .Execute()
```

Fetch backups



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mariadb"
)

func main() {
    backupId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the backup.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.BackupsFindById(context.Background(), backupId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.BackupsFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `BackupsFindById`: BackupResponse
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.BackupsFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**backupId** | **string** | The unique ID of the backup. | |

### Other Parameters

Other parameters are passed through a pointer to an apiBackupsFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**BackupResponse**](../models/BackupResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BackupsApiService.BackupsFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "BackupsApiService.BackupsFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "BackupsApiService.BackupsFindById": {
    "port": "8443",
},
})
```


## BackupsGet

```go
var result BackupList = BackupsGet(ctx)
                      .Limit(limit)
                      .Offset(offset)
                      .Execute()
```

List of backups.



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mariadb"
)

func main() {
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.BackupsGet(context.Background()).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.BackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `BackupsGet`: BackupList
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.BackupsGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiBackupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|

### Return type

[**BackupList**](../models/BackupList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BackupsApiService.BackupsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "BackupsApiService.BackupsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "BackupsApiService.BackupsGet": {
    "port": "8443",
},
})
```


## ClusterBackupsGet

```go
var result BackupList = ClusterBackupsGet(ctx, clusterId)
                      .Limit(limit)
                      .Offset(offset)
                      .Execute()
```

List backups of cluster



### Example

```go
package main

import (
    "context"
    "fmt"
    "os"

    ionoscloud "github.com/ionos-cloud/sdk-go-dbaas-mariadb"
)

func main() {
    clusterId := "498ae72f-411f-11eb-9d07-046c59cc737e" // string | The unique ID of the cluster.
    limit := int32(100) // int32 | The maximum number of elements to return. Use together with 'offset' for pagination. (optional) (default to 100)
    offset := int32(200) // int32 | The first element to return. Use together with 'limit' for pagination. (optional) (default to 0)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.BackupsApi.ClusterBackupsGet(context.Background(), clusterId).Limit(limit).Offset(offset).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `BackupsApi.ClusterBackupsGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClusterBackupsGet`: BackupList
    fmt.Fprintf(os.Stdout, "Response from `BackupsApi.ClusterBackupsGet`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClusterBackupsGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|

### Return type

[**BackupList**](../models/BackupList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"BackupsApiService.ClusterBackupsGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "BackupsApiService.ClusterBackupsGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "BackupsApiService.ClusterBackupsGet": {
    "port": "8443",
},
})
```

