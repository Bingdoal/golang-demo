package model

import (
	"time"

	_ "gorm.io/gorm"
)

type BaseModel struct {
	ID               uint      `json:"id" gorm:"primarykey"`
	CreationTime     time.Time `json:"creation_time"`
	ModificationTime time.Time `json:"modification_time"`
}
