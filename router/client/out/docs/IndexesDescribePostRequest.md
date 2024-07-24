# IndexesDescribePostRequest


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**db_name** | **str** |  | [optional] 
**collection_name** | **str** |  | 
**index_name** | **str** |  | 

## Example

```python
from openapi_client.models.indexes_describe_post_request import IndexesDescribePostRequest

# TODO update the JSON string below
json = "{}"
# create an instance of IndexesDescribePostRequest from a JSON string
indexes_describe_post_request_instance = IndexesDescribePostRequest.from_json(json)
# print the JSON string representation of the object
print(IndexesDescribePostRequest.to_json())

# convert the object into a dict
indexes_describe_post_request_dict = indexes_describe_post_request_instance.to_dict()
# create an instance of IndexesDescribePostRequest from a dict
indexes_describe_post_request_from_dict = IndexesDescribePostRequest.from_dict(indexes_describe_post_request_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


