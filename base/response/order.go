// Copyright 2020 The shop Authors

// Package response implements response body.
package response

import "shop/base"

type OrderSubmitRtnJSON struct {
	OrderID int `json:"orderId"`
}

// OrderListRtnJSON 订单列表返回体
type OrderListRtnJSON struct {
	Total int64          `json:"total"`
	Limit int            `json:"limit"`
	Page  int            `json:"page"`
	List  []OrderRtnJSON `json:"list"`
}

// OrderRtnJSON 订单列表返回数据
type OrderRtnJSON struct {
	ID              int          `json:"id"`
	OrderStatusText string       `json:"orderStatusText"`
	AftersaleStatus int          `json:"aftersaleStatus"`
	OrderSn         string       `json:"orderSn"`
	ActualPrice     float64      `json:"actualPrice"`
	GoodsList       []OrderGoods `json:"goodsList"`
}

type OrderGoods struct {
	ID             int      `json:"id"`
	OrderID        int      `json:"orderId"`
	GoodsID        int      `json:"goodsId"`
	GoodsName      string   `json:"goodsName"`
	GoodsSn        string   `json:"goodsSn"`
	ProductID      int      `json:"productId"`
	Number         int      `json:"number"`
	PicURL         string   `json:"picUrl"`
	Price          float64  `json:"price"`
	Specifications []string `json:"specifications"`
	Comment        int      `json:"comment"`
	AddTime        string   `json:"addTime"`
}

type OrderHandleOption struct {
	Cancel    bool `json:"cancel"`    // 取消操作
	Delete    bool `json:"delete"`    // 删除操作
	Pay       bool `json:"pay"`       // 支付操作
	Comment   bool `json:"comment"`   // 评论操作
	Confirm   bool `json:"confirm"`   // 确认收货操作
	Refund    bool `json:"refund"`    // 取消订单并退款操作
	Rebuy     bool `json:"rebuy"`     // 再次购买
	Aftersale bool `json:"aftersale"` // 售后操作
	Ship      bool `json:"ship"`      // 物流
}

// OrderDetailRtnJSON 订单详情返回体
type OrderDetailRtnJSON struct {
	OrderInfo   OrderInfo             `json:"orderInfo"`
	OrderGoods  []base.ShopOrderGoods `json:"orderGoods"`
	ExpressInfo ExpressRtnInfo        `json:"expressInfo"`
}

type OrderInfo struct {
	ID              int               `json:"id"`
	UserID          int               `json:"userId"`
	OrderSn         string            `json:"orderSn"`
	OrderStatus     int               `json:"orderStatus"`
	AftersaleStatus int               `json:"aftersaleStatus"`
	Consignee       string            `json:"consignee"`
	Mobile          string            `json:"mobile"`
	Address         string            `json:"address"`
	Message         string            `json:"message"`
	GoodsPrice      float64           `json:"goodsPrice"`
	FreightPrice    float64           `json:"freightPrice"`
	CouponPrice     float64           `json:"couponPrice"`
	IntegralPrice   float64           `json:"integralPrice"`
	GrouponPrice    float64           `json:"grouponPrice"`
	OrderPrice      float64           `json:"orderPrice"`
	ActualPrice     float64           `json:"actualPrice"`
	PayID           string            `json:"payId"`
	PayTime         string            `json:"payTime"`
	ShipSn          string            `json:"shipSn"`
	ShipChannel     string            `json:"shipChannel"`
	ShipTime        string            `json:"shipTime"`
	RefundAmount    float64           `json:"refundAmount"`
	RefundType      string            `json:"refundType"`
	RefundContent   string            `json:"refundContent"`
	RefundTime      string            `json:"refundTime"`
	ConfirmTime     string            `json:"confirmTime"`
	Comments        int               `json:"comments"`
	EndTime         string            `json:"endTime"`
	AddTime         string            `json:"addTime"`
	UpdateTime      string            `json:"updateTime"`
	OrderStatusText string            `json:"orderStatusText"` // 订单状态 例如：未付款
	HandleOption    OrderHandleOption `json:"handleOption"`
}

//type OrderGoods struct {
//	ID             int    `json:"id"`
//	GoodsName      string `json:"goodsName"`
//	Specifications string `json:"specifications"`
//	Number         int    `json:"number"`
//	PicURL         string `json:"picURL"`
//}

//type OrderInfo struct {
//	ID              int          `json:"id"`
//	OrderStatusText string       `json:"orderStatusText"` // 订单状态 例如：未付款
//	OrderSn         string       `json:"orderSn"`         // 订单编号
//	ActualPrice     float64      `json:"actualPrice"`     // 实际价格
//	GoodsList       []OrderGoods `json:"goodsList"`       // 商品信息
//}

// OrderResult 查询订单支付结果返回体
type OrderResult struct {
	OutTradeNo    string `json:"outTradeNo"`    // 后台订单好
	TransactionID string `json:"transactionId"` // 微信支付订单号
	TradeState    string `json:"tradeState"`    // 交易状态 SUCCESS—支付成功  REFUND—转入退款 NOTPAY—未支付 CLOSED—已关闭 REVOKED—已撤销（刷卡支付） USERPAYING--用户支付中 PAYERROR--支付失败(其他原因，如银行返回失败)
	CashFee       int    `json:"cashFee"`       // 现金支付金额
}

// UnifiedOrderRtnJSON 统一下单返回体
type UnifiedOrderRtnJSON struct {
	TimeStamp string `json:"timeStamp"`
	NonceStr  string `json:"nonceStr"`
	Package   string `json:"packageValue"`
	SignType  string `json:"signType"`
	PaySign   string `json:"paySign"`
}
