package service

import (
	"news/model"
	"news/utils/exception"
	"news/utils/logging"
	"news/serializer"
)

// ListMediaService 媒体列表服务
type ListMediaService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
}

// List 媒体列表
func (service *ListMediaService) List() serializer.Response {
	media := []model.Media{}

	var total int64 = 0
	code := exception.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	if err := model.DB.Model(model.Media{}).Count(&total).Error; err != nil {
		logging.Info(err)
		code = exception.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&media).Error; err != nil {
		logging.Info(err)
		code = exception.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildMediaList(media), uint(total))
}