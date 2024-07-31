# VectorsService

A list of all methods in the `VectorsService` service. Click on the method name to view detailed information about that method.

| Methods                                         | Description                                                                |
| :---------------------------------------------- | :------------------------------------------------------------------------- |
| [delete](#delete) | This operation deletes entities by their IDs or with a boolean expression. |
| [get](#get)       | This operation gets vectors by their IDs.                                  |
| [insert](#insert) | This operation inserts vectors into a specified collection.                |
| [query](#query)   | This operation queries vectors in a specified collection.                  |
| [search](#search) | This operation searches vectors in a specified collection.                 |
| [upsert](#upsert) | This operation upserts vectors into a specified collection.                |

## delete

This operation deletes entities by their IDs or with a boolean expression.

- HTTP Method: `POST`
- Endpoint: `/vectors/delete`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [VectorsDeleteRequest](../models/VectorsDeleteRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import VectorsDeleteRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = VectorsDeleteRequest(
    db_name="dbName",
    collection_name="collectionName",
    filter="filter",
    partition_name="partitionName"
)

result = sdk.vectors.delete(request_body=request_body)

print(result)
```

## get

This operation gets vectors by their IDs.

- HTTP Method: `POST`
- Endpoint: `/vectors/get`

**Parameters**

| Name         | Type                                                            | Required | Description       |
| :----------- | :-------------------------------------------------------------- | :------- | :---------------- |
| request_body | [VectorsGetRequest](../models/VectorsGetRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import VectorsGetRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = VectorsGetRequest(
    db_name="dbName",
    collection_name="collectionName",
    id_="veniam lab",
    partition_names=[
        "partitionNames"
    ],
    output_fields=[
        "outputFields"
    ]
)

result = sdk.vectors.get(request_body=request_body)

print(result)
```

## insert

This operation inserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/insert`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [VectorsInsertRequest](../models/VectorsInsertRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import VectorsInsertRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = VectorsInsertRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_name="partitionName",
    data={}
)

result = sdk.vectors.insert(request_body=request_body)

print(result)
```

## query

This operation queries vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/query`

**Parameters**

| Name         | Type                                                                | Required | Description       |
| :----------- | :------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [VectorsQueryRequest](../models/VectorsQueryRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import VectorsQueryRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = VectorsQueryRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_names=[
        "partitionNames"
    ],
    filter="filter",
    output_fields=[
        "outputFields"
    ]
)

result = sdk.vectors.query(request_body=request_body)

print(result)
```

## search

This operation searches vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/search`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [VectorsSearchRequest](../models/VectorsSearchRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import VectorsSearchRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = VectorsSearchRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_names=[
        "partitionNames"
    ],
    output_fields=[
        "outputFields"
    ],
    anss_field="anssField",
    limit=3,
    offset=2,
    filter="filter",
    grouping_field="groupingField",
    search_params={
        "metric_type": "metricType",
        "params": {
            "radius": 8,
            "range_filter": 4
        }
    }
)

result = sdk.vectors.search(request_body=request_body)

print(result)
```

## upsert

This operation upserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/upsert`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [VectorsUpsertRequest](../models/VectorsUpsertRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import VectorsUpsertRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = VectorsUpsertRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_name="partitionName",
    data={}
)

result = sdk.vectors.upsert(request_body=request_body)

print(result)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
