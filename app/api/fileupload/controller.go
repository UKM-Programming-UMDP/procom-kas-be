package fileupload

import (
	"backend/app/pkg/validator"
	"strings"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func Controller(router *gin.Engine, db *gorm.DB) {
	router.MaxMultipartMemory = 10 << 20 // 10 MiB
	v1 := router.Group("/v1")

	v1.GET("/file/images", func(c *gin.Context) {
		GetImages(c)
	})

	v1.GET("/file/images/:id", func(c *gin.Context) {
		var reqUri FileGetByID
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		GetImageByID(c, reqUri)
	})

	v1.POST("/file/images", func(c *gin.Context) {
		UploadImage(c)
	})

	v1.DELETE("/file/images/:id", func(c *gin.Context) {
		var reqUri FileDelete
		if err := validator.BindUri(c, &reqUri); err {
			return
		}

		DeleteImage(c, reqUri)
	})
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
