package serializer

import "news/model"

//Sentence 美句序列化器
type Sentence struct {
	Id uint64 `json:"id"`
	Author string `json:"author"`
	String string `json:"string"`
	Tag string `json:"tag"`
	Type string `json:"type"`
	LikeCount uint64 `json:"like_count"`
	DislikeCount uint64 `json:"dislike_count"`
	From string `json:"from"`
}

// BuildSentence 序列化美句
func BuildSentence(item model.Sentence) Sentence {
	return Sentence{
		Id: item.Id,
		Author: item.Author,
		String: item.String,
		Tag: item.Tag,
		Type: item.Type,
		LikeCount: item.LikeCount,
		DislikeCount: item.DislikeCount,
		From: item.From,
	}
}

// BuildSentenceList 序列化美句列表
func BuildSentenceList(items []model.Sentence) (sentenceList []Sentence) {
	for _, item := range items {
		sentence := BuildSentence(item)
		sentenceList = append(sentenceList, sentence)
	}
	return sentenceList
}