// Copyright 2020 The shop Authors

// Package common implements common.
package common

import "errors"

const (
	// 返回体code
	// SCODE 成功码
	SCODE = 0
	// FCODE 失败码
	FCODE = 1
	// AUTH_INVALID_ACCOUNT 账户认证失败
	AUTH_INVALID_ACCOUNT = 501

	// 返回体msg
	// SUCCESEE ...
	SUCCESEE = "success"
	// FAILED ...
	FAILED = "failed"

	// 订单状态
	// STATUS_CREATE 创建订单
	STATUS_CREATE = 101
	// STATUS_PAY 订单付款
	STATUS_PAY = 201
	// STATUS_SHIP 订单发货
	STATUS_SHIP = 301
	// STATUS_CONFIRM 订单确认
	STATUS_CONFIRM = 401
	// STATUS_CANCEL 取消订单
	STATUS_CANCEL = 102
	// STATUS_AUTO_CANCEL 系统自动取消
	STATUS_AUTO_CANCEL = 103
	// STATUS_ADMIN_CANCEL 管理员取消
	STATUS_ADMIN_CANCEL = 104
	// STATUS_REFUND 售后
	STATUS_REFUND = 202
	// STATUS_REFUND_CONFIRM 售后确认
	STATUS_REFUND_CONFIRM = 203
	// STATUS_AUTO_CONFIRM 售后自动确认
	STATUS_AUTO_CONFIRM = 402
	// STATUS_COMMENT 已评论
	STATUS_COMMENT = 501 // 已经评论
)

var (
	// 首页显示设置
	// SHOP_WX_INDEX_NEW ...
	SHOP_WX_INDEX_NEW = "shop_wx_index_new"
	// SHOP_WX_INDEX_HOT ...
	SHOP_WX_INDEX_HOT = "shop_wx_index_hot"
	// SHOP_WX_INDEX_BRAND ...
	SHOP_WX_INDEX_BRAND = "shop_wx_index_brand"
	// SHOP_WX_INDEX_TOPIC ...
	SHOP_WX_INDEX_TOPIC = "shop_wx_index_topic"
	// SHOP_WX_INDEX_CATLOG_LIST ...
	SHOP_WX_INDEX_CATLOG_LIST = "shop_wx_catlog_list"
	// SHOP_WX_INDEX_CATLOG_GOODS ...
	SHOP_WX_INDEX_CATLOG_GOODS = "shop_wx_catlog_goods"
	// SHOP_WX_SHARE ...
	SHOP_WX_SHARE = "shop_wx_share"
	// SHOP_WX_INDEX_COUPON ...
	SHOP_WX_INDEX_COUPON = "shop_wx_index_coupon"
	// 运费相关配置
	// SHOP_EXPRESS_FREIGHT_VALUE ...
	SHOP_EXPRESS_FREIGHT_VALUE = "shop_express_freight_value"
	// SHOP_EXPRESS_FREIGHT_MIN ...
	SHOP_EXPRESS_FREIGHT_MIN = "shop_express_freight_min"
	// 订单相关配置
	// SHOP_ORDER_UNPAID ...
	SHOP_ORDER_UNPAID = "shop_order_unpaid"
	// SHOP_ORDER_UNCONFIRM ...
	SHOP_ORDER_UNCONFIRM = "shop_order_unconfirm"
	// SHOP_ORDER_COMMENT ...
	SHOP_ORDER_COMMENT = "shop_order_comment"
	// 商场相关配置
	// SHOP_MALL_NAME ...
	SHOP_MALL_NAME = "shop_mall_name"
	// SHOP_MALL_ADDRESS ...
	SHOP_MALL_ADDRESS = "shop_mall_address"
	// SHOP_MALL_PHONE ...
	SHOP_MALL_PHONE = "shop_mall_phone"
	// SHOP_MALL_QQ ...
	SHOP_MALL_QQ = "shop_mall_qq"
	// SHOP_MALL_LONGITUDE ...
	SHOP_MALL_LONGITUDE = "shop_mall_longitude"
	// SHOP_MALL_Latitude ...
	SHOP_MALL_Latitude = "shop_mall_latitude"
)

// 错误消息err
var (
	// ErrAuth ...
	ErrAuth = errors.New("account auth failed")
	// ErrNoRow ...
	ErrNoRow = errors.New("<QuerySeter> no row found")
	// ErrInternal ...
	ErrInternal = errors.New("internal service err")
	// ErrStructJSON ...
	ErrStructJSON = errors.New("Failed to Marshal or Unmarshal info")

	// ErrCreateData ...
	ErrCreateData = errors.New("create info failed")
	// ErrGetData ...
	ErrGetData = errors.New("get info failed")
	// ErrAddData ...
	ErrAddData = errors.New("add info failed")
	// ErrUpdateData ...
	ErrUpdateData = errors.New("update info failed")
	// ErrDeleteData ...
	ErrDeleteData = errors.New("delete info failed")

	// ErrNoEnough ...
	ErrNoEnough = errors.New("库存不足")
	// ErrSelect ...
	ErrSelect = errors.New("选择出错")
)
