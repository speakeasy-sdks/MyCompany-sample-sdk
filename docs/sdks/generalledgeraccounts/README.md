# GeneralLedgerAccounts
(*GeneralLedgerAccounts*)

### Available Operations

* [GetGeneralLedgerAccountByID](#getgeneralledgeraccountbyid) - Get a GeneralLedgerAccount by ID (Auth roles: GeneralLedgerAccount,GeneralLedgerAccount_Full)
* [GetGeneralLedgerAccounts](#getgeneralledgeraccounts) - Get a list of general ledger accounts. (Auth roles: GeneralLedgerAccount,GeneralLedgerAccount_Full)
* [DeleteGeneralLedgerAccountsID](#deletegeneralledgeraccountsid) - Deletes a General Ledger Account with a given id. (Auth roles: GeneralLedgerAccount_Full)
* [PatchGeneralLedgerAccountsID](#patchgeneralledgeraccountsid) - Update an existing general ledger account. (Auth roles: GeneralLedgerAccount_Full)
* [PostGeneralLedgerAccounts](#postgeneralledgeraccounts) - Create a new general ledger account with setting. (Auth roles: GeneralLedgerAccount_Full)

## GetGeneralLedgerAccountByID

Get a GeneralLedgerAccount by ID (Auth roles: GeneralLedgerAccount,GeneralLedgerAccount_Full)

### Example Usage

```go
package main

import(
	"context"
	"log"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.GeneralLedgerAccounts.GetGeneralLedgerAccountByID(ctx, operations.GetGeneralLedgerAccountByIDRequest{
        ID: 921626,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GeneralLedgerAccountDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                      | Type                                                                                                           | Required                                                                                                       | Description                                                                                                    |
| -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                          | [context.Context](https://pkg.go.dev/context#Context)                                                          | :heavy_check_mark:                                                                                             | The context to use for the request.                                                                            |
| `request`                                                                                                      | [operations.GetGeneralLedgerAccountByIDRequest](../../models/operations/getgeneralledgeraccountbyidrequest.md) | :heavy_check_mark:                                                                                             | The request object to use for the request.                                                                     |


### Response

**[*operations.GetGeneralLedgerAccountByIDResponse](../../models/operations/getgeneralledgeraccountbyidresponse.md), error**


## GetGeneralLedgerAccounts

Get a list of general ledger accounts. (Auth roles: GeneralLedgerAccount,GeneralLedgerAccount_Full)

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
    res, err := s.GeneralLedgerAccounts.GetGeneralLedgerAccounts(ctx, operations.GetGeneralLedgerAccountsRequest{
        ResourceParameter: &shared.ResourceParametersInput{
            Fields: mycompanysamplesdk.String("deposit"),
            OrderBy: mycompanysamplesdk.String("ugh"),
            PageNumber: mycompanysamplesdk.Int(2),
            PageSize: mycompanysamplesdk.Int(5000),
        },
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GeneralLedgerAccountDtos != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                | Type                                                                                                     | Required                                                                                                 | Description                                                                                              |
| -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- | -------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                                    | :heavy_check_mark:                                                                                       | The context to use for the request.                                                                      |
| `request`                                                                                                | [operations.GetGeneralLedgerAccountsRequest](../../models/operations/getgeneralledgeraccountsrequest.md) | :heavy_check_mark:                                                                                       | The request object to use for the request.                                                               |


### Response

**[*operations.GetGeneralLedgerAccountsResponse](../../models/operations/getgeneralledgeraccountsresponse.md), error**


## DeleteGeneralLedgerAccountsID

Deletes a General Ledger Account with a given id. (Auth roles: GeneralLedgerAccount_Full)

### Example Usage

```go
package main

import(
	"context"
	"log"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.GeneralLedgerAccounts.DeleteGeneralLedgerAccountsID(ctx, operations.DeleteGeneralLedgerAccountsIDRequest{
        ID: 654350,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.StatusCode == http.StatusOK {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                          | Type                                                                                                               | Required                                                                                                           | Description                                                                                                        |
| ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ | ------------------------------------------------------------------------------------------------------------------ |
| `ctx`                                                                                                              | [context.Context](https://pkg.go.dev/context#Context)                                                              | :heavy_check_mark:                                                                                                 | The context to use for the request.                                                                                |
| `request`                                                                                                          | [operations.DeleteGeneralLedgerAccountsIDRequest](../../models/operations/deletegeneralledgeraccountsidrequest.md) | :heavy_check_mark:                                                                                                 | The request object to use for the request.                                                                         |


### Response

**[*operations.DeleteGeneralLedgerAccountsIDResponse](../../models/operations/deletegeneralledgeraccountsidresponse.md), error**


## PatchGeneralLedgerAccountsID

Update an existing general ledger account. (Auth roles: GeneralLedgerAccount_Full)

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
    res, err := s.GeneralLedgerAccounts.PatchGeneralLedgerAccountsID(ctx, operations.PatchGeneralLedgerAccountsIDRequest{
        RequestBody: []shared.Operation{
            shared.Operation{
                Value: &shared.OperationValue{},
            },
        },
        ID: 24189,
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GeneralLedgerAccountDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                                        | Type                                                                                                             | Required                                                                                                         | Description                                                                                                      |
| ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------------------------------- |
| `ctx`                                                                                                            | [context.Context](https://pkg.go.dev/context#Context)                                                            | :heavy_check_mark:                                                                                               | The context to use for the request.                                                                              |
| `request`                                                                                                        | [operations.PatchGeneralLedgerAccountsIDRequest](../../models/operations/patchgeneralledgeraccountsidrequest.md) | :heavy_check_mark:                                                                                               | The request object to use for the request.                                                                       |


### Response

**[*operations.PatchGeneralLedgerAccountsIDResponse](../../models/operations/patchgeneralledgeraccountsidresponse.md), error**


## PostGeneralLedgerAccounts

Create a new general ledger account with setting. (Auth roles: GeneralLedgerAccount_Full)

### Example Usage

```go
package main

import(
	"context"
	"log"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
)

func main() {
    s := mycompanysamplesdk.New()

    ctx := context.Background()
    res, err := s.GeneralLedgerAccounts.PostGeneralLedgerAccounts(ctx, &shared.GeneralLedgerAccountPostDto{
        AccountNo: 3000,
        CurrencyCode: mycompanysamplesdk.String("USD"),
        DepartmentCode: mycompanysamplesdk.String("110"),
        IsActive: true,
        IsDepartmentRequired: mycompanysamplesdk.Bool(true),
        IsProjectRequired: mycompanysamplesdk.Bool(true),
        IsVatCodeLocked: mycompanysamplesdk.Bool(true),
        Name: "Kundefordringer",
        VatCode: "33",
    })
    if err != nil {
        log.Fatal(err)
    }

    if res.GeneralLedgerAccountDto != nil {
        // handle response
    }
}
```

### Parameters

| Parameter                                                                                | Type                                                                                     | Required                                                                                 | Description                                                                              |
| ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- | ---------------------------------------------------------------------------------------- |
| `ctx`                                                                                    | [context.Context](https://pkg.go.dev/context#Context)                                    | :heavy_check_mark:                                                                       | The context to use for the request.                                                      |
| `request`                                                                                | [shared.GeneralLedgerAccountPostDto](../../models/shared/generalledgeraccountpostdto.md) | :heavy_check_mark:                                                                       | The request object to use for the request.                                               |


### Response

**[*operations.PostGeneralLedgerAccountsResponse](../../models/operations/postgeneralledgeraccountsresponse.md), error**

