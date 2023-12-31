// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetPrivateConnectionWorkspaceRequest struct {
	// ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceID"`
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
}

func (o *GetPrivateConnectionWorkspaceRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

func (o *GetPrivateConnectionWorkspaceRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

type GetPrivateConnectionWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PrivateConnections []shared.PrivateConnection
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPrivateConnectionWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPrivateConnectionWorkspaceResponse) GetPrivateConnections() []shared.PrivateConnection {
	if o == nil {
		return nil
	}
	return o.PrivateConnections
}

func (o *GetPrivateConnectionWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPrivateConnectionWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
