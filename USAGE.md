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

	ctx := context.Background()
	res, err := s.Billing.List(ctx, operations.ListBillingUsageRequest{
		EndTime:   "Northeast Metal Canada",
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
<!-- End SDK Example Usage -->