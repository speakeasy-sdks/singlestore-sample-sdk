// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetteWorkspaceRequest struct {
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the workspace
	WorkspaceID string `pathParam:"style=simple,explode=false,name=workspaceID"`
}

func (o *GetteWorkspaceRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetteWorkspaceRequest) GetWorkspaceID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceID
}

type GetteWorkspaceResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	Workspace *shared.Workspace
}

func (o *GetteWorkspaceResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetteWorkspaceResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetteWorkspaceResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *GetteWorkspaceResponse) GetWorkspace() *shared.Workspace {
	if o == nil {
		return nil
	}
	return o.Workspace
}
