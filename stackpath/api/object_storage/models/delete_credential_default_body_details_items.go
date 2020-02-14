// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"bytes"
	"encoding/json"
	"io"
	"io/ioutil"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	strfmt "github.com/go-openapi/strfmt"
	"github.com/go-openapi/validate"
)

// DeleteCredentialDefaultBodyDetailsItems delete credential default body details items
// swagger:discriminator deleteCredentialDefaultBodyDetailsItems @type
type DeleteCredentialDefaultBodyDetailsItems interface {
	runtime.Validatable

	// at type
	// Required: true
	AtType() string
	SetAtType(string)
}

type deleteCredentialDefaultBodyDetailsItems struct {
	atTypeField string
}

// AtType gets the at type of this polymorphic type
func (m *deleteCredentialDefaultBodyDetailsItems) AtType() string {
	return "deleteCredentialDefaultBodyDetailsItems"
}

// SetAtType sets the at type of this polymorphic type
func (m *deleteCredentialDefaultBodyDetailsItems) SetAtType(val string) {

}

// UnmarshalDeleteCredentialDefaultBodyDetailsItemsSlice unmarshals polymorphic slices of DeleteCredentialDefaultBodyDetailsItems
func UnmarshalDeleteCredentialDefaultBodyDetailsItemsSlice(reader io.Reader, consumer runtime.Consumer) ([]DeleteCredentialDefaultBodyDetailsItems, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []DeleteCredentialDefaultBodyDetailsItems
	for _, element := range elements {
		obj, err := unmarshalDeleteCredentialDefaultBodyDetailsItems(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalDeleteCredentialDefaultBodyDetailsItems unmarshals polymorphic DeleteCredentialDefaultBodyDetailsItems
func UnmarshalDeleteCredentialDefaultBodyDetailsItems(reader io.Reader, consumer runtime.Consumer) (DeleteCredentialDefaultBodyDetailsItems, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalDeleteCredentialDefaultBodyDetailsItems(data, consumer)
}

func unmarshalDeleteCredentialDefaultBodyDetailsItems(data []byte, consumer runtime.Consumer) (DeleteCredentialDefaultBodyDetailsItems, error) {
	buf := bytes.NewBuffer(data)
	buf2 := bytes.NewBuffer(data)

	// the first time this is read is to fetch the value of the @type property.
	var getType struct {
		AtType string `json:"@type"`
	}
	if err := consumer.Consume(buf, &getType); err != nil {
		return nil, err
	}

	if err := validate.RequiredString("@type", "body", getType.AtType); err != nil {
		return nil, err
	}

	// The value of @type is used to determine which type to create and unmarshal the data into
	switch getType.AtType {
	case "deleteCredentialDefaultBodyDetailsItems":
		var result deleteCredentialDefaultBodyDetailsItems
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}

// Validate validates this delete credential default body details items
func (m *deleteCredentialDefaultBodyDetailsItems) Validate(formats strfmt.Registry) error {
	return nil
}