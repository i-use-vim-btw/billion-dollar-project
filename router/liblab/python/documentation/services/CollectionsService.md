# CollectionsService

A list of all methods in the `CollectionsService` service. Click on the method name to view detailed information about that method.

| Methods                                                                 | Description                                                                                          |
| :---------------------------------------------------------------------- | :--------------------------------------------------------------------------------------------------- |
| [create](#create)                 | This operation creates a collection in a specified cluster.                                          |
| [describe](#describe)             | Returns the details of a collection.                                                                 |
| [drop](#drop)                     | This operation drops the current collection and all data within the collection.                      |
| [get_load_state](#get_load_state) | Returns the load state of a specific collection.                                                     |
| [get_stats](#get_stats)           | This operation gets the number of entities in a collection.                                          |
| [has](#has)                       | Checks if a collection exists in the database.                                                       |
| [list](#list)                     | Returns a list of all collections in the specified database.                                         |
| [load](#load)                     | Loads a collection into memory.                                                                      |
| [release](#release)               | Releases a collection from memory.                                                                   |
| [rename](#rename)                 | This operation renames an existing collection and optionally moves the collection to a new database. |

## create

This operation creates a collection in a specified cluster.

- HTTP Method: `POST`
- Endpoint: `/collections/create`

**Parameters**

| Name         | Type                                                | Required | Description       |
| :----------- | :-------------------------------------------------- | :------- | :---------------- |
| request_body | [Collection](../models/Collection.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import Collection

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = Collection(
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

result = sdk.collections.create(request_body=request_body)

print(result)
```

## describe

Returns the details of a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/describe`

**Parameters**

| Name         | Type                                                                              | Required | Description       |
| :----------- | :-------------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CollectionDescribeRequest](../models/CollectionDescribeRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionDescribeRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionDescribeRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.describe(request_body=request_body)

print(result)
```

## drop

This operation drops the current collection and all data within the collection.

- HTTP Method: `POST`
- Endpoint: `/collections/drop`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CollectionDropRequest](../models/CollectionDropRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionDropRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionDropRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.drop(request_body=request_body)

print(result)
```

## get_load_state

Returns the load state of a specific collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_load_state`

**Parameters**

| Name         | Type                                                                                      | Required | Description       |
| :----------- | :---------------------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CollectionGetLoadStateRequest](../models/CollectionGetLoadStateRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionGetLoadStateRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionGetLoadStateRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_names="partitionNames"
)

result = sdk.collections.get_load_state(request_body=request_body)

print(result)
```

## get_stats

This operation gets the number of entities in a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_stats`

**Parameters**

| Name         | Type                                                                              | Required | Description       |
| :----------- | :-------------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CollectionGetStatsRequest](../models/CollectionGetStatsRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionGetStatsRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionGetStatsRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.get_stats(request_body=request_body)

print(result)
```

## has

Checks if a collection exists in the database.

- HTTP Method: `POST`
- Endpoint: `/collections/has`

**Parameters**

| Name         | Type                                                                    | Required | Description       |
| :----------- | :---------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CollectionHasRequest](../models/CollectionHasRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionHasRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionHasRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.has(request_body=request_body)

print(result)
```

## list

Returns a list of all collections in the specified database.

- HTTP Method: `POST`
- Endpoint: `/collections/list`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CollectionListRequest](../models/CollectionListRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionListRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionListRequest(
    db_name="dbName"
)

result = sdk.collections.list(request_body=request_body)

print(result)
```

## load

Loads a collection into memory.

- HTTP Method: `POST`
- Endpoint: `/collections/load`

**Parameters**

| Name         | Type                                                                      | Required | Description       |
| :----------- | :------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CollectionLoadRequest](../models/CollectionLoadRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionLoadRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionLoadRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.load(request_body=request_body)

print(result)
```

## release

Releases a collection from memory.

- HTTP Method: `POST`
- Endpoint: `/collections/release`

**Parameters**

| Name         | Type                                                                            | Required | Description       |
| :----------- | :------------------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CollectionReleaseRequest](../models/CollectionReleaseRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionReleaseRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionReleaseRequest(
    db_name="dbName",
    collection_name="collectionName"
)

result = sdk.collections.release(request_body=request_body)

print(result)
```

## rename

This operation renames an existing collection and optionally moves the collection to a new database.

- HTTP Method: `POST`
- Endpoint: `/collections/rename`

**Parameters**

| Name         | Type                                                                          | Required | Description       |
| :----------- | :---------------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CollectionRenameRequest](../models/CollectionRenameRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CollectionRenameRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CollectionRenameRequest(
    db_name="dbName",
    collection_name="collectionName",
    new_collection_name="newCollectionName"
)

result = sdk.collections.rename(request_body=request_body)

print(result)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
