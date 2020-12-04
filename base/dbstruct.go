// Copyright 2020 The shop Authors

// Package base implements response body.
package base

import "time"

// ShopUser 用户表结构
type ShopUser struct {
	ID            int       `json:"id" orm:"column(id)"`
	Username      string    `json:"username"`
	Password      string    `json:"password"`
	Gender        int       `json:"gender"`
	Birthday      time.Time `json:"birthday"`
	LastLoginTime time.Time `json:"lastLoginTime"`
	LastLoginIP   string    `json:"lastLoginIp" orm:"column(last_login_ip)"`
	UserLevel     int       `json:"userLevel"`
	Mobile        string    `json:"mobile"`
	Avatar        string    `json:"avatar"`
	Nickname      string    `json:"nickname"`
	OpenID        string    `json:"openid" orm:"column(openid)"`
	SessionKey    string    `json:"sessionKey"`
	Status        int       `json:"status"`
	AddTime       time.Time `json:"addTime"`
	UpdateTime    time.Time `json:"updateTime"`
	Deleted       int       `json:"deleted"`
}

// ShopAd 首页轮播图
type ShopAd struct {
	ID         int       `json:"id" orm:"column(id)"`
	Name       string    `json:"name"`
	Link       string    `json:"link"`
	URL        string    `json:"url" orm:"column(url)"`
	Position   int       `json:"position"`
	Content    string    `json:"content"`
	StartTime  time.Time `json:"startTime"`
	EndTime    time.Time `json:"endTime"`
	Enabled    bool      `json:"enabled"`
	AddTime    time.Time `json:"addTime"`
	UpdateTime time.Time `json:"updateTime"`
	Deleted    int       `json:"deleted"`
}

// ShopCategory ...
type ShopCategory struct {
	ID         int       `json:"id" orm:"column(id)"`
	Name       string    `json:"name"`
	Keywords   string    `json:"keywords"`
	Desc       string    `json:"desc"`
	PID        int       `json:"pid" orm:"column(pid)"`
	IconURL    string    `json:"iconUrl" orm:"column(icon_url)"`
	PicURL     string    `json:"picUrl" orm:"column(pic_url)"`
	Level      string    `json:"level"`
	SortOrder  int       `json:"sortOrder"`
	AddTime    time.Time `json:"addTime"`
	UpdateTime time.Time `json:"updateTime"`
	Deleted    int       `json:"deleted"`
}

// ShopGoods ...
type ShopGoods struct {
	ID           int       `json:"id" orm:"column(id)"`
	GoodsSn      string    `json:"goodsSn"`
	Name         string    `json:"name"`
	CategoryID   int       `json:"categoryId" orm:"column(category_id)"`
	Gallery      string    `json:"gallery"` // 商品宣传图片列表，采用JSON数组格式
	Keywords     string    `json:"keywords"`
	Brief        string    `json:"brief"`
	IsOnSale     int       `json:"isOnSale"`
	SortOrder    int       `json:"sortOrder"`
	PicURL       string    `json:"picUrl" orm:"column(pic_url)"` // 商品页面商品图片
	ShareURL     string    `json:"shareUrl" orm:"column(share_url)"`
	IsNew        int       `json:"isNew"`
	IsHot        int       `json:"isHot"`
	Unit         string    `json:"unit"`
	CounterPrice float64   `json:"counterPrice"`
	RetailPrice  float64   `json:"retailPrice"`
	Detail       string    `json:"detail"`
	AddTime      time.Time `json:"addTime"`
	UpdateTime   time.Time `json:"updateTime"`
	Deleted      int       `json:"deleted"`
}

// ShopTopic ...
type ShopTopic struct {
	ID         int       `json:"id" orm:"column(id)"`
	Title      string    `json:"title"`
	Subtitle   string    `json:"subtitle"`
	Content    string    `json:"content"`
	ReadCount  string    `json:"readCount"`
	Price      string    `json:"price"`
	PicURL     string    `json:"picUrl" orm:"column(pic_url)"`
	SortOrder  string    `json:"sortOrder" orm:"column(sort_order)"`
	Goods      string    `json:"goods"`
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted    int       `json:"deleted"`
}

