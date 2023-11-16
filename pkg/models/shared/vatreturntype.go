// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// VatReturnType - The type of VAT returns the client should deliver for it's principal activity.<p>Members:</p><ul><li><i>generalIndustry</i> - General Industry (Norwegian: Alminnelig næring)</li><li><i>primaryIndustry</i> - Primary Industry (Norwegian: Primærnmæring)</li><li><i>vatCompensation</i> - Vat Compensation (Norwegian: Merverdiavgiftskompensasjon)ø</li><li><i>reverseVatLiability</i> - Reverse VAT Liability (Norwegian: Omvendt avgiftsplikt)</li><li><i>noReporting</i> - No reporting (Norwegian: Ingen rapportering)</li></ul>
type VatReturnType string

const (
	VatReturnTypeGeneralIndustry     VatReturnType = "generalIndustry"
	VatReturnTypePrimaryIndustry     VatReturnType = "primaryIndustry"
	VatReturnTypeVatCompensation     VatReturnType = "vatCompensation"
	VatReturnTypeReverseVatLiability VatReturnType = "reverseVatLiability"
	VatReturnTypeNoReporting         VatReturnType = "noReporting"
)

func (e VatReturnType) ToPointer() *VatReturnType {
	return &e
}

func (e *VatReturnType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "generalIndustry":
		fallthrough
	case "primaryIndustry":
		fallthrough
	case "vatCompensation":
		fallthrough
	case "reverseVatLiability":
		fallthrough
	case "noReporting":
		*e = VatReturnType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VatReturnType: %v", v)
	}
}
