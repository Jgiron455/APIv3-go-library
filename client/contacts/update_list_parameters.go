// Code generated by go-swagger; DO NOT EDIT.

package contacts

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

// NewUpdateListParams creates a new UpdateListParams object
// with the default values initialized.
func NewUpdateListParams() *UpdateListParams {
	var ()
	return &UpdateListParams{

		timeout: cr.DefaultTimeout,
	}
}

// NewUpdateListParamsWithTimeout creates a new UpdateListParams object
// with the default values initialized, and the ability to set a timeout on a request
func NewUpdateListParamsWithTimeout(timeout time.Duration) *UpdateListParams {
	var ()
	return &UpdateListParams{

		timeout: timeout,
	}
}

// NewUpdateListParamsWithContext creates a new UpdateListParams object
// with the default values initialized, and the ability to set a context for a request
func NewUpdateListParamsWithContext(ctx context.Context) *UpdateListParams {
	var ()
	return &UpdateListParams{

		Context: ctx,
	}
}

// NewUpdateListParamsWithHTTPClient creates a new UpdateListParams object
// with the default values initialized, and the ability to set a custom HTTPClient for a request
func NewUpdateListParamsWithHTTPClient(client *http.Client) *UpdateListParams {
	var ()
	return &UpdateListParams{
		HTTPClient: client,
	}
}

/*UpdateListParams contains all the parameters to send to the API endpoint
for the update list operation typically these are written to a http.Request
*/
type UpdateListParams struct {

	/*ListID
	  Id of the list

	*/
	ListID int64
	/*UpdateList
	  Values to update a list

	*/
	UpdateList *models.UpdateList

	timeout    time.Duration
	Context    context.Context
	HTTPClient *http.Client
}

// WithTimeout adds the timeout to the update list params
func (o *UpdateListParams) WithTimeout(timeout time.Duration) *UpdateListParams {
	o.SetTimeout(timeout)
	return o
}

// SetTimeout adds the timeout to the update list params
func (o *UpdateListParams) SetTimeout(timeout time.Duration) {
	o.timeout = timeout
}

// WithContext adds the context to the update list params
func (o *UpdateListParams) WithContext(ctx context.Context) *UpdateListParams {
	o.SetContext(ctx)
	return o
}

// SetContext adds the context to the update list params
func (o *UpdateListParams) SetContext(ctx context.Context) {
	o.Context = ctx
}

// WithHTTPClient adds the HTTPClient to the update list params
func (o *UpdateListParams) WithHTTPClient(client *http.Client) *UpdateListParams {
	o.SetHTTPClient(client)
	return o
}

// SetHTTPClient adds the HTTPClient to the update list params
func (o *UpdateListParams) SetHTTPClient(client *http.Client) {
	o.HTTPClient = client
}

// WithListID adds the listID to the update list params
func (o *UpdateListParams) WithListID(listID int64) *UpdateListParams {
	o.SetListID(listID)
	return o
}

// SetListID adds the listId to the update list params
func (o *UpdateListParams) SetListID(listID int64) {
	o.ListID = listID
}

// WithUpdateList adds the updateList to the update list params
func (o *UpdateListParams) WithUpdateList(updateList *models.UpdateList) *UpdateListParams {
	o.SetUpdateList(updateList)
	return o
}

// SetUpdateList adds the updateList to the update list params
func (o *UpdateListParams) SetUpdateList(updateList *models.UpdateList) {
	o.UpdateList = updateList
}

// WriteToRequest writes these params to a swagger request
func (o *UpdateListParams) WriteToRequest(r runtime.ClientRequest, reg strfmt.Registry) error {

	if err := r.SetTimeout(o.timeout); err != nil {
		return err
	}
	var res []error

	// path param listId
	if err := r.SetPathParam("listId", swag.FormatInt64(o.ListID)); err != nil {
		return err
	}

	if o.UpdateList != nil {
		if err := r.SetBodyParam(o.UpdateList); err != nil {
			return err
		}
	}

	if len(res) > 0 {
		return errors.CompositeValidationError(res...)
	}
	return nil
}
