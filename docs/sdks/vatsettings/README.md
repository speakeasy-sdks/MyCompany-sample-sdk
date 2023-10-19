# VatSettings
(*VatSettings*)

### Available Operations

* [GetVatSettings](#getvatsettings) - Gets the vat settings on the client. (Auth roles: CommonServices,CommonServices_Full)

## GetVatSettings

Gets the vat settings on the client. (Auth roles: CommonServices,CommonServices_Full)

### Example Usage

```go
package main

import(
	"context"
	"log"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.VatSettings.GetVatSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }

    if res.VatSettingsDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetVatSettingsResponse](../../models/operations/getvatsettingsresponse.md), error**

