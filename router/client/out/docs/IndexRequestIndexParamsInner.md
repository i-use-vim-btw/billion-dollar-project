# IndexRequestIndexParamsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metric_type** | **str** |  | [optional] 
**field_name** | **str** |  | [optional] 
**index_name** | **str** |  | [optional] 
**params** | [**IndexRequestIndexParamsInnerParams**](IndexRequestIndexParamsInnerParams.md) |  | [optional] 

## Example

```python
from openapi_client.models.index_request_index_params_inner import IndexRequestIndexParamsInner

# TODO update the JSON string below
json = "{}"
# create an instance of IndexRequestIndexParamsInner from a JSON string
index_request_index_params_inner_instance = IndexRequestIndexParamsInner.from_json(json)
# print the JSON string representation of the object
print(IndexRequestIndexParamsInner.to_json())

# convert the object into a dict
index_request_index_params_inner_dict = index_request_index_params_inner_instance.to_dict()
# create an instance of IndexRequestIndexParamsInner from a dict
index_request_index_params_inner_from_dict = IndexRequestIndexParamsInner.from_dict(index_request_index_params_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


