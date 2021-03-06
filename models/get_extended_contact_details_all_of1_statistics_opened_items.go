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

// GetExtendedContactDetailsAllOf1StatisticsOpenedItems get extended contact details all of1 statistics opened items
// swagger:model getExtendedContactDetailsAllOf1StatisticsOpenedItems
type GetExtendedContactDetailsAllOf1StatisticsOpenedItems struct {

	// ID of the campaign which generated the event
	// Required: true
	CampaignID *int64 `json:"campaignId"`

	// Number of openings for the campaign
	// Required: true
	Count *int64 `json:"count"`

	// UTC date-time of the event
	// Required: true
	EventTime *strfmt.DateTime `json:"eventTime"`

	// IP from which the user has opened the email
	// Required: true
	IP *string `json:"ip"`
}

// Validate validates this get extended contact details all of1 statistics opened items
func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateCampaignID(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateCount(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateEventTime(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateIP(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) validateCampaignID(formats strfmt.Registry) error {

	if err := validate.Required("campaignId", "body", m.CampaignID); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) validateCount(formats strfmt.Registry) error {

	if err := validate.Required("count", "body", m.Count); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) validateEventTime(formats strfmt.Registry) error {

	if err := validate.Required("eventTime", "body", m.EventTime); err != nil {
		return err
	}

	if err := validate.FormatOf("eventTime", "body", "date-time", m.EventTime.String(), formats); err != nil {
		return err
	}

	return nil
}

func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) validateIP(formats strfmt.Registry) error {

	if err := validate.Required("ip", "body", m.IP); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *GetExtendedContactDetailsAllOf1StatisticsOpenedItems) UnmarshalBinary(b []byte) error {
	var res GetExtendedContactDetailsAllOf1StatisticsOpenedItems
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
