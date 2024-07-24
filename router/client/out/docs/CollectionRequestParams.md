# CollectionRequestParams


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**max_length** | **int** |  | [optional] 
**enable_dynamic_field** | **bool** |  | [optional] 
**shards_num** | **int** |  | [optional] 
**consistency_level** | **int** |  | [optional] 
**partitions_num** | **int** |  | [optional] 
**ttl_seconds** | **int** |  | [optional] 

## Example

```python
from openapi_client.models.collection_request_params import CollectionRequestParams

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionRequestParams from a JSON string
collection_request_params_instance = CollectionRequestParams.from_json(json)
# print the JSON string representation of the object
print(CollectionRequestParams.to_json())

# convert the object into a dict
collection_request_params_dict = collection_request_params_instance.to_dict()
# create an instance of CollectionRequestParams from a dict
collection_request_params_from_dict = CollectionRequestParams.from_dict(collection_request_params_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


