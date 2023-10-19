# GetFinancialSettingsResponse


## Fields

| Field                                                                       | Type                                                                        | Required                                                                    | Description                                                                 |
| --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- | --------------------------------------------------------------------------- |
| `ContentType`                                                               | *string*                                                                    | :heavy_check_mark:                                                          | HTTP response content type for this operation                               |
| `FinancialSettingsDto`                                                      | [*shared.FinancialSettingsDto](../../models/shared/financialsettingsdto.md) | :heavy_minus_sign:                                                          | Given when resource was successfully returned.                              |
| `StatusCode`                                                                | *int*                                                                       | :heavy_check_mark:                                                          | HTTP response status code for this operation                                |
| `RawResponse`                                                               | [*http.Response](https://pkg.go.dev/net/http#Response)                      | :heavy_minus_sign:                                                          | Raw HTTP response; suitable for custom response parsing                     |