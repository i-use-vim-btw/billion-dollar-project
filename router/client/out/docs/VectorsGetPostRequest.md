# VectorsGetPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**id** | [**VectorsGetPostRequestId**](VectorsGetPostRequestId.md) |  | [optional] 
**partition_names** | **List[str]** |  | [optional] 
**output_fields** | **List[str]** |  | [optional] 

## Example

```python
from openapi_client.models.vectors_get_post_request import VectorsGetPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of VectorsGetPostRequest from a JSON string
vectors_get_post_request_instance = VectorsGetPostRequest.from_json(json)
# print the JSON string representation of the object
print(VectorsGetPostRequest.to_json())

# convert the object into a dict
vectors_get_post_request_dict = vectors_get_post_request_instance.to_dict()
# create an instance of VectorsGetPostRequest from a dict
vectors_get_post_request_from_dict = VectorsGetPostRequest.from_dict(vectors_get_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


