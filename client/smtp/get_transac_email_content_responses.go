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

// GetTransacEmailContentReader is a Reader for the GetTransacEmailContent structure.
type GetTransacEmailContentReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetTransacEmailContentReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetTransacEmailContentOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetTransacEmailContentOK creates a GetTransacEmailContentOK with default headers values
func NewGetTransacEmailContentOK() *GetTransacEmailContentOK {
	return &GetTransacEmailContentOK{}
}

/*GetTransacEmailContentOK handles this case with default header values.

Transactional email content
*/
type GetTransacEmailContentOK struct {
	Payload *models.GetTransacEmailContent
}

func (o *GetTransacEmailContentOK) Error() string {
	return fmt.Sprintf("[GET /smtp/emails/{uuid}][%d] getTransacEmailContentOK  %+v", 200, o.Payload)
}

func (o *GetTransacEmailContentOK) GetPayload() *models.GetTransacEmailContent {
	return o.Payload
}

func (o *GetTransacEmailContentOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetTransacEmailContent)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
