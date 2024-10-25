package request

import "encoding/json"

type SetPrizeList struct {
	ID          json.Number `form:"id" json:"id" binding:"required"`                   // 奖品 ID
	Name        string      `form:"name" json:"name" binding:"required"`               // 奖品名称
	ImageURL    string      `form:"image_url" json:"image_url" binding:"required"`     // 奖品图片 URL
	PrizeType   string      `form:"type" json:"type" binding:"required"`               // 奖品类型
	Value       string      `form:"value" json:"value" binding:"required"`             // 奖品值
	Probability json.Number `form:"probability" json:"probability" binding:"required"` // 抽奖概率
	QuotaStr    string      `form:"quotaStr" json:"quotaStr" binding:"required"`       // 抽中名额限制
}

type SetImgList struct {
	Name string `form:"name" json:"name" binding:"required"`
	Url  string `form:"url" json:"url" binding:"required"`
}

// SetTaskList
type SetTaskList struct {
	Name  string      `form:"name" json:"name" binding:"required"`
	Value json.Number `form:"value" json:"value" binding:"required"`
}

// SetTasks
type SetTasks struct {
	ID    int    `form:"id" json:"id" binding:"required"`
	Name  string `form:"name" json:"name" binding:"required"`
	Url   string `form:"url" json:"url" binding:"required"`
	Type  string `form:"type" json:"type" binding:"required"`
	State string `form:"state" json:"state" binding:"required"`
}
