// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"shop/base"
	"shop/base/response"
	"shop/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.ShopCategory))
}

// GetAllCategory 获取分类信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllCategory(query map[string]string, sortby string, order string,
	page int, limit int) (categoryRtnJSONList []response.CategoryRtnJSON, err error) {
	var categoryList []base.ShopCategory
	var categoryRtnJSON response.CategoryRtnJSON
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopCategory))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
	}
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
	_, err = qs.Limit(limit, offset).All(&categoryList)
	if err != nil {
		return nil, err
	}

	for _, v := range categoryList {
		categoryRtnJSON.ID = v.ID
		categoryRtnJSON.Name = v.Name
		categoryRtnJSON.Keywords = v.Keywords
		categoryRtnJSON.Desc = v.Desc
		categoryRtnJSON.PID = v.PID
		categoryRtnJSON.IconURL = v.IconURL
		categoryRtnJSON.PicURL = v.PicURL
		categoryRtnJSON.Level = v.Level
		categoryRtnJSON.SortOrder = v.SortOrder
		categoryRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		categoryRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		categoryRtnJSONList = append(categoryRtnJSONList, categoryRtnJSON)
	}

	return categoryRtnJSONList, nil
}

// GetOneCategory 获取当前分类信息
func GetOneCategory(ID int) (categoryRtnJSON response.CategoryRtnJSON, err error) {
	var category base.ShopCategory
	o := orm.NewOrm()
	goodsTable := new(base.ShopCategory)
	err = o.QueryTable(goodsTable).Filter("ID", ID).One(&category)
	if err != nil {
		return categoryRtnJSON, err
	}

	categoryRtnJSON.ID = category.ID
	categoryRtnJSON.Name = category.Name
	categoryRtnJSON.Keywords = category.Keywords
	categoryRtnJSON.Desc = category.Desc
	categoryRtnJSON.PID = category.PID
	categoryRtnJSON.IconURL = category.IconURL
	categoryRtnJSON.PicURL = category.PicURL
	categoryRtnJSON.Level = category.Level
	categoryRtnJSON.SortOrder = category.SortOrder
	categoryRtnJSON.AddTime = utils.FormatTimestampStr(category.AddTime)
	categoryRtnJSON.UpdateTime = utils.FormatTimestampStr(category.UpdateTime)

	return categoryRtnJSON, nil
}
