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

// GetResellerChildsReader is a Reader for the GetResellerChilds structure.
type GetResellerChildsReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetResellerChildsReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetResellerChildsOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 403:
		result := NewGetResellerChildsForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result

	default:
		return nil, runtime.NewAPIError("unknown error", response, response.Code())
	}
}

// NewGetResellerChildsOK creates a GetResellerChildsOK with default headers values
func NewGetResellerChildsOK() *GetResellerChildsOK {
	return &GetResellerChildsOK{}
}

/*GetResellerChildsOK handles this case with default header values.

list of children
*/
type GetResellerChildsOK struct {
	Payload *models.GetChildrenList
}

func (o *GetResellerChildsOK) Error() string {
	return fmt.Sprintf("[GET /reseller/children][%d] getResellerChildsOK  %+v", 200, o.Payload)
}

func (o *GetResellerChildsOK) GetPayload() *models.GetChildrenList {
	return o.Payload
}

func (o *GetResellerChildsOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.GetChildrenList)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetResellerChildsForbidden creates a GetResellerChildsForbidden with default headers values
func NewGetResellerChildsForbidden() *GetResellerChildsForbidden {
	return &GetResellerChildsForbidden{}
}

/*GetResellerChildsForbidden handles this case with default header values.

Current account is not a reseller
*/
type GetResellerChildsForbidden struct {
	Payload *models.ErrorModel
}

func (o *GetResellerChildsForbidden) Error() string {
	return fmt.Sprintf("[GET /reseller/children][%d] getResellerChildsForbidden  %+v", 403, o.Payload)
}

func (o *GetResellerChildsForbidden) GetPayload() *models.ErrorModel {
	return o.Payload
}

func (o *GetResellerChildsForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorModel)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
