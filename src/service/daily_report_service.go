package service

import (
	"time"
	"fmt"

	"news/model"
	"news/utils/exception"
	"news/utils/logging"
	"news/serializer"
)

const (
	TIMEFORMAT = "2006-01-02"
)

// ShowDailyReportService 每日报告服务
type ShowDailyReportService struct {
	Day string `form:"day" json:"day"`
}

func (service *ShowDailyReportService) Show() serializer.Response {
	if service.Day == "" {
		var cstZone = time.FixedZone("CST", 8*3600) //东八区
		service.Day = time.Now().In(cstZone).Format(TIMEFORMAT)
	}

	var dailyReport model.DailyReport
	code := exception.SUCCESS

	err := model.DB.Where("day = ?", service.Day).First(&dailyReport).Error

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
		Data:   serializer.BuildReport(dailyReport),
	}
}
