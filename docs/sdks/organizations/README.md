# Organizations
(*Organizations*)

## Overview

Operations related to organizations

### Available Operations

* [Get](#get) - Gets information about the current organization

## Get

Returns information about the current authorized user's organization.

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
    res, err := s.Organizations.Get(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.Organization != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetOrganizationInfoResponse](../../pkg/models/operations/getorganizationinforesponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
