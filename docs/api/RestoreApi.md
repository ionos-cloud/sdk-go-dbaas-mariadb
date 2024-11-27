# \RestoreApi

All URIs are relative to *https://mariadb.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersRestore**](RestoreApi.md#ClustersRestore) | **Post** /clusters/{clusterId}/restore | In-place restore of a cluster.|



## ClustersRestore

```go
var result  = ClustersRestore(ctx, clusterId)
                      .RestoreRequest(restoreRequest)
                      .Execute()
```

In-place restore of a cluster.



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
    restoreRequest := *openapiclient.NewRestoreRequest("498ae72f-411f-11eb-9d07-046c59cc737e") // RestoreRequest | The backup to restore from.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.RestoreApi.ClustersRestore(context.Background(), clusterId).RestoreRequest(restoreRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `RestoreApi.ClustersRestore``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersRestoreRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **restoreRequest** | [**RestoreRequest**](../models/RestoreRequest.md) | The backup to restore from. | |

### Return type

 (empty response body)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"RestoreApiService.ClustersRestore"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "RestoreApiService.ClustersRestore": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "RestoreApiService.ClustersRestore": {
    "port": "8443",
},
})
```

