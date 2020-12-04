// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"encoding/json"
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
	"time"
)

func init() {
	orm.RegisterModel(new(base.ShopOrderExpress))
}

// GetLatestOrderExpress 获取最新的订单物流信息
func GetLatestOrderExpress(orderID int) (expressInfo response.ExpressRtnInfo) {
	expressInfo = response.ExpressRtnInfo{
		ShipperCode:  "",
		ShipperName:  "",
		LogisticCode: "",
		IsFinish:     0,
		RequestTime:  "",
		Traces:       make([]response.Traces, 0),
	}

	o := orm.NewOrm()
	orderexpressTable := new(base.ShopOrderExpress)
	var orderexpress base.ShopOrderExpress
	err := o.QueryTable(orderexpressTable).Filter("orderID", orderID).One(&orderexpress)
	if err != nil {
		return expressInfo
	}

	expressInfo.ShipperCode = orderexpress.ShipperCode
	expressInfo.ShipperName = orderexpress.ShipperName
	expressInfo.LogisticCode = orderexpress.LogisticCode
	expressInfo.IsFinish = orderexpress.IsFinish
	expressInfo.RequestTime = utils.FormatTimestampStr(time.Now())
	var res []response.Traces
	err = json.Unmarshal([]byte(orderexpress.Traces), &res)
	if err != nil {
		logger.Logger.Errorf("JSON Unmarshal Traces is failed! err:[%v].", err)
	}
	expressInfo.Traces = res

	return expressInfo
}

// GetExpressByOrderID 根据订单ID查询物流信息
func GetExpressByOrderID(orderID int) (expressInfo base.ShopOrderExpress, err error) {
	o := orm.NewOrm()
	orderexpressTable := new(base.ShopOrderExpress)
	err = o.QueryTable(orderexpressTable).Filter("OrderID", orderID).One(&expressInfo)
	return expressInfo, err
}

// UpdataOrderExpress 更新物流信息
func UpdataOrderExpress(expressInfo base.ShopOrderExpress) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&expressInfo)
	return err
}

// AddAddress 新增物流信息
func AddOrderExpress(express base.ShopOrderExpress) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&express)
	if err != nil {
		return err
	}
	return nil
}
