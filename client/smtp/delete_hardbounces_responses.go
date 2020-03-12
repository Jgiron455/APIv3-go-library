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

// DeleteHardbouncesReader is a Reader for the DeleteHardbounces structure.
type DeleteHardbouncesReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *DeleteHardbouncesReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 204:
		result := NewDeleteHardbouncesNoContent()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 400:
		result := NewDeleteHardbouncesBadRequest()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewDeleteHardbouncesNoContent creates a DeleteHardbouncesNoContent with default headers values
func NewDeleteHardbouncesNoContent() *DeleteHardbouncesNoContent {
	return &DeleteHardbouncesNoContent{}
}

/*DeleteHardbouncesNoContent handles this case with default header values.

Hardbounces deleted
*/
type DeleteHardbouncesNoContent struct {
}

func (o *DeleteHardbouncesNoContent) Error() string {
	return fmt.Sprintf("[POST /smtp/deleteHardbounces][%d] deleteHardbouncesNoContent ", 204)
}

func (o *DeleteHardbouncesNoContent) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewDeleteHardbouncesBadRequest creates a DeleteHardbouncesBadRequest with default headers values
func NewDeleteHardbouncesBadRequest() *DeleteHardbouncesBadRequest {
	return &DeleteHardbouncesBadRequest{}
}

/*DeleteHardbouncesBadRequest handles this case with default header values.

bad request
*/
type DeleteHardbouncesBadRequest struct {
	Payload *models.ErrorModel
}

func (o *DeleteHardbouncesBadRequest) Error() string {
	return fmt.Sprintf("[POST /smtp/deleteHardbounces][%d] deleteHardbouncesBadRequest  %+v", 400, o.Payload)
}

func (o *DeleteHardbouncesBadRequest) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *DeleteHardbouncesBadRequest) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
