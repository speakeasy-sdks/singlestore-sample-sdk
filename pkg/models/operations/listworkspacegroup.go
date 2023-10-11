// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type ListWorkspaceGroupRequest struct {
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// To include any terminated workspace groups, set to `true`
	IncludeTerminated *bool `queryParam:"style=form,explode=true,name=includeTerminated"`
}

func (o *ListWorkspaceGroupRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *ListWorkspaceGroupRequest) GetIncludeTerminated() *bool {
	if o == nil {
		return nil
	}
	return o.IncludeTerminated
}

type ListWorkspaceGroupResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
	// OK
	WorkspaceGroups []shared.WorkspaceGroup
}

func (o *ListWorkspaceGroupResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *ListWorkspaceGroupResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *ListWorkspaceGroupResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}

func (o *ListWorkspaceGroupResponse) GetWorkspaceGroups() []shared.WorkspaceGroup {
	if o == nil {
		return nil
	}
	return o.WorkspaceGroups
}
