# WorkspaceGroups
(*WorkspaceGroups*)

### Available Operations

* [Create](#create) - Creates a new workspace group
* [CreateStorage](#createstorage) - Sets up Storage DR for the workspace group. Backup region and selected databases to be replicated are provided as part of the request.
* [Delete](#delete) - Terminates a workspace group
* [Get](#get) - Gets information about a workspace group
* [GetPrivateConnection](#getprivateconnection) - Gets private connection information for a workspace group
* [GetRecoveryBackup](#getrecoverybackup) - Gets information about which regions you can setup as a disaster recovery backup
* [GetStorageStatus](#getstoragestatus) - Gets information about the storage DR status of the workspace group
* [List](#list) - Lists the workspace groups the user can access
* [Update](#update) - Updates a workspace group
* [UpdateFailback](#updatefailback) - Starts failback to the primary region
* [UpdateFailover](#updatefailover) - Starts failover to the secondary region
* [UpdateStartFailoverTestMode](#updatestartfailovertestmode) - Starts Failover test mode
* [UpdateStopFailoverTestMode](#updatestopfailovertestmode) - Stops Failover test mode

## Create

Creates a new workspace group for the current user. You must specify the name, region ID, and firewall ranges for the workspace group in the API request body.

You may use the admin user password to connect with any workspace of the group. The admin user password can be specified in the request body. If the admin user password is not specified in the API request, a password is generated and returned in the response object.


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
    res, err := s.WorkspaceGroups.Create(ctx, shared.WorkspaceGroupCreate{
        AdminPassword: singlestoresamplesdk.String("online"),
        ExpiresAt: singlestoresamplesdk.String("Configuration"),
        FirewallRanges: []string{
            "192.168.0.1/32",
            "192.168.0.81/12",
        },
        Name: "demo-workspace-group",
        RegionID: "7e7ffd27-0000-1000-9000-e72828a81ac7",
        UpdateWindow: &shared.UpdateWindow{
            Day: 2,
            Hour: 1,
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateWorkspaceGroup != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `request`                                                                  | [shared.WorkspaceGroupCreate](../../models/shared/workspacegroupcreate.md) | :heavy_check_mark:                                                         | The request object to use for the request.                                 |


### Response

**[*operations.CreateWorkspaceGroupResponse](../../models/operations/createworkspacegroupresponse.md), error**


## CreateStorage

You must specify the workspace group ID of the group you are setting up for disaster recovery and the region ID of your secondary region.

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


    storageDRSetup := shared.StorageDRSetup{
        DatabaseNames: []string{
            "x",
            "_",
            "d",
            "b",
        },
        RegionID: "6eaad1a5-ac7d-4864-9f5c-1bc4cadf5345",
    }

    var workspaceGroupID string = "a7ac8961-1881-4fc6-aa1d-b91f19b96abd"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.CreateStorage(ctx, storageDRSetup, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                         | Type                                                              | Required                                                          | Description                                                       |
| ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- | ----------------------------------------------------------------- |
| `ctx`                                                             | [context.Context](https://pkg.go.dev/context#Context)             | :heavy_check_mark:                                                | The context to use for the request.                               |
| `storageDRSetup`                                                  | [shared.StorageDRSetup](../../models/shared/storagedrsetup.md)    | :heavy_check_mark:                                                | Here's a sample of JSON data sent in the request body to the API. |
| `workspaceGroupID`                                                | *string*                                                          | :heavy_check_mark:                                                | ID of the workspace group                                         |


### Response

**[*operations.CreateStorageWorkspaceGroupsResponse](../../models/operations/createstorageworkspacegroupsresponse.md), error**


## Delete

Terminates a workspace group with the specified workspace group ID. You must specify the workspace group ID in the API call.

By default, you may only terminate empty workspace groups (a workspace group without workspaces). To terminate a workspace group with active workspaces, use the `force` parameter.


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


    var workspaceGroupID string = "8db863f6-ef9b-413a-8a70-cb816b33de6b"

    var force *bool = false

    ctx := context.Background()
    res, err := s.WorkspaceGroups.Delete(ctx, workspaceGroupID, force)
    if err != nil {
        log.Fatal(err)
    }

    if res.TerminateWorkspaceGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `workspaceGroupID`                                                             | *string*                                                                       | :heavy_check_mark:                                                             | ID of the workspace group                                                      |
| `force`                                                                        | **bool*                                                                        | :heavy_minus_sign:                                                             | To terminate a workspace group even if it has active workspaces, set to `true` |


### Response

**[*operations.DeleteWorkspaceGroupsResponse](../../models/operations/deleteworkspacegroupsresponse.md), error**


## Get

Returns information for the specified
workspace group ID, in JSON format. You must specify the workspace group ID in the API call.


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


    var workspaceGroupID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var fields *string = "Optional"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.Get(ctx, workspaceGroupID, fields)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkspaceGroup != nil {
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


### Response

**[*operations.GetWorkspaceGroupsResponse](../../models/operations/getworkspacegroupsresponse.md), error**


## GetPrivateConnection

Returns private connection information for the specified
workspace group ID, in JSON format. You must specify the workspace group ID in the API call.


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


    var workspaceGroupID string = "4e6c6827-27ee-4013-85e7-e36151e0fa57"

    var fields *string = "mobile"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.GetPrivateConnection(ctx, workspaceGroupID, fields)
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
| `workspaceGroupID`                                                                       | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the workspace group                                                                |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |


### Response

**[*operations.GetPrivateConnectionWorkspaceGroupsResponse](../../models/operations/getprivateconnectionworkspacegroupsresponse.md), error**


## GetRecoveryBackup

Returns a list of regions with regions IDs in JSON format. You must specify the workspace group ID of the group you are setting up for disaster recovery.

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


    var workspaceGroupID string = "0ebde48c-76bc-474b-a15a-a09cafb361f4"

    var fields *string = "Dinar"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.GetRecoveryBackup(ctx, workspaceGroupID, fields)
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
| `workspaceGroupID`                                                                       | *string*                                                                                 | :heavy_check_mark:                                                                       | ID of the workspace group                                                                |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |


### Response

**[*operations.GetRecoveryBackupWorkspaceGroupsResponse](../../models/operations/getrecoverybackupworkspacegroupsresponse.md), error**


## GetStorageStatus

Returns the replication status of each database and the status of the latest Storage DR operation (Failover, Failback, etc.). You must specify the workspace group ID of the group that you are requesting status information for.

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


    var workspaceGroupID string = "0a28532b-9597-4f57-8ab0-f1d1e4f998a8"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.GetStorageStatus(ctx, workspaceGroupID)
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
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the workspace group                             |


### Response

**[*operations.GetStorageStatusWorkspaceGroupsResponse](../../models/operations/getstoragestatusworkspacegroupsresponse.md), error**


## List

Returns a list of all of the workspace groups accessible to the user. Use the `includeTerminated` parameter to get a list of terminated workspace groups.


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


    var fields *string = "Bicycle"

    var includeTerminated *bool = false

    ctx := context.Background()
    res, err := s.WorkspaceGroups.List(ctx, fields, includeTerminated)
    if err != nil {
        log.Fatal(err)
    }

    if res.WorkspaceGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |
| `includeTerminated`                                                                      | **bool*                                                                                  | :heavy_minus_sign:                                                                       | To include any terminated workspace groups, set to `true`                                |


### Response

**[*operations.ListWorkspaceGroupResponse](../../models/operations/listworkspacegroupresponse.md), error**


## Update

Updates workspace group information for the specified workspace group, including the name, admin password, and firewall ranges. Specify the workspace group's new parameters in the request body. You must specify the workspace group ID in the API call.


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


    workspaceGroupUpdate := shared.WorkspaceGroupUpdate{
        AdminPassword: singlestoresamplesdk.String("Van"),
        ExpiresAt: singlestoresamplesdk.String("East"),
        FirewallRanges: []string{
            "192.168.0.1/32",
            "192.168.0.81/12",
        },
        Name: singlestoresamplesdk.String("new-workspace-group-name"),
        UpdateWindow: &shared.UpdateWindow{
            Day: 2,
            Hour: 1,
        },
    }

    var workspaceGroupID string = "bf4aa77f-204e-4775-8c35-2acfe54077ca"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.Update(ctx, workspaceGroupUpdate, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.UpdateWorkspaceGroups != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                  | Type                                                                       | Required                                                                   | Description                                                                |
| -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- | -------------------------------------------------------------------------- |
| `ctx`                                                                      | [context.Context](https://pkg.go.dev/context#Context)                      | :heavy_check_mark:                                                         | The context to use for the request.                                        |
| `workspaceGroupUpdate`                                                     | [shared.WorkspaceGroupUpdate](../../models/shared/workspacegroupupdate.md) | :heavy_check_mark:                                                         | Here's a sample of JSON data sent to the API in the request body.          |
| `workspaceGroupID`                                                         | *string*                                                                   | :heavy_check_mark:                                                         | ID of the workspace group                                                  |


### Response

**[*operations.UpdateWorkspaceGroupsResponse](../../models/operations/updateworkspacegroupsresponse.md), error**


## UpdateFailback

You must specify the workspace group ID of the group in the standby (secondary) region from which you are triggering the failback.

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


    var workspaceGroupID string = "ee046eb6-b3d2-4b2b-8152-054b3b64d961"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.UpdateFailback(ctx, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the workspace group                             |


### Response

**[*operations.UpdateFailbackWorkspaceGroupsResponse](../../models/operations/updatefailbackworkspacegroupsresponse.md), error**


## UpdateFailover

You must specify the workspace group ID of the group in the inactive (primary) region from which you are triggering the failover.

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


    var workspaceGroupID string = "15aad03f-9196-4964-a967-bbe47f374a9c"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.UpdateFailover(ctx, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the workspace group                             |


### Response

**[*operations.UpdateFailoverWorkspaceGroupsResponse](../../models/operations/updatefailoverworkspacegroupsresponse.md), error**


## UpdateStartFailoverTestMode

You must specify the workspace group ID for the group in the active (primary) region that you will failover from. This will give you an opportunity to setup any configuration on your workspace group in the secondary region.

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


    var workspaceGroupID string = "062601a8-d539-45de-82f7-4bf69a96713d"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.UpdateStartFailoverTestMode(ctx, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the workspace group                             |


### Response

**[*operations.UpdateStartFailoverTestModeWorkspaceGroupsResponse](../../models/operations/updatestartfailovertestmodeworkspacegroupsresponse.md), error**


## UpdateStopFailoverTestMode

You must specify the workspace group ID for the group in the active (primary) region that you will failover from. This will end the Failover test making the workspace group in the secondary region along with its configuration no longer accessible.

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


    var workspaceGroupID string = "e2ad172a-bce4-4aca-90ec-87dc592f3d9f"

    ctx := context.Background()
    res, err := s.WorkspaceGroups.UpdateStopFailoverTestMode(ctx, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the workspace group                             |


### Response

**[*operations.UpdateStopFailoverTestModeWorkspaceGroupsResponse](../../models/operations/updatestopfailovertestmodeworkspacegroupsresponse.md), error**

