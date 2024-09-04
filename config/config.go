package config

import (
	"douya/consts"
	"encoding/json"
	"fmt"
	"github.com/spf13/viper"
	"os"
)

var DefaultConfig *Config

type Config struct {
	// gin
	GinCfg Gin `mapstructure:"gin"`
	// mysql配置信息
	MysqlCfg Mysql `mapstructure:"mysql"`
	// redis配置信息
	RedisCfg Redis `mapstructure:"redis"`
}

type Gin struct {
	AdminPort string `mapstructure:"admin_port" json:"admin_port" yaml:"admin_port"`
	AppPort   string `mapstructure:"app_port" json:"app_port" yaml:"app_port"`
}

type Mysql struct {
	Username string `mapstructure:"username" json:"username" yaml:"username"`
	Password string `mapstructure:"password" json:"password" yaml:"password"`
	Host     string `mapstructure:"host" json:"host" yaml:"host"`
	Port     string `mapstructure:"port" json:"port" yaml:"port"`
	Db       string `mapstructure:"db" json:"db" yaml:"db"`
}

type Redis struct {
	Addr         string `mapstructure:"addr" json:"addr" yaml:"addr"`
	ReadTimeout  int    `mapstructure:"read_timeout" json:"read_timeout" yaml:"read_timeout"`
	WriteTimeout int    `mapstructure:"write_timeout" json:"write_timeout" yaml:"write_timeout"`
	Username     string `mapstructure:"username" json:"username" yaml:"username"`
	Password     string `mapstructure:"password" json:"password" yaml:"password"`
	DB           int    `mapstructure:"components" json:"components" yaml:"components"`
	DialTimeout  int    `mapstructure:"dial_timeout" json:"dial_timeout" yaml:"dial_timeout"`
}

func Init() *Config {
	// 获取对应环境配置文件
	configModule := os.Getenv(consts.ConfigModuleKey)
	if configModule == "" {
		panic("env config module not set")
	}

	// 解析配置文件
	v := viper.New()
	v.SetConfigFile(fmt.Sprintf("./config/config.%s.yaml", configModule))
	if err := v.ReadInConfig(); err != nil {
		panic(err)
	}

	//反解config 配置文件读取所有配置
	if err := v.Unmarshal(&DefaultConfig); err != nil {
		panic(err)
	}

	data, _ := json.Marshal(DefaultConfig)
	fmt.Println(string(data))

	return DefaultConfig
}
