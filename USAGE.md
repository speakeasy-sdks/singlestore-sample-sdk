<!-- Start SDK Example Usage [usage] -->
```go
package main

import (
	"context"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
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
<!-- End SDK Example Usage [usage] -->