type ShopIssue struct {
	ID         int       `json:"id" orm:"column(id)"`
	Question   string    `json:"question"`
	Answer     string    `json:"answer"`
	AddTime    time.Time `json:"addTime"`
	UpdateTime time.Time `json:"updateTime"`
	Deleted    int       `json:"deleted"`
}

// ShopGoodsProduct 商品货品表
type ShopGoodsProduct struct {
	ID             int       `json:"id" orm:"column(id)"`
	GoodsID        int       `json:"goodsId" orm:"column(goods_id)"`
	Specifications string    `json:"specifications"`
	Price          float64   `json:"price"`
	Number         int       `json:"number"`
	URL            string    `json:"url" orm:"column(url)"`
	AddTime        time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime     time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted        int       `json:"deleted"`
}

// 商品规格表
type ShopGoodsSpecification struct {
	ID            int       `json:"id" orm:"column(id)"`
	GoodsID       int       `json:"goodsId" orm:"column(goods_id)"`
	Specification string    `json:"specification"`
	Value         string    `json:"value"`
	PicURL        string    `json:"picUrl" orm:"column(pic_url)"`
	AddTime       time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime    time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted       int       `json:"deleted"`
}

type ShopGoodsAttribute struct {
	ID         int       `json:"id" orm:"column(id)"`
	GoodsID    int       `json:"goodsId" orm:"column(goods_id)"`
	Attribute  string    `json:"attribute"`
	Value      string    `json:"value"`
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted    int       `json:"deleted"`
}

type ShopOrder struct {
	ID              int       `json:"id" orm:"column(id)"`
	UserID          int       `json:"userId" orm:"column(user_id)"`
	OrderSn         string    `json:"orderSn" orm:"column(order_sn)"`
	OrderStatus     int       `json:"orderStatus" orm:"column(order_status)"`
	AftersaleStatus int       `json:"aftersaleStatus" orm:"column(aftersale_status)"`
	Consignee       string    `json:"consignee"`
	Mobile          string    `json:"mobile"`
	Address         string    `json:"address"`
	Message         string    `json:"message"`
	GoodsPrice      float64   `json:"goodsPrice" orm:"column(goods_price)"`
	FreightPrice    float64   `json:"freightPrice" orm:"column(freight_price)"`
	CouponPrice     float64   `json:"couponPrice" orm:"column(coupon_price)"`
	IntegralPrice   float64   `json:"integralPrice" orm:"column(integral_price)"`
	GrouponPrice    float64   `json:"grouponPrice" orm:"column(groupon_price)"`
	OrderPrice      float64   `json:"orderPrice" orm:"column(order_price)"`
	ActualPrice     float64   `json:"actualPrice" orm:"column(actual_price)"`
	PayID           string    `json:"payId" orm:"column(pay_id)"`
	PayTime         time.Time `json:"payTime" orm:"column(pay_time)"`
	ShipSn          string    `json:"shipSn" orm:"column(ship_sn)"`
	ShipChannel     string    `json:"shipChannel" orm:"column(ship_channel)"`
	ShipTime        time.Time `json:"shipTime" orm:"column(ship_time)"`
	RefundAmount    float64   `json:"refundAmount" orm:"column(refund_amount)"`
	RefundType      string    `json:"refundType" orm:"column(refund_type)"`
	RefundContent   string    `json:"refundContent" orm:"column(refund_content)"`
	RefundTime      time.Time `json:"refundTime" orm:"column(refund_time)"`
	ConfirmTime     time.Time `json:"confirmTime" orm:"column(confirm_time)"`
	Comments        int       `json:"comments"`
	EndTime         time.Time `json:"endTime" orm:"column(end_time)"`
	AddTime         time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime      time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted         int       `json:"deleted"`
}

