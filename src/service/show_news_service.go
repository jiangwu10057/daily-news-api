package service

import (
	"news/model"
	"news/utils/exception"
	"news/utils/logging"
	"news/serializer"
)

// ShowNewsService 新闻详情的服务
type ShowNewsService struct {
}

// Show 新闻
func (service *ShowNewsService) Show(id string) serializer.Response {
	var news model.News
	code := exception.SUCCESS
	err := model.DB.First(&news, "id = ?", id).Error

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
		Data:   serializer.BuildNews(news),
	}
}