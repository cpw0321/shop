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
	"strings"
)

func init() {
	orm.RegisterModel(new(base.ShopAftersale))
}

// GetAllAftersale
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllAftersale(query map[string]string, sortby string, order string,
	page int, limit int) (aftersaleResponList []response.AftersaleRtn, err error) {
	var aftersaleList []base.ShopAftersale
	var aftersaleRespon response.AftersaleRespon
	var aftersaleRtn response.AftersaleRtn
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopAftersale))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		switch k {
		case "name":
			qs = qs.Filter("name__contains", v)
		default:
			qs = qs.Filter(k, v)
		}
	}

	// order
	if sortby != "" {
		orderby := ""
		if order == "desc" {
			orderby = "-" + sortby
		} else {
			orderby = sortby
		}
		qs = qs.OrderBy(orderby)
	}
	_, err = qs.Limit(limit, offset).All(&aftersaleList)
	if err != nil {
		return aftersaleResponList, err
	}

	for _, v := range aftersaleList {
		aftersaleRespon.ID = v.ID
		aftersaleRespon.AftersaleSn = v.AftersaleSn
		aftersaleRespon.OrderID = v.OrderID
		aftersaleRespon.UserID = v.UserID
		aftersaleRespon.Type = v.Type
		aftersaleRespon.Reason = v.Reason
		aftersaleRespon.Amount = v.Amount
		var pictures []string
		err = json.Unmarshal([]byte(v.Pictures), &pictures)
		if err != nil {
			logger.Logger.Errorf("get pictures is failed! aftersale:[%v] aftersale.Pictures:[%v].", v, v.Pictures)

		}
		aftersaleRespon.Pictures = pictures
		aftersaleRespon.Comment = v.Comment
		aftersaleRespon.Status = v.Status
		aftersaleRespon.HandleTime = utils.FormatTimestampStr(v.HandleTime)
		aftersaleRespon.AddTime = utils.FormatTimestampStr(v.AddTime)
		aftersaleRespon.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)

		// 根据orderID获取goodsList
		goodsList, err := GetOrderGoodsListByOrderID(v.OrderID)
		if err != nil {
			logger.Logger.Error("GetOrderGoodsListByOrderID is failed! err:[%v].", err)
		}

		aftersaleRtn.AftersaleRespon = aftersaleRespon
		aftersaleRtn.GoodsList = goodsList

		aftersaleResponList = append(aftersaleResponList, aftersaleRtn)
	}

	return aftersaleResponList, nil
}

// GetAftersaleByOrderID 获取售后信息通过订单ID
func GetAftersaleByOrderID(orderID int) (aftersaleRespon response.AftersaleRespon, err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopAftersale)
	var aftersale base.ShopAftersale
	err = o.QueryTable(orderTable).Filter("OrderID", orderID).One(&aftersale)
	if err != nil {
		logger.Logger.Errorf("query aftersale by orderID is failed! err:[%v].", err)
		return aftersaleRespon, err
	}

	aftersaleRespon.ID = aftersale.ID
	aftersaleRespon.AftersaleSn = aftersale.AftersaleSn
	aftersaleRespon.OrderID = aftersale.OrderID
	aftersaleRespon.UserID = aftersale.UserID
	aftersaleRespon.Type = aftersale.Type
	aftersaleRespon.Reason = aftersale.Reason
	aftersaleRespon.Amount = aftersale.Amount
	var pictures []string
	pictures = []string{}
	err = json.Unmarshal([]byte(aftersale.Pictures), &pictures)
	if err != nil {
		logger.Logger.Errorf("get pictures is failed! aftersale:[%v] aftersale.Pictures:[%v].", aftersale, aftersale.Pictures)
	}
	aftersaleRespon.Pictures = pictures
	aftersaleRespon.Comment = aftersale.Comment
	aftersaleRespon.Status = aftersale.Status
	aftersaleRespon.HandleTime = utils.FormatTimestampStr(aftersale.HandleTime)
	aftersaleRespon.AddTime = utils.FormatTimestampStr(aftersale.AddTime)
	aftersaleRespon.UpdateTime = utils.FormatTimestampStr(aftersale.UpdateTime)

	return aftersaleRespon, nil
}

// AddAftersale 添加售后
func AddAftersale(m *base.ShopAftersale) (ID int64, err error) {
	o := orm.NewOrm()
	ID, err = o.Insert(m)
	return
}

// GetAftersaleByOrderIDIsExist 判断订单是否存在售后列表
func GetAftersaleByOrderIDIsExist(orderID int) (isOk bool) {
	o := orm.NewOrm()
	orderTable := new(base.ShopAftersale)
	isOk = o.QueryTable(orderTable).Filter("OrderID", orderID).Filter("Deleted", 0).Exist()
	return
}
