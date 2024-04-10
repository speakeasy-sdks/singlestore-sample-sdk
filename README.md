# Golang SDK for Singlestore Management API

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/singlestore-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/singlestore-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    <a href="https://codespaces.new/speakeasy-sdks/singlestore-sample-sdk.git/tree/main"><img src="https://github.com/codespaces/badge.svg" /></a>
</div>

<!-- Start SDK Installation [installation] -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/singlestore-sample-sdk
```
<!-- End SDK Installation [installation] -->

<!-- Start SDK Example Usage [usage] -->
## SDK Example Usage

### Example

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

<!-- Start Available Resources and Operations [operations] -->
## Available Resources and Operations

### [Billing](docs/sdks/billing/README.md)

* [List](docs/sdks/billing/README.md#list) - Lists the compute and storage usage of a organization

### [Organizations](docs/sdks/organizations/README.md)

* [Get](docs/sdks/organizations/README.md#get) - Gets information about the current organization

### [PrivateConnection](docs/sdks/privateconnection/README.md)

* [Create](docs/sdks/privateconnection/README.md#create) - Creates a new private connection
* [Delete](docs/sdks/privateconnection/README.md#delete) - Deletes a private connection
* [Get](docs/sdks/privateconnection/README.md#get) - Gets information about a private connection
* [Update](docs/sdks/privateconnection/README.md#update) - Updates a private connection

### [Regions](docs/sdks/regions/README.md)

* [List](docs/sdks/regions/README.md#list) - Lists all of the regions for the user that support workspaces

### [Stages](docs/sdks/stages/README.md)

* [Create](docs/sdks/stages/README.md#create) - Creates a new folder or uploads a file
* [Delete](docs/sdks/stages/README.md#delete) - Deletes a file or folder
* [Get](docs/sdks/stages/README.md#get) - Gets information about a folder or downloads a file
* [Update](docs/sdks/stages/README.md#update) - Modifies a file or folder in Stages

### [WorkspaceGroups](docs/sdks/workspacegroups/README.md)

* [Create](docs/sdks/workspacegroups/README.md#create) - Creates a new workspace group
* [CreateStorage](docs/sdks/workspacegroups/README.md#createstorage) - Sets up Storage DR for the workspace group. Backup region and selected databases to be replicated are provided as part of the request.
* [Delete](docs/sdks/workspacegroups/README.md#delete) - Terminates a workspace group
* [Get](docs/sdks/workspacegroups/README.md#get) - Gets information about a workspace group
* [GetPrivateConnection](docs/sdks/workspacegroups/README.md#getprivateconnection) - Gets private connection information for a workspace group
* [GetRecoveryBackup](docs/sdks/workspacegroups/README.md#getrecoverybackup) - Gets information about which regions you can setup as a disaster recovery backup
* [GetStorageStatus](docs/sdks/workspacegroups/README.md#getstoragestatus) - Gets information about the storage DR status of the workspace group
* [List](docs/sdks/workspacegroups/README.md#list) - Lists the workspace groups the user can access
* [Update](docs/sdks/workspacegroups/README.md#update) - Updates a workspace group
* [UpdateFailback](docs/sdks/workspacegroups/README.md#updatefailback) - Starts failback to the primary region
* [UpdateFailover](docs/sdks/workspacegroups/README.md#updatefailover) - Starts failover to the secondary region
* [UpdateStartFailoverTestMode](docs/sdks/workspacegroups/README.md#updatestartfailovertestmode) - Starts Failover test mode
* [UpdateStopFailoverTestMode](docs/sdks/workspacegroups/README.md#updatestopfailovertestmode) - Stops Failover test mode

### [Workspaces](docs/sdks/workspaces/README.md)

* [Create](docs/sdks/workspaces/README.md#create) - Creates a new workspace
* [CreateResume](docs/sdks/workspaces/README.md#createresume) - Resumes a workspace
* [CreateStorage](docs/sdks/workspaces/README.md#createstorage) - Sets up Storage DR for the workspace group of the provided workspace. Backup region and selected databases to be replicated are provided as part of the request.
* [CreateSuspend](docs/sdks/workspaces/README.md#createsuspend) - Suspends a workspace
* [Delete](docs/sdks/workspaces/README.md#delete) - Terminates a workspace
* [Get](docs/sdks/workspaces/README.md#get) - Gets information about a workspace
* [GetOutbound](docs/sdks/workspaces/README.md#getoutbound) - Gets the outbound allow list for a workspace
* [GetPrivateConnection](docs/sdks/workspaces/README.md#getprivateconnection) - Gets private connection information for a workspace
* [GetRecoveryBackup](docs/sdks/workspaces/README.md#getrecoverybackup) - Gets information about which regions you can setup as a disaster recovery backup
* [GetStorageStatus](docs/sdks/workspaces/README.md#getstoragestatus) - Gets information about the storage DR status of the group in which the provided workspace belongs to
* [List](docs/sdks/workspaces/README.md#list) - Lists the workspaces the user can access
* [Update](docs/sdks/workspaces/README.md#update) - Updates information about a workspace
* [UpdateFailback](docs/sdks/workspaces/README.md#updatefailback) - Starts failback to the primary region
* [UpdateFailover](docs/sdks/workspaces/README.md#updatefailover) - Starts failover to the secondary region
* [UpdateStartFailoverTestMode](docs/sdks/workspaces/README.md#updatestartfailovertestmode) - Starts Failover test mode
* [UpdateStopFailoverTestMode](docs/sdks/workspaces/README.md#updatestopfailovertestmode) - Stops Failover test mode
<!-- End Available Resources and Operations [operations] -->

<!-- Start Special Types [types] -->
## Special Types


<!-- End Special Types [types] -->

<!-- Start Error Handling [errors] -->
## Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |

### Example

```go
package main

