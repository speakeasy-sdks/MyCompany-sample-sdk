// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package mycompanysamplesdk

import (
	"bytes"
	"context"
	"fmt"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/operations"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/sdkerrors"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/models/shared"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/utils"
	"io"
	"net/http"
	"strings"
)

type LockDateSettings struct {
	sdkConfiguration sdkConfiguration
}

func newLockDateSettings(sdkConfig sdkConfiguration) *LockDateSettings {
	return &LockDateSettings{
		sdkConfiguration: sdkConfig,
	}
}

// GetLockDate - Gets the lock date on the client. (Auth roles: CommonServices,CommonServices_Full)
func (s *LockDateSettings) GetLockDate(ctx context.Context) (*operations.GetLockDateResponse, error) {
	baseURL := utils.ReplaceParameters(s.sdkConfiguration.GetServerDetails())
	url := strings.TrimSuffix(baseURL, "/") + "/LockDateSettings"

	req, err := http.NewRequestWithContext(ctx, "GET", url, nil)
	if err != nil {
		return nil, fmt.Errorf("error creating request: %w", err)
	}
	req.Header.Set("Accept", "application/json")
	req.Header.Set("user-agent", s.sdkConfiguration.UserAgent)

	client := s.sdkConfiguration.DefaultClient

	httpRes, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("error sending request: %w", err)
	}
	if httpRes == nil {
		return nil, fmt.Errorf("error sending request: no response")
	}

	contentType := httpRes.Header.Get("Content-Type")

	res := &operations.GetLockDateResponse{
		StatusCode:  httpRes.StatusCode,
		ContentType: contentType,
		RawResponse: httpRes,
	}

	rawBody, err := io.ReadAll(httpRes.Body)
	if err != nil {
		return nil, fmt.Errorf("error reading response body: %w", err)
	}
	httpRes.Body.Close()
	httpRes.Body = io.NopCloser(bytes.NewBuffer(rawBody))
	switch {
	case httpRes.StatusCode == 200:
		switch {
		case utils.MatchContentType(contentType, `application/json`):
			var out shared.LockDateSettingsDto
			if err := utils.UnmarshalJsonFromResponseBody(bytes.NewBuffer(rawBody), &out, ""); err != nil {
				return nil, err
			}

			res.LockDateSettingsDto = &out
		default:
			return nil, sdkerrors.NewSDKError(fmt.Sprintf("unknown content-type received: %s", contentType), httpRes.StatusCode, string(rawBody), httpRes)
		}
	case httpRes.StatusCode == 401:
		fallthrough
	case httpRes.StatusCode == 403:
		fallthrough
	case httpRes.StatusCode >= 400 && httpRes.StatusCode < 500:
		fallthrough
	case httpRes.StatusCode >= 500 && httpRes.StatusCode < 600:
		return nil, sdkerrors.NewSDKError("API error occurred", httpRes.StatusCode, string(rawBody), httpRes)
	}

	return res, nil
}
