package server

import (
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
	if err := global.G_DB.Where("file_name=?", file.FileName).First(&ff).Error; err != gorm.ErrRecordNotFound {
		return false
	}
	err := global.G_DB.Create(file).Error
	if err != nil {
		return false
	}
	return true
}

func (f *file) GetFile(fileshal string) *model.TblFile {
	var ff model.TblFile
	if res := global.G_DB.Where("file_shal=?", fileshal).First(&ff).RowsAffected; res == 0 {
		return nil
	}
	return &ff
}
