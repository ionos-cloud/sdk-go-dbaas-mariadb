# \ClustersApi

All URIs are relative to *https://mariadb.de-txl.ionos.com*

|Method | HTTP request | Description|
|------------- | ------------- | -------------|
|[**ClustersDelete**](ClustersApi.md#ClustersDelete) | **Delete** /clusters/{clusterId} | Delete a cluster|
|[**ClustersFindById**](ClustersApi.md#ClustersFindById) | **Get** /clusters/{clusterId} | Fetch a cluster|
|[**ClustersGet**](ClustersApi.md#ClustersGet) | **Get** /clusters | List clusters|
|[**ClustersPost**](ClustersApi.md#ClustersPost) | **Post** /clusters | Create a cluster|



## ClustersDelete

```go
var result ClusterResponse = ClustersDelete(ctx, clusterId)
                      .Execute()
```

Delete a cluster



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resp, err := apiClient.ClustersApi.ClustersDelete(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersDelete``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersDelete`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersDelete`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersDeleteRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersDelete"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersDelete": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersDelete": {
    "port": "8443",
},
})
```


## ClustersFindById

```go
var result ClusterResponse = ClustersFindById(ctx, clusterId)
                      .Execute()
```

Fetch a cluster



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

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersFindById(context.Background(), clusterId).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersFindById``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersFindById`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersFindById`: %v\n", resource)
}
```

### Path Parameters


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
|**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.|
|**clusterId** | **string** | The unique ID of the cluster. | |

### Other Parameters

Other parameters are passed through a pointer to an apiClustersFindByIdRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersFindById"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersFindById": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersFindById": {
    "port": "8443",
},
})
```


## ClustersGet

```go
var result ClusterList = ClustersGet(ctx)
                      .Limit(limit)
                      .Offset(offset)
                      .FilterName(filterName)
                      .Execute()
```

List clusters



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
    filterName := "filterName_example" // string | Response filter to list only the MariaDB clusters that contain the specified name. The value is case insensitive and matched on the 'displayName' field.  (optional)

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersGet(context.Background()).Limit(limit).Offset(offset).FilterName(filterName).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersGet``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersGet`: ClusterList
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersGet`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersGetRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **limit** | **int32** | The maximum number of elements to return. Use together with &#39;offset&#39; for pagination. | [default to 100]|
| **offset** | **int32** | The first element to return. Use together with &#39;limit&#39; for pagination. | [default to 0]|
| **filterName** | **string** | Response filter to list only the MariaDB clusters that contain the specified name. The value is case insensitive and matched on the &#39;displayName&#39; field.  | |

### Return type

[**ClusterList**](../models/ClusterList.md)

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersGet"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersGet": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersGet": {
    "port": "8443",
},
})
```


## ClustersPost

```go
var result ClusterResponse = ClustersPost(ctx)
                      .CreateClusterRequest(createClusterRequest)
                      .Execute()
```

Create a cluster



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
    createClusterRequest := *openapiclient.NewCreateClusterRequest() // CreateClusterRequest | The cluster to be created.

    configuration := ionoscloud.NewConfiguration("USERNAME", "PASSWORD", "TOKEN", "HOST_URL")
    apiClient := ionoscloud.NewAPIClient(configuration)
    resource, resp, err := apiClient.ClustersApi.ClustersPost(context.Background()).CreateClusterRequest(createClusterRequest).Execute()
    if err != nil {
        fmt.Fprintf(os.Stderr, "Error when calling `ClustersApi.ClustersPost``: %v\n", err)
        fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", resp)
    }
    // response from `ClustersPost`: ClusterResponse
    fmt.Fprintf(os.Stdout, "Response from `ClustersApi.ClustersPost`: %v\n", resource)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to an apiClustersPostRequest struct via the builder pattern


|Name | Type | Description  | Notes|
|------------- | ------------- | ------------- | -------------|
| **createClusterRequest** | [**CreateClusterRequest**](../models/CreateClusterRequest.md) | The cluster to be created. | |

### Return type

[**ClusterResponse**](../models/ClusterResponse.md)

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json


### URLs Configuration per Operation
Each operation can use different server URL defined using `OperationServers` map in the `Configuration`.
An operation is uniquely identified by `"ClustersApiService.ClustersPost"` string.
Similar rules for overriding default operation server index and variables apply by using `sw.ContextOperationServerIndices` and `sw.ContextOperationServerVariables` context maps.

```golang
ctx := context.WithValue(context.Background(), {packageName}.ContextOperationServerIndices, map[string]int{
    "ClustersApiService.ClustersPost": 2,
})
ctx = context.WithValue(context.Background(), {packageName}.ContextOperationServerVariables, map[string]map[string]string{
    "ClustersApiService.ClustersPost": {
    "port": "8443",
},
})
```

