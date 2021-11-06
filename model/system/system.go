package system

type Student struct {
	ID   uint   `gorm:"not null;primarykey"`
	Name string `gorm:"not null;comment:User Name"`
	Age  int64  `gorm:"not null;comment:User Age"`
}
