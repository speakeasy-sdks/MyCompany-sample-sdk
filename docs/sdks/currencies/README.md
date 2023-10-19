# Currencies
(*Currencies*)

### Available Operations

* [GetCurrencies](#getcurrencies) - Gets the currencies set on the client. (Auth roles: CommonServices,CommonServices_Full)

## GetCurrencies

Gets a list of all currencies available in Go.
The currencies in active use on the client has the property IsActive = true.
Active currencies can be used in GUI operations in Go, but does not affect which currencies that can be used when
posting transactions using the APIs

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
    res, err := s.Currencies.GetCurrencies(ctx, operations.GetCurrenciesRequest{
        ResourceParameter: &shared.ResourceParametersInput{
            Fields: mycompanysamplesdk.String("teal"),
            OrderBy: mycompanysamplesdk.String("management"),
            PageNumber: mycompanysamplesdk.Int(2),
            PageSize: mycompanysamplesdk.Int(5000),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.CurrencyDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                          | Type                                                                               | Required                                                                           | Description                                                                        |
| ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------- |
| `ctx`                                                                              | [context.Context](https://pkg.go.dev/context#Context)                              | :heavy_check_mark:                                                                 | The context to use for the request.                                                |
| `request`                                                                          | [operations.GetCurrenciesRequest](../../models/operations/getcurrenciesrequest.md) | :heavy_check_mark:                                                                 | The request object to use for the request.                                         |


### Response

**[*operations.GetCurrenciesResponse](../../models/operations/getcurrenciesresponse.md), error**

