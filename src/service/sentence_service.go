package service

import (
	"news/model"
	"news/utils/exception"
	"news/utils/logging"
	"news/serializer"
)

// ListSentenceService 美句列表服务
type ListSentenceService struct {
	Limit      int  `form:"limit" json:"limit"`
	Start      int  `form:"start" json:"start"`
}

// List 美句列表
func (service *ListSentenceService) List() serializer.Response {
	sentence := []model.Sentence{}

	var total int64 = 0
	code := exception.SUCCESS

	if service.Limit == 0 {
		service.Limit = 15
	}

	if err := model.DB.Model(model.Sentence{}).Count(&total).Error; err != nil {
		logging.Info(err)
		code = exception.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
			Error:  err.Error(),
		}
	}

	if err := model.DB.Limit(service.Limit).Offset(service.Start).Find(&sentence).Error; err != nil {
		logging.Info(err)
		code = exception.ERROR_DATABASE
		return serializer.Response{
			Status: code,
			Msg:    exception.GetMsg(code),
			Error:  err.Error(),
		}
	}

	return serializer.BuildListResponse(serializer.BuildSentenceList(sentence), uint(total))
}


// ShowSentenceService 美句详情的服务
type ShowSentenceService struct {
}

// Show 美句
func (service *ShowSentenceService) Show(id string) serializer.Response {
	var sentence model.Sentence
	code := exception.SUCCESS
	err := model.DB.First(&sentence, "id = ?", id).Error
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
		Data:   serializer.BuildSentence(sentence),
	}
}