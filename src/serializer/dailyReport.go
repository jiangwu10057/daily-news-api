package serializer

import (
	"strings"

	"news/model"
	"news/utils/logging"
)

//Report 简报序列化器
type DailyReport struct {
	Id uint64 `json:"id"`
	Day string `json:"day"`
	News []model.News `json:news`
	Sentence model.Sentence `json:sentence`
}

// BuildReport 序列化简报
func BuildReport(item model.DailyReport) DailyReport {
	news := []model.News{}
	newsIds := strings.Split(item.News, ",")
	if len(newsIds) > 0 && newsIds[0] != "" {
		if err := model.DB.Limit(15).Where("id in ?", newsIds).Find(&news).Error; err != nil {
			logging.Warn(err)
		}
	}

	sentence := model.Sentence{}

	err := model.DB.First(&sentence, "id = ?", item.Sentence).Error
	if err != nil {
		logging.Warn(err)
	}
	

	return DailyReport{
		Id: item.Id,
		Day: item.Day,
		News: news,
		Sentence: sentence,
	}
}