// Copyright 2020 The shop Authors

// Package common implements common.
package common

import (
	"shop/logger"
	"shop/models"
	"time"
)

// HandleOrder 处理订单超时未处理
// 101 订单生成未支付订单超过24h自动关闭为103，下单未支付超期系统自动取消，
// 402 用户没有确认收货，但是快递反馈已收货后，超过一定时间，系统自动确认收货，订单结束
// 401 用户确认收货，订单结束，订单超过48小时未评论则自动关闭, 为501
func HandleOrder() {
	for {
		// 处理未付款订单
		err := alterOrder(101, 103, 24*3600)
		if err != nil {
			logger.Logger.Errorf("hander no pay order is failed! err:[%v].", err)
		}

		// 处理未评论订单
		err = alterOrder(401, 501, 2*24*3600)
		if err != nil {
			logger.Logger.Errorf("hander no comment order is failed! err:[%v].", err)
		}

		time.Sleep(60 * time.Minute) // 1小时处理一次
	}
}

// alterOrder 修改订单
func alterOrder(orderStatus int, status int, setTime int64) (err error) {
	orderList, err := models.GetOrderListByStatus(orderStatus)
	if err != nil {
		logger.Logger.Errorf("GetOrderListByStatus is failed! err:[%v].", err)
		return err
	}
	for _, v := range orderList {
		nowTime := time.Now().Unix()
		intervalTime := nowTime - v.AddTime.Unix()
		if intervalTime > setTime { // 3600 1小时
			v.OrderStatus = status
			v.UpdateTime = time.Now()
			err = models.UpdateOrder(v)
			if err != nil {
				logger.Logger.Errorf("UpdateOrder is failed! err:[%v].", err)
			}
		}
	}
	return nil
}
