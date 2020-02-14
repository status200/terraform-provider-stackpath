// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	"github.com/go-openapi/errors"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// PrometheusMetrics A collection of metrics
// swagger:model prometheusMetrics
type PrometheusMetrics struct {

	// data
	Data *PrometheusMetricsData `json:"data,omitempty"`

	// The error encountered when querying for metrics
	Error string `json:"error,omitempty"`

	// The type of error encountered when querying for metrics
	ErrorType string `json:"errorType,omitempty"`

	// A metrics query's resulting status
	// Enum: [SUCCESS ERROR]
	Status *string `json:"status,omitempty"`

	// Warnings encountered when querying for metrics
	Warnings []string `json:"warnings"`
}

// Validate validates this prometheus metrics
func (m *PrometheusMetrics) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateData(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateStatus(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *PrometheusMetrics) validateData(formats strfmt.Registry) error {

	if swag.IsZero(m.Data) { // not required
		return nil
	}

	if m.Data != nil {
		if err := m.Data.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("data")
			}
			return err
		}
	}

	return nil
}

var prometheusMetricsTypeStatusPropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["SUCCESS","ERROR"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		prometheusMetricsTypeStatusPropEnum = append(prometheusMetricsTypeStatusPropEnum, v)
	}
}

const (

	// PrometheusMetricsStatusSUCCESS captures enum value "SUCCESS"
	PrometheusMetricsStatusSUCCESS string = "SUCCESS"

	// PrometheusMetricsStatusERROR captures enum value "ERROR"
	PrometheusMetricsStatusERROR string = "ERROR"
)

// prop value enum
func (m *PrometheusMetrics) validateStatusEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, prometheusMetricsTypeStatusPropEnum); err != nil {
		return err
	}
	return nil
}

func (m *PrometheusMetrics) validateStatus(formats strfmt.Registry) error {

	if swag.IsZero(m.Status) { // not required
		return nil
	}

	// value enum
	if err := m.validateStatusEnum("status", "body", *m.Status); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *PrometheusMetrics) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *PrometheusMetrics) UnmarshalBinary(b []byte) error {
	var res PrometheusMetrics
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}