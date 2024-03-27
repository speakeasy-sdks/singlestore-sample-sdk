# Workspaces
(*Workspaces*)

## Overview

Operations related to workspaces

### Available Operations

* [Create](#create) - Creates a new workspace
* [CreateResume](#createresume) - Resumes a workspace
* [CreateStorage](#createstorage) - Sets up Storage DR for the workspace group of the provided workspace. Backup region and selected databases to be replicated are provided as part of the request.
* [CreateSuspend](#createsuspend) - Suspends a workspace
* [Delete](#delete) - Terminates a workspace
* [Get](#get) - Gets information about a workspace
* [GetOutbound](#getoutbound) - Gets the outbound allow list for a workspace
* [GetPrivateConnection](#getprivateconnection) - Gets private connection information for a workspace
* [GetRecoveryBackup](#getrecoverybackup) - Gets information about which regions you can setup as a disaster recovery backup
* [GetStorageStatus](#getstoragestatus) - Gets information about the storage DR status of the group in which the provided workspace belongs to
* [List](#list) - Lists the workspaces the user can access
* [Update](#update) - Updates information about a workspace
* [UpdateFailback](#updatefailback) - Starts failback to the primary region
* [UpdateFailover](#updatefailover) - Starts failover to the secondary region
* [UpdateStartFailoverTestMode](#updatestartfailovertestmode) - Starts Failover test mode
* [UpdateStopFailoverTestMode](#updatestopfailovertestmode) - Stops Failover test mode

## Create

Creates a new workspace for the current user in the specified workspace group. You must specify the name of the workspace and the workspace group ID in the request body. Once a workspace is created, you can neither change its name nor its workspace group. See [Workspaces](https://docs.singlestore.com/managed-service/en/reference/management-api.html#workspaces-729524) for more information.
This API call does not take any request parameters.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )

    ctx := context.Background()
    res, err := s.Workspaces.Create(ctx, shared.WorkspaceCreate{
        Name: "demo-workspace",
        Size: singlestoresamplesdk.String("S-2"),
        WorkspaceGroupID: "c74bb6a6-0000-1000-9000-1d874fa277b0",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWorkspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `request`                                                            | [shared.WorkspaceCreate](../../pkg/models/shared/workspacecreate.md) | :heavy_check_mark:                                                   | The request object to use for the request.                           |


### Response

**[*operations.CreateWorkspaceResponse](../../pkg/models/operations/createworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateResume

Resumes a workspace with the specified workspace ID. You must specify the workspace ID in the API call.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "fa7d92df-5772-4b0d-b31b-774d5a65a840"

    ctx := context.Background()
    res, err := s.Workspaces.CreateResume(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWorkspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.CreateResumeWorkspaceResponse](../../pkg/models/operations/createresumeworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateStorage

You must specify the workspace ID of a workspace in the group you are setting up for disaster recovery and the region ID of your secondary region.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    storageDRSetup := shared.StorageDRSetup{
        DatabaseNames: []string{
            "x_db",
        },
        RegionID: "6eaad1a5-ac7d-4864-9f5c-1bc4cadf5345",
    }

    var workspaceID string = "a7ac8961-1881-4fc6-aa1d-b91f19b96abd"

    ctx := context.Background()
    res, err := s.Workspaces.CreateStorage(ctx, storageDRSetup, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                          | Type                                                               | Required                                                           | Description                                                        |
| ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ | ------------------------------------------------------------------ |
| `ctx`                                                              | [context.Context](https://pkg.go.dev/context#Context)              | :heavy_check_mark:                                                 | The context to use for the request.                                |
| `storageDRSetup`                                                   | [shared.StorageDRSetup](../../pkg/models/shared/storagedrsetup.md) | :heavy_check_mark:                                                 | Here's a sample of JSON data sent in the request body to the API.  |
| `workspaceID`                                                      | *string*                                                           | :heavy_check_mark:                                                 | ID of the workspace                                                |


### Response

**[*operations.CreateStorageWorkspaceResponse](../../pkg/models/operations/createstorageworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## CreateSuspend

Suspends a workspace with the specified workspace ID. You must specify the workspace ID in the API call.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "c87573e7-0304-4e45-9e6f-720fedb49011"

    ctx := context.Background()
    res, err := s.Workspaces.CreateSuspend(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWorkspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.CreateSuspendWorkspaceResponse](../../pkg/models/operations/createsuspendworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Delete

Terminates a workspace with the specified workspace ID. You must specify the workspace ID in the API call.

All the databases attached to the workspace are detached when the workspace is terminated.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "8db863f6-ef9b-413a-8a70-cb816b33de6b"

    ctx := context.Background()
    res, err := s.Workspaces.Delete(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWorkspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.DeleteWorkspaceResponse](../../pkg/models/operations/deleteworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Get

Returns workspace information for the specified workspace ID, in JSON format. You must specify the workspace ID in the API call.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var fields *string = singlestoresamplesdk.String("<value>")

    ctx := context.Background()
    res, err := s.Workspaces.Get(ctx, workspaceID, fields)
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `workspaceID`                                                                            | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the workspace                                                                      |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |


### Response

**[*operations.GetteWorkspaceResponse](../../pkg/models/operations/getteworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetOutbound

Returns the account ID which must be allowed for outbound connections.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "29f5c6af-4fba-4437-9707-abf378a1660b"

    ctx := context.Background()
    res, err := s.Workspaces.GetOutbound(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateConnectionOutboundAllowLists != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.GetOutboundWorkspaceResponse](../../pkg/models/operations/getoutboundworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetPrivateConnection

Returns private connection information for the specified workspace ID, in JSON format. You must specify the workspace ID in the API call.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "4e6c6827-27ee-4013-85e7-e36151e0fa57"

    var fields *string = singlestoresamplesdk.String("<value>")

    ctx := context.Background()
    res, err := s.Workspaces.GetPrivateConnection(ctx, workspaceID, fields)
    if err != nil {
        log.Fatal(err)
    }
    if res.PrivateConnections != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `workspaceID`                                                                            | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the workspace                                                                      |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |


### Response

**[*operations.GetPrivateConnectionWorkspaceResponse](../../pkg/models/operations/getprivateconnectionworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetRecoveryBackup

Returns a list of regions with regions IDs in JSON format. You must specify the workspace ID of a workspace in the group you are setting up for disaster recovery.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "0ebde48c-76bc-474b-a15a-a09cafb361f4"

    var fields *string = singlestoresamplesdk.String("<value>")

    ctx := context.Background()
    res, err := s.Workspaces.GetRecoveryBackup(ctx, workspaceID, fields)
    if err != nil {
        log.Fatal(err)
    }
    if res.Regions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `workspaceID`                                                                            | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the workspace                                                                      |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |


### Response

**[*operations.GetRecoveryBackupWorkspaceResponse](../../pkg/models/operations/getrecoverybackupworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## GetStorageStatus

Returns the replication status of each database and the status of the latest Storage DR operation (Failover, Failback, etc.). You must specify the workspace ID of a workspace in the group for which you are requesting the status information.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "0a28532b-9597-4f57-8ab0-f1d1e4f998a8"

    ctx := context.Background()
    res, err := s.Workspaces.GetStorageStatus(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res.StorageDRStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.GetStorageStatusWorkspaceResponse](../../pkg/models/operations/getstoragestatusworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## List

Returns a list of all of the workspaces accessible to the user in the specified workspace group. You must specify the workspace group ID in the API call to list the workspaces in the group. Use the `includeTerminated` parameter to list the terminated workspaces.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceGroupID string = "c184a429-302e-4aca-80db-f1718b882a50"

    var fields *string = singlestoresamplesdk.String("<value>")

    var includeTerminated *bool = singlestoresamplesdk.Bool(false)

    ctx := context.Background()
    res, err := s.Workspaces.List(ctx, workspaceGroupID, fields, includeTerminated)
    if err != nil {
        log.Fatal(err)
    }
    if res.Workspaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `workspaceGroupID`                                                                       | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the workspace group                                                                |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |
| `includeTerminated`                                                                      | **bool*                                                                                  | :heavy_minus_sign:                                                                       | To include any terminated workspaces, set to `true`                                      |


### Response

**[*operations.ListWorkspaceResponse](../../pkg/models/operations/listworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## Update

Updates workspace information for the specified workspace, including the size. Specify the workspace's new parameters in the request body. You must specify the workspace ID in the API call.


### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    workspaceUpdate := shared.WorkspaceUpdate{
        Size: singlestoresamplesdk.String("S-2"),
    }

    var workspaceID string = "d0905bf4-aa77-4f20-8e77-54c352acfe54"

    ctx := context.Background()
    res, err := s.Workspaces.Update(ctx, workspaceUpdate, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res.CreateWorkspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                            | Type                                                                 | Required                                                             | Description                                                          |
| -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- | -------------------------------------------------------------------- |
| `ctx`                                                                | [context.Context](https://pkg.go.dev/context#Context)                | :heavy_check_mark:                                                   | The context to use for the request.                                  |
| `workspaceUpdate`                                                    | [shared.WorkspaceUpdate](../../pkg/models/shared/workspaceupdate.md) | :heavy_check_mark:                                                   | Here's a sample of JSON data sent to the API in the request body.    |
| `workspaceID`                                                        | *string*                                                             | :heavy_check_mark:                                                   | ID of the workspace                                                  |


### Response

**[*operations.UpdateWorkspaceResponse](../../pkg/models/operations/updateworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateFailback

You must specify the workspace ID of a workspace in the standby (secondary) region from which you are triggering the failback.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "ee046eb6-b3d2-4b2b-8152-054b3b64d961"

    ctx := context.Background()
    res, err := s.Workspaces.UpdateFailback(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.UpdateFailbackWorkspaceResponse](../../pkg/models/operations/updatefailbackworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateFailover

You must specify the workspace ID of a workspace in the group in the inactive (primary) region from which you are triggering the failover.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "15aad03f-9196-4964-a967-bbe47f374a9c"

    ctx := context.Background()
    res, err := s.Workspaces.UpdateFailover(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.UpdateFailoverWorkspaceResponse](../../pkg/models/operations/updatefailoverworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateStartFailoverTestMode

You must specify the workspace ID of a workspace in the group in the active (primary) region that you will failover from. This will give you an opportunity to setup any configuration on your workspace group in the secondary region.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "062601a8-d539-45de-82f7-4bf69a96713d"

    ctx := context.Background()
    res, err := s.Workspaces.UpdateStartFailoverTestMode(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.UpdateStartFailoverTestModeWorkspaceResponse](../../pkg/models/operations/updatestartfailovertestmodeworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

## UpdateStopFailoverTestMode

You must specify the workspace ID of a workspace in the group in the active (primary) region that you will failover from. This will end the Failover test making the workspace group in the secondary region along with its configuration no longer accessible.

### Example Usage

```go
package main

import(
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var workspaceID string = "e2ad172a-bce4-4aca-90ec-87dc592f3d9f"

    ctx := context.Background()
    res, err := s.Workspaces.UpdateStopFailoverTestMode(ctx, workspaceID)
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceID`                                         | *string*                                              | :heavy_check_mark:                                    | ID of the workspace                                   |


### Response

**[*operations.UpdateStopFailoverTestModeWorkspaceResponse](../../pkg/models/operations/updatestopfailovertestmodeworkspaceresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
