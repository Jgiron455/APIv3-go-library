// Code generated by go-swagger; DO NOT EDIT.

package transactional_s_m_s

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"net/http"
	"time"

	"golang.org/x/net/context"

	"github.com/go-openapi/errors"
	"github.com/go-openapi/runtime"
	cr "github.com/go-openapi/runtime/client"
	"github.com/go-openapi/swag"

	strfmt "github.com/go-openapi/strfmt"
)

// NewGetTransacSmsReportParams creates a new GetTransacSmsReportParams object
// with the default values initialized.
func NewGetTransacSmsReportParams() *GetTransacSmsReportParams {
	var ()
	return &GetTransacSmsReportParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewGetTransacSmsReportParamsWithTimeout creates a new GetTransacSmsReportParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewGetTransacSmsReportParamsWithTimeout(timeout time.Duration) *GetTransacSmsReportParams {
	var ()
	return &GetTransacSmsReportParams{

		timeout: timeout,
	}
}

// NewGetTransacSmsReportParamsWithContext creates a new GetTransacSmsReportParams object
// with the default values initialized, and the ability to set a context for a request
func NewGetTransacSmsReportParamsWithContext(ctx context.Context) *GetTransacSmsReportParams {
	var ()
	return &GetTransacSmsReportParams{

		Context: ctx,
	}
}

// NewGetTransacSmsReportParamsWithHTTPClient creates a new GetTransacSmsReportParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewGetTransacSmsReportParamsWithHTTPClient(client *http.Client) *GetTransacSmsReportParams {
	var ()
	return &GetTransacSmsReportParams{
		HTTPClient: client,
	}
}

/*GetTransacSmsReportParams contains all the parameters to send to the API endpoint
for the get transac sms report operation typically these are written to a http.Request
*/
type GetTransacSmsReportParams struct {

	/*Days
	  Number of days in the past including today (positive integer). Not compatible with 'startDate' and 'endDate'

	*/
	Days *int64
	/*EndDate
	  Mandatory if startDate is used. Ending date (YYYY-MM-DD) of the report

	*/
	EndDate *strfmt.Date
	/*StartDate
	  Mandatory if endDate is used. Starting date (YYYY-MM-DD) of the report

	*/
	StartDate *strfmt.Date
	/*Tag
	  Filter on a tag

	*/
	Tag *string

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the get transac sms report params
func (o *GetTransacSmsReportParams) WithTimeout(timeout time.Duration) *GetTransacSmsReportParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the get transac sms report params
func (o *GetTransacSmsReportParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the get transac sms report params
func (o *GetTransacSmsReportParams) WithContext(ctx context.Context) *GetTransacSmsReportParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the get transac sms report params
func (o *GetTransacSmsReportParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the get transac sms report params
func (o *GetTransacSmsReportParams) WithHTTPClient(client *http.Client) *GetTransacSmsReportParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the get transac sms report params
func (o *GetTransacSmsReportParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithDays adds the days to the get transac sms report params
func (o *GetTransacSmsReportParams) WithDays(days *int64) *GetTransacSmsReportParams {
	o.SetDays(days)
	return o
}

// SetDays adds the days to the get transac sms report params
func (o *GetTransacSmsReportParams) SetDays(days *int64) {
	o.Days = days
}

// WithEndDate adds the endDate to the get transac sms report params
func (o *GetTransacSmsReportParams) WithEndDate(endDate *strfmt.Date) *GetTransacSmsReportParams {
	o.SetEndDate(endDate)
	return o
}

// SetEndDate adds the endDate to the get transac sms report params
func (o *GetTransacSmsReportParams) SetEndDate(endDate *strfmt.Date) {
	o.EndDate = endDate
}

// WithStartDate adds the startDate to the get transac sms report params
func (o *GetTransacSmsReportParams) WithStartDate(startDate *strfmt.Date) *GetTransacSmsReportParams {
	o.SetStartDate(startDate)
	return o
}

// SetStartDate adds the startDate to the get transac sms report params
func (o *GetTransacSmsReportParams) SetStartDate(startDate *strfmt.Date) {
	o.StartDate = startDate
}

// WithTag adds the tag to the get transac sms report params
func (o *GetTransacSmsReportParams) WithTag(tag *string) *GetTransacSmsReportParams {
	o.SetTag(tag)
	return o
}

// SetTag adds the tag to the get transac sms report params
func (o *GetTransacSmsReportParams) SetTag(tag *string) {
	o.Tag = tag
}

// WriteToRequest writes these params to a swagger request
func (o *GetTransacSmsReportParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	if o.Days != nil {

		// query param days
		var qrDays int64
		if o.Days != nil {
			qrDays = *o.Days
		}
		qDays := swag.FormatInt64(qrDays)
		if qDays != "" {
			if err := r.SetQueryParam("days", qDays); err != nil {
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

	if o.StartDate != nil {

		// query param startDate
		var qrStartDate strfmt.Date
		if o.StartDate != nil {
			qrStartDate = *o.StartDate
		}
		qStartDate := qrStartDate.String()
		if qStartDate != "" {
			if err := r.SetQueryParam("startDate", qStartDate); err != nil {
				return err
			}
		}

	}

	if o.Tag != nil {

		// query param tag
		var qrTag string
		if o.Tag != nil {
			qrTag = *o.Tag
		}
		qTag := qrTag
		if qTag != "" {
			if err := r.SetQueryParam("tag", qTag); err != nil {
				return err
			}
		}

	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}