# VectorsQueryPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**partition_names** | **List[str]** |  | [optional] 
**filter** | **str** |  | 
**output_fields** | **List[str]** |  | [optional] 

## Example

```python
from openapi_client.models.vectors_query_post_request import VectorsQueryPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of VectorsQueryPostRequest from a JSON string
vectors_query_post_request_instance = VectorsQueryPostRequest.from_json(json)
# print the JSON string representation of the object
print(VectorsQueryPostRequest.to_json())

# convert the object into a dict
vectors_query_post_request_dict = vectors_query_post_request_instance.to_dict()
# create an instance of VectorsQueryPostRequest from a dict
vectors_query_post_request_from_dict = VectorsQueryPostRequest.from_dict(vectors_query_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


