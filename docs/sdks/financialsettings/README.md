# FinancialSettings
(*FinancialSettings*)

### Available Operations

* [GetFinancialSettings](#getfinancialsettings) - Gets the financial settings on the client, and conversion date. (Auth roles: CommonServices,CommonServices_Full)

## GetFinancialSettings

Changes in agio (gains) and disagio (loss) accounts does not update the LastChangedDateTimeOffset property in FinancialSettings.

### Example Usage

```go
package main

import(
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"context"
	"log"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.FinancialSettings.GetFinancialSettings(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.FinancialSettingsDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetFinancialSettingsResponse](../../pkg/models/operations/getfinancialsettingsresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
