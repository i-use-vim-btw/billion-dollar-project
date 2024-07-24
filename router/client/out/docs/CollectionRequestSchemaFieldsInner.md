# CollectionRequestSchemaFieldsInner


## Properties

Name | Type | Description | Notes
------------ | ------------- | ------------- | -------------
**field_name** | **str** |  | [optional] 
**data_type** | **str** |  | [optional] 
**element_data_type** | **str** |  | [optional] 
**is_primary** | **bool** |  | [optional] 
**is_partition_key** | **bool** |  | [optional] 
**element_type_params** | [**CollectionRequestSchemaFieldsInnerElementTypeParams**](CollectionRequestSchemaFieldsInnerElementTypeParams.md) |  | [optional] 

## Example

```python
from openapi_client.models.collection_request_schema_fields_inner import CollectionRequestSchemaFieldsInner

# TODO update the JSON string below
json = "{}"
# create an instance of CollectionRequestSchemaFieldsInner from a JSON string
collection_request_schema_fields_inner_instance = CollectionRequestSchemaFieldsInner.from_json(json)
# print the JSON string representation of the object
print(CollectionRequestSchemaFieldsInner.to_json())

# convert the object into a dict
collection_request_schema_fields_inner_dict = collection_request_schema_fields_inner_instance.to_dict()
# create an instance of CollectionRequestSchemaFieldsInner from a dict
collection_request_schema_fields_inner_from_dict = CollectionRequestSchemaFieldsInner.from_dict(collection_request_schema_fields_inner_dict)
```
[[Back to Model list]](../README.md#documentation-for-models) [[Back to API list]](../README.md#documentation-for-api-endpoints) [[Back to README]](../README.md)


