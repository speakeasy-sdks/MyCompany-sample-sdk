# github.com/speakeasy-sdks/MyCompany-sample-sdk

<div align="left">
    <a href="https://speakeasyapi.dev/"><img src="https://custom-icon-badges.demolab.com/badge/-Built%20By%20Speakeasy-212015?style=for-the-badge&logoColor=FBE331&logo=speakeasy&labelColor=545454" /></a>
    <a href="https://github.com/speakeasy-sdks/MyCompany-sample-sdk.git/actions"><img src="https://img.shields.io/github/actions/workflow/status/speakeasy-sdks/MyCompany-sample-sdk/speakeasy_sdk_generation.yml?style=for-the-badge" /></a>
    
</div>

<!-- Start SDK Installation -->
## SDK Installation

```bash
go get github.com/speakeasy-sdks/MyCompany-sample-sdk
```
<!-- End SDK Installation -->

## SDK Example Usage
<!-- Start SDK Example Usage -->
```go
package main

import (
	"context"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := mycompanysamplesdk.New()

	ctx := context.Background()
	res, err := s.Currencies.GetCurrencies(ctx, operations.GetCurrenciesRequest{
		ResourceParameter: &shared.ResourceParameters{
			Fields:     mycompanysamplesdk.String("string"),
			OrderBy:    mycompanysamplesdk.String("string"),
			PageNumber: mycompanysamplesdk.Int(2),
			PageSize:   mycompanysamplesdk.Int(5000),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End SDK Example Usage -->

<!-- Start SDK Available Operations -->
## Available Resources and Operations


### [Currencies](docs/sdks/currencies/README.md)

* [GetCurrencies](docs/sdks/currencies/README.md#getcurrencies) - Gets the currencies set on the client. (Auth roles: CommonServices,CommonServices_Full)

### [FinancialSettings](docs/sdks/financialsettings/README.md)

* [GetFinancialSettings](docs/sdks/financialsettings/README.md#getfinancialsettings) - Gets the financial settings on the client, and conversion date. (Auth roles: CommonServices,CommonServices_Full)

### [GeneralLedgerAccounts](docs/sdks/generalledgeraccounts/README.md)

* [GetGeneralLedgerAccountByID](docs/sdks/generalledgeraccounts/README.md#getgeneralledgeraccountbyid) - Get a GeneralLedgerAccount by ID (Auth roles: GeneralLedgerAccount,GeneralLedgerAccount_Full)
* [GetGeneralLedgerAccounts](docs/sdks/generalledgeraccounts/README.md#getgeneralledgeraccounts) - Get a list of general ledger accounts. (Auth roles: GeneralLedgerAccount,GeneralLedgerAccount_Full)
* [DeleteGeneralLedgerAccountsID](docs/sdks/generalledgeraccounts/README.md#deletegeneralledgeraccountsid) - Deletes a General Ledger Account with a given id. (Auth roles: GeneralLedgerAccount_Full)
* [PatchGeneralLedgerAccountsID](docs/sdks/generalledgeraccounts/README.md#patchgeneralledgeraccountsid) - Update an existing general ledger account. (Auth roles: GeneralLedgerAccount_Full)
* [PostGeneralLedgerAccounts](docs/sdks/generalledgeraccounts/README.md#postgeneralledgeraccounts) - Create a new general ledger account with setting. (Auth roles: GeneralLedgerAccount_Full)

### [LockDateSettings](docs/sdks/lockdatesettings/README.md)

* [GetLockDate](docs/sdks/lockdatesettings/README.md#getlockdate) - Gets the lock date on the client. (Auth roles: CommonServices,CommonServices_Full)

### [SubLedgerNumberSeries](docs/sdks/subledgernumberseries/README.md)

* [GetSubLedgerNumberSeries](docs/sdks/subledgernumberseries/README.md#getsubledgernumberseries) - Gets the sub-ledger number series set on the client (Auth roles: CommonServices,CommonServices_Full)
* [GetSubLedgerNumberSeriesID](docs/sdks/subledgernumberseries/README.md#getsubledgernumberseriesid) - Get a SubLedgerNumberSeries by ID. (Auth roles: CommonServices,CommonServices_Full)

### [VatCodes](docs/sdks/vatcodes/README.md)

* [GetVatCodes](docs/sdks/vatcodes/README.md#getvatcodes) - Gets the vat codes on the client. (Auth roles: CommonServices,CommonServices_Full)

### [VatSettings](docs/sdks/vatsettings/README.md)

* [GetVatSettings](docs/sdks/vatsettings/README.md#getvatsettings) - Gets the vat settings on the client. (Auth roles: CommonServices,CommonServices_Full)
<!-- End SDK Available Operations -->

<!-- Start Dev Containers -->

<!-- End Dev Containers -->

<!-- Start Go Types -->
# Special Types

This SDK defines the following custom types to assist with marshalling and unmarshalling data.

## Date

`types.Date` is a wrapper around time.Time that allows for JSON marshaling a date string formatted as "2006-01-02".

### Usage

```go
d1 := types.NewDate(time.Now()) // returns *types.Date

