// Code generated by go-swagger; DO NOT EDIT.

package s_m_s_campaigns

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

	"APIv3_go_wrapper/models"
)

// NewSendSmsReportParams creates a new SendSmsReportParams object
// with the default values initialized.
func NewSendSmsReportParams() *SendSmsReportParams {
	var ()
	return &SendSmsReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewSendSmsReportParamsWithTimeout creates a new SendSmsReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewSendSmsReportParamsWithTimeout(timeout time.Duration) *SendSmsReportParams {
	var ()
	return &SendSmsReportParams{

		timeout: timeout,
	}
}

// NewSendSmsReportParamsWithContext creates a new SendSmsReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewSendSmsReportParamsWithContext(ctx context.Context) *SendSmsReportParams {
	var ()
	return &SendSmsReportParams{

		Context: ctx,
	}
}

// NewSendSmsReportParamsWithHTTPClient creates a new SendSmsReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewSendSmsReportParamsWithHTTPClient(client *http.Client) *SendSmsReportParams {
	var ()
	return &SendSmsReportParams{
		HTTPClient: client,
	}
}

/*SendSmsReportParams contains all the parameters to send to the API endpoint
for the send sms report operation typically these are written to a http.Request
*/
type SendSmsReportParams struct {

	/*CampaignID
	  id of the campaign

	*/
	CampaignID int64
	/*SendReport
	  Values for send a report

	*/
	SendReport *models.SendReport

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the send sms report params
func (o *SendSmsReportParams) WithTimeout(timeout time.Duration) *SendSmsReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the send sms report params
func (o *SendSmsReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the send sms report params
func (o *SendSmsReportParams) WithContext(ctx context.Context) *SendSmsReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the send sms report params
func (o *SendSmsReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the send sms report params
func (o *SendSmsReportParams) WithHTTPClient(client *http.Client) *SendSmsReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the send sms report params
func (o *SendSmsReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithCampaignID adds the campaignID to the send sms report params
func (o *SendSmsReportParams) WithCampaignID(campaignID int64) *SendSmsReportParams {
	o.SetCampaignID(campaignID)
	return o
}

// SetCampaignID adds the campaignId to the send sms report params
func (o *SendSmsReportParams) SetCampaignID(campaignID int64) {
	o.CampaignID = campaignID
}

// WithSendReport adds the sendReport to the send sms report params
func (o *SendSmsReportParams) WithSendReport(sendReport *models.SendReport) *SendSmsReportParams {
	o.SetSendReport(sendReport)
	return o
}

// SetSendReport adds the sendReport to the send sms report params
func (o *SendSmsReportParams) SetSendReport(sendReport *models.SendReport) {
	o.SendReport = sendReport
}

// WriteToRequest writes these params to a swagger request
func (o *SendSmsReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param campaignId
	if err := r.SetPathParam("campaignId", swag.FormatInt64(o.CampaignID)); err != nil {
		return err
	}

	if o.SendReport != nil {
		if err := r.SetBodyParam(o.SendReport); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
