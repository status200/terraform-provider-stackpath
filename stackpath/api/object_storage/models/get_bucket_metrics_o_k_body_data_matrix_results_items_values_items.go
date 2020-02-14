// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// GetBucketMetricsOKBodyDataMatrixResultsItemsValuesItems An individual metric data point
// swagger:model getBucketMetricsOKBodyDataMatrixResultsItemsValuesItems
type GetBucketMetricsOKBodyDataMatrixResultsItemsValuesItems struct {

	// The time that a data point was recorded
	UnixTime string `json:"unixTime,omitempty"`

	// A data point's value
	Value string `json:"value,omitempty"`
}

// Validate validates this get bucket metrics o k body data matrix results items values items
func (m *GetBucketMetricsOKBodyDataMatrixResultsItemsValuesItems) Validate(formats strfmt.Registry) error {
	return nil
}

// MarshalBinary interface implementation
func (m *GetBucketMetricsOKBodyDataMatrixResultsItemsValuesItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetBucketMetricsOKBodyDataMatrixResultsItemsValuesItems) UnmarshalBinary(b []byte) error {
	var res GetBucketMetricsOKBodyDataMatrixResultsItemsValuesItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}