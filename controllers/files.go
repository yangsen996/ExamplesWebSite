package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangsen996/ExamplesWebSite/model"
	"github.com/yangsen996/ExamplesWebSite/server"
	"github.com/yangsen996/ExamplesWebSite/utils"
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

	}

	shal := utils.MD5([]byte(fmta.FileName))
	tbfile := model.TblFile{
		FileShal: shal,
		FileName: fmta.FileName,
		FileSize: int(file.Size),
		FileAddr: "C:\\Users\\83872\\Desktop\\uploadfile\\" + file.Filename,
		Status:   0,
	}
	if !server.NewFile().AddFile(tbfile) {
		c.JSON(http.StatusOK, "上传文件失败")
		return
	}
	c.JSON(http.StatusOK, "上传文件成功")
}

func (uf *File) GetFileMeta(c *gin.Context) {
	filehash := c.Param("filehash")
	//fileMeta := model.GetFileMeta(filehash)
	//value, ok := fileMeta.(model.FileMeta)
	//if !ok {
	//	c.JSON(http.StatusInternalServerError, "服务器错误")
	//	return
	//} else {
	//	c.JSON(http.StatusOK, value)
	//}
	result := server.NewFile().GetFile(filehash)
	if result == nil {
		c.JSON(http.StatusNotFound, "获取文件失败")
		return
	}
	data, _ := json.Marshal(result)
	c.JSON(http.StatusOK, string(data))

}
func (uf *File) DownloadedFile(c *gin.Context) {
	hash := c.Param("filehash")
	fileMeta := model.GetFileMeta(hash)
	value, ok := fileMeta.(model.FileMeta)
	if !ok {
		c.JSON(http.StatusInternalServerError, "服务器错误")
		return
	}
	//file, err := os.Open(value.LocationPath)
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, "服务器错误")
	//	return
	//}
	//defer file.Close()
	//if err != nil {
	//	c.JSON(http.StatusInternalServerError, "服务器错误")
	//	return
	//}
	c.Header("Content-Type", "application/octet-stream")
	c.Header("Content-Disposition", "attachment; filename="+value.FileName)
	c.Header("Content-Transfer-Encoding", "binary")
	c.File(value.LocationPath)
	return
}
