# CollectionsDescribePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 

## Example

```python
from openapi_client.models.collections_describe_post_request import CollectionsDescribePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionsDescribePostRequest from a JSON string
collections_describe_post_request_instance = CollectionsDescribePostRequest.from_json(json)
# print the JSON string representation of the object
print(CollectionsDescribePostRequest.to_json())

# convert the object into a dict
collections_describe_post_request_dict = collections_describe_post_request_instance.to_dict()
# create an instance of CollectionsDescribePostRequest from a dict
collections_describe_post_request_from_dict = CollectionsDescribePostRequest.from_dict(collections_describe_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


