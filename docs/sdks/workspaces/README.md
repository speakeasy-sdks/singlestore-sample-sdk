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

| Parameter                                                        | Type                                                             | Required                                                         | Description                                                      |
| ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- | ---------------------------------------------------------------- |
| `ctx`                                                            | [context.Context](https://pkg.go.dev/context#Context)            | :heavy_check_mark:                                               | The context to use for the request.                              |
| `request`                                                        | [shared.WorkspaceCreate](../../models/shared/workspacecreate.md) | :heavy_check_mark:                                               | The request object to use for the request.                       |


### Response

**[*operations.CreateWorkspaceResponse](../../models/operations/createworkspaceresponse.md), error**


## CreateResume

Resumes a workspace with the specified workspace ID. You must specify the workspace ID in the API call.


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
    res, err := s.Workspaces.CreateResume(ctx, operations.CreateResumeWorkspaceRequest{
        WorkspaceID: "fa7d92df-5772-4b0d-b31b-774d5a65a840",
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

| Parameter                                                                                          | Type                                                                                               | Required                                                                                           | Description                                                                                        |
| -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                              | :heavy_check_mark:                                                                                 | The context to use for the request.                                                                |
| `request`                                                                                          | [operations.CreateResumeWorkspaceRequest](../../models/operations/createresumeworkspacerequest.md) | :heavy_check_mark:                                                                                 | The request object to use for the request.                                                         |


### Response

**[*operations.CreateResumeWorkspaceResponse](../../models/operations/createresumeworkspaceresponse.md), error**


## CreateStorage

You must specify the workspace ID of a workspace in the group you are setting up for disaster recovery and the region ID of your secondary region.

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
    res, err := s.Workspaces.CreateStorage(ctx, operations.CreateStorageWorkspaceRequest{
        StorageDRSetup: shared.StorageDRSetup{
            DatabaseNames: []string{
                "x",
                "_",
                "d",
                "b",
            },
            RegionID: "6eaad1a5-ac7d-4864-9f5c-1bc4cadf5345",
        },
        WorkspaceID: "a7ac8961-1881-4fc6-aa1d-b91f19b96abd",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateStorageWorkspaceRequest](../../models/operations/createstorageworkspacerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateStorageWorkspaceResponse](../../models/operations/createstorageworkspaceresponse.md), error**


## CreateSuspend

Suspends a workspace with the specified workspace ID. You must specify the workspace ID in the API call.

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
    res, err := s.Workspaces.CreateSuspend(ctx, operations.CreateSuspendWorkspaceRequest{
        WorkspaceID: "c87573e7-0304-4e45-9e6f-720fedb49011",
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

| Parameter                                                                                            | Type                                                                                                 | Required                                                                                             | Description                                                                                          |
| ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                                | :heavy_check_mark:                                                                                   | The context to use for the request.                                                                  |
| `request`                                                                                            | [operations.CreateSuspendWorkspaceRequest](../../models/operations/createsuspendworkspacerequest.md) | :heavy_check_mark:                                                                                   | The request object to use for the request.                                                           |


### Response

**[*operations.CreateSuspendWorkspaceResponse](../../models/operations/createsuspendworkspaceresponse.md), error**


## Delete

Terminates a workspace with the specified workspace ID. You must specify the workspace ID in the API call.

All the databases attached to the workspace are detached when the workspace is terminated.


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
    res, err := s.Workspaces.Delete(ctx, operations.DeleteWorkspaceRequest{
        WorkspaceID: "8db863f6-ef9b-413a-8a70-cb816b33de6b",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.DeleteWorkspaceRequest](../../models/operations/deleteworkspacerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.DeleteWorkspaceResponse](../../models/operations/deleteworkspaceresponse.md), error**


## Get

Returns workspace information for the specified workspace ID, in JSON format. You must specify the workspace ID in the API call.

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
    res, err := s.Workspaces.Get(ctx, operations.GetteWorkspaceRequest{
        WorkspaceID: "b18d8d81-fd7b-4764-a31e-475cb1f36591",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Workspace != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                            | Type                                                                                 | Required                                                                             | Description                                                                          |
| ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------ |
| `ctx`                                                                                | [context.Context](https://pkg.go.dev/context#Context)                                | :heavy_check_mark:                                                                   | The context to use for the request.                                                  |
| `request`                                                                            | [operations.GetteWorkspaceRequest](../../models/operations/getteworkspacerequest.md) | :heavy_check_mark:                                                                   | The request object to use for the request.                                           |


### Response

**[*operations.GetteWorkspaceResponse](../../models/operations/getteworkspaceresponse.md), error**


## GetOutbound

Returns the account ID which must be allowed for outbound connections.

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
    res, err := s.Workspaces.GetOutbound(ctx, operations.GetOutboundWorkspaceRequest{
        WorkspaceID: "29f5c6af-4fba-4437-9707-abf378a1660b",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.PrivateConnectionOutboundAllowLists != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                        | Type                                                                                             | Required                                                                                         | Description                                                                                      |
| ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                            | :heavy_check_mark:                                                                               | The context to use for the request.                                                              |
| `request`                                                                                        | [operations.GetOutboundWorkspaceRequest](../../models/operations/getoutboundworkspacerequest.md) | :heavy_check_mark:                                                                               | The request object to use for the request.                                                       |


### Response

**[*operations.GetOutboundWorkspaceResponse](../../models/operations/getoutboundworkspaceresponse.md), error**


## GetPrivateConnection

Returns private connection information for the specified workspace ID, in JSON format. You must specify the workspace ID in the API call.

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
    res, err := s.Workspaces.GetPrivateConnection(ctx, operations.GetPrivateConnectionWorkspaceRequest{
        WorkspaceID: "4e6c6827-27ee-4013-85e7-e36151e0fa57",
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

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.GetPrivateConnectionWorkspaceRequest](../../models/operations/getprivateconnectionworkspacerequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.GetPrivateConnectionWorkspaceResponse](../../models/operations/getprivateconnectionworkspaceresponse.md), error**


## GetRecoveryBackup

Returns a list of regions with regions IDs in JSON format. You must specify the workspace ID of a workspace in the group you are setting up for disaster recovery.

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
    res, err := s.Workspaces.GetRecoveryBackup(ctx, operations.GetRecoveryBackupWorkspaceRequest{
        WorkspaceID: "0ebde48c-76bc-474b-a15a-a09cafb361f4",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Regions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetRecoveryBackupWorkspaceRequest](../../models/operations/getrecoverybackupworkspacerequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetRecoveryBackupWorkspaceResponse](../../models/operations/getrecoverybackupworkspaceresponse.md), error**


## GetStorageStatus

Returns the replication status of each database and the status of the latest Storage DR operation (Failover, Failback, etc.). You must specify the workspace ID of a workspace in the group for which you are requesting the status information.

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
    res, err := s.Workspaces.GetStorageStatus(ctx, operations.GetStorageStatusWorkspaceRequest{
        WorkspaceID: "0a28532b-9597-4f57-8ab0-f1d1e4f998a8",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StorageDRStatus != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                  | Type                                                                                                       | Required                                                                                                   | Description                                                                                                |
| ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                      | [context.Context](https://pkg.go.dev/context#Context)                                                      | :heavy_check_mark:                                                                                         | The context to use for the request.                                                                        |
| `request`                                                                                                  | [operations.GetStorageStatusWorkspaceRequest](../../models/operations/getstoragestatusworkspacerequest.md) | :heavy_check_mark:                                                                                         | The request object to use for the request.                                                                 |


### Response

**[*operations.GetStorageStatusWorkspaceResponse](../../models/operations/getstoragestatusworkspaceresponse.md), error**


## List

Returns a list of all of the workspaces accessible to the user in the specified workspace group. You must specify the workspace group ID in the API call to list the workspaces in the group. Use the `includeTerminated` parameter to list the terminated workspaces.


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
    res, err := s.Workspaces.List(ctx, operations.ListWorkspaceRequest{
        WorkspaceGroupID: "c184a429-302e-4aca-80db-f1718b882a50",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.Workspaces != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.ListWorkspaceRequest](../../models/operations/listworkspacerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.ListWorkspaceResponse](../../models/operations/listworkspaceresponse.md), error**


## Update

Updates workspace information for the specified workspace, including the size. Specify the workspace's new parameters in the request body. You must specify the workspace ID in the API call.


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
    res, err := s.Workspaces.Update(ctx, operations.UpdateWorkspaceRequest{
        WorkspaceUpdate: shared.WorkspaceUpdate{
            Size: singlestoresamplesdk.String("S-2"),
        },
        WorkspaceID: "d0905bf4-aa77-4f20-8e77-54c352acfe54",
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

| Parameter                                                                              | Type                                                                                   | Required                                                                               | Description                                                                            |
| -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------- |
| `ctx`                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                  | :heavy_check_mark:                                                                     | The context to use for the request.                                                    |
| `request`                                                                              | [operations.UpdateWorkspaceRequest](../../models/operations/updateworkspacerequest.md) | :heavy_check_mark:                                                                     | The request object to use for the request.                                             |


### Response

**[*operations.UpdateWorkspaceResponse](../../models/operations/updateworkspaceresponse.md), error**


## UpdateFailback

You must specify the workspace ID of a workspace in the standby (secondary) region from which you are triggering the failback.

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
    res, err := s.Workspaces.UpdateFailback(ctx, operations.UpdateFailbackWorkspaceRequest{
        WorkspaceID: "ee046eb6-b3d2-4b2b-8152-054b3b64d961",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateFailbackWorkspaceRequest](../../models/operations/updatefailbackworkspacerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.UpdateFailbackWorkspaceResponse](../../models/operations/updatefailbackworkspaceresponse.md), error**


## UpdateFailover

You must specify the workspace ID of a workspace in the group in the inactive (primary) region from which you are triggering the failover.

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
    res, err := s.Workspaces.UpdateFailover(ctx, operations.UpdateFailoverWorkspaceRequest{
        WorkspaceID: "15aad03f-9196-4964-a967-bbe47f374a9c",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                              | Type                                                                                                   | Required                                                                                               | Description                                                                                            |
| ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                  | [context.Context](https://pkg.go.dev/context#Context)                                                  | :heavy_check_mark:                                                                                     | The context to use for the request.                                                                    |
| `request`                                                                                              | [operations.UpdateFailoverWorkspaceRequest](../../models/operations/updatefailoverworkspacerequest.md) | :heavy_check_mark:                                                                                     | The request object to use for the request.                                                             |


### Response

**[*operations.UpdateFailoverWorkspaceResponse](../../models/operations/updatefailoverworkspaceresponse.md), error**


## UpdateStartFailoverTestMode

You must specify the workspace ID of a workspace in the group in the active (primary) region that you will failover from. This will give you an opportunity to setup any configuration on your workspace group in the secondary region.

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
    res, err := s.Workspaces.UpdateStartFailoverTestMode(ctx, operations.UpdateStartFailoverTestModeWorkspaceRequest{
        WorkspaceID: "062601a8-d539-45de-82f7-4bf69a96713d",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                        | Type                                                                                                                             | Required                                                                                                                         | Description                                                                                                                      |
| -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                                            | :heavy_check_mark:                                                                                                               | The context to use for the request.                                                                                              |
| `request`                                                                                                                        | [operations.UpdateStartFailoverTestModeWorkspaceRequest](../../models/operations/updatestartfailovertestmodeworkspacerequest.md) | :heavy_check_mark:                                                                                                               | The request object to use for the request.                                                                                       |


### Response

**[*operations.UpdateStartFailoverTestModeWorkspaceResponse](../../models/operations/updatestartfailovertestmodeworkspaceresponse.md), error**


## UpdateStopFailoverTestMode

You must specify the workspace ID of a workspace in the group in the active (primary) region that you will failover from. This will end the Failover test making the workspace group in the secondary region along with its configuration no longer accessible.

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
    res, err := s.Workspaces.UpdateStopFailoverTestMode(ctx, operations.UpdateStopFailoverTestModeWorkspaceRequest{
        WorkspaceID: "e2ad172a-bce4-4aca-90ec-87dc592f3d9f",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                                      | Type                                                                                                                           | Required                                                                                                                       | Description                                                                                                                    |
| ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                                          | :heavy_check_mark:                                                                                                             | The context to use for the request.                                                                                            |
| `request`                                                                                                                      | [operations.UpdateStopFailoverTestModeWorkspaceRequest](../../models/operations/updatestopfailovertestmodeworkspacerequest.md) | :heavy_check_mark:                                                                                                             | The request object to use for the request.                                                                                     |


### Response

**[*operations.UpdateStopFailoverTestModeWorkspaceResponse](../../models/operations/updatestopfailovertestmodeworkspaceresponse.md), error**

