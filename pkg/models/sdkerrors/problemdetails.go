// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package sdkerrors

import (
	"encoding/json"
)

// ProblemDetails - Given when request is badly formatted
type ProblemDetails struct {
	AdditionalProperties map[string]interface{} `additionalProperties:"true" json:"-"`
	Detail               *string                `json:"Detail,omitempty"`
	Instance             *string                `json:"Instance,omitempty"`
	Status               *int                   `json:"Status,omitempty"`
	Title                *string                `json:"Title,omitempty"`
	Type                 *string                `json:"Type,omitempty"`
}

var _ error = &ProblemDetails{}

func (e *ProblemDetails) Error() string {
	data, _ := json.Marshal(e)
	return string(data)
}
