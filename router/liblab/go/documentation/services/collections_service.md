# CollectionsService

A list of all methods in the `CollectionsService` service. Click on the method name to view detailed information about that method.

| Methods                                                         | Description                                                                                          |
| :-------------------------------------------------------------- | :--------------------------------------------------------------------------------------------------- |
| [CreateCollectionsCreate](#createcollectionscreate)             | This operation creates a collection in a specified cluster.                                          |
| [CreateCollectionsDescribe](#createcollectionsdescribe)         | Returns the details of a collection.                                                                 |
| [CreateCollectionsDrop](#createcollectionsdrop)                 | This operation drops the current collection and all data within the collection.                      |
| [CreateCollectionsGetLoadState](#createcollectionsgetloadstate) | Returns the load state of a specific collection.                                                     |
| [CreateCollectionsGetStats](#createcollectionsgetstats)         | This operation gets the number of entities in a collection.                                          |
| [CreateCollectionsHas](#createcollectionshas)                   | Checks if a collection exists in the database.                                                       |
| [CreateCollectionsList](#createcollectionslist)                 | Returns a list of all collections in the specified database.                                         |
| [CreateCollectionsLoad](#createcollectionsload)                 | Loads a collection into memory.                                                                      |
| [CreateCollectionsRelease](#createcollectionsrelease)           | Releases a collection from memory.                                                                   |
| [CreateCollectionsRename](#createcollectionsrename)             | This operation renames an existing collection and optionally moves the collection to a new database. |

## CreateCollectionsCreate

This operation creates a collection in a specified cluster.

- HTTP Method: `POST`
- Endpoint: `/collections/create`

**Parameters**

| Name              | Type              | Required | Description                              |
| :---------------- | :---------------- | :------- | :--------------------------------------- |
| ctx               | Context           | ✅       | Default go language context              |
| collectionRequest | CollectionRequest | ✅       | Collection object that needs to be added |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CollectionRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsCreate(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsDescribe

Returns the details of a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/describe`

**Parameters**

| Name                             | Type                             | Required | Description                                  |
| :------------------------------- | :------------------------------- | :------- | :------------------------------------------- |
| ctx                              | Context                          | ✅       | Default go language context                  |
| createCollectionsDescribeRequest | CreateCollectionsDescribeRequest | ✅       | Collection object that needs to be described |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsDescribeRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsDescribe(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsDrop

This operation drops the current collection and all data within the collection.

- HTTP Method: `POST`
- Endpoint: `/collections/drop`

**Parameters**

| Name                         | Type                         | Required | Description                                |
| :--------------------------- | :--------------------------- | :------- | :----------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context                |
| createCollectionsDropRequest | CreateCollectionsDropRequest | ✅       | Collection object that needs to be deleted |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsDropRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsDrop(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsGetLoadState

Returns the load state of a specific collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_load_state`

**Parameters**

| Name                                 | Type                                 | Required | Description                                |
| :----------------------------------- | :----------------------------------- | :------- | :----------------------------------------- |
| ctx                                  | Context                              | ✅       | Default go language context                |
| createCollectionsGetLoadStateRequest | CreateCollectionsGetLoadStateRequest | ✅       | Collection object that needs to be checked |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsGetLoadStateRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsGetLoadState(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsGetStats

This operation gets the number of entities in a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_stats`

**Parameters**

| Name                             | Type                             | Required | Description                                |
| :------------------------------- | :------------------------------- | :------- | :----------------------------------------- |
| ctx                              | Context                          | ✅       | Default go language context                |
| createCollectionsGetStatsRequest | CreateCollectionsGetStatsRequest | ✅       | Collection object that needs to be checked |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsGetStatsRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsGetStats(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsHas

Checks if a collection exists in the database.

- HTTP Method: `POST`
- Endpoint: `/collections/has`

**Parameters**

| Name                        | Type                        | Required | Description                                |
| :-------------------------- | :-------------------------- | :------- | :----------------------------------------- |
| ctx                         | Context                     | ✅       | Default go language context                |
| createCollectionsHasRequest | CreateCollectionsHasRequest | ✅       | Collection object that needs to be checked |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsHasRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsHas(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsList

Returns a list of all collections in the specified database.

- HTTP Method: `POST`
- Endpoint: `/collections/list`

**Parameters**

| Name                         | Type                         | Required | Description                             |
| :--------------------------- | :--------------------------- | :------- | :-------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context             |
| createCollectionsListRequest | CreateCollectionsListRequest | ✅       | Database object that needs to be listed |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsListRequest{}
request.SetDbName("DbName")

response, err := client.Collections.CreateCollectionsList(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsLoad

Loads a collection into memory.

- HTTP Method: `POST`
- Endpoint: `/collections/load`

**Parameters**

| Name                         | Type                         | Required | Description                               |
| :--------------------------- | :--------------------------- | :------- | :---------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context               |
| createCollectionsLoadRequest | CreateCollectionsLoadRequest | ✅       | Collection object that needs to be loaded |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsLoadRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsLoad(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsRelease

Releases a collection from memory.

- HTTP Method: `POST`
- Endpoint: `/collections/release`

**Parameters**

| Name                            | Type                            | Required | Description                                 |
| :------------------------------ | :------------------------------ | :------- | :------------------------------------------ |
| ctx                             | Context                         | ✅       | Default go language context                 |
| createCollectionsReleaseRequest | CreateCollectionsReleaseRequest | ✅       | Collection object that needs to be released |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsReleaseRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CreateCollectionsRelease(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateCollectionsRename

This operation renames an existing collection and optionally moves the collection to a new database.

- HTTP Method: `POST`
- Endpoint: `/collections/rename`

**Parameters**

| Name                           | Type                           | Required | Description                                |
| :----------------------------- | :----------------------------- | :------- | :----------------------------------------- |
| ctx                            | Context                        | ✅       | Default go language context                |
| createCollectionsRenameRequest | CreateCollectionsRenameRequest | ✅       | Collection object that needs to be renamed |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/collections"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := collections.CreateCollectionsRenameRequest{}
request.SetCollectionName("CollectionName")
request.SetNewCollectionName("NewCollectionName")

response, err := client.Collections.CreateCollectionsRename(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
