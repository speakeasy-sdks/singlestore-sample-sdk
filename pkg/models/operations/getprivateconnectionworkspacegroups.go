// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetPrivateConnectionWorkspaceGroupsRequest struct {
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

func (o *GetPrivateConnectionWorkspaceGroupsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetPrivateConnectionWorkspaceGroupsRequest) GetWorkspaceGroupID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceGroupID
}

type GetPrivateConnectionWorkspaceGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	PrivateConnections []shared.PrivateConnection
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetPrivateConnectionWorkspaceGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetPrivateConnectionWorkspaceGroupsResponse) GetPrivateConnections() []shared.PrivateConnection {
	if o == nil {
		return nil
	}
	return o.PrivateConnections
}

func (o *GetPrivateConnectionWorkspaceGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetPrivateConnectionWorkspaceGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}