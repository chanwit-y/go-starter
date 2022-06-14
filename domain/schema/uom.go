package schema

import "gorm.io/gorm"

type UOM struct {
	gorm.Model
	Client   Client `gorm:"foreignKey:ClientId"`
	ClientId int64
}
