# CollectionRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**dimension** | **int** |  | [optional] 
**metric_type** | **str** |  | [optional] 
**id_type** | **str** |  | [optional] 
**auto_id** | **str** |  | [optional] 
**primary_field_name** | **str** |  | [optional] 
**vector_field_name** | **str** |  | [optional] 
**var_schema** | [**CollectionRequestSchema**](CollectionRequestSchema.md) |  | [optional] 
**index_params** | [**List[IndexRequestIndexParamsInner]**](IndexRequestIndexParamsInner.md) |  | [optional] 
**params** | [**CollectionRequestParams**](CollectionRequestParams.md) |  | [optional] 

## Example

```python
from openapi_client.models.collection_request import CollectionRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionRequest from a JSON string
collection_request_instance = CollectionRequest.from_json(json)
# print the JSON string representation of the object
print(CollectionRequest.to_json())

# convert the object into a dict
collection_request_dict = collection_request_instance.to_dict()
# create an instance of CollectionRequest from a dict
collection_request_from_dict = CollectionRequest.from_dict(collection_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


