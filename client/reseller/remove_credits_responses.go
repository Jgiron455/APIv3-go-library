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

// RemoveCreditsReader is a Reader for the RemoveCredits structure.
type RemoveCreditsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *RemoveCreditsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewRemoveCreditsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewRemoveCreditsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewRemoveCreditsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewRemoveCreditsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewRemoveCreditsOK creates a RemoveCreditsOK with default headers values
func NewRemoveCreditsOK() *RemoveCreditsOK {
	return &RemoveCreditsOK{}
}

/*RemoveCreditsOK handles this case with default header values.

Credits removed
*/
type RemoveCreditsOK struct {
	Payload *models.RemainingCreditModel
}

func (o *RemoveCreditsOK) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childAuthKey}/credits/remove][%d] removeCreditsOK  %+v", 200, o.Payload)
}

func (o *RemoveCreditsOK) GetPayload() *models.RemainingCreditModel {
	return o.Payload
}

func (o *RemoveCreditsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.RemainingCreditModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveCreditsBadRequest creates a RemoveCreditsBadRequest with default headers values
func NewRemoveCreditsBadRequest() *RemoveCreditsBadRequest {
	return &RemoveCreditsBadRequest{}
}

/*RemoveCreditsBadRequest handles this case with default header values.

bad request
*/
type RemoveCreditsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *RemoveCreditsBadRequest) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childAuthKey}/credits/remove][%d] removeCreditsBadRequest  %+v", 400, o.Payload)
}

func (o *RemoveCreditsBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *RemoveCreditsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveCreditsForbidden creates a RemoveCreditsForbidden with default headers values
func NewRemoveCreditsForbidden() *RemoveCreditsForbidden {
	return &RemoveCreditsForbidden{}
}

/*RemoveCreditsForbidden handles this case with default header values.

Current account is not a reseller
*/
type RemoveCreditsForbidden struct {
	Payload *models.ErrorModel
}

func (o *RemoveCreditsForbidden) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childAuthKey}/credits/remove][%d] removeCreditsForbidden  %+v", 403, o.Payload)
}

func (o *RemoveCreditsForbidden) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *RemoveCreditsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewRemoveCreditsNotFound creates a RemoveCreditsNotFound with default headers values
func NewRemoveCreditsNotFound() *RemoveCreditsNotFound {
	return &RemoveCreditsNotFound{}
}

/*RemoveCreditsNotFound handles this case with default header values.

Child ID not found
*/
type RemoveCreditsNotFound struct {
	Payload *models.ErrorModel
}

func (o *RemoveCreditsNotFound) Error() string {
	return fmt.Sprintf("[POST /reseller/children/{childAuthKey}/credits/remove][%d] removeCreditsNotFound  %+v", 404, o.Payload)
}

func (o *RemoveCreditsNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *RemoveCreditsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
