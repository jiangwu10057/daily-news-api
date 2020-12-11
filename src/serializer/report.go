package serializer

//Report 简报序列化器
type Report struct {
	Id uint64 `json:"id"`
	Day string `json:"day"`
	News []News `json:news`
	Sentence Sentence `json:sentence`
}

