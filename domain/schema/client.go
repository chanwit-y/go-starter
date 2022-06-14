package schema

import "gorm.io/gorm"

type Client struct {
	gorm.Model
	Code string `gorm:"column:code"`
	Name string `gorm:"column:name"`
}
