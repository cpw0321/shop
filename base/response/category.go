// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// CategoryRtnJSON ...
type CategoryRtnJSON struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Keywords   string `json:"keywords"`
	Desc       string `json:"desc"`
	PID        int    `json:"pid"`
	IconURL    string `json:"iconUrl"`
	PicURL     string `json:"picUrl"`
	Level      string `json:"level"`
	SortOrder  int    `json:"sortOrder"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
}

// GoodsCategoryRtnJSON ...
type GoodsCategoryRtnJSON struct {
	CurrentCategory CategoryRtnJSON   `json:"currentCategory"`
	ParentCategory  CategoryRtnJSON   `json:"parentCategory"`
	BrotherCategory []CategoryRtnJSON `json:"brotherCategory"`
}
