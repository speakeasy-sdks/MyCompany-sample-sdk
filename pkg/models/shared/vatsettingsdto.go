// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type VatSettingsDto struct {
	// Gets information on whether this this client is registered for Value Added Tax (VAT, norwegian: mva-registrert). If false, vat is not applied to outgoing invoices.
	IsVatRegistered *bool `json:"IsVatRegistered,omitempty"`
	// Gets information on whether this client report vat returns manually. True if the client are obliged to submit more than one type of vat return, are jointly registered or shall deliver vat return for vat compensation or reverse tax liability.
	ReportVatReturnManually *bool `json:"ReportVatReturnManually,omitempty"`
	// Gets information on whether this client are entitled to and uses vat compensation. If true, vat codes with first letter notaion K can be used when relevant for the client.
	UseVatCompensation *bool `json:"UseVatCompensation,omitempty"`
	// Gets the vat number of the client. The vat number will equal the organization number for norwegian clients, but will relevant in vat terms only if the property isVatRegistered is true.
	VatNumber *string `json:"VatNumber,omitempty"`
	// Enum defining the Value Added Tax (VAT) period on this client. The VAT period states how often VAT is reported to the government from this client.<p>Members:</p><ul><li><i>none</i> - No defined VAT period</li><li><i>weekly</i> - Every week</li><li><i>halfMonthly</i> - Half-Monthly</li><li><i>monthly</i> - Every month</li><li><i>biMonthly</i> - Every two months</li><li><i>yearly</i> - Once a year</li><li><i>halfYearly</i> - Twice a year</li><li><i>quarterly</i> - Four times a year</li></ul>
	VatPeriod *VatPeriodType `json:"VatPeriod,omitempty"`
	// Maps to VatReturnCategory<p>Members:</p><ul><li><i>generalIndustry</i> - Alminnelig næring</li><li><i>primaryIndustry</i> - Primærnmæring</li><li><i>vatFeeCompensation</i> - Merverdiavgiftskompensasjon</li><li><i>reverseVatObligation</i> - Omvendt avgiftsplikt</li><li><i>noReporting</i> - Ingen rapportering</li></ul>
	VatReturnType *VatReturnType `json:"VatReturnType,omitempty"`
}

func (o *VatSettingsDto) GetIsVatRegistered() *bool {
	if o == nil {
		return nil
	}
	return o.IsVatRegistered
}

func (o *VatSettingsDto) GetReportVatReturnManually() *bool {
	if o == nil {
		return nil
	}
	return o.ReportVatReturnManually
}

func (o *VatSettingsDto) GetUseVatCompensation() *bool {
	if o == nil {
		return nil
	}
	return o.UseVatCompensation
}

func (o *VatSettingsDto) GetVatNumber() *string {
	if o == nil {
		return nil
	}
	return o.VatNumber
}

func (o *VatSettingsDto) GetVatPeriod() *VatPeriodType {
	if o == nil {
		return nil
	}
	return o.VatPeriod
}

func (o *VatSettingsDto) GetVatReturnType() *VatReturnType {
	if o == nil {
		return nil
	}
	return o.VatReturnType
}
