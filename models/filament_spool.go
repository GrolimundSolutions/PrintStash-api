package models

import (
	"time"
	"gorm.io/gorm"
)

type FilamentSpool struct {
	gorm.Model
	ManufacturerID  int          `gorm:"not null" json:"manufacturer_id"`
	MaterialID      int          `gorm:"not null" json:"material_id"`
	ColorID         int          `gorm:"not null" json:"color_id"`
	WeightTotal     int          `gorm:"not null" json:"weight_total"`
	WeightRemaining int          `gorm:"not null" json:"weight_remaining"`
	PurchaseDate    time.Time    `gorm:"not null" json:"purchase_date"`
	Price           float64      `gorm:"type:decimal(10,2)" json:"price"`
	Rating          int16        `gorm:"check:rating BETWEEN 1 AND 5" json:"rating"`
	Notes           string       `json:"notes"`
	Code            string       `gorm:"not null" json:"code_id"`
	Manufacturer    Manufacturer `gorm:"foreignKey:ManufacturerID"`
	Material        Material     `gorm:"foreignKey:MaterialID"`
	Color           Color        `gorm:"foreignKey:ColorID"`
}

func (f *FilamentSpool) TableName() string {
	return "filament_spools"
}
