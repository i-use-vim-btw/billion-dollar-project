# IndexesService

A list of all methods in the `IndexesService` service. Click on the method name to view detailed information about that method.

| Methods                                         | Description                                                                                          |
| :---------------------------------------------- | :--------------------------------------------------------------------------------------------------- |
| [CreateIndexesCreate](#createindexescreate)     | This creates a named index for a target field, which can either be a vector field or a scalar field. |
| [CreateIndexesDescribe](#createindexesdescribe) | Returns the details of an index.                                                                     |
| [CreateIndexesDrop](#createindexesdrop)         | This operation drops index from a specified collection.                                              |
| [CreateIndexesList](#createindexeslist)         | Returns a list of all indexes in the specified collection.                                           |

## CreateIndexesCreate

This creates a named index for a target field, which can either be a vector field or a scalar field.

- HTTP Method: `POST`
- Endpoint: `/indexes/create`

**Parameters**

| Name         | Type         | Required | Description                         |
| :----------- | :----------- | :------- | :---------------------------------- |
| ctx          | Context      | ✅       | Default go language context         |
| indexRequest | IndexRequest | ✅       | Index object that needs to be added |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/indexes"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := indexes.IndexRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Indexes.CreateIndexesCreate(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateIndexesDescribe

Returns the details of an index.

- HTTP Method: `POST`
- Endpoint: `/indexes/describe`

**Parameters**

| Name                         | Type                         | Required | Description                             |
| :--------------------------- | :--------------------------- | :------- | :-------------------------------------- |
| ctx                          | Context                      | ✅       | Default go language context             |
| createIndexesDescribeRequest | CreateIndexesDescribeRequest | ✅       | Index object that needs to be described |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/indexes"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := indexes.CreateIndexesDescribeRequest{}
request.SetCollectionName("CollectionName")
request.SetIndexName("IndexName")

response, err := client.Indexes.CreateIndexesDescribe(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateIndexesDrop

This operation drops index from a specified collection.

- HTTP Method: `POST`
- Endpoint: `/indexes/drop`

**Parameters**

| Name                     | Type                     | Required | Description                           |
| :----------------------- | :----------------------- | :------- | :------------------------------------ |
| ctx                      | Context                  | ✅       | Default go language context           |
| createIndexesDropRequest | CreateIndexesDropRequest | ✅       | Index object that needs to be deleted |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/indexes"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := indexes.CreateIndexesDropRequest{}
request.SetCollectionName("CollectionName")
request.SetIndexName("IndexName")

response, err := client.Indexes.CreateIndexesDrop(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateIndexesList

Returns a list of all indexes in the specified collection.

- HTTP Method: `POST`
- Endpoint: `/indexes/list`

**Parameters**

| Name                     | Type                     | Required | Description                               |
| :----------------------- | :----------------------- | :------- | :---------------------------------------- |
| ctx                      | Context                  | ✅       | Default go language context               |
| createIndexesListRequest | CreateIndexesListRequest | ✅       | Collection object that needs to be listed |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/indexes"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := indexes.CreateIndexesListRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Indexes.CreateIndexesList(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
