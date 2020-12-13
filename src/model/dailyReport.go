package model

// 简报信息
type DailyReport struct {
	Id uint64
	Day string
	News string
	Sentence string
}
  
func (DailyReport) TableName() string {
	return "daily_report"
}