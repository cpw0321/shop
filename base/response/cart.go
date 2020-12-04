// Copyright 2020 The shop Authors

// Package response implements response body.
package response

import (
	"shop/base"
	"time"
)

// IndexCartRtnJSON 购物车首页返回数据
type IndexCartRtnJSON struct {
	CartList  []Cart    `json:"cartList"`
	CartTotal CartTotal `json:"cartTotal"`
}

// cart 购物车
type Cart struct {
	ID             int       `json:"id"`
	UserID         int       `json:"userId"`
	GoodsID        int       `json:"goodsId"`
	GoodsSn        string    `json:"goodsSn"`
	GoodsName      string    `json:"goodsName"`
	ProductID      int       `json:"productId"`
	Price          float64   `json:"price"`
	Number         int       `json:"number"`
	Specifications string    `json:"specifications"`
	Checked        bool      `json:"checked"`
	PicURL         string    `json:"picUrl"`
	AddTime        time.Time `json:"addTime"`
	UpdateTime     time.Time `json:"updateTime"`
}

// CartTotal 购物车数量
type CartTotal struct {
	GoodsCount         int     `json:"goodsCount"`
	GoodsAmount        float64 `json:"goodsAmount"`
	CheckedGoodsCount  int     `json:"checkedGoodsCount"`
	CheckedGoodsAmount float64 `json:"checkedGoodsAmount"`
}

// CartCheckoutRtnJSON 购物车选择返回体
type CartCheckoutRtnJSON struct {
	CartID                int              `json:"cartId"` //购物车ID
	Address               base.ShopAddress `json:"checkedAddress"`
	ActualPrice           float64          `json:"actualPrice"`
	OrderTotalPrice       float64          `json:"orderTotalPrice"`
	FreightPrice          float64          `json:"freightPrice"` //运费
	CheckedGoodsList      []base.ShopCart  `json:"checkedGoodsList"`
	UserCouponID          int              `json:"userCouponId"` //用户优惠卷ID
	CouponID              int              `json:"couponId"`
	GoodsTotalPrice       float64          `json:"goodsTotalPrice"`
	CouponPrice           float64          `json:"couponPrice"` // 优惠卷价格
	AddressID             int              `json:"addressId"`   //地址ID
	AvailableCouponLength int              `json:"availableCouponLength"`
}
