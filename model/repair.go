package model

import "gorm.io/gorm"

type Repair struct {
	gorm.Model
	Content string `json:"content" gorm:"type:text"`
}
