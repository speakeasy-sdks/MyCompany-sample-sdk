# LockDateSettings
(*LockDateSettings*)

### Available Operations

* [GetLockDate](#getlockdate) - Gets the lock date on the client. (Auth roles: CommonServices,CommonServices_Full)

## GetLockDate

Gets the lock date on the client. (Auth roles: CommonServices,CommonServices_Full)

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
    res, err := s.LockDateSettings.GetLockDate(ctx)
    if err != nil {
        log.Fatal(err)
    }
    if res.LockDateSettingsDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                             | Type                                                  | Required                                              | Description                                           |
| ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- | ----------------------------------------------------- |
| `ctx`                                                 | [context.Context](https://pkg.go.dev/context#Context) | :heavy_check_mark:                                    | The context to use for the request.                   |


### Response

**[*operations.GetLockDateResponse](../../pkg/models/operations/getlockdateresponse.md), error**
| Error Object       | Status Code        | Content Type       |
| ------------------ | ------------------ | ------------------ |
| sdkerrors.SDKError | 4xx-5xx            | */*                |
