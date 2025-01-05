package models

import "gorm.io/gorm"

type Color struct {
	gorm.Model
	Name string `gorm:"size:50;unique;not null"`
}

func (c *Color) TableName() string {
	return "colors"
}