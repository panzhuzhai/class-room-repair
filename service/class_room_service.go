package service

import (
	"class-room-repair/dto"
	"class-room-repair/middleware/database"
	"class-room-repair/model"
)

func ClassRoom(param dto.PageDto) (*database.Paginator, error) {
	db := database.DB.Table("class_rooms").Order("id asc")
	var records []model.ClassRoom
	paginator := database.Paging(&database.Param{
		DB:      db,
		Page:    int(param.PageNumber),
		Limit:   param.PageSize,
		ShowSQL: true,
	}, &records)
	return paginator, nil
}
