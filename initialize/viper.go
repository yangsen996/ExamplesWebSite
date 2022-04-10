package initialize

import (
	"fmt"
	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
	"github.com/yangsen996/ExamplesWebSite/global"
)

func Viper() *viper.Viper {
	v := viper.New()
	v.SetConfigName("config")
	v.SetConfigType("yaml")
	v.AddConfigPath("./")
	if err := v.ReadInConfig(); err != nil {
		if _, ok := err.(viper.ConfigFileNotFoundError); ok {
			fmt.Println("配置文件未找到:", err)
		} else {
			fmt.Println("读取配置文件失败:", err)
		}
	}
	v.WatchConfig()
	v.OnConfigChange(func(e fsnotify.Event) {
		fmt.Println("配置文件改变", e.Name)
		//序列化
		if err := v.Unmarshal(&global.G_CONF); err != nil {
			fmt.Println("序列化失败", err)
		}
	})
	if err := v.Unmarshal(&global.G_CONF); err != nil {
		fmt.Println("序列化失败", err)
	}
	return v
}
