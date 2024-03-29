// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// GeneralLedgerAccountPostDto - The General Ledger Account Post DTO contains only properties relevant to a general ledger account in Go that can be altered directly via POST-method in API.
type GeneralLedgerAccountPostDto struct {
	// The number of the account.
	AccountNo int64 `json:"AccountNo"`
	// Enum representing the different agriculture departments available for agriculture clients.<p>Members:</p><ul><li><i>None</i> - Unspecified</li><li><i>Finance</i> - Financial (Norsk: Finans)</li><li><i>Farming</i> - Farming (Norsk: Jordbruk)</li><li><i>Forestry</i> - Forestry (Norsk: Skogbruk)</li><li><i>FurAnimals</i> - Fur animals (Norsk: Pelsdyr)</li><li><i>OtherIndustries</i> - Other industries (Norsk: Andre næringer)</li><li><i>Private</i> - Private (Norsk: Privat)</li></ul>
	AgricultureDepartment *AgricultureDepartment `json:"AgricultureDepartment,omitempty"`
	// The standard code of the currency associated with the account.
	CurrencyCode *string `json:"CurrencyCode,omitempty"`
	// The code of the department associated with the account.
	DepartmentCode *string `json:"DepartmentCode,omitempty"`
	// The id of the department associated with the account.
	DepartmentID *int64 `json:"DepartmentId,omitempty"`
	// Value indicating whether this account is active.
	// Accounts that are not active will not be available for users in the journal entry GUI.
	IsActive bool `json:"IsActive"`
	// Value indicating whether department is a required property to set on transactions on this account.
	IsDepartmentRequired *bool `json:"IsDepartmentRequired,omitempty"`
	// Value indicating whether project is a required property to set on transactions on this account.
	IsProjectRequired *bool `json:"IsProjectRequired,omitempty"`
	// Value indicating whether the account is locked to a VAT code. If true, only VatCode can be used on transactions on this account.
	IsVatCodeLocked *bool `json:"IsVatCodeLocked,omitempty"`
	// The name of the account.
	Name string `json:"Name"`
	// Enum representing various unit of measures.<p>Members:</p><ul><li><i>BX</i> - Box (Norwegian: Eske)</li><li><i>MTR</i> - Meter (m) (Norwegian: Meter)</li><li><i>KMT</i> - Kilometer (km) (Norwegian: Kilometer)</li><li><i>KGM</i> - Kilogram (kg) (Norwegian: Kilogram)</li><li><i>EA</i> - Each (Norwegian: Stykke (stk))</li><li><i>LTR</i> - Liter (L) (Norwegian: Liter)</li><li><i>HUR</i> - Hour (Norwegian: Time)</li><li><i>DAY</i> - Day (Norwegian: Dag)</li><li><i>MTK</i> - Square meter (m^2) (Norwegian: Kvadratmeter)</li><li><i>MTQ</i> - Cubic meter (m^3) (Norwegian: Kubikkmeter)</li><li><i>TNE</i> - Metric ton (t) (Norwegian: Metrisk tonn)</li><li><i>MON</i> - Month (Norwegian: Måned)</li><li><i>ANN</i> - Year (Norwegian: År)</li><li><i>QAN</i> - Quarter of a year (Norwegian: Kvartal)</li><li><i>NL</i> - Load (Norwegian: Lass)</li><li><i>XRO</i> - Roll (Norwegian: Rull)</li><li><i>XBA</i> - Barrel (Norwegian: Tønne)</li><li><i>XPL</i> - Pail (Norwegian: Spann)</li><li><i>XPC</i> - Parcel (Norwegian: Kolli)</li><li><i>PR</i> - Pair (Norwegian: Par)</li><li><i>XCI</i> - Canister (Norwegian: Boks)</li><li><i>XBG</i> - Bag (Norwegian: Pose)</li><li><i>SET</i> - Set (Norwegian: Sett)</li><li><i>XTY</i> - Tank (Norwegian: Tank)</li><li><i>XOF</i> - Pallet (Norwegian: Pall)</li><li><i>FTK</i> - Square foot (ft^2) (Norwegian: Kvadratfot)</li><li><i>KWH</i> - Kilowatt hour (kWh) (Norwegian: Kilowattime)</li><li><i>MWH</i> - Megawatt hour (mWh) (Norwegian: Megawattime)</li><li><i>LBR</i> - Pound (lb) (Norwegian: Pund)</li><li><i>CMT</i> - Centimeter (cm) (Norwegian: Centimeter)</li><li><i>DMT</i> - Decimeter (dm) (Norwegian: Desimeter)</li><li><i>LM</i> - Linear meter (Norwegian: Lineær meter)</li><li><i>XPK</i> - Package (Norwegian: Pakke (pk))</li><li><i>GRM</i> - Gram (g) (Norwegian: Gram)</li><li><i>HGM</i> - Hectogram (hg) (Norwegian: Hektogram)</li><li><i>XFL</i> - Flask (Norwegian: Flaske)</li><li><i>XBE</i> - Bundle (Norwegian: Bunt)</li><li><i>E54</i> - Trip (Norwegian: Tur)</li><li><i>MMT</i> - Millimeter (mm) (Norwegian: Millimeter)</li><li><i>DAA</i> - Decare (Norwegian: Dekar)</li><li><i>H18</i> - Hectare (Norwegian: Hektar)</li><li><i>MLT</i> - Milliliter (mL) (Norwegian: Milliliter)</li><li><i>HLT</i> - Hectoliter (hL) (Norwegian: Hektoliter)</li><li><i>DLT</i> - Deciliter (dL) (Norwegian: Desiliter)</li><li><i>AK</i> - Fathom (Norwegian: Favn)</li><li><i>XCR</i> - Crate (Norwegian: Kasse)</li><li><i>E14</i> - Kilocalorie (kcal) (Norwegian: Kilokalori)</li><li><i>MJ</i> - Megajoule (MJ) (Norwegian: Megajoule)</li><li><i>J57</i> - Barrel (petroleum) (Norwegian: Fat)</li><li><i>XJG</i> - Jug (Norwegian: Kanne)</li><li><i>XCT</i> - Carton (Norwegian: Kartong)</li><li><i>XSA</i> - Sack (Norwegian: Sekk)</li><li><i>XTU</i> - Tube (Norwegian: Tube)</li><li><i>WEE</i> - Week (Norwegian: Uke)</li><li><i>XCA</i> - Can (Rectangular) (Norwegian: Boks (Rektangulær))</li><li><i>XCN</i> - Container (Norwegian: Konteiner)</li><li><i>NAR</i> - Number of articles (Norwegian: Antall artikler)</li><li><i>M4</i> - Monetary value (Norwegian: Pengeverdi)</li><li><i>XVQ</i> - Bulk (Norwegian: Bulk)</li><li><i>P1</i> - Percent (%) (Norwegian: Prosent)</li><li><i>MFU</i> - Milk Forage Unit (Norwegian: Forenhet melk (FEm))</li><li><i>KMK</i> - Square kilometer (km^2) (Norwegian: Kvadratkilometer)</li><li><i>LM3</i> - Loose cubic meter (Norwegian: Løskubikkmeter)</li><li><i>FOT</i> - Foot (ft) (Norwegian: Fot)</li><li><i>FM3</i> - Solid cubic meter (Norwegian: Fastkubikkmeter)</li></ul>
	Unit1 *UnitOfMeasureType `json:"Unit1,omitempty"`
	// Enum representing various unit of measures.<p>Members:</p><ul><li><i>BX</i> - Box (Norwegian: Eske)</li><li><i>MTR</i> - Meter (m) (Norwegian: Meter)</li><li><i>KMT</i> - Kilometer (km) (Norwegian: Kilometer)</li><li><i>KGM</i> - Kilogram (kg) (Norwegian: Kilogram)</li><li><i>EA</i> - Each (Norwegian: Stykke (stk))</li><li><i>LTR</i> - Liter (L) (Norwegian: Liter)</li><li><i>HUR</i> - Hour (Norwegian: Time)</li><li><i>DAY</i> - Day (Norwegian: Dag)</li><li><i>MTK</i> - Square meter (m^2) (Norwegian: Kvadratmeter)</li><li><i>MTQ</i> - Cubic meter (m^3) (Norwegian: Kubikkmeter)</li><li><i>TNE</i> - Metric ton (t) (Norwegian: Metrisk tonn)</li><li><i>MON</i> - Month (Norwegian: Måned)</li><li><i>ANN</i> - Year (Norwegian: År)</li><li><i>QAN</i> - Quarter of a year (Norwegian: Kvartal)</li><li><i>NL</i> - Load (Norwegian: Lass)</li><li><i>XRO</i> - Roll (Norwegian: Rull)</li><li><i>XBA</i> - Barrel (Norwegian: Tønne)</li><li><i>XPL</i> - Pail (Norwegian: Spann)</li><li><i>XPC</i> - Parcel (Norwegian: Kolli)</li><li><i>PR</i> - Pair (Norwegian: Par)</li><li><i>XCI</i> - Canister (Norwegian: Boks)</li><li><i>XBG</i> - Bag (Norwegian: Pose)</li><li><i>SET</i> - Set (Norwegian: Sett)</li><li><i>XTY</i> - Tank (Norwegian: Tank)</li><li><i>XOF</i> - Pallet (Norwegian: Pall)</li><li><i>FTK</i> - Square foot (ft^2) (Norwegian: Kvadratfot)</li><li><i>KWH</i> - Kilowatt hour (kWh) (Norwegian: Kilowattime)</li><li><i>MWH</i> - Megawatt hour (mWh) (Norwegian: Megawattime)</li><li><i>LBR</i> - Pound (lb) (Norwegian: Pund)</li><li><i>CMT</i> - Centimeter (cm) (Norwegian: Centimeter)</li><li><i>DMT</i> - Decimeter (dm) (Norwegian: Desimeter)</li><li><i>LM</i> - Linear meter (Norwegian: Lineær meter)</li><li><i>XPK</i> - Package (Norwegian: Pakke (pk))</li><li><i>GRM</i> - Gram (g) (Norwegian: Gram)</li><li><i>HGM</i> - Hectogram (hg) (Norwegian: Hektogram)</li><li><i>XFL</i> - Flask (Norwegian: Flaske)</li><li><i>XBE</i> - Bundle (Norwegian: Bunt)</li><li><i>E54</i> - Trip (Norwegian: Tur)</li><li><i>MMT</i> - Millimeter (mm) (Norwegian: Millimeter)</li><li><i>DAA</i> - Decare (Norwegian: Dekar)</li><li><i>H18</i> - Hectare (Norwegian: Hektar)</li><li><i>MLT</i> - Milliliter (mL) (Norwegian: Milliliter)</li><li><i>HLT</i> - Hectoliter (hL) (Norwegian: Hektoliter)</li><li><i>DLT</i> - Deciliter (dL) (Norwegian: Desiliter)</li><li><i>AK</i> - Fathom (Norwegian: Favn)</li><li><i>XCR</i> - Crate (Norwegian: Kasse)</li><li><i>E14</i> - Kilocalorie (kcal) (Norwegian: Kilokalori)</li><li><i>MJ</i> - Megajoule (MJ) (Norwegian: Megajoule)</li><li><i>J57</i> - Barrel (petroleum) (Norwegian: Fat)</li><li><i>XJG</i> - Jug (Norwegian: Kanne)</li><li><i>XCT</i> - Carton (Norwegian: Kartong)</li><li><i>XSA</i> - Sack (Norwegian: Sekk)</li><li><i>XTU</i> - Tube (Norwegian: Tube)</li><li><i>WEE</i> - Week (Norwegian: Uke)</li><li><i>XCA</i> - Can (Rectangular) (Norwegian: Boks (Rektangulær))</li><li><i>XCN</i> - Container (Norwegian: Konteiner)</li><li><i>NAR</i> - Number of articles (Norwegian: Antall artikler)</li><li><i>M4</i> - Monetary value (Norwegian: Pengeverdi)</li><li><i>XVQ</i> - Bulk (Norwegian: Bulk)</li><li><i>P1</i> - Percent (%) (Norwegian: Prosent)</li><li><i>MFU</i> - Milk Forage Unit (Norwegian: Forenhet melk (FEm))</li><li><i>KMK</i> - Square kilometer (km^2) (Norwegian: Kvadratkilometer)</li><li><i>LM3</i> - Loose cubic meter (Norwegian: Løskubikkmeter)</li><li><i>FOT</i> - Foot (ft) (Norwegian: Fot)</li><li><i>FM3</i> - Solid cubic meter (Norwegian: Fastkubikkmeter)</li></ul>
	Unit2 *UnitOfMeasureType `json:"Unit2,omitempty"`
	// The VAT code associated with the account. VatCode or VatCodeId must be supplied.
	// Standard codes in Go is based on the SAF-T standard.
	// Codes and any custom codes can be queried and identified using the VatCode service.
	VatCode *string `json:"VatCode,omitempty"`
	// The id of the VAT Code associated with the account.
	// VatCode or VatCodeId must be supplied when creating a general ledger account.
	VatCodeID *int `json:"VatCodeId,omitempty"`
	// Vat return specification used for giving extra vat information about the transaction. Will be reported from 2022 in the norwegian vat return.<p>Members:</p><ul><li><i>None</i> - None</li><li><i>Adjustment</i> - Adjustments (Norwegian: Justering)</li><li><i>LossesOnClaims</i> - Losses on claims (Norwegian: Tap på krav)</li><li><i>ReversalOfInputValueAddedTax</i> - Reversal of input value added tax (Norwegian: Tilbakeføring av inngående merverdiavgift)</li><li><i>Withdrawals</i> - Withdrawals (Norwegian: Uttak)</li><li><i>AdjustmentVatCompensationRealProperty</i> - Adjustment of value added tax (VAT) compensation for real estate (Norwegian: Justering av merverdiavgiftskompensasjon for fast eiendom)</li><li><i>PurchasesEligibleForCompensation</i> - Purchases eligible for compensation (Norwegian: Kjøp med kompensasjonsrett). This specification can not be used by the API directly when creating vouchers/account transactions, but will appear on account transactions on outgoing vat on purchases on client's eligible for compensation.</li></ul>
	VatReturnSpecification *VatReturnSpecification `json:"VatReturnSpecification,omitempty"`
}

