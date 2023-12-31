// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetCurrenciesRequest struct {
	// Filter by currency code.
	CurrencyCode *string `queryParam:"style=form,explode=true,name=currencyCode"`
	// Filter by isActive parameter..
	IsActive *bool `queryParam:"style=form,explode=true,name=isActive"`
	// Structure containing various resource-filter options
	ResourceParameter *shared.ResourceParameters `queryParam:"style=form,explode=true,name=resourceParameter"`
}

func (o *GetCurrenciesRequest) GetCurrencyCode() *string {
	if o == nil {
		return nil
	}
	return o.CurrencyCode
}

func (o *GetCurrenciesRequest) GetIsActive() *bool {
	if o == nil {
		return nil
	}
	return o.IsActive
}

func (o *GetCurrenciesRequest) GetResourceParameter() *shared.ResourceParameters {
	if o == nil {
		return nil
	}
	return o.ResourceParameter
}

type GetCurrenciesResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// Given when resource returned successfully
	Classes []shared.CurrencyDto
}

func (o *GetCurrenciesResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetCurrenciesResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetCurrenciesResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetCurrenciesResponse) GetClasses() []shared.CurrencyDto {
	if o == nil {
		return nil
	}
	return o.Classes
}
