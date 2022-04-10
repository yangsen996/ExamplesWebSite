package controllers

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangsen996/ExamplesWebSite/model"
	"net/http"
	"time"
)

type File struct{}

func NewFile() *File {
	return &File{}
}

func (uf *File) UploadFile(c *gin.Context) {

	file, err := c.FormFile("upload_file")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmta := model.FileMeta{
		FileName:     file.Filename,
		FileSize:     file.Size,
		LocationPath: "C:\\Users\\83872\\Desktop\\uploadfile\\" + file.Filename,
		UploadAt:     time.Now().Format("2006-01-02 15:05"),
	}
	err = c.SaveUploadedFile(file, fmta.LocationPath)
	if err != nil {
		c.JSON(http.StatusOK, "上传文件失败")
		fmt.Println(err)
		return
	} else {
		model.AddFileMeta(fmta)
		c.JSON(http.StatusOK, "上传文件成功")
	}
}
