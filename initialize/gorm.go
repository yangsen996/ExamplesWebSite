package initialize

import (
	"github.com/yangsen996/ExamplesWebSite/global"
	"github.com/yangsen996/ExamplesWebSite/model"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"log"
	"os"
)

func GormDB() *gorm.DB {
	m := global.G_CONF.Mysql

	if db, err := gorm.Open(mysql.Open(m.Dsn()), &gorm.Config{}); err != nil {
		return nil
	} else {
		sqlDb, _ := db.DB()
		sqlDb.SetMaxIdleConns(m.MaxIdleConns)
		sqlDb.SetMaxOpenConns(m.MaxOpenConns)
		return db
	}
}

func RegisterTables(db *gorm.DB) {
	err := db.AutoMigrate(
		model.User{},
		model.TblFile{},
		model.UserFile{},
	)
	if err != nil {
		log.Fatal("初始化失败", err)
		os.Exit(0)
	}
	log.Println("初始化成功")
}
