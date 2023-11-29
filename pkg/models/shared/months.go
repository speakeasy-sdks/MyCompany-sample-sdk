// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"fmt"
)

// Months - Enum defining the months of a year.<p>Members:</p><ul><li><i>January</i> - January</li><li><i>February</i> - February</li><li><i>March</i> - March</li><li><i>April</i> - April</li><li><i>May</i> - May</li><li><i>June</i> - June</li><li><i>July</i> - July</li><li><i>August</i> - August</li><li><i>September</i> - September</li><li><i>October</i> - October</li><li><i>November</i> - November</li><li><i>December</i> - December</li></ul>
type Months string

const (
	MonthsJanuary   Months = "January"
	MonthsFebruary  Months = "February"
	MonthsMarch     Months = "March"
	MonthsApril     Months = "April"
	MonthsMay       Months = "May"
	MonthsJune      Months = "June"
	MonthsJuly      Months = "July"
	MonthsAugust    Months = "August"
	MonthsSeptember Months = "September"
	MonthsOctober   Months = "October"
	MonthsNovember  Months = "November"
	MonthsDecember  Months = "December"
)

func (e Months) ToPointer() *Months {
	return &e
}

func (e *Months) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "January":
		fallthrough
	case "February":
		fallthrough
	case "March":
		fallthrough
	case "April":
		fallthrough
	case "May":
		fallthrough
	case "June":
		fallthrough
	case "July":
		fallthrough
	case "August":
		fallthrough
	case "September":
		fallthrough
	case "October":
		fallthrough
	case "November":
		fallthrough
	case "December":
		*e = Months(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Months: %v", v)
	}
}
