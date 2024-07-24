# CollectionRequestSchema


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**auto_id** | **str** |  | [optional] 
**enable_dynamic_field** | **str** |  | [optional] 
**fields** | [**List[CollectionRequestSchemaFieldsInner]**](CollectionRequestSchemaFieldsInner.md) |  | [optional] 

## Example

```python
from openapi_client.models.collection_request_schema import CollectionRequestSchema

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionRequestSchema from a JSON string
collection_request_schema_instance = CollectionRequestSchema.from_json(json)
# print the JSON string representation of the object
print(CollectionRequestSchema.to_json())

# convert the object into a dict
collection_request_schema_dict = collection_request_schema_instance.to_dict()
# create an instance of CollectionRequestSchema from a dict
collection_request_schema_from_dict = CollectionRequestSchema.from_dict(collection_request_schema_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


