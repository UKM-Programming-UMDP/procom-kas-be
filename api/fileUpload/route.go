package fileUpload

import (
	"github.com/gin-gonic/gin"
)

func Routes(router *gin.Engine) {
	router.MaxMultipartMemory = 5 << 20 // 5 MiB

	r := router.Group("/api/file")
	r.GET("/images", GetAllImages)
	r.GET("/image/:url_id", GetImage)
	r.POST("/image", UploadImage)
	r.DELETE("/image/:url_id", DeleteImage)
}
