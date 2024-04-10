// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

import (
	"encoding/json"
	"errors"
	"fmt"
	"github.com/speakeasy-sdks/singlestore-sample-sdk/pkg/utils"
)

type ContentType string

const (
	ContentTypeStr                         ContentType = "str"
	ContentTypeArrayOfStagesObjectMetadata ContentType = "arrayOfStagesObjectMetadata"
)

type Content struct {
	Str                         *string
	ArrayOfStagesObjectMetadata []StagesObjectMetadata

	Type ContentType
}

func CreateContentStr(str string) Content {
	typ := ContentTypeStr

	return Content{
		Str:  &str,
		Type: typ,
	}
}

func CreateContentArrayOfStagesObjectMetadata(arrayOfStagesObjectMetadata []StagesObjectMetadata) Content {
	typ := ContentTypeArrayOfStagesObjectMetadata

	return Content{
		ArrayOfStagesObjectMetadata: arrayOfStagesObjectMetadata,
		Type:                        typ,
	}
}

func (u *Content) UnmarshalJSON(data []byte) error {

	str := ""
	if err := utils.UnmarshalJSON(data, &str, "", true, true); err == nil {
		u.Str = &str
		u.Type = ContentTypeStr
		return nil
	}

	arrayOfStagesObjectMetadata := []StagesObjectMetadata{}
	if err := utils.UnmarshalJSON(data, &arrayOfStagesObjectMetadata, "", true, true); err == nil {
		u.ArrayOfStagesObjectMetadata = arrayOfStagesObjectMetadata
		u.Type = ContentTypeArrayOfStagesObjectMetadata
		return nil
	}

	return errors.New("could not unmarshal into supported union types")
}

func (u Content) MarshalJSON() ([]byte, error) {
	if u.Str != nil {
		return utils.MarshalJSON(u.Str, "", true)
	}

	if u.ArrayOfStagesObjectMetadata != nil {
		return utils.MarshalJSON(u.ArrayOfStagesObjectMetadata, "", true)
	}

	return nil, errors.New("could not marshal union type: all fields are null")
}

// Format of the response
type Format string

const (
	FormatJSON Format = "json"
)

func (e Format) ToPointer() *Format {
	return &e
}

func (e *Format) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "json":
		*e = Format(v)
		return nil
	default:
		return fmt.Errorf("invalid value for Format: %v", v)
	}
}

// StagesObjectMetadataType - Object type
type StagesObjectMetadataType string

const (
	StagesObjectMetadataTypeUnknown   StagesObjectMetadataType = ""
	StagesObjectMetadataTypeJSON      StagesObjectMetadataType = "json"
	StagesObjectMetadataTypeDirectory StagesObjectMetadataType = "directory"
)

func (e StagesObjectMetadataType) ToPointer() *StagesObjectMetadataType {
	return &e
}

func (e *StagesObjectMetadataType) UnmarshalJSON(data []byte) error {
	var v string
	if err := json.Unmarshal(data, &v); err != nil {
		return err
	}
	switch v {
	case "":
		fallthrough
	case "json":
		fallthrough
	case "directory":
		*e = StagesObjectMetadataType(v)
		return nil
	default:
		return fmt.Errorf("invalid value for StagesObjectMetadataType: %v", v)
	}
}

// StagesObjectMetadata - Represents the metadata corresponding to an object in Stages
type StagesObjectMetadata struct {
	Content *Content `json:"content,omitempty"`
	Created *string  `json:"created,omitempty"`
	// Format of the response
	Format       *Format `json:"format,omitempty"`
	LastModified *string `json:"last_modified,omitempty"`
	Mimetype     *string `json:"mimetype,omitempty"`
	// Name of the Stages object
	Name *string `json:"name,omitempty"`
	// Path of the Stages object
	Path *string `json:"path,omitempty"`
	Size *int64  `json:"size,omitempty"`
	// Object type
	Type     *StagesObjectMetadataType `json:"type,omitempty"`
	Writable *bool                     `json:"writable,omitempty"`
}

func (o *StagesObjectMetadata) GetContent() *Content {
	if o == nil {
		return nil
	}
	return o.Content
}

func (o *StagesObjectMetadata) GetCreated() *string {
	if o == nil {
		return nil
	}
	return o.Created
}

func (o *StagesObjectMetadata) GetFormat() *Format {
	if o == nil {
		return nil
	}
	return o.Format
}

func (o *StagesObjectMetadata) GetLastModified() *string {
	if o == nil {
		return nil
	}
	return o.LastModified
}

func (o *StagesObjectMetadata) GetMimetype() *string {
	if o == nil {
		return nil
	}
	return o.Mimetype
}

func (o *StagesObjectMetadata) GetName() *string {
	if o == nil {
		return nil
	}
	return o.Name
}

func (o *StagesObjectMetadata) GetPath() *string {
	if o == nil {
		return nil
	}
	return o.Path
}

func (o *StagesObjectMetadata) GetSize() *int64 {
	if o == nil {
		return nil
	}
	return o.Size
}

func (o *StagesObjectMetadata) GetType() *StagesObjectMetadataType {
	if o == nil {
		return nil
	}
	return o.Type
}

func (o *StagesObjectMetadata) GetWritable() *bool {
	if o == nil {
		return nil
	}
	return o.Writable
}
