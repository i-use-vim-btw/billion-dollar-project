# VectorsService

A list of all methods in the `VectorsService` service. Click on the method name to view detailed information about that method.

| Methods                                         | Description                                                                |
| :---------------------------------------------- | :------------------------------------------------------------------------- |
| [create_vectors_delete](#create_vectors_delete) | This operation deletes entities by their IDs or with a boolean expression. |
| [create_vectors_get](#create_vectors_get)       | This operation gets vectors by their IDs.                                  |
| [create_vectors_insert](#create_vectors_insert) | This operation inserts vectors into a specified collection.                |
| [create_vectors_query](#create_vectors_query)   | This operation queries vectors in a specified collection.                  |
| [create_vectors_search](#create_vectors_search) | This operation searches vectors in a specified collection.                 |
| [create_vectors_upsert](#create_vectors_upsert) | This operation upserts vectors into a specified collection.                |

## create_vectors_delete

This operation deletes entities by their IDs or with a boolean expression.

- HTTP Method: `POST`
- Endpoint: `/vectors/delete`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateVectorsDeleteRequest](../models/CreateVectorsDeleteRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateVectorsDeleteRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateVectorsDeleteRequest(
    db_name="dbName",
    collection_name="collectionName",
    filter="filter",
    partition_name="partitionName"
)

result = sdk.vectors.create_vectors_delete(request_body=request_body)

print(result)
```

## create_vectors_get

This operation gets vectors by their IDs.

- HTTP Method: `POST`
- Endpoint: `/vectors/get`

**Parameters**

| Name         | Type                                                            | Required | Description       |
| :----------- | :-------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateVectorsGetRequest](../models/CreateVectorsGetRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateVectorsGetRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateVectorsGetRequest(
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

result = sdk.vectors.create_vectors_get(request_body=request_body)

print(result)
```

## create_vectors_insert

This operation inserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/insert`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateVectorsInsertRequest](../models/CreateVectorsInsertRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateVectorsInsertRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateVectorsInsertRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_name="partitionName",
    data={}
)

result = sdk.vectors.create_vectors_insert(request_body=request_body)

print(result)
```

## create_vectors_query

This operation queries vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/query`

**Parameters**

| Name         | Type                                                                | Required | Description       |
| :----------- | :------------------------------------------------------------------ | :------- | :---------------- |
| request_body | [CreateVectorsQueryRequest](../models/CreateVectorsQueryRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateVectorsQueryRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateVectorsQueryRequest(
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

result = sdk.vectors.create_vectors_query(request_body=request_body)

print(result)
```

## create_vectors_search

This operation searches vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/search`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateVectorsSearchRequest](../models/CreateVectorsSearchRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateVectorsSearchRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateVectorsSearchRequest(
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

result = sdk.vectors.create_vectors_search(request_body=request_body)

print(result)
```

## create_vectors_upsert

This operation upserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/upsert`

**Parameters**

| Name         | Type                                                                  | Required | Description       |
| :----------- | :-------------------------------------------------------------------- | :------- | :---------------- |
| request_body | [CreateVectorsUpsertRequest](../models/CreateVectorsUpsertRequest.md) | ✅       | The request body. |

**Example Usage Code Snippet**

```python
from gandi import Gandi, Environment
from gandi.models import CreateVectorsUpsertRequest

sdk = Gandi(
    access_token="YOUR_ACCESS_TOKEN",
    base_url=Environment.DEFAULT.value
)

request_body = CreateVectorsUpsertRequest(
    db_name="dbName",
    collection_name="collectionName",
    partition_name="partitionName",
    data={}
)

result = sdk.vectors.create_vectors_upsert(request_body=request_body)

print(result)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
