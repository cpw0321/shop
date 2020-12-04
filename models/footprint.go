// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"strings"
	"time"
)

func init() {
	orm.RegisterModel(new(base.ShopFootprint))
}

// GetAllFootprint 获取足迹，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllFootprint(query map[string]string, sortby string, order string,
	page int, limit int) (footprintList []base.ShopFootprint, err error) {
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopFootprint))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
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

	_, err = qs.Limit(limit, offset).All(&footprintList)
	if err != nil {
		return footprintList, err
	}
	return footprintList, nil
}

// DeleteFootprint 删除足迹
func DeleteFootprint(ID int) (err error) {
	o := orm.NewOrm()
	footprintTable := new(base.ShopFootprint)
	_, err = o.QueryTable(footprintTable).Filter("ID", ID).Update(orm.Params{
		"deleted": 1})
	if err != nil {
		return err
	}
	return nil
}

// AddFootprint 添加足迹
func AddFootprint(footprint base.ShopFootprint) (err error) {
	o := orm.NewOrm()
	footprintTable := new(base.ShopFootprint)
	// 查询到更新，没有查询到新增
	err = o.QueryTable(footprintTable).Filter("UserID", footprint.UserID).Filter("GoodsID", footprint.GoodsID).One(&footprint)
	if err != nil {
		_, err = o.Insert(&footprint)
		if err != nil {
			return err
		}
	}
	_, err = o.QueryTable(footprintTable).Filter("UserID", footprint.UserID).Filter("GoodsID", footprint.GoodsID).Update(orm.Params{
		"UpdateTime": time.Now()})
	if err != nil {
		return err
	}
	return nil
}
