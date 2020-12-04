// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// FootprintListRtnJSON 足迹返回体
type FootprintListRtnJSON struct {
	Page  int              `json:"page"`
	Limit int              `json:"limit"`
	Total int64            `json:"total"`
	List  []FootprintGoods `json:"list"`
}

// FootprintGoods ...
type FootprintGoods struct {
	ID          int     `json:"id"`
	GoodsID     int     `json:"goodsId"`
	Name        string  `json:"name"`
	Brief       string  `json:"brief"`
	PicURL      string  `json:"picUrl"`
	RetailPrice float64 `json:"retailPrice"`
	AddTime     string  `json:"addTime"`
}
