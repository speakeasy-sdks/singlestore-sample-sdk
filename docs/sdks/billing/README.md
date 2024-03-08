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
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
	"context"
	"log"
)

func main() {
    s := singlestoresamplesdk.New(
        singlestoresamplesdk.WithSecurity("<YOUR_BEARER_TOKEN_HERE>"),
    )


    var endTime string = "<value>"

    var startTime string = "<value>"

    var aggregateBy *operations.AggregateBy = operations.AggregateByMonth.ToPointer()

    var metric *operations.Metric = operations.MetricComputeCredit.ToPointer()

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

| Parameter                                                                                   | Type                                                                                        | Required                                                                                    | Description                                                                                 |
| ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- | ------------------------------------------------------------------------------------------- |
| `ctx`                                                                                       | [context.Context](https://pkg.go.dev/context#Context)                                       | :heavy_check_mark:                                                                          | The context to use for the request.                                                         |
| `endTime`                                                                                   | *string*                                                                                    | :heavy_check_mark:                                                                          | The end time for the usage interval valid UTC ISO8601 format e.g. 2023-07-30T18:30:00Z      |
| `startTime`                                                                                 | *string*                                                                                    | :heavy_check_mark:                                                                          | The start time for the usage interval in valid UTC ISO8601 format e.g. 2023-07-30T18:30:00Z |
| `aggregateBy`                                                                               | [*operations.AggregateBy](../../pkg/models/operations/aggregateby.md)                       | :heavy_minus_sign:                                                                          | The aggregate type used to group usage which includes hour, day and month. default is hour  |
| `metric`                                                                                    | [*operations.Metric](../../pkg/models/operations/metric.md)                                 | :heavy_minus_sign:                                                                          | Metrics include ComputeCredit, StorageAvgByte. default is all                               |


### Response

**[*operations.ListBillingUsageResponse](../../pkg/models/operations/listbillingusageresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
