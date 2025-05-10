package files

import "github.com/gin-gonic/gin"

func Register(engine *gin.Engine) {
	router := engine.Group("/api/files")
	{
		router.GET("/get-list", GetList)
		router.POST("/upload/single", UploadSingleFile)
		router.POST("/upload/multi", UploadMultipleFile)
		router.GET("/download", DownloadFile)
	}
}
