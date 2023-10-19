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
	"context"
	"log"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.VatCodes.GetVatCodes(ctx, operations.GetVatCodesRequest{
        ResourceParameter: &shared.ResourceParametersInput{
            Fields: mycompanysamplesdk.String("Gorgeous"),
            OrderBy: mycompanysamplesdk.String("yellow"),
            PageNumber: mycompanysamplesdk.Int(2),
            PageSize: mycompanysamplesdk.Int(5000),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.VatCodeDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                      | Type                                                                           | Required                                                                       | Description                                                                    |
| ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ | ------------------------------------------------------------------------------ |
| `ctx`                                                                          | [context.Context](https://pkg.go.dev/context#Context)                          | :heavy_check_mark:                                                             | The context to use for the request.                                            |
| `request`                                                                      | [operations.GetVatCodesRequest](../../models/operations/getvatcodesrequest.md) | :heavy_check_mark:                                                             | The request object to use for the request.                                     |


### Response

**[*operations.GetVatCodesResponse](../../models/operations/getvatcodesresponse.md), error**

