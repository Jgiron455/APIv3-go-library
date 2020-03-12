// Code generated by go-swagger; DO NOT EDIT.

package smtp

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"APIv3_go_wrapper/models"
)

// GetTransacBlockedContactsReader is a Reader for the GetTransacBlockedContacts structure.
type GetTransacBlockedContactsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransacBlockedContactsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransacBlockedContactsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewGetTransacBlockedContactsBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransacBlockedContactsOK creates a GetTransacBlockedContactsOK with default headers values
func NewGetTransacBlockedContactsOK() *GetTransacBlockedContactsOK {
	return &GetTransacBlockedContactsOK{}
}

/*GetTransacBlockedContactsOK handles this case with default header values.

List of blocked or unsubscribed transactional contacts
*/
type GetTransacBlockedContactsOK struct {
	Payload *models.GetTransacBlockedContacts
}

func (o *GetTransacBlockedContactsOK) Error() string {
	return fmt.Sprintf("[GET /smtp/blockedContacts][%d] getTransacBlockedContactsOK  %+v", 200, o.Payload)
}

func (o *GetTransacBlockedContactsOK) GetPayload() *models.GetTransacBlockedContacts {
	return o.Payload
}

func (o *GetTransacBlockedContactsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTransacBlockedContacts)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetTransacBlockedContactsBadRequest creates a GetTransacBlockedContactsBadRequest with default headers values
func NewGetTransacBlockedContactsBadRequest() *GetTransacBlockedContactsBadRequest {
	return &GetTransacBlockedContactsBadRequest{}
}

/*GetTransacBlockedContactsBadRequest handles this case with default header values.

bad request
*/
type GetTransacBlockedContactsBadRequest struct {
	Payload *models.ErrorModel
}

func (o *GetTransacBlockedContactsBadRequest) Error() string {
	return fmt.Sprintf("[GET /smtp/blockedContacts][%d] getTransacBlockedContactsBadRequest  %+v", 400, o.Payload)
}

func (o *GetTransacBlockedContactsBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetTransacBlockedContactsBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
