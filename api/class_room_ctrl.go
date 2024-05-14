package api

import (
	"class-room-repair/dto"
	"class-room-repair/serializer"
	"class-room-repair/service"
	"github.com/gin-gonic/gin"
)

// @Summary web-oriented class-room
// @Description  ido
// @Accept  json
// @Produce  json
// @Param  request body  dto.PageDto  true "IdoCompleted parameters"
// @Success 200 {object} service.IdoCompletedResp
// @Router /api/v1/web-oriented/class-room[get]
func ClassRoom(c *gin.Context) {
	var param dto.PageDto
	err := c.BindQuery(&param)
	if err != nil {
		serializer.WriteData2Front(c, nil, err, "args is err")
		return
	}
	paginator, err := service.ClassRoom(param)
	serializer.WriteData2Front(c, paginator, err, "")
}
