// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/utils"
	"time"
)

// GeneralLedgerAccountDto - The General Ledger Account DTO contains all properties relevant to a GeneralLedgerAccount-object in Go
type GeneralLedgerAccountDto struct {
	// Gets or sets the account number.
	AccountNo int64 `json:"AccountNo"`
	// Gets or sets the standard code of the currency associated with the account.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// Gets the code of the department associated with the account.
	DepartmentCode *string `json:"DepartmentCode,omitempty"`
	// Gets the id of the department associated with the account.
	DepartmentID *int64 `json:"DepartmentId,omitempty"`
	// Gets the identifier. This identifier is unique and assigned by PowerOffice Go when a new entity is saved, and should be provided when an entity should be edited. If this identifier is not provided in POST, PowerOffice Go will try to create a new entity
	ID *int64 `json:"Id,omitempty"`
	// Gets or sets a value indicating whether this account is active. Accounts that are not active will not be available for users in the journal entry GUI
	IsActive bool `json:"IsActive"`
	// Gets or sets a value indicating whether department is a required property to set on transactions on this account
	IsDepartmentRequired *bool `json:"IsDepartmentRequired,omitempty"`
	// Gets or sets a value indicating whether project is a required property to set on transactions on this account
	IsProjectRequired *bool `json:"IsProjectRequired,omitempty"`
	// Gets or sets a value indicating whether the account is locked to the set vatcode. If true, only this vat code can be used on transactions on this account
	IsVatCodeLocked *bool `json:"IsVatCodeLocked,omitempty"`
	// Gets the timestamp of the last change of settings for this account.
	LastChangedDateTimeOffset *time.Time `json:"LastChangedDateTimeOffset,omitempty"`
	// Gets or sets the name of the account
	Name string `json:"Name"`
	// Gets or sets the vat code associated with the account.
	// Standard codes in Go is based on the SAF-T standard. Codes and any custom codes can be queried and identified using the VatCode service.
	VatCode string `json:"VatCode"`
	// Gets the id of the vatCode associated with the account. This property cannot be set, only retrieved from GET calls or in the response form POST/PATCH.
	VatCodeID *int64 `json:"VatCodeId,omitempty"`
	// Vat return specification used for giving extra vat information about the transaction. Will be reported from 2022 in the norwegian vat return.<p>Members:</p><ul><li><i>None</i> - None</li><li><i>Adjustment</i> - Adjustments (Norwegian: Justering)</li><li><i>LossesOnClaims</i> - Losses on claims (Norwegian: Tap på krav)</li><li><i>ReversalOfInputValueAddedTax</i> - Reversal of input value added tax (VAT) (Norwegian: Tilbakeføring av inngående merverdiavgift)</li><li><i>Withdrawals</i> - Withdrawals (Norwegian: Uttak)</li><li><i>AdjustmentVatCompensationRealProperty</i> - Adjustment of value added tax (VAT) compensation for real estate (Norwegian: Justering av merverdiavgiftskompensasjon for fast eiendom)</li><li><i>PurchasesEligibleForCompensation</i> - Purchases eligible for compensation (Norwegian: Kjøp med kompensasjonsrett)</li></ul>
	VatReturnSpecification *VatReturnSpecification `json:"VatReturnSpecification,omitempty"`
}

func (g GeneralLedgerAccountDto) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(g, "", false)
}

func (g *GeneralLedgerAccountDto) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &g, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *GeneralLedgerAccountDto) GetAccountNo() int64 {
	if o == nil {
		return 0
	}
	return o.AccountNo
}

func (o *GeneralLedgerAccountDto) GetCurrencyCode() *string {
	if o == nil {
		return nil
	}
	return o.CurrencyCode
}

func (o *GeneralLedgerAccountDto) GetDepartmentCode() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentCode
}

func (o *GeneralLedgerAccountDto) GetDepartmentID() *int64 {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *GeneralLedgerAccountDto) GetID() *int64 {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *GeneralLedgerAccountDto) GetIsActive() bool {
	if o == nil {
		return false
	}
	return o.IsActive
}

func (o *GeneralLedgerAccountDto) GetIsDepartmentRequired() *bool {
	if o == nil {
		return nil
	}
	return o.IsDepartmentRequired
}

func (o *GeneralLedgerAccountDto) GetIsProjectRequired() *bool {
	if o == nil {
		return nil
	}
	return o.IsProjectRequired
}

func (o *GeneralLedgerAccountDto) GetIsVatCodeLocked() *bool {
	if o == nil {
		return nil
	}
	return o.IsVatCodeLocked
}

func (o *GeneralLedgerAccountDto) GetLastChangedDateTimeOffset() *time.Time {
	if o == nil {
		return nil
	}
	return o.LastChangedDateTimeOffset
}

func (o *GeneralLedgerAccountDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GeneralLedgerAccountDto) GetVatCode() string {
	if o == nil {
		return ""
	}
	return o.VatCode
}

func (o *GeneralLedgerAccountDto) GetVatCodeID() *int64 {
	if o == nil {
		return nil
	}
	return o.VatCodeID
}

func (o *GeneralLedgerAccountDto) GetVatReturnSpecification() *VatReturnSpecification {
	if o == nil {
		return nil
	}
	return o.VatReturnSpecification
}
