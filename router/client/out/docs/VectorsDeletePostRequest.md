# VectorsDeletePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**filter** | **str** |  | 
**partition_name** | **str** |  | [optional] 

## Example

```python
from openapi_client.models.vectors_delete_post_request import VectorsDeletePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of VectorsDeletePostRequest from a JSON string
vectors_delete_post_request_instance = VectorsDeletePostRequest.from_json(json)
# print the JSON string representation of the object
print(VectorsDeletePostRequest.to_json())

# convert the object into a dict
vectors_delete_post_request_dict = vectors_delete_post_request_instance.to_dict()
# create an instance of VectorsDeletePostRequest from a dict
vectors_delete_post_request_from_dict = VectorsDeletePostRequest.from_dict(vectors_delete_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


