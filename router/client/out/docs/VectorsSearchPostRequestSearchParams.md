# VectorsSearchPostRequestSearchParams


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**metric_type** | **str** |  | [optional] 
**params** | [**VectorsSearchPostRequestSearchParamsParams**](VectorsSearchPostRequestSearchParamsParams.md) |  | [optional] 

## Example

```python
from openapi_client.models.vectors_search_post_request_search_params import VectorsSearchPostRequestSearchParams

# TODO update the JSON string below
json = "{}"
# create an instance of VectorsSearchPostRequestSearchParams from a JSON string
vectors_search_post_request_search_params_instance = VectorsSearchPostRequestSearchParams.from_json(json)
# print the JSON string representation of the object
print(VectorsSearchPostRequestSearchParams.to_json())

# convert the object into a dict
vectors_search_post_request_search_params_dict = vectors_search_post_request_search_params_instance.to_dict()
# create an instance of VectorsSearchPostRequestSearchParams from a dict
vectors_search_post_request_search_params_from_dict = VectorsSearchPostRequestSearchParams.from_dict(vectors_search_post_request_search_params_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


