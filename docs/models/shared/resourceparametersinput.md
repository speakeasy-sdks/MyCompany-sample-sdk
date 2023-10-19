# ResourceParametersInput

Structure containing various resource-filter options


## Fields

| Field                                                                                                                                                                                                                                            | Type                                                                                                                                                                                                                                             | Required                                                                                                                                                                                                                                         | Description                                                                                                                                                                                                                                      | Example                                                                                                                                                                                                                                          |
| ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------------ |
| `Fields`                                                                                                                                                                                                                                         | **string*                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                               | Define desired fields that should be returned in Json (All fields by default).                                                                                                                                                                   |                                                                                                                                                                                                                                                  |
| `OrderBy`                                                                                                                                                                                                                                        | **string*                                                                                                                                                                                                                                        | :heavy_minus_sign:                                                                                                                                                                                                                               | Define fields to order by, and if it should order asc (default) or desc).                                                                                                                                                                        |                                                                                                                                                                                                                                                  |
| `PageNumber`                                                                                                                                                                                                                                     | **int*                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                               | The desired page to return.                                                                                                                                                                                                                      | 2                                                                                                                                                                                                                                                |
| `PageSize`                                                                                                                                                                                                                                       | **int*                                                                                                                                                                                                                                           | :heavy_minus_sign:                                                                                                                                                                                                                               | The amount of objects to return in a Page (5000 by default). Not allowed to be larger than MaxPageSize.                                                                                                                                          | 5000                                                                                                                                                                                                                                             |
| `UseDatabaseValidation`                                                                                                                                                                                                                          | **bool*                                                                                                                                                                                                                                          | :heavy_minus_sign:                                                                                                                                                                                                                               | If set to true, all relevant filter-values will be validated against the database. NOTE: This will make the api-call much slower, but only for the initial request, as the validation-results are cached. Useful for debugging/support purposes. |                                                                                                                                                                                                                                                  |