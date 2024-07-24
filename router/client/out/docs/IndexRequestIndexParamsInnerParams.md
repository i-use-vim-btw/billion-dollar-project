# IndexRequestIndexParamsInnerParams


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**index_type** | **str** |  | [optional] 
**m** | **int** |  | [optional] 
**ef_construction** | **int** |  | [optional] 
**nlist** | **int** |  | [optional] 

## Example

```python
from openapi_client.models.index_request_index_params_inner_params import IndexRequestIndexParamsInnerParams

# TODO update the JSON string below
json = "{}"
# create an instance of IndexRequestIndexParamsInnerParams from a JSON string
index_request_index_params_inner_params_instance = IndexRequestIndexParamsInnerParams.from_json(json)
# print the JSON string representation of the object
print(IndexRequestIndexParamsInnerParams.to_json())

# convert the object into a dict
index_request_index_params_inner_params_dict = index_request_index_params_inner_params_instance.to_dict()
# create an instance of IndexRequestIndexParamsInnerParams from a dict
index_request_index_params_inner_params_from_dict = IndexRequestIndexParamsInnerParams.from_dict(index_request_index_params_inner_params_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


