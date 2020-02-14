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

// GetBucketsUnauthorizedBodyDetailsItems get buckets unauthorized body details items
// swagger:discriminator getBucketsUnauthorizedBodyDetailsItems @type
type GetBucketsUnauthorizedBodyDetailsItems interface {
	runtime.Validatable

	// at type
	// Required: true
	AtType() string
	SetAtType(string)
}

type getBucketsUnauthorizedBodyDetailsItems struct {
	atTypeField string
}

// AtType gets the at type of this polymorphic type
func (m *getBucketsUnauthorizedBodyDetailsItems) AtType() string {
	return "getBucketsUnauthorizedBodyDetailsItems"
}

// SetAtType sets the at type of this polymorphic type
func (m *getBucketsUnauthorizedBodyDetailsItems) SetAtType(val string) {

}

// UnmarshalGetBucketsUnauthorizedBodyDetailsItemsSlice unmarshals polymorphic slices of GetBucketsUnauthorizedBodyDetailsItems
func UnmarshalGetBucketsUnauthorizedBodyDetailsItemsSlice(reader io.Reader, consumer runtime.Consumer) ([]GetBucketsUnauthorizedBodyDetailsItems, error) {
	var elements []json.RawMessage
	if err := consumer.Consume(reader, &elements); err != nil {
		return nil, err
	}

	var result []GetBucketsUnauthorizedBodyDetailsItems
	for _, element := range elements {
		obj, err := unmarshalGetBucketsUnauthorizedBodyDetailsItems(element, consumer)
		if err != nil {
			return nil, err
		}
		result = append(result, obj)
	}
	return result, nil
}

// UnmarshalGetBucketsUnauthorizedBodyDetailsItems unmarshals polymorphic GetBucketsUnauthorizedBodyDetailsItems
func UnmarshalGetBucketsUnauthorizedBodyDetailsItems(reader io.Reader, consumer runtime.Consumer) (GetBucketsUnauthorizedBodyDetailsItems, error) {
	// we need to read this twice, so first into a buffer
	data, err := ioutil.ReadAll(reader)
	if err != nil {
		return nil, err
	}
	return unmarshalGetBucketsUnauthorizedBodyDetailsItems(data, consumer)
}

func unmarshalGetBucketsUnauthorizedBodyDetailsItems(data []byte, consumer runtime.Consumer) (GetBucketsUnauthorizedBodyDetailsItems, error) {
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
	case "getBucketsUnauthorizedBodyDetailsItems":
		var result getBucketsUnauthorizedBodyDetailsItems
		if err := consumer.Consume(buf2, &result); err != nil {
			return nil, err
		}
		return &result, nil

	}
	return nil, errors.New(422, "invalid @type value: %q", getType.AtType)

}

// Validate validates this get buckets unauthorized body details items
func (m *getBucketsUnauthorizedBodyDetailsItems) Validate(formats strfmt.Registry) error {
	return nil
}