package models

type ProductLog struct {
	Code       string `json:"code"`
	ActionType string `json:"action_type"`
	Quantity   int    `json:"quantity"`
}

func (ProductLog) TableName() string {
	return "product_log"
}

type ProductLogs []*ProductLog
