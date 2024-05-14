package database

import (
	"class-room-repair/model"
)

func migration() {
	_ = DB.AutoMigrate(&model.ClassRoom{})
	_ = DB.AutoMigrate(&model.Repair{})
}
