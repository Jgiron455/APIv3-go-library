// Code generated by go-swagger; DO NOT EDIT.

package models

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"encoding/json"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/swag"
	"github.com/go-openapi/validate"
)

// EmailExportRecipients email export recipients
// swagger:model emailExportRecipients
type EmailExportRecipients struct {

	// Webhook called once the export process is finished
	NotifyURL string `json:"notifyURL,omitempty"`

	// Type of recipients to export for a campaign
	// Required: true
	RecipientsType *string `json:"recipientsType"`
}

// Validate validates this email export recipients
func (m *EmailExportRecipients) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateRecipientsType(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

var emailExportRecipientsTypeRecipientsTypePropEnum []interface{}

func init() {
	var res []string
	if err := json.Unmarshal([]byte(`["all","nonClickers","nonOpeners","clickers","openers","softBounces","hardBounces","unsubscribed"]`), &res); err != nil {
		panic(err)
	}
	for _, v := range res {
		emailExportRecipientsTypeRecipientsTypePropEnum = append(emailExportRecipientsTypeRecipientsTypePropEnum, v)
	}
}

const (
	// EmailExportRecipientsRecipientsTypeAll captures enum value "all"
	EmailExportRecipientsRecipientsTypeAll string = "all"
	// EmailExportRecipientsRecipientsTypeNonClickers captures enum value "nonClickers"
	EmailExportRecipientsRecipientsTypeNonClickers string = "nonClickers"
	// EmailExportRecipientsRecipientsTypeNonOpeners captures enum value "nonOpeners"
	EmailExportRecipientsRecipientsTypeNonOpeners string = "nonOpeners"
	// EmailExportRecipientsRecipientsTypeClickers captures enum value "clickers"
	EmailExportRecipientsRecipientsTypeClickers string = "clickers"
	// EmailExportRecipientsRecipientsTypeOpeners captures enum value "openers"
	EmailExportRecipientsRecipientsTypeOpeners string = "openers"
	// EmailExportRecipientsRecipientsTypeSoftBounces captures enum value "softBounces"
	EmailExportRecipientsRecipientsTypeSoftBounces string = "softBounces"
	// EmailExportRecipientsRecipientsTypeHardBounces captures enum value "hardBounces"
	EmailExportRecipientsRecipientsTypeHardBounces string = "hardBounces"
	// EmailExportRecipientsRecipientsTypeUnsubscribed captures enum value "unsubscribed"
	EmailExportRecipientsRecipientsTypeUnsubscribed string = "unsubscribed"
)

// prop value enum
func (m *EmailExportRecipients) validateRecipientsTypeEnum(path, location string, value string) error {
	if err := validate.Enum(path, location, value, emailExportRecipientsTypeRecipientsTypePropEnum); err != nil {
		return err
	}
	return nil
}

func (m *EmailExportRecipients) validateRecipientsType(formats strfmt.Registry) error {

	if err := validate.Required("recipientsType", "body", m.RecipientsType); err != nil {
		return err
	}

	// value enum
	if err := m.validateRecipientsTypeEnum("recipientsType", "body", *m.RecipientsType); err != nil {
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *EmailExportRecipients) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *EmailExportRecipients) UnmarshalBinary(b []byte) error {
	var res EmailExportRecipients
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
