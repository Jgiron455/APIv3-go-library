// Code generated by go-swagger; DO NOT EDIT.

package webhooks

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"APIv3_go_wrapper/models"
)

// GetWebhookReader is a Reader for the GetWebhook structure.
type GetWebhookReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetWebhookReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetWebhookOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetWebhookBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 404:
		result := NewGetWebhookNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetWebhookOK creates a GetWebhookOK with default headers values
func NewGetWebhookOK() *GetWebhookOK {
	return &GetWebhookOK{}
}

/*GetWebhookOK handles this case with default header values.

Webhook informations
*/
type GetWebhookOK struct {
	Payload *models.GetWebhook
}

func (o *GetWebhookOK) Error() string {
	return fmt.Sprintf("[GET /webhooks/{webhookId}][%d] getWebhookOK  %+v", 200, o.Payload)
}

func (o *GetWebhookOK) GetPayload() *models.GetWebhook {
	return o.Payload
}

func (o *GetWebhookOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetWebhook)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookBadRequest creates a GetWebhookBadRequest with default headers values
func NewGetWebhookBadRequest() *GetWebhookBadRequest {
	return &GetWebhookBadRequest{}
}

/*GetWebhookBadRequest handles this case with default header values.

bad request
*/
type GetWebhookBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetWebhookBadRequest) Error() string {
	return fmt.Sprintf("[GET /webhooks/{webhookId}][%d] getWebhookBadRequest  %+v", 400, o.Payload)
}

func (o *GetWebhookBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetWebhookBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetWebhookNotFound creates a GetWebhookNotFound with default headers values
func NewGetWebhookNotFound() *GetWebhookNotFound {
	return &GetWebhookNotFound{}
}

/*GetWebhookNotFound handles this case with default header values.

Webhook ID not found
*/
type GetWebhookNotFound struct {
	Payload *models.ErrorModel
}

func (o *GetWebhookNotFound) Error() string {
	return fmt.Sprintf("[GET /webhooks/{webhookId}][%d] getWebhookNotFound  %+v", 404, o.Payload)
}

func (o *GetWebhookNotFound) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetWebhookNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
