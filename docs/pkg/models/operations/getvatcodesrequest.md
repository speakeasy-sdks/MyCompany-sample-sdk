# GetVatCodesRequest


## Fields

| Field                                                                                                                                                                 | Type                                                                                                                                                                  | Required                                                                                                                                                              | Description                                                                                                                                                           |
| --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- | --------------------------------------------------------------------------------------------------------------------------------------------------------------------- |
| `IncludeExpired`                                                                                                                                                      | **bool*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                    | Include expired VAT codes. Expired codes has a ValidTo date and can still be used for accounting events in the validFrom - validTo timeframe as long as it is active. |
| `IsActive`                                                                                                                                                            | **bool*                                                                                                                                                               | :heavy_minus_sign:                                                                                                                                                    | Returns both active and inactive codes as default. True returns active codes, False returns only inactive codes.                                                      |
| `ResourceParameter`                                                                                                                                                   | [*shared.ResourceParameters](../../../pkg/models/shared/resourceparameters.md)                                                                                        | :heavy_minus_sign:                                                                                                                                                    | Structure containing various resource-filter options                                                                                                                  |