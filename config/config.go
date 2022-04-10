package config

type Server struct {
	Mysql Mysql `mapstructure:"mysql" json:"mysql"`
	Zap   Zap   `mapstructure:"zap" json:"zap"`
	JWT   JWT   `mapstructure:"jwt" json:"jwt"`
}
