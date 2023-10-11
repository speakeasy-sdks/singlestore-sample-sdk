// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type DeleteWorkspaceGroupsRequest struct {
	// To terminate a workspace group even if it has active workspaces, set to `true`
	Force *bool `queryParam:"style=form,explode=true,name=force"`
	// ID of the workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

func (o *DeleteWorkspaceGroupsRequest) GetForce() *bool {
	if o == nil {
		return nil
	}
	return o.Force
}

func (o *DeleteWorkspaceGroupsRequest) GetWorkspaceGroupID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceGroupID
}

type DeleteWorkspaceGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	TerminateWorkspaceGroups *shared.TerminateWorkspaceGroups
}

func (o *DeleteWorkspaceGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteWorkspaceGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteWorkspaceGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *DeleteWorkspaceGroupsResponse) GetTerminateWorkspaceGroups() *shared.TerminateWorkspaceGroups {
	if o == nil {
		return nil
	}
	return o.TerminateWorkspaceGroups
}