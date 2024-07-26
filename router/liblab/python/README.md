# Gandi Python SDK 1.0.0

Welcome to the Gandi SDK documentation. This guide will help you get started with integrating and using the Gandi SDK in your project.

## Versions

- API version: `1.0.0`
- SDK version: `1.0.0`

## About the API

API for Gandi Cloud vector database service

## Table of Contents

- [Setup & Configuration](#setup--configuration)
  - [Supported Language Versions](#supported-language-versions)
  - [Installation](#installation)
- [Authentication](#authentication)
  - [Access Token Authentication](#access-token-authentication)
- [Using Union Types](#using-union-types)
- [Sample Usage](#sample-usage)
- [Services](#services)
- [Models](#models)
- [License](#license)

## Setup & Configuration

### Supported Language Versions

This SDK is compatible with the following versions: `Python >= 3.7`

### Installation

To get started with the SDK, we recommend installing using `pip`:

```bash
pip install gandi
```

## Authentication

### Access Token Authentication

The Gandi API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```py
Gandi(
    access_token="YOUR_ACCESS_TOKEN"
)
```

If you need to set or update the access token after initializing the SDK, you can use:

```py
sdk.set_access_token("YOUR_ACCESS_TOKEN")
```

## Using Union Types

Union types allow you to specify that a variable can have more than one type. This is particularly useful when a function can accept multiple types of inputs. The Union type hint is used for this purpose.

### Example Function with Union Types

You can call service method with an instance of `TypeA`, `TypeB`, or a dictionary that can be converted to an instance of either type.

```python
# Model Definition
ParamType = Union[TypeA, TypeB]

# Service Method
def service_method(param: ParamType):
...

## Usage
type_a = TypeA(key="value")
type_b = TypeB(key="value")

sdk.service.service_method(type_a)
sdk.service.service_method(type_b)
sdk.service.service_method({"key": "value"})
```

You cannot create an instance of a `Union` type itself. Instead, pass an instance of one of the types in the `Union`, or a dictionary that can be converted to one of those types.

# Sample Usage

Below is a comprehensive example demonstrating how to authenticate and call a simple endpoint:

```py
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

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                               |
| :----------------------------------------------------------------- |
| [CollectionsService](documentation/services/CollectionsService.md) |
| [IndexesService](documentation/services/IndexesService.md)         |
| [VectorsService](documentation/services/VectorsService.md)         |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details> 
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                                 | Description |
| :--------------------------------------------------------------------------------------------------- | :---------- |
| [CollectionRequest](documentation/models/CollectionRequest.md)                                       |             |
| [CreateCollectionsDescribeRequest](documentation/models/CreateCollectionsDescribeRequest.md)         |             |
| [CreateCollectionsDropRequest](documentation/models/CreateCollectionsDropRequest.md)                 |             |
| [CreateCollectionsGetLoadStateRequest](documentation/models/CreateCollectionsGetLoadStateRequest.md) |             |
| [CreateCollectionsGetStatsRequest](documentation/models/CreateCollectionsGetStatsRequest.md)         |             |
| [CreateCollectionsHasRequest](documentation/models/CreateCollectionsHasRequest.md)                   |             |
| [CreateCollectionsListRequest](documentation/models/CreateCollectionsListRequest.md)                 |             |
| [CreateCollectionsLoadRequest](documentation/models/CreateCollectionsLoadRequest.md)                 |             |
| [CreateCollectionsReleaseRequest](documentation/models/CreateCollectionsReleaseRequest.md)           |             |
| [CreateCollectionsRenameRequest](documentation/models/CreateCollectionsRenameRequest.md)             |             |
| [IndexRequest](documentation/models/IndexRequest.md)                                                 |             |
| [CreateIndexesDescribeRequest](documentation/models/CreateIndexesDescribeRequest.md)                 |             |
| [CreateIndexesDropRequest](documentation/models/CreateIndexesDropRequest.md)                         |             |
| [CreateIndexesListRequest](documentation/models/CreateIndexesListRequest.md)                         |             |
| [CreateVectorsDeleteRequest](documentation/models/CreateVectorsDeleteRequest.md)                     |             |
| [CreateVectorsGetRequest](documentation/models/CreateVectorsGetRequest.md)                           |             |
| [CreateVectorsInsertRequest](documentation/models/CreateVectorsInsertRequest.md)                     |             |
| [CreateVectorsQueryRequest](documentation/models/CreateVectorsQueryRequest.md)                       |             |
| [CreateVectorsSearchRequest](documentation/models/CreateVectorsSearchRequest.md)                     |             |
| [CreateVectorsUpsertRequest](documentation/models/CreateVectorsUpsertRequest.md)                     |             |
| [Schema](documentation/models/Schema.md)                                                             |             |
| [CollectionRequestIndexParams](documentation/models/CollectionRequestIndexParams.md)                 |             |
| [CollectionRequestParams](documentation/models/CollectionRequestParams.md)                           |             |
| [Fields](documentation/models/Fields.md)                                                             |             |
| [ElementTypeParams](documentation/models/ElementTypeParams.md)                                       |             |
| [IndexParamsParams_1](documentation/models/IndexParamsParams1.md)                                    |             |
| [IndexRequestIndexParams](documentation/models/IndexRequestIndexParams.md)                           |             |
| [IndexParamsParams_2](documentation/models/IndexParamsParams2.md)                                    |             |
| [CreateVectorsInsertRequestData](documentation/models/CreateVectorsInsertRequestData.md)             |             |
| [SearchParams](documentation/models/SearchParams.md)                                                 |             |
| [SearchParamsParams](documentation/models/SearchParamsParams.md)                                     |             |
| [CreateVectorsUpsertRequestData](documentation/models/CreateVectorsUpsertRequestData.md)             |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.

<!-- This file was generated by liblab | https://liblab.com/ -->
