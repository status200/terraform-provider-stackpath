// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// DataVectorResultValue An individual metric data point
// swagger:model dataVectorResultValue
type DataVectorResultValue struct {

	// The time that a data point was recorded
	UnixTime string `json:"unixTime,omitempty"`

	// A data point's value
	Value string `json:"value,omitempty"`
}

// Validate validates this data vector result value
func (m *DataVectorResultValue) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *DataVectorResultValue) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *DataVectorResultValue) UnmarshalBinary(b []byte) error {
	var res DataVectorResultValue
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}