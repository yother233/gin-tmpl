package model

type Config struct {
	System System
	Zap    Zap
	Mysql  Mysql
}

type System struct {
	Port string
	Mode string
}

type Zap struct {
	FileName   string
	MaxSize    int
	MaxBackups int
	MaxAge     int
	Compress   bool
	LocalTime  bool
}

type Mysql struct {
	Addr   string
	Port   string
	User   string
	Passwd string
	Db     string
}