func (o *GeneralLedgerAccountPostDto) GetAccountNo() int64 {
	if o == nil {
		return 0
	}
	return o.AccountNo
}

func (o *GeneralLedgerAccountPostDto) GetAgricultureDepartment() *AgricultureDepartment {
	if o == nil {
		return nil
	}
	return o.AgricultureDepartment
}

func (o *GeneralLedgerAccountPostDto) GetCurrencyCode() *string {
	if o == nil {
		return nil
	}
	return o.CurrencyCode
}

func (o *GeneralLedgerAccountPostDto) GetDepartmentCode() *string {
	if o == nil {
		return nil
	}
	return o.DepartmentCode
}

func (o *GeneralLedgerAccountPostDto) GetDepartmentID() *int64 {
	if o == nil {
		return nil
	}
	return o.DepartmentID
}

func (o *GeneralLedgerAccountPostDto) GetIsActive() bool {
	if o == nil {
		return false
	}
	return o.IsActive
}

func (o *GeneralLedgerAccountPostDto) GetIsDepartmentRequired() *bool {
	if o == nil {
		return nil
	}
	return o.IsDepartmentRequired
}

func (o *GeneralLedgerAccountPostDto) GetIsProjectRequired() *bool {
	if o == nil {
		return nil
	}
	return o.IsProjectRequired
}

func (o *GeneralLedgerAccountPostDto) GetIsVatCodeLocked() *bool {
	if o == nil {
		return nil
	}
	return o.IsVatCodeLocked
}

func (o *GeneralLedgerAccountPostDto) GetName() string {
	if o == nil {
		return ""
	}
	return o.Name
}

func (o *GeneralLedgerAccountPostDto) GetUnit1() *UnitOfMeasureType {
	if o == nil {
		return nil
	}
	return o.Unit1
}

func (o *GeneralLedgerAccountPostDto) GetUnit2() *UnitOfMeasureType {
	if o == nil {
		return nil
	}
	return o.Unit2
}

func (o *GeneralLedgerAccountPostDto) GetVatCode() *string {
	if o == nil {
		return nil
	}
	return o.VatCode
}

func (o *GeneralLedgerAccountPostDto) GetVatCodeID() *int {
	if o == nil {
		return nil
	}
	return o.VatCodeID
}

func (o *GeneralLedgerAccountPostDto) GetVatReturnSpecification() *VatReturnSpecification {
	if o == nil {
		return nil
	}
	return o.VatReturnSpecification
}
