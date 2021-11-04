package model

import "time"

type Example struct {
	ID uint `json:"id"`
	Name string `json:"name"`
	Age uint `json:"age"`
	CreateTime time.Time `json:"create_time"`
	UpdateTime time.Time `json:"update_time"`
}

func (Example) TableName() string {
	return "example"
}
