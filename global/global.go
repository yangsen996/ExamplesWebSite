package global

import (
	"github.com/spf13/viper"
	"github.com/yangsen996/ExamplesWebSite/config"
	"go.uber.org/zap"
	"gorm.io/gorm"
)

var (
	G_DB   *gorm.DB
	G_LOG  *zap.Logger
	G_VP   *viper.Viper
	G_CONF *config.Server
)
