package model

// 美句信息
type Sentence struct {
	Id string
	Author string
	String string
	Tag string
	Type string
	LikeCount uint64
	DislikeCount uint64
	From string
}
  
func (Sentence) TableName() string {
	return "baidu_sentence"
}