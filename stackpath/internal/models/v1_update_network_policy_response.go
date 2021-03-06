// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// V1UpdateNetworkPolicyResponse A response from a request to update a network policy
// swagger:model v1UpdateNetworkPolicyResponse
type V1UpdateNetworkPolicyResponse struct {

	// network policy
	NetworkPolicy *V1NetworkPolicy `json:"networkPolicy,omitempty"`
}

// Validate validates this v1 update network policy response
func (m *V1UpdateNetworkPolicyResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateNetworkPolicy(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *V1UpdateNetworkPolicyResponse) validateNetworkPolicy(formats strfmt.Registry) error {

	if swag.IsZero(m.NetworkPolicy) { // not required
		return nil
	}

	if m.NetworkPolicy != nil {
		if err := m.NetworkPolicy.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("networkPolicy")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *V1UpdateNetworkPolicyResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *V1UpdateNetworkPolicyResponse) UnmarshalBinary(b []byte) error {
	var res V1UpdateNetworkPolicyResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
