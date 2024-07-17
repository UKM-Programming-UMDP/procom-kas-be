package fileupload

import (
	"mime/multipart"
)

type FileForm struct {
	File *multipart.FileHeader `form:"file" binding:"required"`
}

type FileGetByID struct {
	UrlID string `json:"url_id" uri:"id" binding:"required"`
}

type FileDelete struct {
	UrlID string `json:"url_id" uri:"id" binding:"required"`
}

type FileResponse struct {
	UrlID string `json:"url_id"`
	Name  string `json:"name"`
}
