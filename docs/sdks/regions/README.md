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
    res, err := s.Regions.List(ctx, operations.ListRegionsRequest{})
    if err != nil {
        log.Fatal(err)
    }

    if res.Regions != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.ListRegionsRequest](../../models/operations/listregionsrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.ListRegionsResponse](../../models/operations/listregionsresponse.md), error**

