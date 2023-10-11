// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package operations

import (
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/models/shared"
	"net/http"
)

type UpdateStagesFileRequest struct {
	StagesPatch *shared.StagesPatch `request:"mediaType=application/json"`
	// Path in Stages to modify
	Path string `pathParam:"style=simple,explode=false,name=path"`
	// ID of the Stages-enabled workspace group
	WorkspaceGroupID string `pathParam:"style=simple,explode=false,name=workspaceGroupID"`
}

func (o *UpdateStagesFileRequest) GetStagesPatch() *shared.StagesPatch {
	if o == nil {
		return nil
	}
	return o.StagesPatch
}

func (o *UpdateStagesFileRequest) GetPath() string {
	if o == nil {
		return ""
	}
	return o.Path
}

func (o *UpdateStagesFileRequest) GetWorkspaceGroupID() string {
	if o == nil {
		return ""
	}
	return o.WorkspaceGroupID
}

type UpdateStagesFileResponse struct {
	// HTTP response content type for this operation
	ContentType string
	// OK
	ModifyStagesFile *shared.ModifyStagesFile
	// HTTP response status code for this operation
	StatusCode int
	// Raw HTTP response; suitable for custom response parsing
	RawResponse *http.Response
}

func (o *UpdateStagesFileResponse) GetContentType() string {
	if o == nil {
		return ""
	}
	return o.ContentType
}

func (o *UpdateStagesFileResponse) GetModifyStagesFile() *shared.ModifyStagesFile {
	if o == nil {
		return nil
	}
	return o.ModifyStagesFile
}

func (o *UpdateStagesFileResponse) GetStatusCode() int {
	if o == nil {
		return 0
	}
	return o.StatusCode
}

func (o *UpdateStagesFileResponse) GetRawResponse() *http.Response {
	if o == nil {
		return nil
	}
	return o.RawResponse
}
