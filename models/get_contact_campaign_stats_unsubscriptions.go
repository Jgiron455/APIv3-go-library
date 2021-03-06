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

// GetContactCampaignStatsUnsubscriptions get contact campaign stats unsubscriptions
// swagger:model getContactCampaignStatsUnsubscriptions
type GetContactCampaignStatsUnsubscriptions struct {

	// admin unsubscription
	// Required: true
	AdminUnsubscription GetContactCampaignStatsUnsubscriptionsAdminUnsubscription `json:"adminUnsubscription"`

	// user unsubscription
	// Required: true
	UserUnsubscription GetContactCampaignStatsUnsubscriptionsUserUnsubscription `json:"userUnsubscription"`
}

// Validate validates this get contact campaign stats unsubscriptions
func (m *GetContactCampaignStatsUnsubscriptions) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateAdminUnsubscription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateUserUnsubscription(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetContactCampaignStatsUnsubscriptions) validateAdminUnsubscription(formats strfmt.Registry) error {

	if err := validate.Required("adminUnsubscription", "body", m.AdminUnsubscription); err != nil {
		return err
	}

	if err := m.AdminUnsubscription.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("adminUnsubscription")
		}
		return err
	}

	return nil
}

func (m *GetContactCampaignStatsUnsubscriptions) validateUserUnsubscription(formats strfmt.Registry) error {

	if err := validate.Required("userUnsubscription", "body", m.UserUnsubscription); err != nil {
		return err
	}

	if err := m.UserUnsubscription.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("userUnsubscription")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetContactCampaignStatsUnsubscriptions) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetContactCampaignStatsUnsubscriptions) UnmarshalBinary(b []byte) error {
	var res GetContactCampaignStatsUnsubscriptions
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
