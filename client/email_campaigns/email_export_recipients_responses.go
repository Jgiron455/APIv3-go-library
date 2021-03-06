// Code generated by go-swagger; DO NOT EDIT.

package email_campaigns

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	models "github.com/sendinblue/APIv3-go-library/models"
)

// EmailExportRecipientsReader is a Reader for the EmailExportRecipients structure.
type EmailExportRecipientsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *EmailExportRecipientsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {

	case 202:
		result := NewEmailExportRecipientsAccepted()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	case 400:
		result := NewEmailExportRecipientsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	case 404:
		result := NewEmailExportRecipientsNotFound()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewEmailExportRecipientsAccepted creates a EmailExportRecipientsAccepted with default headers values
func NewEmailExportRecipientsAccepted() *EmailExportRecipientsAccepted {
	return &EmailExportRecipientsAccepted{}
}

/*EmailExportRecipientsAccepted handles this case with default header values.

Recipient export request has been accepted
*/
type EmailExportRecipientsAccepted struct {
	Payload *models.CreatedProcessID
}

func (o *EmailExportRecipientsAccepted) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/exportRecipients][%d] emailExportRecipientsAccepted  %+v", 202, o.Payload)
}

func (o *EmailExportRecipientsAccepted) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.CreatedProcessID)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailExportRecipientsBadRequest creates a EmailExportRecipientsBadRequest with default headers values
func NewEmailExportRecipientsBadRequest() *EmailExportRecipientsBadRequest {
	return &EmailExportRecipientsBadRequest{}
}

/*EmailExportRecipientsBadRequest handles this case with default header values.

bad request
*/
type EmailExportRecipientsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *EmailExportRecipientsBadRequest) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/exportRecipients][%d] emailExportRecipientsBadRequest  %+v", 400, o.Payload)
}

func (o *EmailExportRecipientsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewEmailExportRecipientsNotFound creates a EmailExportRecipientsNotFound with default headers values
func NewEmailExportRecipientsNotFound() *EmailExportRecipientsNotFound {
	return &EmailExportRecipientsNotFound{}
}

/*EmailExportRecipientsNotFound handles this case with default header values.

Campaign ID not found
*/
type EmailExportRecipientsNotFound struct {
	Payload *models.ErrorModel
}

func (o *EmailExportRecipientsNotFound) Error() string {
	return fmt.Sprintf("[POST /emailCampaigns/{campaignId}/exportRecipients][%d] emailExportRecipientsNotFound  %+v", 404, o.Payload)
}

func (o *EmailExportRecipientsNotFound) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
