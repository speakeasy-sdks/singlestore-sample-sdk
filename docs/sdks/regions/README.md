# Regions
(*Regions*)

## Overview

Operations related to regions

### Available Operations

* [List](#list) - Lists all of the regions for the user that support workspaces

## List

Returns a list of valid regions for the user that support workspaces, including the region ID and provider for each region.

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


    var fields *string = singlestoresamplesdk.String("<value>")

    ctx := context.Background()
    res, err := s.Regions.List(ctx, fields)
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
| `fields`                                                                                 | **string*                                                                                | :heavy_minus_sign:                                                                       | Comma-separated values list that correspond to the filtered fields for returned entities |


### Response

**[*operations.ListRegionsResponse](../../pkg/models/operations/listregionsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
