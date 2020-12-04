// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// OrderSubmitBody 订单提交体
type OrderSubmitBody struct {
	CartID       int    `json:"cartId"`
	AddressID    int    `json:"addressId"`
	CouponID     int    `json:"couponId"`
	UserCouponID int    `json:"userCouponId"`
	Message      string `json:"message"`
}

// OrderCancelBody 取消订单请求体
type OrderCancelBody struct {
	OrderID int `json:"orderId"`
}

// OrderRefundBody 订单退款请求体
type OrderRefundBody struct {
	OrderID int `json:"orderId"`
}

// OrderConfirmBody 订单确认请求体
type OrderConfirmBody struct {
	OrderID int `json:"orderId"`
}

// OrderCommentBody 评价商品信息请求体
type OrderCommentBody struct {
	OrderGoodsID int      `json:"orderGoodsId"`
	Content      string   `json:"content"`
	Star         int      `json:"star"`
	HasPicture   bool     `json:"hasPicture"`
	PicURLs      []string `json:"picUrls"`
}

// OrderPrepayBody 订单支付请求体
type OrderPrepayBody struct {
	OrderID int `json:"orderId"`
}

// OrderDeleteBody 删除订单请求体
type OrderDeleteBody struct {
	OrderID int `json:"orderId"`
}
