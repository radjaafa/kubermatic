// Code generated by go-swagger; DO NOT EDIT.

package admin

// This file was generated by the swagger tool.
// Editing this file might prove futile when you re-run the swagger generate command

import (
	"fmt"
	"io"

	"github.com/go-openapi/runtime"

	strfmt "github.com/go-openapi/strfmt"

	"github.com/kubermatic/kubermatic/api/pkg/test/e2e/api/utils/apiclient/models"
)

// GetSeedReader is a Reader for the GetSeed structure.
type GetSeedReader struct {
	formats strfmt.Registry
}

// ReadResponse reads a server response into the received o.
func (o *GetSeedReader) ReadResponse(response runtime.ClientResponse, consumer runtime.Consumer) (interface{}, error) {
	switch response.Code() {
	case 200:
		result := NewGetSeedOK()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return result, nil
	case 401:
		result := NewGetSeedUnauthorized()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	case 403:
		result := NewGetSeedForbidden()
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		return nil, result
	default:
		result := NewGetSeedDefault(response.Code())
		if err := result.readResponse(response, consumer, o.formats); err != nil {
			return nil, err
		}
		if response.Code()/100 == 2 {
			return result, nil
		}
		return nil, result
	}
}

// NewGetSeedOK creates a GetSeedOK with default headers values
func NewGetSeedOK() *GetSeedOK {
	return &GetSeedOK{}
}

/*GetSeedOK handles this case with default header values.

Seed
*/
type GetSeedOK struct {
	Payload *models.Seed
}

func (o *GetSeedOK) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedOK  %+v", 200, o.Payload)
}

func (o *GetSeedOK) GetPayload() *models.Seed {
	return o.Payload
}

func (o *GetSeedOK) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.Seed)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}

// NewGetSeedUnauthorized creates a GetSeedUnauthorized with default headers values
func NewGetSeedUnauthorized() *GetSeedUnauthorized {
	return &GetSeedUnauthorized{}
}

/*GetSeedUnauthorized handles this case with default header values.

EmptyResponse is a empty response
*/
type GetSeedUnauthorized struct {
}

func (o *GetSeedUnauthorized) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedUnauthorized ", 401)
}

func (o *GetSeedUnauthorized) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeedForbidden creates a GetSeedForbidden with default headers values
func NewGetSeedForbidden() *GetSeedForbidden {
	return &GetSeedForbidden{}
}

/*GetSeedForbidden handles this case with default header values.

EmptyResponse is a empty response
*/
type GetSeedForbidden struct {
}

func (o *GetSeedForbidden) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeedForbidden ", 403)
}

func (o *GetSeedForbidden) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	return nil
}

// NewGetSeedDefault creates a GetSeedDefault with default headers values
func NewGetSeedDefault(code int) *GetSeedDefault {
	return &GetSeedDefault{
		_statusCode: code,
	}
}

/*GetSeedDefault handles this case with default header values.

errorResponse
*/
type GetSeedDefault struct {
	_statusCode int

	Payload *models.ErrorResponse
}

// Code gets the status code for the get seed default response
func (o *GetSeedDefault) Code() int {
	return o._statusCode
}

func (o *GetSeedDefault) Error() string {
	return fmt.Sprintf("[GET /api/v1/admin/seeds/{seed_name}][%d] getSeed default  %+v", o._statusCode, o.Payload)
}

func (o *GetSeedDefault) GetPayload() *models.ErrorResponse {
	return o.Payload
}

func (o *GetSeedDefault) readResponse(response runtime.ClientResponse, consumer runtime.Consumer, formats strfmt.Registry) error {

	o.Payload = new(models.ErrorResponse)

	// response payload
	if err := consumer.Consume(response.Body(), o.Payload); err != nil && err != io.EOF {
		return err
	}

	return nil
}
