package main

import (
	"fmt"
	"gin-tmpl/constant/global"
	"gin-tmpl/internal/pkg/db"
	"gin-tmpl/internal/pkg/viper"
	"gin-tmpl/internal/pkg/zap"
)

func main() {
	// core.Start()
	// 配置读取
	var v viper.Viper
	err := v.InitViper()
	if err != nil {
		panic(fmt.Sprintf("配置读取错误：%s", err.Error()))
	}
	// 日志配置
	z := zap.Logger{}
	z.InitLogger()
	// 数据库连接
	g := db.GtGorm{}
	db, err := g.InitGorm()
	if err != nil {
		panic(fmt.Sprintf("配置读取错误：%s", err.Error()))
	}
	if err := g.InitMigrate(db); err != nil {
		if d, err := db.DB(); err != nil {
			defer d.Close()
		}
	} else {
		global.GtLogger.Info("数据库迁移完成")
	}
	// gin服务注册

}
