// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetGeneralLedgerAccountByIDRequest struct {
	ID int64 `pathParam:"style=simple,explode=false,name=id"`
}

func (o *GetGeneralLedgerAccountByIDRequest) GetID() int64 {
	if o == nil {
		return 0
	}
	return o.ID
}

type GetGeneralLedgerAccountByIDResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Given when resource returned successfully
	GeneralLedgerAccountDto *shared.GeneralLedgerAccountDto
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetGeneralLedgerAccountByIDResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetGeneralLedgerAccountByIDResponse) GetGeneralLedgerAccountDto() *shared.GeneralLedgerAccountDto {
	if o == nil {
		return nil
	}
	return o.GeneralLedgerAccountDto
}

func (o *GetGeneralLedgerAccountByIDResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetGeneralLedgerAccountByIDResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
