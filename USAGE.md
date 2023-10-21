<!-- Start SDK Example Usage -->


```go
package main

import (
	"context"
	mycompanysamplesdk "github.com/speakeasy-sdks/MyCompany-sample-sdk"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"log"
)

func main() {
	s := mycompanysamplesdk.New()

	ctx := context.Background()
	res, err := s.Currencies.GetCurrencies(ctx, operations.GetCurrenciesRequest{
		ResourceParameter: &shared.ResourceParametersInput{
			Fields:     mycompanysamplesdk.String("string"),
			OrderBy:    mycompanysamplesdk.String("string"),
			PageNumber: mycompanysamplesdk.Int(2),
			PageSize:   mycompanysamplesdk.Int(5000),
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
<!-- End SDK Example Usage -->