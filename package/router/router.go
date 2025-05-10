package router

import (
	"File_Management_Portal/modules/files"
	"github.com/gin-gonic/gin"
)

func Register(engine *gin.Engine) {
	files.Register(engine)
}
