<!-- Start SDK Example Usage -->


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
<!-- End SDK Example Usage -->