type ShopOrderGoods struct {
	ID             int       `json:"id" orm:"column(id)"`
	OrderID        int       `json:"orderId"  orm:"column(order_id)"`
	GoodsID        int       `json:"goodsId"  orm:"column(goods_id)"`
	GoodsName      string    `json:"goodsName"  orm:"column(goods_name)"`
	GoodsSn        string    `json:"goodsSn"  orm:"column(goods_sn)"`
	ProductID      int       `json:"productId"  orm:"column(product_id)"`
	Number         int       `json:"number"`
	Price          float64   `json:"price"`
	Specifications string    `json:"specifications"`
	PicURL         string    `json:"picUrl"  orm:"column(pic_url)"`
	Comment        int       `json:"comment"`
	AddTime        time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime     time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted        int       `json:"deleted"`
}

type ShopCart struct {
	ID             int       `json:"id" orm:"column(id)"`
	UserID         int       `json:"userId"  orm:"column(user_id)"`
	GoodsID        int       `json:"goodsId"  orm:"column(goods_id)"`
	GoodsSn        string    `json:"goodsSn"  orm:"column(goods_sn)"`
	GoodsName      string    `json:"goodsName"  orm:"column(goods_name)"`
	ProductID      int       `json:"productId"  orm:"column(product_id)"`
	Price          float64   `json:"price"`
	Number         int       `json:"number"`
	Specifications string    `json:"specifications"`
	Checked        int       `json:"checked"`
	PicURL         string    `json:"picUrl"  orm:"column(pic_url)"`
	AddTime        time.Time `json:"addTime" orm:"column(add_time)"`
	UpdateTime     time.Time `json:"updateTime" orm:"column(update_time)"`
	Deleted        int       `json:"deleted"`
}
type ShopAddress struct {
	ID            int       `json:"id" orm:"column(id)"`
	Name          string    `json:"name"`                                 // 收货人名称
	UserID        int       `json:"userId"  orm:"column(user_id)"`        // 用户ID
	Province      string    `json:"province"`                             // 省ID
	City          string    `json:"city"`                                 // 市ID
	County        string    `json:"county"`                               // 区县ID
	AddressDetail string    `json:"addressDetail"`                        // 详细收货地址
	AreaCode      string    `json:"areaCode"`                             // 地区编码
	PostalCode    string    `json:"postalCode"`                           // 邮政编码
	Tel           string    `json:"tel"`                                  // 手机号码
	IsDefault     int       `json:"isDefault"`                            // 是否默认地址
	AddTime       time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime    time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted       int       `json:"deleted"`                              // 逻辑删除
}

type ShopCoupon struct {
	ID         int       `json:"id" orm:"column(id)"`
	Name       string    `json:"name"`                                 // 优惠券名称
	Desc       string    `json:"desc"`                                 // 优惠券介绍，通常是显示优惠券使用限制文字
	Tag        string    `json:"tag"`                                  // 优惠券标签，例如新人专用
	Total      int       `json:"total"`                                // 优惠券数量，如果是0，则是无限量
	Discount   float64   `json:"discount"`                             // 优惠金额
	Min        float64   `json:"min"`                                  // 最少消费金额才能使用优惠券
	Limit      int       `json:"limit"`                                // 用户领券限制数量，如果是0，则是不限制；默认是1，限领一张
	Type       int       `json:"type"`                                 // 优惠券赠送类型，如果是0则通用券，用户领取；如果是1，则是注册赠券；如果是2，则是优惠券码兑换
	Status     int       `json:"Status"`                               // 优惠券状态，如果是0则是正常可用；如果是1则是过期; 如果是2则是下架
	GoodsType  int       `json:"goodsType"`                            // 商品限制类型，如果0则全商品，如果是1则是类目限制，如果是2则是商品限制
	GoodsValue string    `json:"goodsValue"`                           // 商品限制值，goods_type如果是0则空集合，如果是1则是类目集合，如果是2则是商品集合
	Code       string    `json:"code"`                                 // 优惠券兑换码
	TimeType   int       `json:"timeType"`                             // 有效时间限制，如果是0，则基于领取时间的有效天数days；如果是1，则start_time和end_time是优惠券有效期
	Days       int       `json:"days"`                                 // 基于领取时间的有效天数
	StartTime  time.Time `json:"startTime"`                            // 使用券开始时间
	EndTime    time.Time `json:"endTime"`                              // 使用券截至时间
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`                              // 逻辑删除
}

type ShopCouponUser struct {
	ID         int       `json:"id" orm:"column(id)"`
	UserID     int       `json:"userId" orm:"column(user_id)"`         // 用户ID
	CouponID   int       `json:"couponId" orm:"column(coupon_id)"`     // 优惠卷ID
	Status     int       `json:"status"`                               // 使用状态, 如果是0则未使用；如果是1则已使用；如果是2则已过期；如果是3则已经下架
	UsedTime   time.Time `json:"usedTime"`                             // 使用时间
	StartTime  time.Time `json:"startTime"`                            // 有效期开始时间
	EndTime    time.Time `json:"endTime"`                              // 有效期截至时间
	OrderID    int       `json:"orderId" orm:"column(order_id)"`       // 订单ID
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`                              // 逻辑删除
}

