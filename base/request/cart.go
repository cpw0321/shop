// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// CartAddBody 添加购物车请求参数
type CartAddBody struct {
	GoodsID   int `json:"goodsId"`
	ProductID int `json:"productId"`
	Number    int `json:"number"`
}

// CartUpdateBody 更新购物车消息请求体
type CartUpdateBody struct {
	GoodsID   int `json:"goodsId"`
	ProductID int `json:"productId"`
	Number    int `json:"number"`
	ID        int `json:"id"`
}

// 删除购物车消息请求体
type CartDeleteBody struct {
	ProductIDs []int `json:"productIds"`
}

// CartCheckedBody 选择或取消选择商品请求体
type CartCheckedBody struct {
	IsChecked  int   `json:"isChecked"`
	ProductIDs []int `json:"productIds"`
}
