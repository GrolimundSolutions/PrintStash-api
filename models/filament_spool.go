package models

import (
	"gorm.io/gorm"
	"time"
)

type FilamentSpool struct {
	gorm.Model
	ManufacturerID  int16 `gorm:"not null"`
	MaterialID      int16 `gorm:"not null"`
	ColorID         int16 `gorm:"not null"`
	WeightTotal     int   `gorm:"not null"`
	WeightRemaining int
	PurchaseDate    time.Time
	Price           float64 `gorm:"type:decimal(10,2)"`
	Rating          int16   `gorm:"check:rating BETWEEN 1 AND 5"`
	Notes           string
	Code            string       `gorm:"size:12;<-:false"` // Generated field, read-only
	Manufacturer    Manufacturer `gorm:"foreignKey:ManufacturerID"`
	Material        Material     `gorm:"foreignKey:MaterialID"`
	Color           Color        `gorm:"foreignKey:ColorID"`
}

func (f *FilamentSpool) TableName() string {
	return "filament_spools"
}
