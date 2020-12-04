// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// AftersaleRtnJSON 售后列表返回体
type AftersaleRtnJSON struct {
	List  []AftersaleRtn `json:"list"`
	Total int64          `json:"total"`
	Page  int            `json:"page"`
	Limit int            `json:"limit"`
}

// AftersaleRtn 售后列表数据
type AftersaleRtn struct {
	AftersaleRespon AftersaleRespon `json:"aftersale"`
	GoodsList       []OrderGoods    `json:"goodsList"`
}

// AftersaleRespon 售后数据
type AftersaleRespon struct {
	ID          int      `json:"id"`
	AftersaleSn string   `json:"aftersaleSn"` // 售后编号
	OrderID     int      `json:"orderId"`     // 订单ID
	UserID      int      `json:"userId"`      // 用户ID
	Type        int      `json:"type"`        // 售后类型，0是未收货退款，1是已收货（无需退货）退款，2用户退货退款
	Reason      string   `json:"reason"`      // 退款原因
	Amount      float64  `json:"amount"`      // 退款金额
	Pictures    []string `json:"pictures"`    // 退款凭证图片链接数组
	Comment     string   `json:"comment"`     // 退款说明
	Status      int      `json:"status"`      // 售后状态，0是可申请，1是用户已申请，2是管理员审核通过，3是管理员退款成功，4是管理员审核拒绝，5是用户已取消
	HandleTime  string   `json:"handleTime"`  // 管理员操作时间
	AddTime     string   `json:"addTime"`     // 创建时间
	UpdateTime  string   `json:"updateTime"`  // 更新时间
}

// AftersaleDetailRtnJSON 售后详情返回体
type AftersaleDetailRtnJSON struct {
	Aftersale  AftersaleRespon `json:"aftersale"`
	Order      OrderInfo       `json:"order"`
	OrderGoods []OrderGoods    `json:"orderGoods"`
}
