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
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )


    var path string = "online"

    var workspaceGroupID string = "ad642c1f-c6fe-4072-81bc-dd89dc7fa504"

    createStagesFileRequest := &shared.CreateStagesFileRequest{
        File: &shared.CreateStagesFileRequestFile{
            Content: []byte("t\"Q644c'n?"),
            File: "Grocery Borders Northwest",
        },
    }

    var isFile *bool = false

    ctx := context.Background()
    res, err := s.Stages.Create(ctx, path, workspaceGroupID, createStagesFileRequest, isFile)
    if err != nil {
        log.Fatal(err)
    }

    if res.CreateStagesFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                         | Type                                                                              | Required                                                                          | Description                                                                       |
| --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- | --------------------------------------------------------------------------------- |
| `ctx`                                                                             | [context.Context](https://pkg.go.dev/context#Context)                             | :heavy_check_mark:                                                                | The context to use for the request.                                               |
| `path`                                                                            | *string*                                                                          | :heavy_check_mark:                                                                | Path in Stages                                                                    |
| `workspaceGroupID`                                                                | *string*                                                                          | :heavy_check_mark:                                                                | ID of the Stages-enabled workspace group                                          |
| `createStagesFileRequest`                                                         | [*shared.CreateStagesFileRequest](../../models/shared/createstagesfilerequest.md) | :heavy_minus_sign:                                                                | N/A                                                                               |
| `isFile`                                                                          | **bool*                                                                           | :heavy_minus_sign:                                                                | If set to `true`, forces creation of an empty file                                |


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
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )


    var path string = "program"

    var workspaceGroupID string = "b863f6ef-9b13-4aca-b0cb-816b33de6bc7"

    ctx := context.Background()
    res, err := s.Stages.Delete(ctx, path, workspaceGroupID)
    if err != nil {
        log.Fatal(err)
    }

    if res.DeleteStagesFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `path`                                                | *string*                                              | :heavy_check_mark:                                    | Path in Stages to a file or folder to delete          |
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the Stages-enabled workspace group              |


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
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )


    var path string = "female"

    var workspaceGroupID string = "8d8d81fd-7b76-44e3-9e47-5cb1f3659158"

    var metadata *bool = false

    ctx := context.Background()
    res, err := s.Stages.Get(ctx, path, workspaceGroupID, metadata)
    if err != nil {
        log.Fatal(err)
    }

    if res.StagesObjectMetadata != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |
| `path`                                                | *string*                                              | :heavy_check_mark:                                    | Path in Stages to a file or folder                    |
| `workspaceGroupID`                                    | *string*                                              | :heavy_check_mark:                                    | ID of the Stages-enabled workspace group              |
| `metadata`                                            | **bool*                                               | :heavy_minus_sign:                                    | If enabled, the API request returns only metadata     |


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
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity(""),
    )


    var path string = "Van"

    var workspaceGroupID string = "05bf4aa7-7f20-44e7-b54c-352acfe54077"

    stagesPatch := &shared.StagesPatch{
        NewPath: singlestoresamplesdk.String("parent_folder/new_sample_folder"),
    }

    ctx := context.Background()
    res, err := s.Stages.Update(ctx, path, workspaceGroupID, stagesPatch)
    if err != nil {
        log.Fatal(err)
    }

    if res.ModifyStagesFile != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                 | Type                                                      | Required                                                  | Description                                               |
| --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- | --------------------------------------------------------- |
| `ctx`                                                     | [context.Context](https://pkg.go.dev/context#Context)     | :heavy_check_mark:                                        | The context to use for the request.                       |
| `path`                                                    | *string*                                                  | :heavy_check_mark:                                        | Path in Stages to modify                                  |
| `workspaceGroupID`                                        | *string*                                                  | :heavy_check_mark:                                        | ID of the Stages-enabled workspace group                  |
| `stagesPatch`                                             | [*shared.StagesPatch](../../models/shared/stagespatch.md) | :heavy_minus_sign:                                        | N/A                                                       |


### Response

**[*operations.UpdateStagesFileResponse](../../models/operations/updatestagesfileresponse.md), error**

