# CollectionsService

A list of all methods in the `CollectionsService` service. Click on the method name to view detailed information about that method.

| Methods                                                                 | Description                                                                                          |
| :---------------------------------------------------------------------- | :--------------------------------------------------------------------------------------------------- |
| [create_collections_create](#create_collections_create)                 | This operation creates a collection in a specified cluster.                                          |
| [create_collections_describe](#create_collections_describe)             | Returns the details of a collection.                                                                 |
| [create_collections_drop](#create_collections_drop)                     | This operation drops the current collection and all data within the collection.                      |
| [create_collections_get_load_state](#create_collections_get_load_state) | Returns the load state of a specific collection.                                                     |
| [create_collections_get_stats](#create_collections_get_stats)           | This operation gets the number of entities in a collection.                                          |
| [create_collections_has](#create_collections_has)                       | Checks if a collection exists in the database.                                                       |
| [create_collections_list](#create_collections_list)                     | Returns a list of all collections in the specified database.                                         |
| [create_collections_load](#create_collections_load)                     | Loads a collection into memory.                                                                      |
| [create_collections_release](#create_collections_release)               | Releases a collection from memory.                                                                   |
| [create_collections_rename](#create_collections_rename)                 | This operation renames an existing collection and optionally moves the collection to a new database. |

## create_collections_create

This operation creates a collection in a specified cluster.

- HTTP Method: `POST`
- Endpoint: `/collections/create`

**Parameters**

| Name         | Type                                                | Required | Description       |
| :----------- | :-------------------------------------------------- | :------- | :---------------- |
| request_body | [CollectionRequest](../models/CollectionRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionRequest(
    db_name="dbName",
    collection_name="collectionName",
    dimension=8,
    metric_type="metricType",
    id_type="idType",
    auto_id="autoID",
    primary_field_name="primaryFieldName",
    vector_field_name="vectorFieldName",
    schema={
        "auto_id": "autoID",
        "enable_dynamic_field": "enableDynamicField",
        "fields": [
            {
                "field_name": "fieldName",
                "data_type": "dataType",
                "element_data_type": "elementDataType",
                "is_primary": False,
                "is_partition_key": True,
                "element_type_params": {
                    "max_length": 8,
                    "dim": 2,
                    "max_capacity": 3
                }
            }
        ]
    },
    index_params=[
        {
            "metric_type": "metricType",
            "field_name": "fieldName",
            "index_name": "indexName",
            "params": {
                "index_type": "index_type",
                "m": 6,
                "ef_construction": 7,
                "nlist": 1
            }
        }
    ],
    params={
        "max_length": 2,
        "enable_dynamic_field": True,
        "shards_num": 1,
        "consistency_level": 3,
        "partitions_num": 8,
        "ttl_seconds": 4
    }
)

result = sdk.collections.create_collections_create(request_body=request_body)

print(result)
```

## create_collections_describe

Returns the details of a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/describe`

**Parameters**

| Name         | Type                                                                              | Required | Description       |
| :----------- | :-------------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateCollectionsDescribeRequest](../models/CreateCollectionsDescribeRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsDescribeRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsDescribeRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.create_collections_describe(request_body=request_body)

print(result)
```

## create_collections_drop

This operation drops the current collection and all data within the collection.

- HTTP Method: `POST`
- Endpoint: `/collections/drop`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CreateCollectionsDropRequest](../models/CreateCollectionsDropRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsDropRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsDropRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.create_collections_drop(request_body=request_body)

print(result)
```

## create_collections_get_load_state

Returns the load state of a specific collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_load_state`

**Parameters**

| Name         | Type                                                                                      | Required | Description       |
| :----------- | :---------------------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateCollectionsGetLoadStateRequest](../models/CreateCollectionsGetLoadStateRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsGetLoadStateRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsGetLoadStateRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_names="partitionNames"
)

result = sdk.collections.create_collections_get_load_state(request_body=request_body)

print(result)
```

## create_collections_get_stats

This operation gets the number of entities in a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_stats`

**Parameters**

| Name         | Type                                                                              | Required | Description       |
| :----------- | :-------------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateCollectionsGetStatsRequest](../models/CreateCollectionsGetStatsRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsGetStatsRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsGetStatsRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.create_collections_get_stats(request_body=request_body)

print(result)
```

## create_collections_has

Checks if a collection exists in the database.

- HTTP Method: `POST`
- Endpoint: `/collections/has`

**Parameters**

| Name         | Type                                                                    | Required | Description       |
| :----------- | :---------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateCollectionsHasRequest](../models/CreateCollectionsHasRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsHasRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsHasRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.create_collections_has(request_body=request_body)

print(result)
```

## create_collections_list

Returns a list of all collections in the specified database.

- HTTP Method: `POST`
- Endpoint: `/collections/list`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CreateCollectionsListRequest](../models/CreateCollectionsListRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsListRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsListRequest(
    db_name="dbName"
)

result = sdk.collections.create_collections_list(request_body=request_body)

print(result)
```

## create_collections_load

Loads a collection into memory.

- HTTP Method: `POST`
- Endpoint: `/collections/load`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CreateCollectionsLoadRequest](../models/CreateCollectionsLoadRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsLoadRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsLoadRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.create_collections_load(request_body=request_body)

print(result)
```

## create_collections_release

Releases a collection from memory.

- HTTP Method: `POST`
- Endpoint: `/collections/release`

**Parameters**

| Name         | Type                                                                            | Required | Description       |
| :----------- | :------------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CreateCollectionsReleaseRequest](../models/CreateCollectionsReleaseRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsReleaseRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsReleaseRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.create_collections_release(request_body=request_body)

print(result)
```

## create_collections_rename

This operation renames an existing collection and optionally moves the collection to a new database.

- HTTP Method: `POST`
- Endpoint: `/collections/rename`

**Parameters**

| Name         | Type                                                                          | Required | Description       |
| :----------- | :---------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateCollectionsRenameRequest](../models/CreateCollectionsRenameRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateCollectionsRenameRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateCollectionsRenameRequest(
    db_name="dbName",
    collection_name="collectionName",
    new_collection_name="newCollectionName"
)

result = sdk.collections.create_collections_rename(request_body=request_body)

print(result)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
