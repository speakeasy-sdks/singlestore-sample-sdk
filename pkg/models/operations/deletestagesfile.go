// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type DeleteStagesFileRequest struct {
	// Path in Stages to a file or folder to delete
	Path string `pathParam:"style=simple,explode=false,name=path"`
	// ID of the Stages-enabled workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

func (o *DeleteStagesFileRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *DeleteStagesFileRequest) GetWorkspaceGroupID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceGroupID
}

type DeleteStagesFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	DeleteStagesFile *shared.DeleteStagesFile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *DeleteStagesFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *DeleteStagesFileResponse) GetDeleteStagesFile() *shared.DeleteStagesFile {
	if o == nil {
		return nil
	}
	return o.DeleteStagesFile
}

func (o *DeleteStagesFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *DeleteStagesFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
