# VatCodes
(*VatCodes*)

### Available Operations

* [GetVatCodes](#getvatcodes) - Gets the vat codes on the client. (Auth roles: CommonServices,CommonServices_Full)

## GetVatCodes

Gets the SAF-T vat codes and any custom codes set on the client.
Custom codes will have a prefix with the letter 'C', and will be based on the SAF-T codes, but used if the client have mixed vat conditions, for instance.
Compensation codes has a prefix with the letter 'K'.

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
    res, err := s.VatCodes.GetVatCodes(ctx, operations.GetVatCodesRequest{})
    if err != nil {
        log.Fatal(err)
    }
    if res.Classes != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetVatCodesRequest](../../pkg/models/operations/getvatcodesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetVatCodesResponse](../../pkg/models/operations/getvatcodesresponse.md), error**
| Error Object             | Status Code              | Content Type             |
| ------------------------ | ------------------------ | ------------------------ |
| sdkerrors.ProblemDetails | 400                      | application/json         |
| sdkerrors.SDKError       | 4xx-5xx                  | */*                      |