import (
	"context"
	"errors"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/sdkerrors"
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

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling [errors] -->

<!-- Start Server Selection [server] -->
## Server Selection

### Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://api.singlestore.com` | None |

#### Example

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
		singlestoresamplesdk.WithServerIndex(0),
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


### Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:
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
		singlestoresamplesdk.WithServerURL("https://api.singlestore.com"),
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
<!-- End Server Selection [server] -->

<!-- Start Custom HTTP Client [http-client] -->
## Custom HTTP Client

The Go SDK makes API calls that wrap an internal HTTP client. The requirements for the HTTP client are very simple. It must match this interface:

```go
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}
```

The built-in `net/http` client satisfies this interface and a default client based on the built-in is provided by default. To replace this default with a client of your own, you can implement this interface yourself or provide your own client configured as desired. Here's a simple example, which adds a client with a 30 second timeout.

```go
import (
	"net/http"
	"time"
	"github.com/myorg/your-go-sdk"
)

var (
	httpClient = &http.Client{Timeout: 30 * time.Second}
	sdkClient  = sdk.New(sdk.WithClient(httpClient))
)
```

This can be a convenient way to configure timeouts, cookies, proxies, custom headers, and other low-level configuration.
<!-- End Custom HTTP Client [http-client] -->

<!-- Start Authentication [security] -->
## Authentication

### Per-Client Security Schemes

This SDK supports the following security scheme globally:

| Name         | Type         | Scheme       |
| ------------ | ------------ | ------------ |
| `APIKeyAuth` | http         | HTTP Bearer  |

You can configure it using the `WithSecurity` option when initializing the SDK client instance. For example:
```go
package main

import (
	"context"
	singlestoresamplesdk "github.com/speakeasy-sdks/singlestore-sample-sdk"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/operations"
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
<!-- End Authentication [security] -->

<!-- Placeholder for Future Speakeasy SDK Sections -->

# Development

## Maturity

This SDK is in beta, and there may be breaking changes between versions without a major version update. Therefore, we recommend pinning usage
to a specific package version. This way, you can install the same version each time without breaking changes unless you are intentionally
looking for the latest version.

## Contributions

While we value open-source contributions to this SDK, this library is generated programmatically.
Feel free to open a PR or a Github issue as a proof of concept and we'll do our best to include it in a future release!

### SDK Created by [Speakeasy](https://docs.speakeasyapi.dev/docs/using-speakeasy/client-sdks)
