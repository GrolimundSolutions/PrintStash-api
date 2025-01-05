package models

type Material struct {
	ID   int16  `gorm:"primaryKey;column:id"`
	Name string `gorm:"size:50;unique;not null"`
}

func (m *Material) TableName() string {
	return "materials"
}