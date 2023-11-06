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


    var endTime string = "string"

    var startTime string = "string"

    var aggregateBy *operations.ListBillingUsageAggregateBy = operations.ListBillingUsageAggregateByMonth

    var metric *operations.ListBillingUsageMetric = operations.ListBillingUsageMetricComputeCredit

    ctx := context.Background()
    res, err := s.Billing.List(ctx, endTime, startTime, aggregateBy, metric)
    if err != nil {
        log.Fatal(err)
    }

    if res.BillingUsage != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                         | Type                                                                                              | Required                                                                                          | Description                                                                                       |
| ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                             | [context.Context](https://pkg.go.dev/context#Context)                                             | :heavy_check_mark:                                                                                | The context to use for the request.                                                               |
| `endTime`                                                                                         | *string*                                                                                          | :heavy_check_mark:                                                                                | The end time for the usage interval valid UTC ISO8601 format e.g. 2023-07-30T18:30:00Z            |
| `startTime`                                                                                       | *string*                                                                                          | :heavy_check_mark:                                                                                | The start time for the usage interval in valid UTC ISO8601 format e.g. 2023-07-30T18:30:00Z       |
| `aggregateBy`                                                                                     | [*operations.ListBillingUsageAggregateBy](../../models/operations/listbillingusageaggregateby.md) | :heavy_minus_sign:                                                                                | The aggregate type used to group usage which includes hour, day and month. default is hour        |
| `metric`                                                                                          | [*operations.ListBillingUsageMetric](../../models/operations/listbillingusagemetric.md)           | :heavy_minus_sign:                                                                                | Metrics include ComputeCredit, StorageAvgByte. default is all                                     |


### Response

**[*operations.ListBillingUsageResponse](../../models/operations/listbillingusageresponse.md), error**

