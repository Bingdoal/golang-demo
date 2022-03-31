package model

import (
	"time"

	_ "gorm.io/gorm"
)

type BaseModel struct {
	ID               uint64    `json:"id" gorm:"primarykey"`
	CreationTime     time.Time `json:"creationTime"`
	ModificationTime time.Time `json:"modificationTime"`
}
