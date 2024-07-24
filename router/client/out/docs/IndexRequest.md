# IndexRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**index_params** | [**List[IndexRequestIndexParamsInner]**](IndexRequestIndexParamsInner.md) |  | [optional] 

## Example

```python
from openapi_client.models.index_request import IndexRequest

# TODO update the JSON string below
json = "{}"
# create an instance of IndexRequest from a JSON string
index_request_instance = IndexRequest.from_json(json)
# print the JSON string representation of the object
print(IndexRequest.to_json())

# convert the object into a dict
index_request_dict = index_request_instance.to_dict()
# create an instance of IndexRequest from a dict
index_request_from_dict = IndexRequest.from_dict(index_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


