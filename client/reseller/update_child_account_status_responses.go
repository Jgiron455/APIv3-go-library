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

// UpdateChildAccountStatusReader is a Reader for the UpdateChildAccountStatus structure.
type UpdateChildAccountStatusReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *UpdateChildAccountStatusReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewUpdateChildAccountStatusNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewUpdateChildAccountStatusBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewUpdateChildAccountStatusForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewUpdateChildAccountStatusNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewUpdateChildAccountStatusNoContent creates a UpdateChildAccountStatusNoContent with default headers values
func NewUpdateChildAccountStatusNoContent() *UpdateChildAccountStatusNoContent {
	return &UpdateChildAccountStatusNoContent{}
}

/*UpdateChildAccountStatusNoContent handles this case with default header values.

reseller's child account status updated
*/
type UpdateChildAccountStatusNoContent struct {
}

func (o *UpdateChildAccountStatusNoContent) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}/accountStatus][%d] updateChildAccountStatusNoContent ", 204)
}

func (o *UpdateChildAccountStatusNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewUpdateChildAccountStatusBadRequest creates a UpdateChildAccountStatusBadRequest with default headers values
func NewUpdateChildAccountStatusBadRequest() *UpdateChildAccountStatusBadRequest {
	return &UpdateChildAccountStatusBadRequest{}
}

/*UpdateChildAccountStatusBadRequest handles this case with default header values.

bad request
*/
type UpdateChildAccountStatusBadRequest struct {
	Payload *models.ErrorModel
}

func (o *UpdateChildAccountStatusBadRequest) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}/accountStatus][%d] updateChildAccountStatusBadRequest  %+v", 400, o.Payload)
}

func (o *UpdateChildAccountStatusBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *UpdateChildAccountStatusBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChildAccountStatusForbidden creates a UpdateChildAccountStatusForbidden with default headers values
func NewUpdateChildAccountStatusForbidden() *UpdateChildAccountStatusForbidden {
	return &UpdateChildAccountStatusForbidden{}
}

/*UpdateChildAccountStatusForbidden handles this case with default header values.

Current account is not a reseller
*/
type UpdateChildAccountStatusForbidden struct {
	Payload *models.ErrorModel
}

func (o *UpdateChildAccountStatusForbidden) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}/accountStatus][%d] updateChildAccountStatusForbidden  %+v", 403, o.Payload)
}

func (o *UpdateChildAccountStatusForbidden) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *UpdateChildAccountStatusForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewUpdateChildAccountStatusNotFound creates a UpdateChildAccountStatusNotFound with default headers values
func NewUpdateChildAccountStatusNotFound() *UpdateChildAccountStatusNotFound {
	return &UpdateChildAccountStatusNotFound{}
}

/*UpdateChildAccountStatusNotFound handles this case with default header values.

Child auth key not found
*/
type UpdateChildAccountStatusNotFound struct {
	Payload *models.ErrorModel
}

func (o *UpdateChildAccountStatusNotFound) Error() string {
	return fmt.Sprintf("[PUT /reseller/children/{childAuthKey}/accountStatus][%d] updateChildAccountStatusNotFound  %+v", 404, o.Payload)
}

func (o *UpdateChildAccountStatusNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *UpdateChildAccountStatusNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
