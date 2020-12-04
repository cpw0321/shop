// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// CollectBody 收藏请求体
type CollectBody struct {
	ValueID int `json:"valueId"`
	Type    int `json:"type"`
}
