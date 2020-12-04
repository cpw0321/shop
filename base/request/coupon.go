// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// CouponReceiveBody 接收优惠劵请求体
type CouponReceiveBody struct {
	CouponID int `json:"couponId"`
}

// CouponExchangeBody 兑换请求体
type CouponExchangeBody struct {
	Code string `json:"code"`
}
