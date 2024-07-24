# CollectionsGetLoadStatePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**partition_names** | **str** |  | [optional] 

## Example

```python
from openapi_client.models.collections_get_load_state_post_request import CollectionsGetLoadStatePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionsGetLoadStatePostRequest from a JSON string
collections_get_load_state_post_request_instance = CollectionsGetLoadStatePostRequest.from_json(json)
# print the JSON string representation of the object
print(CollectionsGetLoadStatePostRequest.to_json())

# convert the object into a dict
collections_get_load_state_post_request_dict = collections_get_load_state_post_request_instance.to_dict()
# create an instance of CollectionsGetLoadStatePostRequest from a dict
collections_get_load_state_post_request_from_dict = CollectionsGetLoadStatePostRequest.from_dict(collections_get_load_state_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


