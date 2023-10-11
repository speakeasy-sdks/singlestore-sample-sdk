// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type GetRecoveryBackupWorkspaceGroupsRequest struct {
	// Comma-separated values list that correspond to the filtered fields for returned entities
	Fields *string `queryParam:"style=form,explode=true,name=fields"`
	// ID of the workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

func (o *GetRecoveryBackupWorkspaceGroupsRequest) GetFields() *string {
	if o == nil {
		return nil
	}
	return o.Fields
}

func (o *GetRecoveryBackupWorkspaceGroupsRequest) GetWorkspaceGroupID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceGroupID
}

type GetRecoveryBackupWorkspaceGroupsResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	Regions []shared.Region
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *GetRecoveryBackupWorkspaceGroupsResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *GetRecoveryBackupWorkspaceGroupsResponse) GetRegions() []shared.Region {
	if o == nil {
		return nil
	}
	return o.Regions
}

func (o *GetRecoveryBackupWorkspaceGroupsResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *GetRecoveryBackupWorkspaceGroupsResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
