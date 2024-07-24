# VectorsSearchPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**partition_names** | **List[str]** |  | [optional] 
**output_fields** | **List[str]** |  | [optional] 
**anss_field** | **str** |  | [optional] 
**limit** | **int** |  | [optional] 
**offset** | **int** |  | [optional] 
**filter** | **str** |  | [optional] 
**grouping_field** | **str** |  | [optional] 
**search_params** | [**VectorsSearchPostRequestSearchParams**](VectorsSearchPostRequestSearchParams.md) |  | [optional] 

## Example

```python
from openapi_client.models.vectors_search_post_request import VectorsSearchPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of VectorsSearchPostRequest from a JSON string
vectors_search_post_request_instance = VectorsSearchPostRequest.from_json(json)
# print the JSON string representation of the object
print(VectorsSearchPostRequest.to_json())

# convert the object into a dict
vectors_search_post_request_dict = vectors_search_post_request_instance.to_dict()
# create an instance of VectorsSearchPostRequest from a dict
vectors_search_post_request_from_dict = VectorsSearchPostRequest.from_dict(vectors_search_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


