// Copyright 2020 The shop Authors

// Package response implements response body.
package response

type SystemRespon struct {
	ID         int    `json:"id"`
	KeyName    string `json:"KeyName"`    // 设置类型
	KeyValue   string `json:"KeyValue"`   // 设置的值
	AddTime    string `json:"addTime"`    // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
