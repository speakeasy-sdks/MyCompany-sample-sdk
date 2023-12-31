// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/types"
	"github.com/speakeasy-sdks/MyCompany-sample-sdk/pkg/utils"
)

// VatBasisRatios - Custom VAT code basis ratios and valid from dates.
type VatBasisRatios struct {
	// The VAT basis ratio for a custom VatCode.
	VatBasisRatio *float64 `json:"VatBasisRatio,omitempty"`
	// The from date for the VAT basis ratio.
	VatBasisRatioValidFrom *types.Date `json:"VatBasisRatioValidFrom,omitempty"`
}

func (v VatBasisRatios) MarshalJSON() ([]byte, error) {
	return utils.MarshalJSON(v, "", false)
}

func (v *VatBasisRatios) UnmarshalJSON(data []byte) error {
	if err := utils.UnmarshalJSON(data, &v, "", false, false); err != nil {
		return err
	}
	return nil
}

func (o *VatBasisRatios) GetVatBasisRatio() *float64 {
	if o == nil {
		return nil
	}
	return o.VatBasisRatio
}

func (o *VatBasisRatios) GetVatBasisRatioValidFrom() *types.Date {
	if o == nil {
		return nil
	}
	return o.VatBasisRatioValidFrom
}
