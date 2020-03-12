// Code generated by go-swagger; DO NOT EDIT.

package reseller

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"APIv3_go_wrapper/models"
)

// UpdateResellerChildReader is a Reader for the UpdateResellerChild structure.
type UpdateResellerChildReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateResellerChildReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateResellerChildNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateResellerChildBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateResellerChildForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateResellerChildNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateResellerChildNoContent creates a UpdateResellerChildNoContent with default headers values
func NewUpdateResellerChildNoContent() *UpdateResellerChildNoContent {
	return &UpdateResellerChildNoContent{}
}

/*UpdateResellerChildNoContent handles this case with default header values.

reseller's child updated
*/
type UpdateResellerChildNoContent struct {
}

func (o *UpdateResellerChildNoContent) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}][%d] updateResellerChildNoContent ", 204)
}

func (o *UpdateResellerChildNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateResellerChildBadRequest creates a UpdateResellerChildBadRequest with default headers values
func NewUpdateResellerChildBadRequest() *UpdateResellerChildBadRequest {
	return &UpdateResellerChildBadRequest{}
}

/*UpdateResellerChildBadRequest handles this case with default header values.

bad request
*/
type UpdateResellerChildBadRequest struct {
	Payload *models.ErrorModel
}

func (o *UpdateResellerChildBadRequest) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}][%d] updateResellerChildBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateResellerChildBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *UpdateResellerChildBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResellerChildForbidden creates a UpdateResellerChildForbidden with default headers values
func NewUpdateResellerChildForbidden() *UpdateResellerChildForbidden {
	return &UpdateResellerChildForbidden{}
}

/*UpdateResellerChildForbidden handles this case with default header values.

Current account is not a reseller
*/
type UpdateResellerChildForbidden struct {
	Payload *models.ErrorModel
}

func (o *UpdateResellerChildForbidden) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}][%d] updateResellerChildForbidden  %+v", 403, o.Payload)
}

func (o *UpdateResellerChildForbidden) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *UpdateResellerChildForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateResellerChildNotFound creates a UpdateResellerChildNotFound with default headers values
func NewUpdateResellerChildNotFound() *UpdateResellerChildNotFound {
	return &UpdateResellerChildNotFound{}
}

/*UpdateResellerChildNotFound handles this case with default header values.

Child auth key not found
*/
type UpdateResellerChildNotFound struct {
	Payload *models.ErrorModel
}

func (o *UpdateResellerChildNotFound) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}][%d] updateResellerChildNotFound  %+v", 404, o.Payload)
}

func (o *UpdateResellerChildNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *UpdateResellerChildNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
