# CollectionsService

A list of all methods in the `CollectionsService` service. Click on the method name to view detailed information about that method.

| Methods                                                         | Description                                                                                          |
| :-------------------------------------------------------------- | :--------------------------------------------------------------------------------------------------- |
| [CollectionCreate](#createcollectionscreate)             | This operation creates a collection in a specified cluster.                                          |
| [CollectionDescribe](#createcollectionsdescribe)         | Returns the details of a collection.                                                                 |
| [CollectionDrop](#createcollectionsdrop)                 | This operation drops the current collection and all data within the collection.                      |
| [CollectionGetLoadState](#createcollectionsgetloadstate) | Returns the load state of a specific collection.                                                     |
| [CollectionGetStats](#createcollectionsgetstats)         | This operation gets the number of entities in a collection.                                          |
| [CollectionHas](#createcollectionshas)                   | Checks if a collection exists in the database.                                                       |
| [CollectionList](#createcollectionslist)                 | Returns a list of all collections in the specified database.                                         |
| [CollectionLoad](#createcollectionsload)                 | Loads a collection into memory.                                                                      |
| [CollectionRelease](#createcollectionsrelease)           | Releases a collection from memory.                                                                   |
| [CollectionRename](#createcollectionsrename)             | This operation renames an existing collection and optionally moves the collection to a new database. |

## CollectionCreate

This operation creates a collection in a specified cluster.

- HTTP Method: `POST`
- Endpoint: `/collections/create`

**Parameters**

| Name              | Type              | Required | Description                              |
| :---------------- | :---------------- | :------- | :--------------------------------------- |
| ctx               | Context           | ✅       | Default go language context              |
| Collection | Collection | ✅       | Collection object that needs to be added |

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


request := collections.Collection{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionCreate(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionDescribe

Returns the details of a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/describe`

**Parameters**

| Name                             | Type                             | Required | Description                                  |
| :------------------------------- | :------------------------------- | :------- | :------------------------------------------- |
| ctx                              | Context                          | ✅       | Default go language context                  |
| createCollectionsDescribeRequest | CollectionDescribeRequest | ✅       | Collection object that needs to be described |

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


request := collections.CollectionDescribeRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionDescribe(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionDrop

This operation drops the current collection and all data within the collection.

- HTTP Method: `POST`
- Endpoint: `/collections/drop`

**Parameters**

| Name                         | Type                         | Required | Description                                |
| :--------------------------- | :--------------------------- | :------- | :----------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context                |
| createCollectionsDropRequest | CollectionDropRequest | ✅       | Collection object that needs to be deleted |

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


request := collections.CollectionDropRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionDrop(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionGetLoadState

Returns the load state of a specific collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_load_state`

**Parameters**

| Name                                 | Type                                 | Required | Description                                |
| :----------------------------------- | :----------------------------------- | :------- | :----------------------------------------- |
| ctx                                  | Context                              | ✅       | Default go language context                |
| createCollectionsGetLoadStateRequest | CollectionGetLoadStateRequest | ✅       | Collection object that needs to be checked |

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


request := collections.CollectionGetLoadStateRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionGetLoadState(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionGetStats

This operation gets the number of entities in a collection.

- HTTP Method: `POST`
- Endpoint: `/collections/get_stats`

**Parameters**

| Name                             | Type                             | Required | Description                                |
| :------------------------------- | :------------------------------- | :------- | :----------------------------------------- |
| ctx                              | Context                          | ✅       | Default go language context                |
| createCollectionsGetStatsRequest | CollectionGetStatsRequest | ✅       | Collection object that needs to be checked |

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


request := collections.CollectionGetStatsRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionGetStats(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionHas

Checks if a collection exists in the database.

- HTTP Method: `POST`
- Endpoint: `/collections/has`

**Parameters**

| Name                        | Type                        | Required | Description                                |
| :-------------------------- | :-------------------------- | :------- | :----------------------------------------- |
| ctx                         | Context                     | ✅       | Default go language context                |
| createCollectionsHasRequest | CollectionHasRequest | ✅       | Collection object that needs to be checked |

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


request := collections.CollectionHasRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionHas(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionList

Returns a list of all collections in the specified database.

- HTTP Method: `POST`
- Endpoint: `/collections/list`

**Parameters**

| Name                         | Type                         | Required | Description                             |
| :--------------------------- | :--------------------------- | :------- | :-------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context             |
| createCollectionsListRequest | CollectionListRequest | ✅       | Database object that needs to be listed |

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


request := collections.CollectionListRequest{}
request.SetDbName("DbName")

response, err := client.Collections.CollectionList(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionLoad

Loads a collection into memory.

- HTTP Method: `POST`
- Endpoint: `/collections/load`

**Parameters**

| Name                         | Type                         | Required | Description                               |
| :--------------------------- | :--------------------------- | :------- | :---------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context               |
| createCollectionsLoadRequest | CollectionLoadRequest | ✅       | Collection object that needs to be loaded |

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


request := collections.CollectionLoadRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionLoad(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionRelease

Releases a collection from memory.

- HTTP Method: `POST`
- Endpoint: `/collections/release`

**Parameters**

| Name                            | Type                            | Required | Description                                 |
| :------------------------------ | :------------------------------ | :------- | :------------------------------------------ |
| ctx                             | Context                         | ✅       | Default go language context                 |
| createCollectionsReleaseRequest | CollectionReleaseRequest | ✅       | Collection object that needs to be released |

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


request := collections.CollectionReleaseRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Collections.CollectionRelease(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CollectionRename

This operation renames an existing collection and optionally moves the collection to a new database.

- HTTP Method: `POST`
- Endpoint: `/collections/rename`

**Parameters**

| Name                           | Type                           | Required | Description                                |
| :----------------------------- | :----------------------------- | :------- | :----------------------------------------- |
| ctx                            | Context                        | ✅       | Default go language context                |
| createCollectionsRenameRequest | CollectionRenameRequest | ✅       | Collection object that needs to be renamed |

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


request := collections.CollectionRenameRequest{}
request.SetCollectionName("CollectionName")
request.SetNewCollectionName("NewCollectionName")

response, err := client.Collections.CollectionRename(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
