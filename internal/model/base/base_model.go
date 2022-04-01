package base

import (
	"time"

	"gorm.io/gorm"
)

type BaseModel struct {
	ID               uint64    `json:"id" gorm:"primarykey"`
	CreationTime     time.Time `json:"creationTime"`
	ModificationTime time.Time `json:"modificationTime"`
}

func (model *BaseModel) BeforeCreate(db *gorm.DB) (err error) {
	model.CreationTime = time.Now()
	model.ModificationTime = time.Now()
	return
}

func (model *BaseModel) BeforeUpdate(db *gorm.DB) (err error) {
	model.ModificationTime = time.Now()
	return
}
