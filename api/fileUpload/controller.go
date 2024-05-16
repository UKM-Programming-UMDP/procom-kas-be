package fileUpload

import (
	"backend/handler"
	"backend/utils"
	"log"
	"net/http"
	"os"
	"strings"

	"github.com/gin-gonic/gin"
)

func GetAllImages(c *gin.Context) {
	entries, err := os.ReadDir("upload/images")
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

func GetImage(c *gin.Context) {
	var fileDetails FileDetails

	if err := c.ShouldBindUri(&fileDetails); err != nil {
		handler.Error(c, http.StatusBadRequest, "Failed to get image")
		return
	}

	if _, err := os.Stat("upload/images/" + fileDetails.UrlID); os.IsNotExist(err) {
		handler.Error(c, http.StatusNotFound, "Image not found")
		return
	}

	c.File("upload/images/" + fileDetails.UrlID)
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

	if fileForm.File.Size > 5<<20 { // 5 MiB
		handler.Error(c, http.StatusBadRequest, "Image size too large")
		return
	}

	var fileDetails FileDetails
	fileDetails.UrlID = utils.RandomString(40) + "-" + fileForm.File.Filename

	if err := c.SaveUploadedFile(fileForm.File, "upload/images/"+fileDetails.UrlID); err != nil {
		handler.Error(c, http.StatusBadRequest, "Failed to saving file into server")
		return
	}

	handler.Success(c, http.StatusCreated, "Success uploading an image", fileDetails)
}

func DeleteImage(c *gin.Context) {
	var fileDetails FileDelete

	if err := c.ShouldBindUri(&fileDetails); err != nil {
		handler.Error(c, http.StatusBadRequest, "Failed to delete image")
		return
	}

	if _, err := os.Stat("upload/images/" + fileDetails.UrlID); os.IsNotExist(err) {
		handler.Error(c, http.StatusNotFound, "Image not found")
		return
	}

	if err := os.Remove("upload/images/" + fileDetails.UrlID); err != nil {
		handler.Error(c, http.StatusBadRequest, "Failed to delete image")
		return
	}

	handler.Success(c, http.StatusOK, "Success deleting an image", nil)
}

func getImageName(input string) string {
	parts := strings.Split(input, "-")
	parts = parts[1:]
	result := strings.Join(parts, "-")
	parts = strings.Split(result, ".")
	parts = parts[:len(parts)-1]
	result = strings.Join(parts, ".")
	return result
}
