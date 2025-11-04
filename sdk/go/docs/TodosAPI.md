# \TodosAPI

All URIs are relative to *http://localhost:8081*

Method | HTTP request | Description
------------- | ------------- | -------------
[**TodosCreate**](TodosAPI.md#TodosCreate) | **Post** /v1/todos | Create a new todo
[**TodosDelete**](TodosAPI.md#TodosDelete) | **Delete** /v1/todos/{id} | Delete a todo
[**TodosGet**](TodosAPI.md#TodosGet) | **Get** /v1/todos | Get all todos
[**TodosGetByID**](TodosAPI.md#TodosGetByID) | **Get** /v1/todos/{id} | Get a todo by ID
[**TodosUpdate**](TodosAPI.md#TodosUpdate) | **Put** /v1/todos/{id} | Update an existing todo



## TodosCreate

> ModelsTodo TodosCreate(ctx).Todo(todo).Execute()

Create a new todo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	todo := *openapiclient.NewModelsTodo() // ModelsTodo | Todo item details

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodosAPI.TodosCreate(context.Background()).Todo(todo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodosAPI.TodosCreate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TodosCreate`: ModelsTodo
	fmt.Fprintf(os.Stdout, "Response from `TodosAPI.TodosCreate`: %v\n", resp)
}
```

### Path Parameters



### Other Parameters

Other parameters are passed through a pointer to a apiTodosCreateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **todo** | [**ModelsTodo**](ModelsTodo.md) | Todo item details | 

### Return type

[**ModelsTodo**](ModelsTodo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TodosDelete

> string TodosDelete(ctx, id).Execute()

Delete a todo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Todo ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodosAPI.TodosDelete(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodosAPI.TodosDelete``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TodosDelete`: string
	fmt.Fprintf(os.Stdout, "Response from `TodosAPI.TodosDelete`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Todo ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTodosDeleteRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

**string**

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TodosGet

> []ModelsTodo TodosGet(ctx).Execute()

Get all todos



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodosAPI.TodosGet(context.Background()).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodosAPI.TodosGet``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TodosGet`: []ModelsTodo
	fmt.Fprintf(os.Stdout, "Response from `TodosAPI.TodosGet`: %v\n", resp)
}
```

### Path Parameters

This endpoint does not need any parameter.

### Other Parameters

Other parameters are passed through a pointer to a apiTodosGetRequest struct via the builder pattern


### Return type

[**[]ModelsTodo**](ModelsTodo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TodosGetByID

> ModelsTodo TodosGetByID(ctx, id).Execute()

Get a todo by ID



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Todo ID

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodosAPI.TodosGetByID(context.Background(), id).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodosAPI.TodosGetByID``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TodosGetByID`: ModelsTodo
	fmt.Fprintf(os.Stdout, "Response from `TodosAPI.TodosGetByID`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Todo ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTodosGetByIDRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------


### Return type

[**ModelsTodo**](ModelsTodo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: Not defined
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)


## TodosUpdate

> ModelsTodo TodosUpdate(ctx, id).Todo(todo).Execute()

Update an existing todo



### Example

```go
package main

import (
	"context"
	"fmt"
	"os"
	openapiclient "github.com/GIT_USER_ID/GIT_REPO_ID"
)

func main() {
	id := "id_example" // string | Todo ID
	todo := *openapiclient.NewModelsTodo() // ModelsTodo | Todo item details to update

	configuration := openapiclient.NewConfiguration()
	apiClient := openapiclient.NewAPIClient(configuration)
	resp, r, err := apiClient.TodosAPI.TodosUpdate(context.Background(), id).Todo(todo).Execute()
	if err != nil {
		fmt.Fprintf(os.Stderr, "Error when calling `TodosAPI.TodosUpdate``: %v\n", err)
		fmt.Fprintf(os.Stderr, "Full HTTP response: %v\n", r)
	}
	// response from `TodosUpdate`: ModelsTodo
	fmt.Fprintf(os.Stdout, "Response from `TodosAPI.TodosUpdate`: %v\n", resp)
}
```

### Path Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
**ctx** | **context.Context** | context for authentication, logging, cancellation, deadlines, tracing, etc.
**id** | **string** | Todo ID | 

### Other Parameters

Other parameters are passed through a pointer to a apiTodosUpdateRequest struct via the builder pattern


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------

 **todo** | [**ModelsTodo**](ModelsTodo.md) | Todo item details to update | 

### Return type

[**ModelsTodo**](ModelsTodo.md)

### Authorization

No authorization required

### HTTP request headers

- **Content-Type**: application/json
- **Accept**: application/json

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints)
[[Back to Model list]](../README.md#documentation-for-models)
[[Back to README]](../README.md)

