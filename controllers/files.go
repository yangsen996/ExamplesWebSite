package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/yangsen996/ExamplesWebSite/model"
	"github.com/yangsen996/ExamplesWebSite/server"
	"github.com/yangsen996/ExamplesWebSite/utils"
	"log"
	"net/http"
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

	shal := utils.MD5([]byte(file.Filename))
	log.Println("shal", shal)
	tbfile := model.TblFile{
		FileShal: shal,
		FileName: file.Filename,
		FileSize: int(file.Size),
		FileAddr: "C:\\Users\\83872\\Desktop\\uploadfile\\" + file.Filename,
		Status:   0,
	}

	//todo 上传文件和上传用户文件表做事务
	//tx := global.G_DB.Begin()

	if !server.NewFile().AddFile(tbfile) {
		c.JSON(http.StatusOK, "上传文件失败")
		return
	} else {
		err = c.SaveUploadedFile(file, tbfile.FileAddr)
		if err != nil {
			c.JSON(http.StatusOK, "上传文件失败")
			fmt.Println(err)
			return
		}
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
