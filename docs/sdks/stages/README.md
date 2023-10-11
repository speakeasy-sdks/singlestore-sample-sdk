# Stages
(*Stages*)

## Overview

Operations related to stages

### Available Operations

* [Create](#create) - Creates a new folder or uploads a file
* [Delete](#delete) - Deletes a file or folder
* [Get](#get) - Gets information about a folder or downloads a file
* [Update](#update) - Modifies a file or folder in Stages

## Create

Creates a new folder or uploads a new file at the specified path in Stages as follows:
 - If the request body contains a file, the file is uploaded at the specified path in Stages.
 - If the request body does not contain a file and the `isFile` parameter is disabled, a folder is created at the specified path.
 - If the request body does not contain a file to upload and the `isFile` parameter is enabled, an empty file is created at the specified path.

This endpoint supports both URL-encoded and regular URL paths. Here are some examples of valid requests:
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder/sample_path
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder%2Fsample_path

You must specify the `workspaceGroupID` and the folder/file path in the API call.


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
    res, err := s.Stages.Create(ctx, operations.CreateStagesFileRequest{
        CreateStagesFileRequest: &shared.CreateStagesFileRequest{
            File: &shared.CreateStagesFileRequestFile{
                Content: []byte("NN\pG;-j'}"),
                File: "abnormally deposit evolve",
            },
        },
        Path: "/usr/bin",
        WorkspaceGroupID: "cdd89dc7-fa50-44e0-8333-b1d5e261915a",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateStagesFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.CreateStagesFileRequest](../../models/operations/createstagesfilerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.CreateStagesFileResponse](../../models/operations/createstagesfileresponse.md), error**


## Delete

Deletes the file or folder at the specified path in Stages.

This endpoint supports both URL-encoded and regular URL paths. Here are some examples of valid requests:
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder/sample_folder
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder%2Fsample_folder

You must specify the `workspaceGroupID` and the folder/file path in the API call.


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
    res, err := s.Stages.Delete(ctx, operations.DeleteStagesFileRequest{
        Path: "/root",
        WorkspaceGroupID: "db863f6e-f9b1-43ac-a70c-b816b33de6bc",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteStagesFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.DeleteStagesFileRequest](../../models/operations/deletestagesfilerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.DeleteStagesFileResponse](../../models/operations/deletestagesfileresponse.md), error**


## Get

If the specified path is a folder, the API returns the list of files and folders inside this folder in Stages. If the specified path is a file, the API call is redirected to a download URL.

This endpoint supports both URL-encoded and regular URL paths. Here are some examples of valid requests:
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder/sample_path
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder%2Fsample_path

You must specify the `workspaceGroupID` and the folder/file path in the API call.


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
    res, err := s.Stages.Get(ctx, operations.GetStagesFileRequest{
        Path: "/usr/bin",
        WorkspaceGroupID: "18d8d81f-d7b7-464e-b1e4-75cb1f365915",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StagesObjectMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetStagesFileRequest](../../models/operations/getstagesfilerequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetStagesFileResponse](../../models/operations/getstagesfileresponse.md), error**


## Update

Modifies the path of the existing file or folder in Stages to the new path specified in the request body.

This endpoint supports both URL-encoded and regular URL paths, so both of the following examples are valid requests:
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder/sample_folder
- /v1/stages/68af2f46-0000-1000-9000-3f6f5365d878/fs/parent_folder%2Fsample_folder

You must specify the `workspaceGroupID` and the folder/file path in the API call.


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
    res, err := s.Stages.Update(ctx, operations.UpdateStagesFileRequest{
        StagesPatch: &shared.StagesPatch{
            NewPath: singlestoresamplesdk.String("parent_folder/new_sample_folder"),
        },
        Path: "/usr/sbin",
        WorkspaceGroupID: "0905bf4a-a77f-4204-a775-4c352acfe540",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.ModifyStagesFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.UpdateStagesFileRequest](../../models/operations/updatestagesfilerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.UpdateStagesFileResponse](../../models/operations/updatestagesfileresponse.md), error**

