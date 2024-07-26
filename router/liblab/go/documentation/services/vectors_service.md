# VectorsService

A list of all methods in the `VectorsService` service. Click on the method name to view detailed information about that method.

| Methods                                     | Description                                                                |
| :------------------------------------------ | :------------------------------------------------------------------------- |
| [CreateVectorsDelete](#createvectorsdelete) | This operation deletes entities by their IDs or with a boolean expression. |
| [CreateVectorsGet](#createvectorsget)       | This operation gets vectors by their IDs.                                  |
| [CreateVectorsInsert](#createvectorsinsert) | This operation inserts vectors into a specified collection.                |
| [CreateVectorsQuery](#createvectorsquery)   | This operation queries vectors in a specified collection.                  |
| [CreateVectorsSearch](#createvectorssearch) | This operation searches vectors in a specified collection.                 |
| [CreateVectorsUpsert](#createvectorsupsert) | This operation upserts vectors into a specified collection.                |

## CreateVectorsDelete

This operation deletes entities by their IDs or with a boolean expression.

- HTTP Method: `POST`
- Endpoint: `/vectors/delete`

**Parameters**

| Name                       | Type                       | Required | Description                             |
| :------------------------- | :------------------------- | :------- | :-------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context             |
| createVectorsDeleteRequest | CreateVectorsDeleteRequest | ✅       | Vectors object that needs to be deleted |

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


request := vectors.CreateVectorsDeleteRequest{}
request.SetCollectionName("CollectionName")
request.SetFilter("Filter")

response, err := client.Vectors.CreateVectorsDelete(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateVectorsGet

This operation gets vectors by their IDs.

- HTTP Method: `POST`
- Endpoint: `/vectors/get`

**Parameters**

| Name                    | Type                    | Required | Description                               |
| :---------------------- | :---------------------- | :------- | :---------------------------------------- |
| ctx                     | Context                 | ✅       | Default go language context               |
| createVectorsGetRequest | CreateVectorsGetRequest | ✅       | Vectors object that needs to be retrieved |

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


request := vectors.CreateVectorsGetRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.CreateVectorsGet(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateVectorsInsert

This operation inserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/insert`

**Parameters**

| Name                       | Type                       | Required | Description                              |
| :------------------------- | :------------------------- | :------- | :--------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context              |
| createVectorsInsertRequest | CreateVectorsInsertRequest | ✅       | Vectors object that needs to be inserted |

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


request := vectors.CreateVectorsInsertRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.CreateVectorsInsert(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateVectorsQuery

This operation queries vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/query`

**Parameters**

| Name                      | Type                      | Required | Description                             |
| :------------------------ | :------------------------ | :------- | :-------------------------------------- |
| ctx                       | Context                   | ✅       | Default go language context             |
| createVectorsQueryRequest | CreateVectorsQueryRequest | ✅       | Vectors object that needs to be queried |

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


request := vectors.CreateVectorsQueryRequest{}
request.SetCollectionName("CollectionName")
request.SetFilter("Filter")

response, err := client.Vectors.CreateVectorsQuery(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateVectorsSearch

This operation searches vectors in a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/search`

**Parameters**

| Name                       | Type                       | Required | Description                              |
| :------------------------- | :------------------------- | :------- | :--------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context              |
| createVectorsSearchRequest | CreateVectorsSearchRequest | ✅       | Vectors object that needs to be searched |

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


request := vectors.CreateVectorsSearchRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.CreateVectorsSearch(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

## CreateVectorsUpsert

This operation upserts vectors into a specified collection.

- HTTP Method: `POST`
- Endpoint: `/vectors/upsert`

**Parameters**

| Name                       | Type                       | Required | Description                              |
| :------------------------- | :------------------------- | :------- | :--------------------------------------- |
| ctx                        | Context                    | ✅       | Default go language context              |
| createVectorsUpsertRequest | CreateVectorsUpsertRequest | ✅       | Vectors object that needs to be upserted |

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


request := vectors.CreateVectorsUpsertRequest{}
request.SetCollectionName("CollectionName")

response, err := client.Vectors.CreateVectorsUpsert(context.Background(), request)
if err != nil {
  panic(err)
}

fmt.Print(response)
```

<!-- This file was generated by liblab | https://liblab.com/ -->
