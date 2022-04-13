package model

type UserFile struct {
	Model
	UserName string `json:"user_name" gorm:"not null;index:idx_user_id;index:idx_user_file"`
	FileShal string `json:"file_shal" gorm:"type:char(40);not null;default:'';index:idx_user_file;comment:文件hash"`
	FileName string `json:"file_name" gorm:"not null;default:'';comment:文件名"`
	FileSize int    `json:"file_size" gorm:"default:0;comment:文件大小"`
	Status   int    `json:"status" gorm:"not null;default:0;index:idx_status;comment:文件状态(启用/禁用/锁定/标记删除)"`
}

func (*UserFile) TableName() string {
	return "user_file"
}
