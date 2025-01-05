package models

type Manufacturer struct {
	ID   int16  `gorm:"primaryKey;column:id"`
	Name string `gorm:"size:50;unique;not null"`
}

func (m *Manufacturer) TableName() string {
	return "manufacturers"
}