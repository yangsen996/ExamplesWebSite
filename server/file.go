package server

import (
	"fmt"
	"github.com/yangsen996/ExamplesWebSite/global"
	"github.com/yangsen996/ExamplesWebSite/model"
	"gorm.io/gorm"
)

type file struct{}

func NewFile() *file {
	return &file{}
}

func (f *file) AddFile(file model.TblFile) bool {
	var ff model.TblFile
	if err := global.G_DB.Where("file_name=?", file.FileName).First(&ff).Error; err == gorm.ErrRecordNotFound {
		global.G_DB.Create(&file)
		fmt.Println("上传文件")
		return true
	}
	return false
}

func (f *file) GetFile(fileshal string) *model.TblFile {
	var ff model.TblFile
	if res := global.G_DB.Where("file_shal=?", fileshal).First(&ff).RowsAffected; res == 0 {
		return nil
	}
	return &ff
}

func (f *file) UploadUserFile(file model.UserFile) bool {
	var ff model.UserFile
	if err := global.G_DB.Where("user_share=? and file_shal=?", file.UserName, file.FileShal).First(&ff).Error; err != gorm.ErrRecordNotFound {
		return false
	}
	global.G_DB.Create(&ff)
	return true
}
