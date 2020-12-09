package service

import (
	"news/model"
	"news/utils/exception"
	"news/utils/logging"
	"news/serializer"
)

// ShowMediaService 媒体详情的服务
type ShowMediaService struct {
}

// Show 媒体
func (service *ShowMediaService) Show(id string) serializer.Response {
	var media model.Media
	code := exception.SUCCESS
	err := model.DB.First(&media, id).Error
	if err != nil {
		logging.Info(err)
		code = exception.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.Response{
		Status: code,
		Msg:    exception.GetMsg(code),
		Data:   serializer.BuildMedia(media),
	}
}