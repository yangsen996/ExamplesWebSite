package router

import (
	"github.com/gin-gonic/gin"
	"github.com/yangsen996/ExamplesWebSite/controllers"
)

func Router() *gin.Engine {
	r := gin.Default()

	r.POST("/user", controllers.NewUserController().Register)
	r.POST("/upload_file", controllers.NewFile().UploadFile)
	r.GET("/file/:filehash", controllers.NewFile().GetFileMeta)
	r.GET("/downloadfile/:filehash", controllers.NewFile().GetFileMeta)
	return r
}
