// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"encoding/json"
	"shop/base"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
	"strings"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.ShopGoods))
}

// GetAllGoods 获取商品信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllGoods(query map[string]string, sortby string, order string,
	page int, limit int) (goodsRtnJSONList []response.GoodsRtnJSON, err error) {
	goodsRtnJSONList = []response.GoodsRtnJSON{}
	var goods []base.ShopGoods
	var goodsRtnJSON response.GoodsRtnJSON
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopGoods))
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

	_, err = qs.Limit(limit, offset).All(&goods)
	if err != nil {
		return goodsRtnJSONList, err
	}

	// 将商品中的宣传图片列表，采用JSON数组格式
	for _, v := range goods {
		var gallery []string
		err = json.Unmarshal([]byte(v.Gallery), &gallery)
		if err != nil {
			logger.Logger.Errorf("get goods info is failed! goods:[%v].", v)
		}

		goodsRtnJSON.ID = v.ID
		goodsRtnJSON.CategoryID = v.CategoryID
		goodsRtnJSON.Name = v.Name
		goodsRtnJSON.Brief = v.Brief
		goodsRtnJSON.PicURL = v.PicURL
		goodsRtnJSON.CounterPrice = v.CounterPrice
		goodsRtnJSON.RetailPrice = v.RetailPrice
		goodsRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		goodsRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		goodsRtnJSONList = append(goodsRtnJSONList, goodsRtnJSON)
	}

	return goodsRtnJSONList, nil
}

// GetOneGoods 获取单个商品详情
func GetOneGoods(ID int) (goodsRtnJSON response.GoodsRtnJSON, err error) {
	var goods base.ShopGoods
	o := orm.NewOrm()
	goodsTable := new(base.ShopGoods)
	err = o.QueryTable(goodsTable).Filter("ID", ID).One(&goods)
	if err != nil {
		return goodsRtnJSON, err
	}

	var gallery []string
	err = json.Unmarshal([]byte(goods.Gallery), &gallery)
	if err != nil {
		logger.Logger.Errorf("get goods info is failed! goods:[%v].", goods)
	}
	goodsRtnJSON.ID = goods.ID
	goodsRtnJSON.GoodsSn = goods.GoodsSn
	goodsRtnJSON.Name = goods.Name
	goodsRtnJSON.CategoryID = goods.CategoryID
	goodsRtnJSON.Gallery = gallery
	goodsRtnJSON.Keywords = goods.Keywords
	goodsRtnJSON.Brief = goods.Brief
	goodsRtnJSON.PicURL = goods.PicURL
	goodsRtnJSON.ShareURL = goods.ShareURL
	goodsRtnJSON.Unit = goods.Unit
	goodsRtnJSON.CounterPrice = goods.CounterPrice
	goodsRtnJSON.RetailPrice = goods.RetailPrice
	goodsRtnJSON.Detail = goods.Detail
	goodsRtnJSON.AddTime = utils.FormatTimestampStr(goods.AddTime)
	goodsRtnJSON.UpdateTime = utils.FormatTimestampStr(goods.UpdateTime)
	return goodsRtnJSON, nil
}

// GetGoodsCount 获取在售商品总数
func GetGoodsCount() (count int64, err error) {
	o := orm.NewOrm()
	goodsTable := new(base.ShopGoods)

	count, err = o.QueryTable(goodsTable).Filter("IsOnSale", 1).Count()
	if err != nil {
		return count, err
	}
	return count, nil
}