type ShopCollect struct {
	ID         int       `json:"id" orm:"column(id)"`
	UserID     int       `json:"userId"  orm:"column(user_id)"`        // 用户ID
	ValueID    int       `json:"valueId" orm:"column(value_id)"`       // 商品ID
	Type       int       `json:"type"`                                 // 收藏类型，如果type=0，则是商品ID；如果type=1，则是专题ID
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`                              // 逻辑删除
}

type ShopComment struct {
	ID           int       `json:"id" orm:"column(id)"`
	ValueID      int       `json:"valueId" orm:"column(value_id)"`       // 商品ID
	Type         int       `json:"type"`                                 // 评论类型，如果type=0，则是商品评论；如果是type=1，则是专题评论；如果type=3，则是订单商品评论。
	Content      string    `json:"content"`                              // 评论内容
	AdminContent string    `json:"adminContent"`                         // 管理员回复
	UserID       int       `json:"userId" orm:"column(user_id)"`         // 用户ID
	HasPicture   int       `json:"hasPicture"`                           // 是否含有图片
	PicURLs      string    `json:"picUrls" orm:"column(pic_urls)"`       // 图片地址列表，采用JSON数组格式
	Star         int       `json:"star"`                                 // 评分， 1-5
	AddTime      time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime   time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted      int       `json:"deleted"`                              // 逻辑删除
}

type ShopKeyword struct {
	ID         int       `json:"id" orm:"column(id)"`
	Keyword    string    `json:"keyword"`                              // 关键字
	URL        string    `json:"url" orm:"column(url)"`                // 关键字的跳转链接
	IsHot      int       `json:"isHot"`                                // 是否是热门关键字
	IsDefault  int       `json:"IsDefault"`                            // 是否是默认关键字
	SortOrder  int       `json:"sortOrder"`                            // 排序
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`                              // 逻辑删除
}

type ShopSearchHistory struct {
	ID         int       `json:"id" orm:"column(id)"`
	UserID     int       `json:"userId" orm:"column(user_id)"`         // 用户ID
	Keyword    string    `json:"keyword"`                              // 关键字
	From       string    `json:"from"`                                 // 搜索来源，如pc、wx、app
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`                              // 逻辑删除
}

type ShopRegion struct {
	ID   int    `json:"id" orm:"column(id)"`
	PID  int    `json:"pid" orm:"column(pid)"` // 行政区域父ID，例如区县的pid指向市，市的pid指向省，省的pid则是0
	Name string `json:"name"`                  // 行政区域名称
	Type int    `json:"type"`                  // 行政区域类型，如如1则是省， 如果是2则是市，如果是3则是区县
	Code int    `json:"code"`                  // 行政区域编码
}

type ShopFootprint struct {
	ID         int       `json:"id" orm:"column(id)"`
	UserID     int       `json:"userId" orm:"column(user_id)"`         // 用户ID
	GoodsID    int       `json:"goodsId" orm:"column(goods_id)"`       // 浏览商品ID
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`
}

