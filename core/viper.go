package core

import (
	"fmt"
	"gin-tmp/constant/global"
	"gin-tmp/model"
	"log"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

type Viper struct {
}

func (v *Viper) InitViper() {
	fileName := "./config.yaml"
	viper.SetConfigFile(fileName)
	viper.SetConfigType("yaml")
	err := viper.ReadInConfig() // Find and read the config file
	if err != nil {             // Handle errors reading the config file
		panic(err.Error())
	}
	// 配置映射
	_ = v.ParseSetting()
	// global.GtConfig = config
}

func (v *Viper) ParseSetting() (config *model.Config) {
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("Config file changed:", in.Name)
	})
	viper.WatchConfig()
	err := viper.Unmarshal(&global.GtConfig)
	if err != nil {
		log.Panic(err.Error())
	}
	log.Println("配置读取成功！")
	return
}
