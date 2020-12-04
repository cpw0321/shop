// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// CollectListRtnJSON 收藏夹商品返回列表
type CollectListRtnJSON struct {
	List []CollectRespon `json:"list"`
}

// CollectRespon 收藏返回数据
type CollectRespon struct {
	ID          int     `json:"id"`
	ValueID     int     `json:"valueId"` // 商品ID
	Type        int     `json:"type"`    // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
	Name        string  `json:"name"`
	Brief       string  `json:"brief"`
	PicURL      string  `json:"picUrl"` // 商品页面商品图片
	RetailPrice float64 `json:"retailPrice"`
}
