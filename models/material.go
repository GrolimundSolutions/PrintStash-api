package models

import "gorm.io/gorm"

type Material struct {
	gorm.Model
	Name string `gorm:"size:50;unique;not null"`
}

func (m *Material) TableName() string {
	return "materials"
}
