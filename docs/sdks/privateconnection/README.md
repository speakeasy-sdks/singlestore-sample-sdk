# PrivateConnection
(*PrivateConnection*)

### Available Operations

* [Create](#create) - Creates a new private connection
* [Delete](#delete) - Deletes a private connection
* [Get](#get) - Gets information about a private connection
* [Update](#update) - Updates a private connection

## Create

Creates a new private connection. Upon successful completion of the request, a private connection is scheduled for creation. To query the private connection status, use the endpoints for the workspace group and workspace (if provided).


### Example Usage

```go
package main

import(
	"context"
	"log"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PrivateConnection.Create(ctx, shared.PrivateConnectionCreate{
        AllowList: singlestoresamplesdk.String("my-allow-list"),
        ServiceName: singlestoresamplesdk.String("My private link"),
        Type: shared.PrivateConnectionCreateTypeInbound.ToPointer(),
        WorkspaceGroupID: "68af2f46-0000-1000-9000-3f6f5365d878",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreatePrivateConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                        | Type                                                                             | Required                                                                         | Description                                                                      |
| -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- | -------------------------------------------------------------------------------- |
| `ctx`                                                                            | [context.Context](https://pkg.go.dev/context#Context)                            | :heavy_check_mark:                                                               | The context to use for the request.                                              |
| `request`                                                                        | [shared.PrivateConnectionCreate](../../models/shared/privateconnectioncreate.md) | :heavy_check_mark:                                                               | The request object to use for the request.                                       |


### Response

**[*operations.CreatePrivateConnectionResponse](../../models/operations/createprivateconnectionresponse.md), error**


## Delete

Deletes a private connection for the specified connection ID. Upon successful completion, a private connection is scheduled for deletion.

### Example Usage

```go
package main

import(
	"context"
	"log"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PrivateConnection.Delete(ctx, operations.DeletePrivateConnectionRequest{
        ConnectionID: "8db863f6-ef9b-413a-8a70-cb816b33de6b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeletePrivateConnection != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.DeletePrivateConnectionRequest](../../models/operations/deleteprivateconnectionrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.DeletePrivateConnectionResponse](../../models/operations/deleteprivateconnectionresponse.md), error**


## Get

Returns private connection information for the specified connection ID, in JSON format. You must specify the connection ID in the API call.

### Example Usage

```go
package main

import(
	"context"
	"log"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PrivateConnection.Get(ctx, operations.GetPrivateConnectionRequest{
        ConnectionID: "b18d8d81-fd7b-4764-a31e-475cb1f36591",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PrivateConnections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetPrivateConnectionRequest](../../models/operations/getprivateconnectionrequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetPrivateConnectionResponse](../../models/operations/getprivateconnectionresponse.md), error**


## Update

Updates a private connection. You must specify the connection ID in the API call.

### Example Usage

```go
package main

import(
	"context"
	"log"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )

    ctx := context.Background()
    res, err := s.PrivateConnection.Update(ctx, operations.UpdatePrivateConnectionRequest{
        UpdatePrivateConnection: shared.UpdatePrivateConnection{
            AllowList: singlestoresamplesdk.String("my-allow-list"),
        },
        ConnectionID: "d0905bf4-aa77-4f20-8e77-54c352acfe54",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdatePrivateConnectionResponse != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdatePrivateConnectionRequest](../../models/operations/updateprivateconnectionrequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.UpdatePrivateConnectionResponse](../../models/operations/updateprivateconnectionresponse.md), error**

