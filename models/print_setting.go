package models

type PrintSetting struct {
    ID                 int32   `gorm:"primaryKey;autoIncrement:true"`
    FilamentSpoolID    int32   `gorm:"not null"`
    NozzleTemperature  int     `gorm:"not null"`
    BedTemperature     int     `gorm:"not null"`
    PrintSpeed         int     `gorm:"not null"`
    RetractionDistance float64 `gorm:"type:decimal(4,1)"`
    RetractionSpeed    int
    FlowRate           int16   `gorm:"default:100"`
    FanSpeed           int16   `gorm:"default:100"`
    Notes              string
    FilamentSpool      FilamentSpool `gorm:"foreignKey:FilamentSpoolID"`
}

func (p *PrintSetting) TableName() string {
    return "print_settings"
}