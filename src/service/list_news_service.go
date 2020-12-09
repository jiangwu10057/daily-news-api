package service

import (
	"news/model"
	"news/utils/exception"
	"news/utils/logging"
	"news/serializer"
)

// ListNewsService 新闻列表服务
type ListNewsService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
	MediaId uint `form:"media_id" json:"media_id"`
}

// List 新闻列表
func (service *ListNewsService) List() serializer.Response {
	news := []model.News{}

	var total int64 = 0
	code := exception.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}
	if service.MediaId == 0 {
		if err := model.DB.Model(model.News{}).Count(&total).Error; err != nil {
			logging.Info(err)
			code = exception.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    exception.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&news).Error; err != nil {
			logging.Info(err)
			code = exception.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    exception.GetMsg(code),
				Error:  err.Error(),
			}
		}
	} else {
		if err := model.DB.Model(model.News{}).Where("media_id=?", service.MediaId).Count(&total).Error; err != nil {
			logging.Info(err)
			code = exception.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    exception.GetMsg(code),
				Error:  err.Error(),
			}
		}

		if err := model.DB.Where("media_id=?", service.MediaId).Limit(service.Limit).Offset(service.Start).Find(&news).Error; err != nil {
			logging.Info(err)
			code = exception.ERROR_DATABASE
			return serializer.Response{
				Status: code,
				Msg:    exception.GetMsg(code),
				Error:  err.Error(),
			}
		}
	}

	return serializer.BuildListResponse(serializer.BuildNewsList(news), uint(total))
}