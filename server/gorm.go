package server

import (
	"fmt"
	"gin-tmp/constant/global"
	"gin-tmp/model/system"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

type GtGorm struct {
}

func (g *GtGorm) InitGorm() *gorm.DB {
	config := global.GtConfig.Mysql
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		config.User,
		config.Passwd,
		config.Addr,
		config.Port,
		config.Db,
	)
	db, err := gorm.Open(mysql.New(mysql.Config{
		DSN:                       dsn,   // DSN data source name
		DefaultStringSize:         256,   // string 类型字段的默认长度
		DisableDatetimePrecision:  true,  // 禁用 datetime 精度，MySQL 5.6 之前的数据库不支持
		DontSupportRenameIndex:    true,  // 重命名索引时采用删除并新建的方式，MySQL 5.7 之前的数据库和 MariaDB 不支持重命名索引
		DontSupportRenameColumn:   true,  // 用 `change` 重命名列，MySQL 8 之前的数据库和 MariaDB 不支持重命名列
		SkipInitializeWithVersion: false, // 根据当前 MySQL 版本自动配置
	}), &gorm.Config{})
	if err != nil {
		panic(err.Error())
	}
	global.GtDB = db
	return db
}

func (g *GtGorm) InitMigrate() {
	// 迁移 schema
	global.GtDB.AutoMigrate(system.Student{})
}
