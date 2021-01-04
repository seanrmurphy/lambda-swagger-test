// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"github.com/go-openapi/errors"
	"github.com/go-openapi/strfmt"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// SimpleObjectTwo simple object two
//
// swagger:model SimpleObjectTwo
type SimpleObjectTwo struct {

	// descriptor
	// Required: true
	Descriptor *string `json:"descriptor"`

	// int array
	// Required: true
	IntArray []int64 `json:"int_array"`
}

// Validate validates this simple object two
func (m *SimpleObjectTwo) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateDescriptor(formats); err != nil {
		res = append(res, err)
	}

	if err := m.validateIntArray(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SimpleObjectTwo) validateDescriptor(formats strfmt.Registry) error {

	if err := validate.Required("descriptor", "body", m.Descriptor); err != nil {
		return err
	}

	return nil
}

func (m *SimpleObjectTwo) validateIntArray(formats strfmt.Registry) error {

	if err := validate.Required("int_array", "body", m.IntArray); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SimpleObjectTwo) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SimpleObjectTwo) UnmarshalBinary(b []byte) error {
	var res SimpleObjectTwo
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
