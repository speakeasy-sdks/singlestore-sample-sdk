# Billing
(*Billing*)

## Overview

Operations related to billing

### Available Operations

* [List](#list) - Lists the compute and storage usage of a organization

## List

Lists the compute and storage usage of an organization in a given timeframe and aggregate type. The usage entries also contains details such as name, type and usage value.


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
    res, err := s.Billing.List(ctx, operations.ListBillingUsageRequest{
        EndTime: "Northeast Metal Canada",
        StartTime: "Data Response West",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingUsage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [operations.ListBillingUsageRequest](../../models/operations/listbillingusagerequest.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.ListBillingUsageResponse](../../models/operations/listbillingusageresponse.md), error**

