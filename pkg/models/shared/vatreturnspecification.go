// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// VatReturnSpecification - Vat return specification used for giving extra vat information about the transaction. Will be reported from 2022 in the norwegian vat return.<p>Members:</p><ul><li><i>None</i> - None</li><li><i>Adjustment</i> - Adjustments (Norwegian: Justering)</li><li><i>LossesOnClaims</i> - Losses on claims (Norwegian: Tap på krav)</li><li><i>ReversalOfInputValueAddedTax</i> - Reversal of input value added tax (Norwegian: Tilbakeføring av inngående merverdiavgift)</li><li><i>Withdrawals</i> - Withdrawals (Norwegian: Uttak)</li><li><i>AdjustmentVatCompensationRealProperty</i> - Adjustment of value added tax (VAT) compensation for real estate (Norwegian: Justering av merverdiavgiftskompensasjon for fast eiendom)</li><li><i>PurchasesEligibleForCompensation</i> - Purchases eligible for compensation (Norwegian: Kjøp med kompensasjonsrett). This specification can not be used by the API directly when creating vouchers/account transactions, but will appear on account transactions on outgoing vat on purchases on client's eligible for compensation.</li></ul>
type VatReturnSpecification string

const (
	VatReturnSpecificationNone                                  VatReturnSpecification = "None"
	VatReturnSpecificationAdjustment                            VatReturnSpecification = "Adjustment"
	VatReturnSpecificationLossesOnClaims                        VatReturnSpecification = "LossesOnClaims"
	VatReturnSpecificationReversalOfInputValueAddedTax          VatReturnSpecification = "ReversalOfInputValueAddedTax"
	VatReturnSpecificationWithdrawals                           VatReturnSpecification = "Withdrawals"
	VatReturnSpecificationAdjustmentVatCompensationRealProperty VatReturnSpecification = "AdjustmentVatCompensationRealProperty"
	VatReturnSpecificationPurchasesEligibleForCompensation      VatReturnSpecification = "PurchasesEligibleForCompensation"
)

func (e VatReturnSpecification) ToPointer() *VatReturnSpecification {
	return &e
}

func (e *VatReturnSpecification) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "None":
		fallthrough
	case "Adjustment":
		fallthrough
	case "LossesOnClaims":
		fallthrough
	case "ReversalOfInputValueAddedTax":
		fallthrough
	case "Withdrawals":
		fallthrough
	case "AdjustmentVatCompensationRealProperty":
		fallthrough
	case "PurchasesEligibleForCompensation":
		*e = VatReturnSpecification(v)
		return nil
	default:
		return fmt.Errorf("invalid value for VatReturnSpecification: %v", v)
	}
}
