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

// SendSMTPEmail send Smtp email
// swagger:model sendSmtpEmail
type SendSMTPEmail struct {

	// attachment
	Attachment SendSMTPEmailAttachment `json:"attachment"`

	// bcc
	Bcc SendSMTPEmailBcc `json:"bcc"`

	// cc
	Cc SendSMTPEmailCc `json:"cc"`

	// headers
	Headers map[string]string `json:"headers,omitempty"`

	// HTML body of the message
	// Required: true
	HTMLContent *string `json:"htmlContent"`

	// reply to
	ReplyTo *SendSMTPEmailReplyTo `json:"replyTo,omitempty"`

	// sender
	// Required: true
	Sender *SendSMTPEmailSender `json:"sender"`

	// Subject of the message
	// Required: true
	Subject *string `json:"subject"`

	// Plain Text body of the message
	TextContent string `json:"textContent,omitempty"`

	// to
	// Required: true
	To SendSMTPEmailTo `json:"to"`
}

// Validate validates this send Smtp email
func (m *SendSMTPEmail) Validate(formats strfmt.Registry) error {
	var res []error

	if err := m.validateHTMLContent(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateReplyTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSender(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateSubject(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if err := m.validateTo(formats); err != nil {
		// prop
		res = append(res, err)
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}

func (m *SendSMTPEmail) validateHTMLContent(formats strfmt.Registry) error {

	if err := validate.Required("htmlContent", "body", m.HTMLContent); err != nil {
		return err
	}

	return nil
}

func (m *SendSMTPEmail) validateReplyTo(formats strfmt.Registry) error {

	if swag.IsZero(m.ReplyTo) { // not required
		return nil
	}

	if m.ReplyTo != nil {

		if err := m.ReplyTo.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("replyTo")
			}
			return err
		}
	}

	return nil
}

func (m *SendSMTPEmail) validateSender(formats strfmt.Registry) error {

	if err := validate.Required("sender", "body", m.Sender); err != nil {
		return err
	}

	if m.Sender != nil {

		if err := m.Sender.Validate(formats); err != nil {
			if ve, ok := err.(*errors.Validation); ok {
				return ve.ValidateName("sender")
			}
			return err
		}
	}

	return nil
}

func (m *SendSMTPEmail) validateSubject(formats strfmt.Registry) error {

	if err := validate.Required("subject", "body", m.Subject); err != nil {
		return err
	}

	return nil
}

func (m *SendSMTPEmail) validateTo(formats strfmt.Registry) error {

	if err := validate.Required("to", "body", m.To); err != nil {
		return err
	}

	if err := m.To.Validate(formats); err != nil {
		if ve, ok := err.(*errors.Validation); ok {
			return ve.ValidateName("to")
		}
		return err
	}

	return nil
}

// MarshalBinary interface implementation
func (m *SendSMTPEmail) MarshalBinary() ([]byte, error) {
	if m == nil {
		return nil, nil
	}
	return swag.WriteJSON(m)
}

// UnmarshalBinary interface implementation
func (m *SendSMTPEmail) UnmarshalBinary(b []byte) error {
	var res SendSMTPEmail
	if err := swag.ReadJSON(b, &res); err != nil {
		return err
	}
	*m = res
	return nil
}
