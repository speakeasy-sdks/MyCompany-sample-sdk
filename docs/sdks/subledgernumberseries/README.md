# SubLedgerNumberSeries
(*SubLedgerNumberSeries*)

### Available Operations

* [GetSubLedgerNumberSeries](#getsubledgernumberseries) - Gets the sub-ledger number series set on the client (Auth roles: CommonServices,CommonServices_Full)
* [GetSubLedgerNumberSeriesID](#getsubledgernumberseriesid) - Get a SubLedgerNumberSeries by ID. (Auth roles: CommonServices,CommonServices_Full)

## GetSubLedgerNumberSeries

Gets the sub-ledger number series set on the client. Sub-ledgers are sub-accounts of a general ledger account, used
for entries related to either customers, suppliers or employees.

### Example Usage

```go
package main

import(
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.SubLedgerNumberSeries.GetSubLedgerNumberSeries(ctx, operations.GetSubLedgerNumberSeriesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                    | Type                                                                                                         | Required                                                                                                     | Description                                                                                                  |
| ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                        | [context.Context](https://pkg.go.dev/context#Context)                                                        | :heavy_check_mark:                                                                                           | The context to use for the request.                                                                          |
| `request`                                                                                                    | [operations.GetSubLedgerNumberSeriesRequest](../../pkg/models/operations/getsubledgernumberseriesrequest.md) | :heavy_check_mark:                                                                                           | The request object to use for the request.                                                                   |


### Response

**[*operations.GetSubLedgerNumberSeriesResponse](../../pkg/models/operations/getsubledgernumberseriesresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ProblemDetails | 400                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |

## GetSubLedgerNumberSeriesID

Get a SubLedgerNumberSeries by ID. (Auth roles: CommonServices,CommonServices_Full)

### Example Usage

```go
package main

import(
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"context"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"log"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.SubLedgerNumberSeries.GetSubLedgerNumberSeriesID(ctx, operations.GetSubLedgerNumberSeriesIDRequest{
        ID: "19c1da1d-ccfd-4602-ab2a-5c8d858f5ee3",
    })
    if err != nil {
        log.Fatal(err)
    }
    if res != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.GetSubLedgerNumberSeriesIDRequest](../../pkg/models/operations/getsubledgernumberseriesidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.GetSubLedgerNumberSeriesIDResponse](../../pkg/models/operations/getsubledgernumberseriesidresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ProblemDetails | 404                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
