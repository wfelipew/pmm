// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
)

// InventoryChangeGenericNodeResponse inventory change generic node response
// swagger:model inventoryChangeGenericNodeResponse
type InventoryChangeGenericNodeResponse struct {

	// generic
	Generic *InventoryGenericNode `json:"generic,omitempty"`
}

// Validate validates this inventory change generic node response
func (m *InventoryChangeGenericNodeResponse) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateGeneric(formats); err != nil {
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *InventoryChangeGenericNodeResponse) validateGeneric(formats strfmt.Registry) error {

	if swag.IsZero(m.Generic) { // not required
		return nil
	}

	if m.Generic != nil {
		if err := m.Generic.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("generic")
			}
			return err
		}
	}

	return nil
}

// MarshalBinary interface implementation
func (m *InventoryChangeGenericNodeResponse) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *InventoryChangeGenericNodeResponse) UnmarshalBinary(b []byte) error {
	var res InventoryChangeGenericNodeResponse
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}