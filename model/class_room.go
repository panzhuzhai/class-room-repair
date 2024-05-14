package model

import "gorm.io/gorm"

type ClassRoom struct {
	gorm.Model
	Name        string `json:"name" gorm:"type:varchar(100)"`
	ClassRoomNo string `json:"classRoomNo" gorm:"type:varchar(100);"`
}
