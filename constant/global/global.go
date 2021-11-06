package global

import (
	"gin-tmp/model"

	"go.uber.org/zap"
	"gorm.io/gorm"
)

// 全局变量

var (
	GtConfig *model.Config // 配置信息
	GtLogger *zap.Logger   // 日志
	GtDB     *gorm.DB      // Gorm
)
