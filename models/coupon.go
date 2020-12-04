// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"shop/base"
	"shop/base/response"
	"shop/utils"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.ShopCoupon))
}

// GetAllCoupon 获取分类信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllCoupon(query map[string]string, sortby string, order string,
	page int, limit int) (couponRtnJSONList []response.CouponRtnJSON, err error) {
	couponRtnJSONList = []response.CouponRtnJSON{}
	var couponList []base.ShopCoupon
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopCoupon))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
	}

	// order
	orderby := ""
	if order == "desc" {
		orderby = "-" + sortby
	} else {
		orderby = sortby
	}

	qs = qs.OrderBy(orderby)
	_, err = qs.Limit(limit, offset).All(&couponList)
	if err != nil {
		return couponRtnJSONList, err
	}

	var couponRtnJSON response.CouponRtnJSON
	for _, v1 := range couponList {
		couponRtnJSON.ID = v1.ID
		couponRtnJSON.Name = v1.Name
		couponRtnJSON.Desc = v1.Desc
		couponRtnJSON.Tag = v1.Tag
		couponRtnJSON.Discount = v1.Discount
		couponRtnJSON.Min = v1.Min
		couponRtnJSON.Type = v1.Type
		couponRtnJSON.Status = v1.Status
		couponRtnJSON.GoodsType = v1.GoodsType
		couponRtnJSON.GoodsValue = v1.GoodsValue
		couponRtnJSON.TimeType = v1.TimeType
		couponRtnJSON.Days = v1.Days
		couponRtnJSON.StartTime = utils.FormatTimestampStr(v1.StartTime)
		couponRtnJSON.EndTime = utils.FormatTimestampStr(v1.EndTime)
		couponRtnJSON.AddTime = utils.FormatTimestampStr(v1.AddTime)
		couponRtnJSON.UpdateTime = utils.FormatTimestampStr(v1.UpdateTime)
		couponRtnJSONList = append(couponRtnJSONList, couponRtnJSON)
	}

	return couponRtnJSONList, nil
}

// GetOneCoupon 获取一个优惠卷的详细信息
func GetOneCoupon(ID int) (coupon base.ShopCoupon, err error) {
	o := orm.NewOrm()
	couponTable := new(base.ShopCoupon)
	err = o.QueryTable(couponTable).Filter("ID", ID).One(&coupon)
	if err != nil {
		return coupon, err
	}
	return coupon, nil
}

// CouponReceive 优惠券领取
func CouponReceive(UserID int, CouponID int) error {
	var coupon base.ShopCoupon
	o := orm.NewOrm()
	couponTable := new(base.ShopCoupon)
	err := o.QueryTable(couponTable).Filter("ID", CouponID).One(&coupon)
	if err != nil {
		return err
	}

	couponUserData := base.ShopCouponUser{
		UserID:     UserID,
		CouponID:   coupon.ID,
		Status:     coupon.Status,
		StartTime:  time.Now(),
		EndTime:    time.Now().AddDate(0, 0, coupon.Days),
		AddTime:    time.Now(),
		UpdateTime: time.Now(),
	}
	_, err = o.Insert(&couponUserData)
	if err != nil {
		return err
	}

	return nil
}

// CouponExchange 优惠券兑换
func CouponExchange(userID int, code string) error {
	var coupon base.ShopCoupon
	o := orm.NewOrm()
	couponTable := new(base.ShopCoupon)
	err := o.QueryTable(couponTable).Filter("Code", code).One(&coupon)
	if err != nil {
		return err
	}

	couponUserData := base.ShopCouponUser{
		UserID:    userID,
		CouponID:  coupon.ID,
		Status:    coupon.Status,
		StartTime: coupon.StartTime,
		EndTime:   coupon.EndTime,
		AddTime:   time.Now(),
	}
	_, err = o.Insert(&couponUserData)
	if err != nil {
		return err
	}

	return nil
}
