// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetLockDateResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// Given when resource was successfully returned.
	LockDateSettingsDto *shared.LockDateSettingsDto
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetLockDateResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetLockDateResponse) GetLockDateSettingsDto() *shared.LockDateSettingsDto {
	if o == nil {
		return nil
	}
	return o.LockDateSettingsDto
}

func (o *GetLockDateResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetLockDateResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
