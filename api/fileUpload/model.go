package fileUpload

import (
	"mime/multipart"
)

type FileForm struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type FileDetails struct {
	FileUniqueKey
}

type FileDelete struct {
	FileUniqueKey
}

type FileResponse struct {
	UrlID string `json:"url_id"`
	Name  string `json:"name"`
}

type FileUniqueKey struct {
	UrlID string `json:"url_id" uri:"url_id" binding:"required"`
}
