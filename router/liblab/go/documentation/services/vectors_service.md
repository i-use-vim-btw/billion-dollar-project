# VectorsService

A list of all methods in the `VectorsService` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description                                                                |
| :------------------------------------------ | :------------------------------------------------------------------------- |
| [VectorsDelete](#createvectorsdelete) | This operation deletes entities by their IDs or with a boolean expression. |
| [VectorsGet](#createvectorsget)       | This operation gets vectors by their IDs.                                  |
| [VectorsInsert](#createvectorsinsert) | This operation inserts vectors into a specified collection.                |
| [VectorsQuery](#createvectorsquery)   | This operation queries vectors in a specified collection.                  |
| [VectorsSearch](#createvectorssearch) | This operation searches vectors in a specified collection.                 |
| [VectorsUpsert](#createvectorsupsert) | This operation upserts vectors into a specified collection.                |

## VectorsDelete

This operation deletes entities by their IDs or with a boolean expression.

- HTTP Method: `POST`
- Endpoint: `/vectors/delete`

**Parameters**

| Name                       | Type                       | Required | Description                             |
| :------------------------- | :------------------------- | :------- | :-------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context             |
| createVectorsDeleteRequest | VectorsDeleteRequest | ✅       | Vectors object that needs to be deleted |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/vectors"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := vectors.VectorsDeleteRequest{}
request.SetCollectionName("CollectionName")
request.SetFilter("Filter")

response, err := client.Vectors.VectorsDelete(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## VectorsGet

This operation gets vectors by their IDs.

- HTTP Method: `POST`
- Endpoint: `/vectors/get`

**Parameters**

| Name                    | Type                    | Required | Description                               |
| :---------------------- | :---------------------- | :------- | :---------------------------------------- |
| ctx                     | Context                 | ✅       | Default go language context               |
| createVectorsGetRequest | VectorsGetRequest | ✅       | Vectors object that needs to be retrieved |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/vectors"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := vectors.VectorsGetRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.VectorsGet(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## VectorsInsert

This operation inserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/insert`

**Parameters**

| Name                       | Type                       | Required | Description                              |
| :------------------------- | :------------------------- | :------- | :--------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context              |
| createVectorsInsertRequest | VectorsInsertRequest | ✅       | Vectors object that needs to be inserted |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/vectors"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := vectors.VectorsInsertRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.VectorsInsert(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## VectorsQuery

This operation queries vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/query`

**Parameters**

| Name                      | Type                      | Required | Description                             |
| :------------------------ | :------------------------ | :------- | :-------------------------------------- |
| ctx                       | Context                   | ✅       | Default go language context             |
| createVectorsQueryRequest | VectorsQueryRequest | ✅       | Vectors object that needs to be queried |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/vectors"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := vectors.VectorsQueryRequest{}
request.SetCollectionName("CollectionName")
request.SetFilter("Filter")

response, err := client.Vectors.VectorsQuery(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## VectorsSearch

This operation searches vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/search`

**Parameters**

| Name                       | Type                       | Required | Description                              |
| :------------------------- | :------------------------- | :------- | :--------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context              |
| createVectorsSearchRequest | VectorsSearchRequest | ✅       | Vectors object that needs to be searched |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/vectors"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := vectors.VectorsSearchRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.VectorsSearch(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## VectorsUpsert

This operation upserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/upsert`

**Parameters**

| Name                       | Type                       | Required | Description                              |
| :------------------------- | :------------------------- | :------- | :--------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context              |
| createVectorsUpsertRequest | VectorsUpsertRequest | ✅       | Vectors object that needs to be upserted |

**Return Type**

`any`

**Example Usage Code Snippet**

```go
import (
  "fmt"
  "encoding/json"
  "github.com/swagger-api/swagger-petstore/pkg/gandiconfig"
  "github.com/swagger-api/swagger-petstore/pkg/gandi"
  "github.com/swagger-api/swagger-petstore/pkg/vectors"
)

config := gandiconfig.NewConfig()
client := gandi.NewGandi(config)


request := vectors.VectorsUpsertRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.VectorsUpsert(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
