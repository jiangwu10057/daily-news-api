package serializer

import "news/model"

//News 新闻序列化器
type News struct {
	Id string `json:"id"`
	Abstract string `json:"abstract"`
	Media string `json:"media"`
	MediaId uint `json:"media_id"`
	Thumbnails string `json:"thumbnails"`
	Title string `json:"title"`
	Url string `json:"url"`
}

// BuildNews 序列化新闻
func BuildNews(item model.News) News {
	return News{
		Id: item.Id,
		Abstract: item.Abstract,
		Media: item.Media,
		MediaId: item.MediaId,
		Thumbnails: item.Thumbnails,
		Title: item.Title,
		Url: item.Url,
	}
}

// BuildNewsList 序列化新闻列表
func BuildNewsList(items []model.News) (newsList []News) {
	for _, item := range items {
		news := BuildNews(item)
		newsList = append(newsList, news)
	}
	return newsList
}