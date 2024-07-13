package fileupload

import (
	"backend/app/common/utils"
	"backend/app/pkg/handler"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

const BASE_UPLOAD_PATH = "upload/images"

func GetImages(c *gin.Context) {
	entries, err := os.ReadDir(BASE_UPLOAD_PATH)
	if err != nil {
		log.Fatal(err)
	}

	images := []FileResponse{}
	for _, e := range entries {
		if e.IsDir() {
			continue
		}
		images = append(images, FileResponse{
			UrlID: e.Name(),
			Name:  getImageName(e.Name()),
		})
	}
	handler.Success(c, http.StatusOK, "Success get all images", images)
}

func GetImageByID(c *gin.Context, reqUri FileGetByID) {
	if _, err := os.Stat(BASE_UPLOAD_PATH + "/" + reqUri.UrlID); os.IsNotExist(err) {
		handler.Error(c, http.StatusNotFound, "Image not found")
		return
	}

	c.File(BASE_UPLOAD_PATH + "/" + reqUri.UrlID)
}

func UploadImage(c *gin.Context) {
	var fileForm FileForm
	if err := c.ShouldBind(&fileForm); err != nil {
		errors := make([]handler.ApiError, 1)
		errors[0] = handler.ApiError{
			Field:   "file",
			Message: "This field is required",
		}

		if err.Error() == "unexpected end of JSON input" {
			handler.Error(c, http.StatusBadRequest, "No image uploaded", errors...)
			return
		} else if strings.Contains(err.Error(), "'required'") {
			handler.Error(c, http.StatusBadRequest, "Image size too large", errors...)
			return
		}

		handler.Error(c, http.StatusBadRequest, "Failed to upload image")
		return
	}

	if fileForm.File == nil {
		handler.Error(c, http.StatusBadRequest, "No image uploaded")
		return
	}

	if fileForm.File.Size > 10<<20 { // 10 MiB
		handler.Error(c, http.StatusBadRequest, "Image size too large")
		return
	}

	var fileDetails FileGetByID
	fileDetails.UrlID = utils.GenerateRandomString(40) + "-" + fileForm.File.Filename

	if err := c.SaveUploadedFile(fileForm.File, BASE_UPLOAD_PATH+"/"+fileDetails.UrlID); err != nil {
		handler.Error(c, http.StatusBadRequest, "Failed to saving file into server")
		return
	}

	handler.Success(c, http.StatusCreated, "Success uploading an image", fileDetails)
}

func DeleteImage(c *gin.Context, reqUri FileDelete) {
	if _, err := os.Stat(BASE_UPLOAD_PATH + "/" + reqUri.UrlID); os.IsNotExist(err) {
		handler.Error(c, http.StatusNotFound, "Image not found")
		return
	}

	if err := os.Remove(BASE_UPLOAD_PATH + "/" + reqUri.UrlID); err != nil {
		handler.Error(c, http.StatusBadRequest, "Failed to delete image")
		return
	}

	handler.Success(c, http.StatusNoContent, "", nil)
}
