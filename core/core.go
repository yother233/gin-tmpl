package core

import (
	"gin-tmp/constant/global"
	"gin-tmp/server"

	"go.uber.org/zap"
)

type Core struct {
	V *Viper
	L *Logger
	G *server.GtGorm
	S *server.Serve
}

func Start() {
	s := Core{}
	s.NewCore()
	s.InitCore()
}

func (s *Core) NewCore() *Core {
	return &Core{
		V: &Viper{},
		L: &Logger{},
		G: &server.GtGorm{},
		S: &server.Serve{},
	}
}

func (s *Core) InitCore() {
	s.V.InitViper()
	s.L.InitLogger()
	_ = s.G.InitGorm()
	if global.GtDB != nil {
		s.G.InitMigrate()
		if db, err := global.GtDB.DB(); err != nil {
			defer db.Close()
		}
	} else {
		global.GtLogger.Error("Gorm init error", zap.String("message", "gorm is nil!"))
		panic("")
	}
	s.S.RunServe()
}
