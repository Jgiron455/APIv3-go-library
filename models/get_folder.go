// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// GetFolder get folder
// swagger:model getFolder
type GetFolder struct {

	// ID of the folder
	// Required: true
	ID *int64 `json:"id"`

	// Name of the folder
	// Required: true
	Name *string `json:"name"`

	// Number of blacklisted contacts in the folder
	// Required: true
	TotalBlacklisted *int64 `json:"totalBlacklisted"`

	// Number of contacts in the folder
	// Required: true
	TotalSubscribers *int64 `json:"totalSubscribers"`

	// Number of unique contacts in the folder
	// Required: true
	UniqueSubscribers *int64 `json:"uniqueSubscribers"`
}

// Validate validates this get folder
func (m *GetFolder) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateName(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalBlacklisted(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTotalSubscribers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUniqueSubscribers(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetFolder) validateID(formats strfmt.Registry) error {

	if err := validate.Required("id", "body", m.ID); err != nil {
		return err
	}

	return nil
}

func (m *GetFolder) validateName(formats strfmt.Registry) error {

	if err := validate.Required("name", "body", m.Name); err != nil {
		return err
	}

	return nil
}

func (m *GetFolder) validateTotalBlacklisted(formats strfmt.Registry) error {

	if err := validate.Required("totalBlacklisted", "body", m.TotalBlacklisted); err != nil {
		return err
	}

	return nil
}

func (m *GetFolder) validateTotalSubscribers(formats strfmt.Registry) error {

	if err := validate.Required("totalSubscribers", "body", m.TotalSubscribers); err != nil {
		return err
	}

	return nil
}

func (m *GetFolder) validateUniqueSubscribers(formats strfmt.Registry) error {

	if err := validate.Required("uniqueSubscribers", "body", m.UniqueSubscribers); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetFolder) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetFolder) UnmarshalBinary(b []byte) error {
	var res GetFolder
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