type ShopFeedback struct {
	ID         int       `json:"id" orm:"column(id)"`
	UserID     int       `json:"userId" orm:"column(user_id)"`         // 用户ID
	Username   string    `json:"username"`                             // 用户名称
	Mobile     string    `json:"mobile"`                               // 手机号
	FeedType   string    `json:"feedType"`                             // 反馈类型
	Content    string    `json:"content"`                              // 反馈内容
	Status     int       `json:"status"`                               // 状态
	HasPicture int       `json:"hasPicture"`                           // 是否含有图片
	PicURLs    string    `json:"picUrls" orm:"column(pic_urls)"`       // 图片地址列表，采用JSON数组格式
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`
}

// ShopAbout 关于我们
type ShopAbout struct {
	ID         int       `json:"id" orm:"column(id)"`
	Title      string    `json:"title"`
	Context    string    `json:"context"`
	AddTime    time.Time `json:"addTime"`
	UpdateTime time.Time `json:"updateTime"`
	Deleted    int       `json:"deleted"`
}

// ShopOrderExpress 物流
type ShopOrderExpress struct {
	ID           int       `json:"id" orm:"column(id)"`
	IsFinish     int       `json:"isFinish"`
	LogisticCode string    `json:"logisticCode"`
	OrderID      int       `json:"orderId" orm:"column(order_id)"`
	RequestCount int       `json:"requestCount"`
	RequestTime  time.Time `json:"requestTime"`
	ShipperCode  string    `json:"shipperCode"`
	ShipperID    int       `json:"shipperId" orm:"column(shipper_id)"`
	ShipperName  string    `json:"shipperName"`
	Traces       string    `json:"traces"`
	AddTime      time.Time `json:"addTime"`
	UpdateTime   time.Time `json:"updateTime"`
}

type ShopAftersale struct {
	ID          int       `json:"id" orm:"column(id)"`
	AftersaleSn string    `json:"aftersaleSn"`                          // 售后编号
	OrderID     int       `json:"orderId" orm:"column(order_id)"`       // 订单ID
	UserID      int       `json:"userId" orm:"column(user_id)"`         // 用户ID
	Type        int       `json:"type"`                                 // 售后类型，0是未收货退款，1是已收货（无需退货）退款，2用户退货退款
	Reason      string    `json:"reason"`                               // 退款原因
	Amount      float64   `json:"amount"`                               // 退款金额
	Pictures    string    `json:"pictures"`                             // 退款凭证图片链接数组
	Comment     string    `json:"comment"`                              // 退款说明
	Status      int       `json:"status"`                               // 售后状态，0是可申请，1是用户已申请，2是管理员审核通过，3是管理员退款成功，4是管理员审核拒绝，5是用户已取消
	HandleTime  time.Time `json:"handleTime"`                           // 管理员操作时间
	AddTime     time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime  time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted     int       `json:"deleted"`
}

type ShopSystem struct {
	ID         int       `json:"id" orm:"column(id)"`
	KeyName    string    `json:"KeyName"`                              // 设置类型
	KeyValue   string    `json:"KeyValue"`                             // 设置的值
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`
}

type ShopBrand struct {
	ID         int       `json:"id" orm:"column(id)"`
	Name       string    `json:"name"`
	Desc       string    `json:"desc"`
	PicURL     string    `json:"picUrl" orm:"column(pic_url)"`
	SortOrder  int       `json:"sortOrder"`
	FloorPrice float64   `json:"floorPrice"`
	AddTime    time.Time `json:"addTime" orm:"column(add_time)"`       // 创建时间
	UpdateTime time.Time `json:"updateTime" orm:"column(update_time)"` // 更新时间
	Deleted    int       `json:"deleted"`
}
