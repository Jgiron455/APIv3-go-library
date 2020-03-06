// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"context"
	"net/http"
	"time"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTransacEmailsListParams creates a new GetTransacEmailsListParams object
// with the default values initialized.
func NewGetTransacEmailsListParams() *GetTransacEmailsListParams {
	var ()
	return &GetTransacEmailsListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransacEmailsListParamsWithTimeout creates a new GetTransacEmailsListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTransacEmailsListParamsWithTimeout(timeout time.Duration) *GetTransacEmailsListParams {
	var ()
	return &GetTransacEmailsListParams{

		timeout: timeout,
	}
}

// NewGetTransacEmailsListParamsWithContext creates a new GetTransacEmailsListParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTransacEmailsListParamsWithContext(ctx context.Context) *GetTransacEmailsListParams {
	var ()
	return &GetTransacEmailsListParams{

		Context: ctx,
	}
}

// NewGetTransacEmailsListParamsWithHTTPClient creates a new GetTransacEmailsListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTransacEmailsListParamsWithHTTPClient(client *http.Client) *GetTransacEmailsListParams {
	var ()
	return &GetTransacEmailsListParams{
		HTTPClient: client,
	}
}

/*GetTransacEmailsListParams contains all the parameters to send to the API endpoint
for the get transac emails list operation typically these are written to a http.Request
*/
type GetTransacEmailsListParams struct {

	/*Email
	  Mandatory if templateId and messageId are not passed in query filters. Email address to which transactional email has been sent.

	*/
	Email *string
	/*EndDate
	  Mandatory if startDate is used. Ending date (YYYY-MM-DD) till which you want to fetch the list. Maximum time period that can be selected is one month.

	*/
	EndDate *strfmt.Date
	/*MessageID
	  Mandatory if templateId and email are not passed in query filters. Message ID of the transactional email sent.

	*/
	MessageID *string
	/*StartDate
	  Mandatory if endDate is used. Starting date (YYYY-MM-DD) from which you want to fetch the list. Maximum time period that can be selected is one month.

	*/
	StartDate *string
	/*TemplateID
	  Mandatory if email and messageId are not passed in query filters. Id of the template that was used to compose transactional email.

	*/
	TemplateID *int64

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get transac emails list params
func (o *GetTransacEmailsListParams) WithTimeout(timeout time.Duration) *GetTransacEmailsListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transac emails list params
func (o *GetTransacEmailsListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transac emails list params
func (o *GetTransacEmailsListParams) WithContext(ctx context.Context) *GetTransacEmailsListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transac emails list params
func (o *GetTransacEmailsListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transac emails list params
func (o *GetTransacEmailsListParams) WithHTTPClient(client *http.Client) *GetTransacEmailsListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transac emails list params
func (o *GetTransacEmailsListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithEmail adds the email to the get transac emails list params
func (o *GetTransacEmailsListParams) WithEmail(email *string) *GetTransacEmailsListParams {
	o.SetEmail(email)
	return o
}

// SetEmail adds the email to the get transac emails list params
func (o *GetTransacEmailsListParams) SetEmail(email *string) {
	o.Email = email
}

// WithEndDate adds the endDate to the get transac emails list params
func (o *GetTransacEmailsListParams) WithEndDate(endDate *strfmt.Date) *GetTransacEmailsListParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get transac emails list params
func (o *GetTransacEmailsListParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithMessageID adds the messageID to the get transac emails list params
func (o *GetTransacEmailsListParams) WithMessageID(messageID *string) *GetTransacEmailsListParams {
	o.SetMessageID(messageID)
	return o
}

// SetMessageID adds the messageId to the get transac emails list params
func (o *GetTransacEmailsListParams) SetMessageID(messageID *string) {
	o.MessageID = messageID
}

// WithStartDate adds the startDate to the get transac emails list params
func (o *GetTransacEmailsListParams) WithStartDate(startDate *string) *GetTransacEmailsListParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get transac emails list params
func (o *GetTransacEmailsListParams) SetStartDate(startDate *string) {
	o.StartDate = startDate
}

// WithTemplateID adds the templateID to the get transac emails list params
func (o *GetTransacEmailsListParams) WithTemplateID(templateID *int64) *GetTransacEmailsListParams {
	o.SetTemplateID(templateID)
	return o
}

// SetTemplateID adds the templateId to the get transac emails list params
func (o *GetTransacEmailsListParams) SetTemplateID(templateID *int64) {
	o.TemplateID = templateID
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransacEmailsListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Email != nil {

		// query param email
		var qrEmail string
		if o.Email != nil {
			qrEmail = *o.Email
		}
		qEmail := qrEmail
		if qEmail != "" {
			if err := r.SetQueryParam("email", qEmail); err != nil {
				return err
			}
		}

	}

	if o.EndDate != nil {

		// query param endDate
		var qrEndDate strfmt.Date
		if o.EndDate != nil {
			qrEndDate = *o.EndDate
		}
		qEndDate := qrEndDate.String()
		if qEndDate != "" {
			if err := r.SetQueryParam("endDate", qEndDate); err != nil {
				return err
			}
		}

	}

	if o.MessageID != nil {

		// query param messageId
		var qrMessageID string
		if o.MessageID != nil {
			qrMessageID = *o.MessageID
		}
		qMessageID := qrMessageID
		if qMessageID != "" {
			if err := r.SetQueryParam("messageId", qMessageID); err != nil {
				return err
			}
		}

	}

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate string
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.TemplateID != nil {

		// query param templateId
		var qrTemplateID int64
		if o.TemplateID != nil {
			qrTemplateID = *o.TemplateID
		}
		qTemplateID := swag.FormatInt64(qrTemplateID)
		if qTemplateID != "" {
			if err := r.SetQueryParam("templateId", qTemplateID); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}