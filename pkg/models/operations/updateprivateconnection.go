// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type UpdatePrivateConnectionRequest struct {
	// Here's a sample of JSON data sent in the request body to the API.
	UpdatePrivateConnection shared.UpdatePrivateConnection `request:"mediaType=application/json"`
	// ID of the private connection
	ConnectionID string `pathParam:"style=simple,explode=false,name=connectionID"`
}

func (o *UpdatePrivateConnectionRequest) GetUpdatePrivateConnection() shared.UpdatePrivateConnection {
	if o == nil {
		return shared.UpdatePrivateConnection{}
	}
	return o.UpdatePrivateConnection
}

func (o *UpdatePrivateConnectionRequest) GetConnectionID() string {
	if o == nil {
		return ""
	}
	return o.ConnectionID
}

type UpdatePrivateConnectionResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	UpdatePrivateConnectionResponse *shared.UpdatePrivateConnectionResponse
}

func (o *UpdatePrivateConnectionResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdatePrivateConnectionResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdatePrivateConnectionResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *UpdatePrivateConnectionResponse) GetUpdatePrivateConnectionResponse() *shared.UpdatePrivateConnectionResponse {
	if o == nil {
		return nil
	}
	return o.UpdatePrivateConnectionResponse
}
