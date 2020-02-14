// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"strconv"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
)

// CreateBucketUnauthorizedBody create bucket unauthorized body
// swagger:model createBucketUnauthorizedBody
type CreateBucketUnauthorizedBody struct {

	// code
	Code int32 `json:"code,omitempty"`

	detailsField []CreateBucketUnauthorizedBodyDetailsItems

	// error
	Error string `json:"error,omitempty"`
}

// Details gets the details of this base type
func (m *CreateBucketUnauthorizedBody) Details() []CreateBucketUnauthorizedBodyDetailsItems {
	return m.detailsField
}

// SetDetails sets the details of this base type
func (m *CreateBucketUnauthorizedBody) SetDetails(val []CreateBucketUnauthorizedBodyDetailsItems) {
	m.detailsField = val
}

// UnmarshalJSON unmarshals this object with a polymorphic type from a JSON structure
func (m *CreateBucketUnauthorizedBody) UnmarshalJSON(raw []byte) error {
	var data struct {
		Code int32 `json:"code,omitempty"`

		Details json.RawMessage `json:"details"`

		Error string `json:"error,omitempty"`
	}
	buf := bytes.NewBuffer(raw)
	dec := json.NewDecoder(buf)
	dec.UseNumber()

	if err := dec.Decode(&data); err != nil {
		return err
	}

	var propDetails []CreateBucketUnauthorizedBodyDetailsItems
	if string(data.Details) != "null" {
		details, err := UnmarshalCreateBucketUnauthorizedBodyDetailsItemsSlice(bytes.NewBuffer(data.Details), runtime.JSONConsumer())
		if err != nil && err != io.EOF {
			return err
		}
		propDetails = details
	}

	var result CreateBucketUnauthorizedBody

	// code
	result.Code = data.Code

	// details
	result.detailsField = propDetails

	// error
	result.Error = data.Error

	*m = result

	return nil
}

// MarshalJSON marshals this object with a polymorphic type to a JSON structure
func (m CreateBucketUnauthorizedBody) MarshalJSON() ([]byte, error) {
	var b1, b2, b3 []byte
	var err error
	b1, err = json.Marshal(struct {
		Code int32 `json:"code,omitempty"`

		Error string `json:"error,omitempty"`
	}{

		Code: m.Code,

		Error: m.Error,
	},
	)
	if err != nil {
		return nil, err
	}
	b2, err = json.Marshal(struct {
		Details []CreateBucketUnauthorizedBodyDetailsItems `json:"details"`
	}{

		Details: m.detailsField,
	},
	)
	if err != nil {
		return nil, err
	}

	return swag.ConcatJSON(b1, b2, b3), nil
}

// Validate validates this create bucket unauthorized body
func (m *CreateBucketUnauthorizedBody) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDetails(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *CreateBucketUnauthorizedBody) validateDetails(formats strfmt.Registry) error {

	if swag.IsZero(m.Details()) { // not required
		return nil
	}

	for i := 0; i < len(m.Details()); i++ {

		if err := m.detailsField[i].Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("details" + "." + strconv.Itoa(i))
			}
			return err
		}

	}

	return nil
}

// MarshalBinary interface implementation
func (m *CreateBucketUnauthorizedBody) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *CreateBucketUnauthorizedBody) UnmarshalBinary(b []byte) error {
	var res CreateBucketUnauthorizedBody
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}