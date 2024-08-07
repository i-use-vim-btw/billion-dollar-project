# Collection

Collection object for creating a collection.

**Properties**

| Name            | Type              | Required | Description  |
| :-------------- | :---------------- | :------- | :----------  |
| collection_name | str               | ✅       | The name of the collection to create.|
| db_name         | str               | ❌       | The name of the database.|
| dimension       | int               | ❌       | The number of dimensions a vector value should have. This is required if dtype of this field is set to DataType.FLOAT_VECTOR.|
| schema          | Schema            | ❌       | The schema is responsible for organizing data in the target collection. A valid schema should have multiple fields, which must include a primary key, a vector field, and several scalar fields.|
| index_params    | List[IndexParams] | ❌       | The parameters that apply to the index-building process.|
| params          | CollectionParams  | ❌       | Extra parameters for the collection.|

Parameters below are designed for the quick-setup of a collection and will be ignored if schema is defined.
| Name | Type | Required | Description |
| :----------------- | :--------------------------------- | :------- | :---------- |
| metric_type | str | ❌ | The metric type applied to this operation. Possible values are L2, IP, and COSINE. The value defaults to COSINE.|
| id_type | str | ❌ | The data type of the primary field. |
| auto_id | str | ❌ | Whether the primary field automatically increments. |
| primary_field_name | str | ❌ | The name of the primary field. |
| vector_field_name | str | ❌ | The name of the vector field. |

## Schema

**Properties**

| Name                 | Type        | Required | Description                                                                                                                                                                                                                                                                                   |
| :------------------- | :---------- | :------- | :-------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| auto_id              | str         | ❌       | Whether allows the primary field to automatically increment. Setting this to True makes the primary field automatically increment. In this case, the primary field should not be included in the data to insert to avoid errors. Set this parameter in the field with is_primary set to True. |
| enable_dynamic_field | str         | ❌       | Whether allows to use the reserved $meta field to hold non-schema-defined fields in key-value pairs.                                                                                                                                                                                          |
| fields               | List[Field] | ❌       | A list of field objects.                                                                                                                                                                                                                                                                      |

## Field

**Properties**

| Name                | Type              | Required | Description                                                                                                                     |
| :------------------ | :---------------- | :------- | :------------------------------------------------------------------------------------------------------------------------------ |
| field_name          | str               | ❌       | The name of the field to create in the target collection.|
| data_type           | str               | ❌       | The data type of the field values.|
| element_data_type   | str               | ❌       | The data type of the elements in an array field.|
| is_primary          | bool              | ❌       | Whether the current field is the primary field. Setting this to True makes the current field the primary field.|
| is_partition_key    | bool              | ❌       | Whether the current field serves as the partition key. Setting this to True makes the current field serve as the partition key.|
| element_type_params | ElementTypeParams | ❌       | Extra field parameters.|

## ElementTypeParams

**Properties**

| Name         | Type | Required | Description                                                                                                             |
| :----------- | :--- | :------- | :---------------------------------------------------------------------------------------------------------------------- |
| max_length   | int  | ❌       | An optional parameter for VarChar values that determines the maximum length of the value in the current field.|
| dim          | int  | ❌       | An optional parameter for FloatVector or BinaryVector fields that determines the vector dimension.|
| max_capacity | int  | ❌       | An optional parameter for Array field values that determines the maximum number of elements in the current array field. |

## IndexParams

**Properties**

| Name        | Type        | Required | Description                                                                       |
| :---------- | :---------- | :------- | :-------------------------------------------------------------------------------- |
| metric_type | str         | ❌       | The similarity metric type used to build the index. The value defaults to COSINE. |
| field_name  | str         | ❌       | The name of the target field on which an index is to be created.                  |
| index_name  | str         | ❌       | The name of the index to create, the value defaults to the target field name.     |
| params      | IndexConfig | ❌       | The index type and related settings.                                              |

## IndexConfig

**Properties**

| Name            | Type | Required | Description                                                                     |
| :-------------- | :--- | :------- | :------------------------------------------------------------------------------ |
| index_type      | str  | ❌       | The type of the index to create.                                                |
| m               | int  | ❌       | The maximum degree of the node and applies only when index_type is set to HNSW. |
| ef_construction | int  | ❌       | The search scope. This applies only when index_type is set to HNSW.             |
| nlist           | int  | ❌       | The number of cluster units. This applies to IVF-related index types.           |

## CollectionParams

**Properties**

| Name                 | Type | Required | Description                                                                                                                                                                |
| :------------------- | :--- | :------- | :------------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| max_length           | int  | ❌       | The maximum number of characters in a VarChar field. This parameter is mandatory when the current field type is VarChar.|
| enable_dynamic_field | bool | ❌       | Whether to enable the reserved dynamic field. If set to true, non-schema-defined fields are saved in the reserved dynamic field as key-value pairs.|
| shards_num           | int  | ❌       | The number of shards to create along with the current collection.|
| consistency_level    | int  | ❌       | The consistency level of the collection. Possible values are STRONG, BOUNDED, SESSION, and EVENTUALLY.|
| partitions_num       | int  | ❌       | The number of partitions to create along with the current collection. This parameter is mandatory if one field of the collection has been designated as the partition key.|
| ttl_seconds          | int  | ❌       | The time-to-live (TTL) period of the collection. If set, the collection is to be dropped once the period ends.|

<!-- This file was generated by liblab | https://liblab.com/ -->
