// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// FeedbackBody 反馈返回体
type FeedbackBody struct {
	HasPicture bool     `json:"hasPicture"`
	Mobile     string   `json:"mobile"`
	FeedType   string   `json:"feedType"`
	Content    string   `json:"content"`
	PicURLs    []string `json:"picUrls"`
}
