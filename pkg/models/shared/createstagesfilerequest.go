// Code generated by Speakeasy (https://speakeasyapi.dev). DO NOT EDIT.

package shared

type File struct {
	Content  []byte `multipartForm:"content"`
	FileName string `multipartForm:"name=file"`
}

func (o *File) GetContent() []byte {
	if o == nil {
		return []byte{}
	}
	return o.Content
}

func (o *File) GetFileName() string {
	if o == nil {
		return ""
	}
	return o.FileName
}

type CreateStagesFileRequest struct {
	// File to upload
	File *File `multipartForm:"file"`
}

func (o *CreateStagesFileRequest) GetFile() *File {
	if o == nil {
		return nil
	}
	return o.File
}
