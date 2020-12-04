// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// CommentBody 评论请求体
type CommentBody struct {
	ValueID    int      `json:"valueId"`
	Type       int      `json:"type"`
	Content    string   `json:"content"`
	HasPicture bool     `json:"hasPicture"`
	PicURLs    []string `json:"picUrls"`
	Star       int      `json:"star"`
}
