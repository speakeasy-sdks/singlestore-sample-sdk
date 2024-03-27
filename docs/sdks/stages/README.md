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
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var path string = "<value>"

    var workspaceGroupID string = "77ad642c-1fc6-4fe0-b241-bcdd89dc7fa5"

    createStagesFileRequest := &shared.CreateStagesFileRequest{}

    var isFile *bool = singlestoresamplesdk.Bool(false)

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

| Parameter                                                                             | Type                                                                                  | Required                                                                              | Description                                                                           |
| ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------- |
| `ctx`                                                                                 | [context.Context](https://pkg.go.dev/context#Context)                                 | :heavy_check_mark:                                                                    | The context to use for the request.                                                   |
| `path`                                                                                | *string*                                                                              | :heavy_check_mark:                                                                    | Path in Stages                                                                        |
| `workspaceGroupID`                                                                    | *string*                                                                              | :heavy_check_mark:                                                                    | ID of the Stages-enabled workspace group                                              |
| `createStagesFileRequest`                                                             | [*shared.CreateStagesFileRequest](../../pkg/models/shared/createstagesfilerequest.md) | :heavy_minus_sign:                                                                    | N/A                                                                                   |
| `isFile`                                                                              | **bool*                                                                               | :heavy_minus_sign:                                                                    | If set to `true`, forces creation of an empty file                                    |


### Response

**[*operations.CreateStagesFileResponse](../../pkg/models/operations/createstagesfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var path string = "<value>"

    var workspaceGroupID string = "8db863f6-ef9b-413a-8a70-cb816b33de6b"

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

**[*operations.DeleteStagesFileResponse](../../pkg/models/operations/deletestagesfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var path string = "<value>"

    var workspaceGroupID string = "b18d8d81-fd7b-4764-a31e-475cb1f36591"

    var metadata *bool = singlestoresamplesdk.Bool(false)

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

**[*operations.GetStagesFileResponse](../../pkg/models/operations/getstagesfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

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
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var path string = "<value>"

    var workspaceGroupID string = "d0905bf4-aa77-4f20-8e77-54c352acfe54"

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

| Parameter                                                     | Type                                                          | Required                                                      | Description                                                   |
| ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- | ------------------------------------------------------------- |
| `ctx`                                                         | [context.Context](https://pkg.go.dev/context#Context)         | :heavy_check_mark:                                            | The context to use for the request.                           |
| `path`                                                        | *string*                                                      | :heavy_check_mark:                                            | Path in Stages to modify                                      |
| `workspaceGroupID`                                            | *string*                                                      | :heavy_check_mark:                                            | ID of the Stages-enabled workspace group                      |
| `stagesPatch`                                                 | [*shared.StagesPatch](../../pkg/models/shared/stagespatch.md) | :heavy_minus_sign:                                            | N/A                                                           |


### Response

**[*operations.UpdateStagesFileResponse](../../pkg/models/operations/updatestagesfileresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
