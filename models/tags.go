package models

type Tags struct {
	Id   int    `gorm:"column:id;type:int;primaryKey;autoIncrement"`
	Name string `gorm:"column:name;type:varchar(255)"`
}
