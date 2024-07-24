# CollectionsRenamePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**new_collection_name** | **str** |  | 

## Example

```python
from openapi_client.models.collections_rename_post_request import CollectionsRenamePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionsRenamePostRequest from a JSON string
collections_rename_post_request_instance = CollectionsRenamePostRequest.from_json(json)
# print the JSON string representation of the object
print(CollectionsRenamePostRequest.to_json())

# convert the object into a dict
collections_rename_post_request_dict = collections_rename_post_request_instance.to_dict()
# create an instance of CollectionsRenamePostRequest from a dict
collections_rename_post_request_from_dict = CollectionsRenamePostRequest.from_dict(collections_rename_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


