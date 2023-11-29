// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/utils"
	"time"
)

// FinancialSettingsDto - Financial settings on a client.
type FinancialSettingsDto struct {
	// Gets the start date of processing in Go. This is the date the client started using Go. Transactions can only be posted on or after this date. The date can also be used to identify the date of the startbalance in Go, which will be this date -1 day
	ConversionDate *types.Date `json:"ConversionDate,omitempty"`
	// Gets the standard code of the currency the client use. As of now, only NOK is supported and this property will allways return NOK
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// Gets the general ledger account code of the account used for posting currency gains (agio). More information of accounts can be retrieved using the GeneralLedgerAccounts endpoint.
	CurrencyGainsAccountNo *int64 `json:"CurrencyGainsAccountNo,omitempty"`
	// Gets the general ledger account code of the accountused for posting currency losses (disagio). More information of accounts can be retrieved using the GeneralLedgerAccounts endpoint.
	CurrencyLossAccountNo *int64 `json:"CurrencyLossAccountNo,omitempty"`
	// Enum defining the months of a year.<p>Members:</p><ul><li><i>January</i> - January</li><li><i>February</i> - February</li><li><i>March</i> - March</li><li><i>April</i> - April</li><li><i>May</i> - May</li><li><i>June</i> - June</li><li><i>July</i> - July</li><li><i>August</i> - August</li><li><i>September</i> - September</li><li><i>October</i> - October</li><li><i>November</i> - November</li><li><i>December</i> - December</li></ul>
	FinancialYearEndMonth *Months `json:"FinancialYearEndMonth,omitempty"`
	// Gets the last changed timestamp. Last changed will update when the settings are created, or whenever the settings are changed.
	LastChangedDateTimeOffset *time.Time `json:"LastChangedDateTimeOffset,omitempty"`
	// Gets a flag indication whether the client use the trust account management functionality in Go.
	// Default to false, as the funcionality is optional for law firms, brokers and other firms subject of trust accounts.
	// If true, the client have dedicated trust bank accounts, and use projects in account transactions.
	UseTrustAccountManagement *bool `json:"UseTrustAccountManagement,omitempty"`
}

func (f FinancialSettingsDto) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(f, "", false)
}

func (f *FinancialSettingsDto) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &f, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *FinancialSettingsDto) GetConversionDate() *types.Date {
	if o == nil {
		return nil
	}
	return o.ConversionDate
}

func (o *FinancialSettingsDto) GetCurrencyCode() *string {
	if o == nil {
		return nil
	}
	return o.CurrencyCode
}

func (o *FinancialSettingsDto) GetCurrencyGainsAccountNo() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrencyGainsAccountNo
}

func (o *FinancialSettingsDto) GetCurrencyLossAccountNo() *int64 {
	if o == nil {
		return nil
	}
	return o.CurrencyLossAccountNo
}

func (o *FinancialSettingsDto) GetFinancialYearEndMonth() *Months {
	if o == nil {
		return nil
	}
	return o.FinancialYearEndMonth
}

func (o *FinancialSettingsDto) GetLastChangedDateTimeOffset() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastChangedDateTimeOffset
}

func (o *FinancialSettingsDto) GetUseTrustAccountManagement() *bool {
	if o == nil {
		return nil
	}
	return o.UseTrustAccountManagement
}
