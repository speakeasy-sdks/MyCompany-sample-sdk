// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package mycompanysamplesdk

import (
	"fmt"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/internal/hooks"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/utils"
	"net/http"
	"time"
)

// ServerList contains the list of servers available to the SDK
var ServerList = []string{
	// API Management PROD
	"https://goapi.poweroffice.net/v2",
}

// HTTPClient provides an interface for suplying the SDK with a custom HTTP client
type HTTPClient interface {
	Do(req *http.Request) (*http.Response, error)
}

// String provides a helper function to return a pointer to a string
func String(s string) *string { return &s }

// Bool provides a helper function to return a pointer to a bool
func Bool(b bool) *bool { return &b }

// Int provides a helper function to return a pointer to an int
func Int(i int) *int { return &i }

// Int64 provides a helper function to return a pointer to an int64
func Int64(i int64) *int64 { return &i }

// Float32 provides a helper function to return a pointer to a float32
func Float32(f float32) *float32 { return &f }

// Float64 provides a helper function to return a pointer to a float64
func Float64(f float64) *float64 { return &f }

type sdkConfiguration struct {
	Client HTTPClient

	ServerURL         string
	ServerIndex       int
	Language          string
	OpenAPIDocVersion string
	SDKVersion        string
	GenVersion        string
	UserAgent         string
	RetryConfig       *utils.RetryConfig
	Hooks             *hooks.Hooks
}

func (c *sdkConfiguration) GetServerDetails() (string, map[string]string) {
	if c.ServerURL != "" {
		return c.ServerURL, nil
	}

	return ServerList[c.ServerIndex], nil
}

// AccountingSettings - Accounting Settings: Service to query client accounting settings.
//
//	**Required access privileges:**
//
//	+ *CommonServices*
//	 + *read/full*: permission to GET
//
//	+ *GeneralLedgerAccount*
//	  + *read*: permission to GET
//	  + *full*: permission to DELETE, GET, PATCH and POST
//
//	**Base url:** *https://goapi.poweroffice.net/v2/*
type AccountingSettings struct {
	Currencies            *Currencies
	FinancialSettings     *FinancialSettings
	GeneralLedgerAccounts *GeneralLedgerAccounts
	LockDateSettings      *LockDateSettings
	SubLedgerNumberSeries *SubLedgerNumberSeries
	VatCodes              *VatCodes
	VatSettings           *VatSettings

	sdkConfiguration sdkConfiguration
}

type SDKOption func(*AccountingSettings)

// WithServerURL allows the overriding of the default server URL
func WithServerURL(serverURL string) SDKOption {
	return func(sdk *AccountingSettings) {
		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithTemplatedServerURL allows the overriding of the default server URL with a templated URL populated with the provided parameters
func WithTemplatedServerURL(serverURL string, params map[string]string) SDKOption {
	return func(sdk *AccountingSettings) {
		if params != nil {
			serverURL = utils.ReplaceParameters(serverURL, params)
		}

		sdk.sdkConfiguration.ServerURL = serverURL
	}
}

// WithServerIndex allows the overriding of the default server by index
func WithServerIndex(serverIndex int) SDKOption {
	return func(sdk *AccountingSettings) {
		if serverIndex < 0 || serverIndex >= len(ServerList) {
			panic(fmt.Errorf("server index %d out of range", serverIndex))
		}

		sdk.sdkConfiguration.ServerIndex = serverIndex
	}
}

// WithClient allows the overriding of the default HTTP client used by the SDK
func WithClient(client HTTPClient) SDKOption {
	return func(sdk *AccountingSettings) {
		sdk.sdkConfiguration.Client = client
	}
}

func WithRetryConfig(retryConfig utils.RetryConfig) SDKOption {
	return func(sdk *AccountingSettings) {
		sdk.sdkConfiguration.RetryConfig = &retryConfig
	}
}

// New creates a new instance of the SDK with the provided options
func New(opts ...SDKOption) *AccountingSettings {
	sdk := &AccountingSettings{
		sdkConfiguration: sdkConfiguration{
			Language:          "go",
			OpenAPIDocVersion: "v2",
			SDKVersion:        "0.10.0",
			GenVersion:        "2.279.1",
			UserAgent:         "speakeasy-sdk/go 0.10.0 2.279.1 v2 github.com/speakeasy-sdks/MyCompany-sample-sdk",
			Hooks:             hooks.New(),
		},
	}
	for _, opt := range opts {
		opt(sdk)
	}

	// Use WithClient to override the default client if you would like to customize the timeout
	if sdk.sdkConfiguration.Client == nil {
		sdk.sdkConfiguration.Client = &http.Client{Timeout: 60 * time.Second}
	}

	currentServerURL, _ := sdk.sdkConfiguration.GetServerDetails()
	serverURL := currentServerURL
	serverURL, sdk.sdkConfiguration.Client = sdk.sdkConfiguration.Hooks.SDKInit(currentServerURL, sdk.sdkConfiguration.Client)
	if serverURL != currentServerURL {
		sdk.sdkConfiguration.ServerURL = serverURL
	}

	sdk.Currencies = newCurrencies(sdk.sdkConfiguration)

	sdk.FinancialSettings = newFinancialSettings(sdk.sdkConfiguration)

	sdk.GeneralLedgerAccounts = newGeneralLedgerAccounts(sdk.sdkConfiguration)

	sdk.LockDateSettings = newLockDateSettings(sdk.sdkConfiguration)

	sdk.SubLedgerNumberSeries = newSubLedgerNumberSeries(sdk.sdkConfiguration)

	sdk.VatCodes = newVatCodes(sdk.sdkConfiguration)

	sdk.VatSettings = newVatSettings(sdk.sdkConfiguration)

	return sdk
}
