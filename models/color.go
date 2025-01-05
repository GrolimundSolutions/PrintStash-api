package models

type Color struct {
	ID   int16  `gorm:"primaryKey;column:id"`
	Name string `gorm:"size:50;unique;not null"`
}

func (c *Color) TableName() string {
	return "colors"
}