// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

// SubLedgerNumberSeriesDto - The subledger number series DTO contains all properties relevant to subledger number series.
type SubLedgerNumberSeriesDto struct {
	// The minimum number the subledger accounts in this series can have
	FromInclusive *int64 `json:"FromInclusive,omitempty"`
	// Gets the underlying general ledger account ID of the subledger series.
	GeneralLedgerAccountID *int64 `json:"GeneralLedgerAccountId,omitempty"`
	// Gets the underlying general ledger account code of this subledger series.
	// Subledger transactions in Go will be posted using the subledger number, but the overall accounting effect will be on this underlying account.
	// More information of accounts can be retrieved using the GeneralLedgerAccount service.
	GeneralLedgerAccountNo *int64 `json:"GeneralLedgerAccountNo,omitempty"`
	// Gets the unique identifier of the subledger number series.
	ID *string `json:"Id,omitempty"`
	// Gets a value indicating whether this number series is used by accounts that contain client trust funds.
	// Client trust accounts are accounts where realtors or lawyers, for instance, handles their clients' money.
	// Can be true if the client have the FinancialSettings.UseTrustAccountManagement set true.
	IsClientTrust *bool `json:"IsClientTrust,omitempty"`
	// Gets a flag indicating whether this instance is the default number series for that subledger numberSeriesType.
	IsDefault *bool `json:"IsDefault,omitempty"`
	// Gets the name of the subledger number series.
	Name *string `json:"Name,omitempty"`
	// Enum SubLedgerNumberSeriesType indicating what kind of sub ledger accounts are represented in the number series.<p>Members:</p><ul><li><i>Customer</i> - The sub ledger number series contains Customer</li><li><i>Supplier</i> - The sub ledger number series contains Supplier</li><li><i>Employee</i> - The sub ledger number series contains Employee</li></ul>
	SubLedgerNumberSeriesType *SubLedgerNumberSeriesType `json:"SubLedgerNumberSeriesType,omitempty"`
	// The maximum number the subledger accounts in this series can have.
	ToInclusive *int64 `json:"ToInclusive,omitempty"`
}

func (o *SubLedgerNumberSeriesDto) GetFromInclusive() *int64 {
	if o == nil {
		return nil
	}
	return o.FromInclusive
}

func (o *SubLedgerNumberSeriesDto) GetGeneralLedgerAccountID() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralLedgerAccountID
}

func (o *SubLedgerNumberSeriesDto) GetGeneralLedgerAccountNo() *int64 {
	if o == nil {
		return nil
	}
	return o.GeneralLedgerAccountNo
}

func (o *SubLedgerNumberSeriesDto) GetID() *string {
	if o == nil {
		return nil
	}
	return o.ID
}

func (o *SubLedgerNumberSeriesDto) GetIsClientTrust() *bool {
	if o == nil {
		return nil
	}
	return o.IsClientTrust
}

func (o *SubLedgerNumberSeriesDto) GetIsDefault() *bool {
	if o == nil {
		return nil
	}
	return o.IsDefault
}

func (o *SubLedgerNumberSeriesDto) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *SubLedgerNumberSeriesDto) GetSubLedgerNumberSeriesType() *SubLedgerNumberSeriesType {
	if o == nil {
		return nil
	}
	return o.SubLedgerNumberSeriesType
}

func (o *SubLedgerNumberSeriesDto) GetToInclusive() *int64 {
	if o == nil {
		return nil
	}
	return o.ToInclusive
}
