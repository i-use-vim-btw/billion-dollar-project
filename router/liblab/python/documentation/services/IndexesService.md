# IndexesService

A list of all methods in the `IndexesService` service. Click on the method name to view detailed information about that method.

| Methods                                             | Description                                                                                          |
| :-------------------------------------------------- | :--------------------------------------------------------------------------------------------------- |
| [create_indexes_create](#create_indexes_create)     | This creates a named index for a target field, which can either be a vector field or a scalar field. |
| [create_indexes_describe](#create_indexes_describe) | Returns the details of an index.                                                                     |
| [create_indexes_drop](#create_indexes_drop)         | This operation drops index from a specified collection.                                              |
| [create_indexes_list](#create_indexes_list)         | Returns a list of all indexes in the specified collection.                                           |

## create_indexes_create

This creates a named index for a target field, which can either be a vector field or a scalar field.

- HTTP Method: `POST`
- Endpoint: `/indexes/create`

**Parameters**

| Name         | Type                                      | Required | Description       |
| :----------- | :---------------------------------------- | :------- | :---------------- |
| request_body | [IndexRequest](../models/IndexRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import IndexRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = IndexRequest(
    db_name="dbName",
    collection_name="collectionName",
    index_params=[
        {
            "metric_type": "metricType",
            "field_name": "fieldName",
            "index_name": "indexName",
            "params": {
                "index_type": "index_type",
                "m": 7,
                "ef_construction": 8,
                "nlist": 9
            }
        }
    ]
)

result = sdk.indexes.create_indexes_create(request_body=request_body)

print(result)
```

## create_indexes_describe

Returns the details of an index.

- HTTP Method: `POST`
- Endpoint: `/indexes/describe`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CreateIndexesDescribeRequest](../models/CreateIndexesDescribeRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateIndexesDescribeRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateIndexesDescribeRequest(
    db_name="dbName",
    collection_name="collectionName",
    index_name="indexName"
)

result = sdk.indexes.create_indexes_describe(request_body=request_body)

print(result)
```

## create_indexes_drop

This operation drops index from a specified collection.

- HTTP Method: `POST`
- Endpoint: `/indexes/drop`

**Parameters**

| Name         | Type                                                              | Required | Description       |
| :----------- | :---------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateIndexesDropRequest](../models/CreateIndexesDropRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateIndexesDropRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateIndexesDropRequest(
    db_name="dbName",
    collection_name="collectionName",
    index_name="indexName"
)

result = sdk.indexes.create_indexes_drop(request_body=request_body)

print(result)
```

## create_indexes_list

Returns a list of all indexes in the specified collection.

- HTTP Method: `POST`
- Endpoint: `/indexes/list`

**Parameters**

| Name         | Type                                                              | Required | Description       |
| :----------- | :---------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateIndexesListRequest](../models/CreateIndexesListRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateIndexesListRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateIndexesListRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.indexes.create_indexes_list(request_body=request_body)

print(result)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
