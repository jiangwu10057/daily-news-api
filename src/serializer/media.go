package serializer

import "news/model"

//Media 新闻序列化器
type Media struct {
	Id uint `json:"id"`
	Name string `json:"name"`
	Icon string `json:"icon"`
	Intro string `json:"intro"`
}

// BuildMedia 序列化新闻
func BuildMedia(item model.Media) Media {
	return Media{
		Id: item.Id,
		Name: item.Name,
		Icon: item.Icon,
		Intro: item.Intro,
	}
}

// BuildMediaList 序列化新闻列表
func BuildMediaList(items []model.Media) (mediaList []Media) {
	for _, item := range items {
		media := BuildMedia(item)
		mediaList = append(mediaList, media)
	}
	return mediaList
}