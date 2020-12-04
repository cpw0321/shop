// Copyright 2020 The shop Authors

// Package response implements response body.
package response

import "shop/base"

// TopicRtnJSON 主题返回体
type TopicRtnJSON struct {
	List  []TopicJSON `json:"list"`
	Total int64       `json:"total"`
	Page  int         `json:"page"`
	Limit int         `json:"limit"`
}

// TopicJSON 主题结构体
type TopicJSON struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	PicURL     string `json:"picUrl"`
	SortOrder  string `json:"sortOrder"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
}

// TopicDetailRtnJSON 主题详情返回体
type TopicDetailRtnJSON struct {
	Topic TopicDetailJSON `json:"topic"`
}

// TopicDetailJSON ...
type TopicDetailJSON struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Subtitle   string `json:"subtitle"`
	Content    string `json:"content"`
	ReadCount  string `json:"readCount"`
	Price      string `json:"price"`
	PicURL     string `json:"picUrl"`
	SortOrder  string `json:"sortOrder"`
	Goods      string `json:"goods"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
}

// TopicRelatedRtnJSON 相关主题返回体
type TopicRelatedRtnJSON struct {
	List []base.ShopTopic `json:"list"`
}
