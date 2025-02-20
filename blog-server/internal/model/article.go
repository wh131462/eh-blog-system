package model

import (
    "time"
    "gorm.io/gorm"
)

type Article struct {
    ID        uint           `gorm:"primaryKey"`
    Title     string         `gorm:"size:255"`
    Content   string         `gorm:"type:text"`
    CreatedAt time.Time
    UpdatedAt time.Time
    DeletedAt gorm.DeletedAt `gorm:"index"`
}

