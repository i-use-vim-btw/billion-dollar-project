# openapi_client.DefaultApi

All URIs are relative to *https://api.gandi.io/v1*

Method | HTTP request | Description
------------- | ------------- | -------------
[**collections_create_post**](DefaultApi.md#collections_create_post) | **POST** /collections/create | Create a new collection
[**collections_describe_post**](DefaultApi.md#collections_describe_post) | **POST** /collections/describe | Describe a collection
[**collections_drop_post**](DefaultApi.md#collections_drop_post) | **POST** /collections/drop | Drop a collection
[**collections_get_load_state_post**](DefaultApi.md#collections_get_load_state_post) | **POST** /collections/get_load_state | Get the load state of a collection
[**collections_get_stats_post**](DefaultApi.md#collections_get_stats_post) | **POST** /collections/get_stats | Get the stats of a collection
[**collections_has_post**](DefaultApi.md#collections_has_post) | **POST** /collections/has | Check if a collection exists
[**collections_list_post**](DefaultApi.md#collections_list_post) | **POST** /collections/list | List all collections
[**collections_load_post**](DefaultApi.md#collections_load_post) | **POST** /collections/load | Load a collection
[**collections_release_post**](DefaultApi.md#collections_release_post) | **POST** /collections/release | Release a collection
[**collections_rename_post**](DefaultApi.md#collections_rename_post) | **POST** /collections/rename | Rename a collection
[**indexes_create_post**](DefaultApi.md#indexes_create_post) | **POST** /indexes/create | Create a new index
[**indexes_describe_post**](DefaultApi.md#indexes_describe_post) | **POST** /indexes/describe | Describe an index
[**indexes_drop_post**](DefaultApi.md#indexes_drop_post) | **POST** /indexes/drop | Drop an index
[**indexes_list_post**](DefaultApi.md#indexes_list_post) | **POST** /indexes/list | List all indexes
[**vectors_delete_post**](DefaultApi.md#vectors_delete_post) | **POST** /vectors/delete | Delete vectors
[**vectors_get_post**](DefaultApi.md#vectors_get_post) | **POST** /vectors/get | Get vectors
[**vectors_insert_post**](DefaultApi.md#vectors_insert_post) | **POST** /vectors/insert | Insert vectors
[**vectors_query_post**](DefaultApi.md#vectors_query_post) | **POST** /vectors/query | Query vectors
[**vectors_search_post**](DefaultApi.md#vectors_search_post) | **POST** /vectors/search | Search vectors
[**vectors_upsert_post**](DefaultApi.md#vectors_upsert_post) | **POST** /vectors/upsert | Upsert vectors


# **collections_create_post**
> collections_create_post(collection_request)

Create a new collection

This operation creates a collection in a specified cluster.

### Example


```python
import openapi_client
from openapi_client.models.collection_request import CollectionRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collection_request = openapi_client.CollectionRequest() # CollectionRequest | Collection object that needs to be added

    try:
        # Create a new collection
        api_instance.collections_create_post(collection_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_create_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collection_request** | [**CollectionRequest**](CollectionRequest.md)| Collection object that needs to be added | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Collection created successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_describe_post**
> collections_describe_post(collections_describe_post_request)

Describe a collection

Returns the details of a collection.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be described

    try:
        # Describe a collection
        api_instance.collections_describe_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_describe_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be described | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection schema returned successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_drop_post**
> collections_drop_post(collections_describe_post_request)

Drop a collection

This operation drops the current collection and all data within the collection.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be deleted

    try:
        # Drop a collection
        api_instance.collections_drop_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_drop_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be deleted | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection dropped successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_get_load_state_post**
> collections_get_load_state_post(collections_get_load_state_post_request)

Get the load state of a collection

Returns the load state of a specific collection.

### Example


```python
import openapi_client
from openapi_client.models.collections_get_load_state_post_request import CollectionsGetLoadStatePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_get_load_state_post_request = openapi_client.CollectionsGetLoadStatePostRequest() # CollectionsGetLoadStatePostRequest | Collection object that needs to be checked

    try:
        # Get the load state of a collection
        api_instance.collections_get_load_state_post(collections_get_load_state_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_get_load_state_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_get_load_state_post_request** | [**CollectionsGetLoadStatePostRequest**](CollectionsGetLoadStatePostRequest.md)| Collection object that needs to be checked | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Load state returned successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_get_stats_post**
> collections_get_stats_post(collections_describe_post_request)

Get the stats of a collection

This operation gets the number of entities in a collection.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be checked

    try:
        # Get the stats of a collection
        api_instance.collections_get_stats_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_get_stats_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be checked | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Stats returned successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_has_post**
> collections_has_post(collections_describe_post_request)

Check if a collection exists

Checks if a collection exists in the database.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be checked

    try:
        # Check if a collection exists
        api_instance.collections_has_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_has_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be checked | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection exists. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_list_post**
> collections_list_post(collections_list_post_request)

List all collections

Returns a list of all collections in the specified database.

### Example


```python
import openapi_client
from openapi_client.models.collections_list_post_request import CollectionsListPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_list_post_request = openapi_client.CollectionsListPostRequest() # CollectionsListPostRequest | Database object that needs to be listed

    try:
        # List all collections
        api_instance.collections_list_post(collections_list_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_list_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_list_post_request** | [**CollectionsListPostRequest**](CollectionsListPostRequest.md)| Database object that needs to be listed | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collections listed successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_load_post**
> collections_load_post(collections_describe_post_request)

Load a collection

Loads a collection into memory.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be loaded

    try:
        # Load a collection
        api_instance.collections_load_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_load_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be loaded | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection loaded successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_release_post**
> collections_release_post(collections_describe_post_request)

Release a collection

Releases a collection from memory.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be released

    try:
        # Release a collection
        api_instance.collections_release_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_release_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be released | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection released successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **collections_rename_post**
> collections_rename_post(collections_rename_post_request)

Rename a collection

This operation renames an existing collection and optionally moves the collection to a new database.

### Example


```python
import openapi_client
from openapi_client.models.collections_rename_post_request import CollectionsRenamePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_rename_post_request = openapi_client.CollectionsRenamePostRequest() # CollectionsRenamePostRequest | Collection object that needs to be renamed

    try:
        # Rename a collection
        api_instance.collections_rename_post(collections_rename_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->collections_rename_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_rename_post_request** | [**CollectionsRenamePostRequest**](CollectionsRenamePostRequest.md)| Collection object that needs to be renamed | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Collection renamed successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **indexes_create_post**
> indexes_create_post(index_request)

Create a new index

This creates a named index for a target field, which can either be a vector field or a scalar field.

### Example


```python
import openapi_client
from openapi_client.models.index_request import IndexRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    index_request = openapi_client.IndexRequest() # IndexRequest | Index object that needs to be added

    try:
        # Create a new index
        api_instance.indexes_create_post(index_request)
    except Exception as e:
        print("Exception when calling DefaultApi->indexes_create_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **index_request** | [**IndexRequest**](IndexRequest.md)| Index object that needs to be added | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Index created successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **indexes_describe_post**
> indexes_describe_post(indexes_describe_post_request)

Describe an index

Returns the details of an index.

### Example


```python
import openapi_client
from openapi_client.models.indexes_describe_post_request import IndexesDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    indexes_describe_post_request = openapi_client.IndexesDescribePostRequest() # IndexesDescribePostRequest | Index object that needs to be described

    try:
        # Describe an index
        api_instance.indexes_describe_post(indexes_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->indexes_describe_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexes_describe_post_request** | [**IndexesDescribePostRequest**](IndexesDescribePostRequest.md)| Index object that needs to be described | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Index schema returned successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **indexes_drop_post**
> indexes_drop_post(indexes_describe_post_request)

Drop an index

This operation drops index from a specified collection.

### Example


```python
import openapi_client
from openapi_client.models.indexes_describe_post_request import IndexesDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    indexes_describe_post_request = openapi_client.IndexesDescribePostRequest() # IndexesDescribePostRequest | Index object that needs to be deleted

    try:
        # Drop an index
        api_instance.indexes_drop_post(indexes_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->indexes_drop_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **indexes_describe_post_request** | [**IndexesDescribePostRequest**](IndexesDescribePostRequest.md)| Index object that needs to be deleted | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Index dropped successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **indexes_list_post**
> indexes_list_post(collections_describe_post_request)

List all indexes

Returns a list of all indexes in the specified collection.

### Example


```python
import openapi_client
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    collections_describe_post_request = openapi_client.CollectionsDescribePostRequest() # CollectionsDescribePostRequest | Collection object that needs to be listed

    try:
        # List all indexes
        api_instance.indexes_list_post(collections_describe_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->indexes_list_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **collections_describe_post_request** | [**CollectionsDescribePostRequest**](CollectionsDescribePostRequest.md)| Collection object that needs to be listed | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Indexes listed successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vectors_delete_post**
> vectors_delete_post(vectors_delete_post_request)

Delete vectors

This operation deletes entities by their IDs or with a boolean expression.

### Example


```python
import openapi_client
from openapi_client.models.vectors_delete_post_request import VectorsDeletePostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    vectors_delete_post_request = openapi_client.VectorsDeletePostRequest() # VectorsDeletePostRequest | Vectors object that needs to be deleted

    try:
        # Delete vectors
        api_instance.vectors_delete_post(vectors_delete_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->vectors_delete_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectors_delete_post_request** | [**VectorsDeletePostRequest**](VectorsDeletePostRequest.md)| Vectors object that needs to be deleted | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Vectors deleted successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vectors_get_post**
> vectors_get_post(vectors_get_post_request)

Get vectors

This operation gets vectors by their IDs.

### Example


```python
import openapi_client
from openapi_client.models.vectors_get_post_request import VectorsGetPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    vectors_get_post_request = openapi_client.VectorsGetPostRequest() # VectorsGetPostRequest | Vectors object that needs to be retrieved

    try:
        # Get vectors
        api_instance.vectors_get_post(vectors_get_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->vectors_get_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectors_get_post_request** | [**VectorsGetPostRequest**](VectorsGetPostRequest.md)| Vectors object that needs to be retrieved | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Vectors retrieved successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vectors_insert_post**
> vectors_insert_post(vectors_insert_post_request)

Insert vectors

This operation inserts vectors into a specified collection.

### Example


```python
import openapi_client
from openapi_client.models.vectors_insert_post_request import VectorsInsertPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    vectors_insert_post_request = openapi_client.VectorsInsertPostRequest() # VectorsInsertPostRequest | Vectors object that needs to be inserted

    try:
        # Insert vectors
        api_instance.vectors_insert_post(vectors_insert_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->vectors_insert_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectors_insert_post_request** | [**VectorsInsertPostRequest**](VectorsInsertPostRequest.md)| Vectors object that needs to be inserted | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Vectors inserted successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vectors_query_post**
> vectors_query_post(vectors_query_post_request)

Query vectors

This operation queries vectors in a specified collection.

### Example


```python
import openapi_client
from openapi_client.models.vectors_query_post_request import VectorsQueryPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    vectors_query_post_request = openapi_client.VectorsQueryPostRequest() # VectorsQueryPostRequest | Vectors object that needs to be queried

    try:
        # Query vectors
        api_instance.vectors_query_post(vectors_query_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->vectors_query_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectors_query_post_request** | [**VectorsQueryPostRequest**](VectorsQueryPostRequest.md)| Vectors object that needs to be queried | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Vectors queried successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vectors_search_post**
> vectors_search_post(vectors_search_post_request)

Search vectors

This operation searches vectors in a specified collection.

### Example


```python
import openapi_client
from openapi_client.models.vectors_search_post_request import VectorsSearchPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    vectors_search_post_request = openapi_client.VectorsSearchPostRequest() # VectorsSearchPostRequest | Vectors object that needs to be searched

    try:
        # Search vectors
        api_instance.vectors_search_post(vectors_search_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->vectors_search_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectors_search_post_request** | [**VectorsSearchPostRequest**](VectorsSearchPostRequest.md)| Vectors object that needs to be searched | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**200** | Vectors searched successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

# **vectors_upsert_post**
> vectors_upsert_post(vectors_insert_post_request)

Upsert vectors

This operation upserts vectors into a specified collection.

### Example


```python
import openapi_client
from openapi_client.models.vectors_insert_post_request import VectorsInsertPostRequest
from openapi_client.rest import ApiException
from pprint import pprint

# Defining the host is optional and defaults to https://api.gandi.io/v1
# See configuration.py for a list of all supported configuration parameters.
configuration = openapi_client.Configuration(
    host = "https://api.gandi.io/v1"
)


# Enter a context with an instance of the API client
with openapi_client.ApiClient(configuration) as api_client:
    # Create an instance of the API class
    api_instance = openapi_client.DefaultApi(api_client)
    vectors_insert_post_request = openapi_client.VectorsInsertPostRequest() # VectorsInsertPostRequest | Vectors object that needs to be upserted

    try:
        # Upsert vectors
        api_instance.vectors_upsert_post(vectors_insert_post_request)
    except Exception as e:
        print("Exception when calling DefaultApi->vectors_upsert_post: %s\n" % e)
```



### Parameters


Name | Type | Description  | Notes
------------- | ------------- | ------------- | -------------
 **vectors_insert_post_request** | [**VectorsInsertPostRequest**](VectorsInsertPostRequest.md)| Vectors object that needs to be upserted | 

### Return type

void (empty response body)

### Authorization

No authorization required

### HTTP request headers

 - **Content-Type**: application/json
 - **Accept**: Not defined

### HTTP response details

| Status code | Description | Response headers |
|-------------|-------------|------------------|
**201** | Vectors upserted successfully. |  -  |
**400** | Bad Request |  -  |

[[Back to top]](#) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to Model list]](../README.md#documentation-for-models) [[Back to README]](../README.md)

