// Copyright 2020 The shop Authors

// Package response implements response body.
package response

type BrandRespon struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`
	Desc       string  `json:"desc"`
	PicURL     string  `json:"picUrl"`
	SortOrder  int     `json:"sortOrder"`
	FloorPrice float64 `json:"floorPrice"`
	AddTime    string  `json:"addTime"`    // 创建时间
	UpdateTime string  `json:"updateTime"` // 更新时间
}
