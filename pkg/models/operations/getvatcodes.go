// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetVatCodesRequest struct {
	// Include expired VAT codes. Expired codes has a ValidTo date and can still be used for accounting events in the validFrom - validTo timeframe as long as it is active.
	IncludeExpired *bool `queryParam:"style=form,explode=true,name=includeExpired"`
	// Returns both active and inactive codes as default. True returns active codes, False returns only inactive codes.
	IsActive *bool `queryParam:"style=form,explode=true,name=isActive"`
	// Structure containing various resource-filter options
	ResourceParameter *shared.ResourceParametersInput `queryParam:"style=form,explode=true,name=resourceParameter"`
}

func (o *GetVatCodesRequest) GetIncludeExpired() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeExpired
}

func (o *GetVatCodesRequest) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *GetVatCodesRequest) GetResourceParameter() *shared.ResourceParametersInput {
	if o == nil {
		return nil
	}
	return o.ResourceParameter
}

type GetVatCodesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Given when resource returned successfully.
	VatCodeDtos []shared.VatCodeDto
}

func (o *GetVatCodesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetVatCodesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetVatCodesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetVatCodesResponse) GetVatCodeDtos() []shared.VatCodeDto {
	if o == nil {
		return nil
	}
	return o.VatCodeDtos
}