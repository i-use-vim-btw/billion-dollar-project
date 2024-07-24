# VectorsInsertPostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**partition_name** | **str** |  | [optional] 
**data** | [**Dict[str, VectorsInsertPostRequestDataValue]**](VectorsInsertPostRequestDataValue.md) |  | [optional] 

## Example

```python
from openapi_client.models.vectors_insert_post_request import VectorsInsertPostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of VectorsInsertPostRequest from a JSON string
vectors_insert_post_request_instance = VectorsInsertPostRequest.from_json(json)
# print the JSON string representation of the object
print(VectorsInsertPostRequest.to_json())

# convert the object into a dict
vectors_insert_post_request_dict = vectors_insert_post_request_instance.to_dict()
# create an instance of VectorsInsertPostRequest from a dict
vectors_insert_post_request_from_dict = VectorsInsertPostRequest.from_dict(vectors_insert_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


