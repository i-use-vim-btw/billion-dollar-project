# Gandi Go SDK 1.0.0

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
- [Services](#services)
- [Models](#models)
- [License](#license)

# Setup & Configuration

## Supported Language Versions

This SDK is compatible with the following versions: `Go >= 1.19.0`

## Authentication

### Access Token Authentication

The gandi API uses an Access Token for authentication.

This token must be provided to authenticate your requests to the API.

#### Setting the Access Token

When you initialize the SDK, you can set the access token as follows:

```go
import (
    "github.com/swagger-api/swagger-petstore/pkg/gandi"
    "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  )

config := gandiconfig.NewConfig{}
config.SetBearerToken("YOUR-TOKEN")

sdk := gandi.NewGandi(config)
```

If you need to set or update the access token after initializing the SDK, you can use:

```go
import (
    "github.com/swagger-api/swagger-petstore/pkg/gandi"
    "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  )

config := gandiconfig.NewConfig{}

sdk := gandi.NewGandi(config)
sdk.SetBearerToken("YOUR-TOKEN")
```

## Services

The SDK provides various services to interact with the API.

<details> 
<summary>Below is a list of all available services with links to their detailed documentation:</summary>

| Name                                                                |
| :------------------------------------------------------------------ |
| [CollectionsService](documentation/services/collections_service.md) |
| [IndexesService](documentation/services/indexes_service.md)         |
| [VectorsService](documentation/services/vectors_service.md)         |

</details>

## Models

The SDK includes several models that represent the data structures used in API requests and responses. These models help in organizing and managing the data efficiently.

<details> 
<summary>Below is a list of all available models with links to their detailed documentation:</summary>

| Name                                                                                                      | Description |
| :-------------------------------------------------------------------------------------------------------- | :---------- |
| [Collection](documentation/models/collection.md)                                           |             |
| [CollectionDescribeRequest](documentation/models/collection_describe_request.md)           |             |
| [CollectionDropRequest](documentation/models/collections_drop_request.md)                   |             |
| [CollectionGetLoadStateRequest](documentation/models/collections_get_load_state_request.md) |             |
| [CollectionGetStatsRequest](documentation/models/collections_get_stats_request.md)          |             |
| [CollectionHasRequest](documentation/models/collections_has_request.md)                     |             |
| [CollectionListRequest](documentation/models/collections_list_request.md)                   |             |
| [CollectionLoadRequest](documentation/models/collections_load_request.md)                   |             |
| [CollectionReleaseRequest](documentation/models/collections_release_request.md)             |             |
| [CollectionRenameRequest](documentation/models/collections_rename_request.md)               |             |
| [IndexRequest](documentation/models/index_request.md)                                                     |             |
| [IndexDescribeRequest](documentation/models/indexes_describe_request.md)                   |             |
| [IndexDropRequest](documentation/models/indexes_drop_request.md)                           |             |
| [IndexListRequest](documentation/models/indexes_list_request.md)                           |             |
| [VectorsDeleteRequest](documentation/models/vectors_delete_request.md)                       |             |
| [VectorsGetRequest](documentation/models/vectors_get_request.md)                             |             |
| [VectorsInsertRequest](documentation/models/vectors_insert_request.md)                       |             |
| [VectorsQueryRequest](documentation/models/vectors_query_request.md)                         |             |
| [VectorsSearchRequest](documentation/models/vectors_search_request.md)                       |             |
| [VectorsUpsertRequest](documentation/models/vectors_upsert_request.md)                       |             |

</details>

## License

This SDK is licensed under the MIT License.

See the [LICENSE](LICENSE) file for more details.

<!-- This file was generated by liblab | https://liblab.com/ -->
