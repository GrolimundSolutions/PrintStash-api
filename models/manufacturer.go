package models

import "gorm.io/gorm"

type Manufacturer struct {
	gorm.Model
	Name string `gorm:"size:50;unique;not null"`
}

func (m *Manufacturer) TableName() string {
	return "manufacturers"
}