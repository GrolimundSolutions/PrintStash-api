package models

import "gorm.io/gorm"

type PrintSetting struct {
	gorm.Model
	FilamentSpoolID   int32 `gorm:"not null"`
	NozzleTemperature int   `gorm:"not null"`
	BedTemperature    int   `gorm:"not null"`
	FlowRate          int16 `gorm:"default:100"`
	Notes             string
	FilamentSpool     FilamentSpool `gorm:"foreignKey:FilamentSpoolID"`
}

func (p *PrintSetting) TableName() string {
	return "print_settings"
}