d2 := types.DateFromTime(time.Now()) // returns types.Date

d3, err := types.NewDateFromString("2019-01-01") // returns *types.Date, error

d4, err := types.DateFromString("2019-01-01") // returns types.Date, error

d5 := types.MustNewDateFromString("2019-01-01") // returns *types.Date and panics on error

d6 := types.MustDateFromString("2019-01-01") // returns types.Date and panics on error
```
<!-- End Go Types -->



<!-- Start Error Handling -->
# Error Handling

Handling errors in this SDK should largely match your expectations.  All operations return a response object or an error, they will never return both.  When specified by the OpenAPI spec document, the SDK will return the appropriate subclass.

| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ProblemDetails | 400                      | application/json         |
| sdkerrors.SDKError       | 400-600                  | */*                      |


## Example

```go
package main

import (
	"context"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := mycompanysamplesdk.New()

	ctx := context.Background()
	res, err := s.Currencies.GetCurrencies(ctx, operations.GetCurrenciesRequest{
		ResourceParameter: &shared.ResourceParameters{
			Fields:     mycompanysamplesdk.String("string"),
			OrderBy:    mycompanysamplesdk.String("string"),
			PageNumber: mycompanysamplesdk.Int(2),
			PageSize:   mycompanysamplesdk.Int(5000),
		},
	})
	if err != nil {

		var e *sdkerrors.ProblemDetails
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}

		var e *sdkerrors.SDKError
		if errors.As(err, &e) {
			// handle error
			log.Fatal(e.Error())
		}
	}
}

```
<!-- End Error Handling -->



<!-- Start Server Selection -->
# Server Selection

## Select Server by Index

You can override the default server globally using the `WithServerIndex` option when initializing the SDK client instance. The selected server will then be used as the default on the operations that use it. This table lists the indexes associated with the available servers:

| # | Server | Variables |
| - | ------ | --------- |
| 0 | `https://goapi.poweroffice.net/v2` | None |

For example:

```go
package main

import (
	"context"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := mycompanysamplesdk.New(
		mycompanysamplesdk.WithServerIndex(0),
	)

	ctx := context.Background()
	res, err := s.Currencies.GetCurrencies(ctx, operations.GetCurrenciesRequest{
		ResourceParameter: &shared.ResourceParameters{
			Fields:     mycompanysamplesdk.String("string"),
			OrderBy:    mycompanysamplesdk.String("string"),
			PageNumber: mycompanysamplesdk.Int(2),
			PageSize:   mycompanysamplesdk.Int(5000),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```


## Override Server URL Per-Client

The default server can also be overridden globally using the `WithServerURL` option when initializing the SDK client instance. For example:

```go
package main

import (
	"context"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := mycompanysamplesdk.New(
		mycompanysamplesdk.WithServerURL("https://goapi.poweroffice.net/v2"),
	)

	ctx := context.Background()
	res, err := s.Currencies.GetCurrencies(ctx, operations.GetCurrenciesRequest{
		ResourceParameter: &shared.ResourceParameters{
			Fields:     mycompanysamplesdk.String("string"),
			OrderBy:    mycompanysamplesdk.String("string"),
			PageNumber: mycompanysamplesdk.Int(2),
			PageSize:   mycompanysamplesdk.Int(5000),
		},
	})
	if err != nil {
		log.Fatal(err)
	}

	if res.Classes != nil {
		// handle response
	}
}

```
<!-- End Server Selection -->



<!-- Start Custom HTTP Client -->
# Custom HTTP Client

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
<!-- End Custom HTTP Client -->

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
