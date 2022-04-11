package model

type TblFile struct {
	Model
	FileShal string `json:"file_shal" gorm:"type:char(40);not null;default:'';uniqueIndex:idx_file_shal;comment:文件hash"`
	FileName string `json:"file_name" gorm:"not null;default:'';comment:文件名"`
	FileSize int    `json:"file_size" gorm:"default:0;comment:文件大小"`
	FileAddr string `json:"file_addr" gorm:"not null;default:'';comment:文件存储地址"`
	Status   int    `json:"status" gorm:"not null;default:0;index:idx_status;comment:文件状态(启用/禁用/锁定/标记删除)"`
	Ext1     int    `json:"ext1" gorm:"default:0;comment:备用字段1"`
	Ext2     int    `json:"ext2" gorm:"default:0;comment:备用字段2"`
}
