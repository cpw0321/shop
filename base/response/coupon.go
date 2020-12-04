// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// CouponListRtnJSON 优惠劵列表返回体
type CouponListRtnJSON struct {
	List  []CouponRtnJSON `json:"list"`
	Total int64           `json:"total"`
	Page  int             `json:"page"`
	Limit int             `json:"limit"`
}

// CouponRtnJSON 优惠劵返回体
type CouponRtnJSON struct {
	ID         int     `json:"id"`
	Name       string  `json:"name"`       // 优惠券名称
	Desc       string  `json:"desc"`       // 优惠券介绍，通常是显示优惠券使用限制文字
	Tag        string  `json:"tag"`        // 优惠券标签，例如新人专用
	Discount   float64 `json:"discount"`   // 优惠金额
	Min        float64 `json:"min"`        // 最少消费金额才能使用优惠券
	Type       int     `json:"type"`       // 优惠券赠送类型，如果是0则通用券，用户领取；如果是1，则是注册赠券；如果是2，则是优惠券码兑换
	Status     int     `json:"Status"`     // 优惠券状态，如果是0则是正常可用；如果是1则是过期; 如果是2则是下架
	GoodsType  int     `json:"goodsType"`  // 商品限制类型，如果0则全商品，如果是1则是类目限制，如果是2则是商品限制
	GoodsValue string  `json:"goodsValue"` // 商品限制值，goods_type如果是0则空集合，如果是1则是类目集合，如果是2则是商品集合
	TimeType   int     `json:"timeType"`   // 有效时间限制，如果是0，则基于领取时间的有效天数days；如果是1，则start_time和end_time是优惠券有效期
	Days       int     `json:"days"`       // 基于领取时间的有效天数
	StartTime  string  `json:"startTime"`  // 使用券开始时间
	EndTime    string  `json:"endTime"`    // 使用券截至时间
	AddTime    string  `json:"addTime"`    // 创建时间
	UpdateTime string  `json:"updateTime"` // 更新时间
}

// CouponUserListRtnJSON 我的优惠劵列表返回体
type CouponUserListRtnJSON struct {
	List  []CouponUserRtnJSON `json:"list"`
	Total int64               `json:"total"`
	Page  int                 `json:"page"`
	Limit int                 `json:"limit"`
}

// CouponUserRtnJSON ...
type CouponUserRtnJSON struct {
	ID        int     `json:"id"`
	CID       int     `json:"cid"`
	Name      string  `json:"name"`      // 优惠券名称
	Desc      string  `json:"desc"`      // 优惠券介绍，通常是显示优惠券使用限制文字
	Tag       string  `json:"tag"`       // 优惠券标签，例如新人专用
	Discount  float64 `json:"discount"`  // 优惠金额
	Min       float64 `json:"min"`       // 最少消费金额才能使用优惠券
	Type      int     `json:"type"`      // 优惠券赠送类型，如果是0则通用券，用户领取；如果是1，则是注册赠券；如果是2，则是优惠券码兑换
	StartTime string  `json:"startTime"` // 使用券开始时间
	EndTime   string  `json:"endTime"`   // 使用券截至时间
	Available bool    `json:"available"` // 是否可以使用
}

// CouponUserRespon ...
type CouponUserRespon struct {
	ID         int    `json:"id"`
	UserID     int    `json:"userId"`     // 用户ID
	CouponID   int    `json:"couponId"`   // 优惠卷ID
	Status     int    `json:"status"`     // 使用状态, 如果是0则未使用；如果是1则已使用；如果是2则已过期；如果是3则已经下架
	UsedTime   string `json:"usedTime"`   // 使用时间
	StartTime  string `json:"startTime"`  // 有效期开始时间
	EndTime    string `json:"endTime"`    // 有效期截至时间
	OrderID    int    `json:"orderId"`    // 订单ID
	AddTime    string `json:"addTime""`   // 创建时间
	UpdateTime string `json:"updateTime"` // 更新时间
}
