// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/utils"
)

func init() {
	orm.RegisterModel(new(base.ShopGoodsAttribute))
}

//  GetGoodsAttrList 根据商品ID获取商品属性
func GetGoodsAttrList(goodsID int) (goodsAttrList []response.GoodsAttributeRtnJSON, err error) {
	var goodsAttr []base.ShopGoodsAttribute
	var goodsAttrRtnJSON response.GoodsAttributeRtnJSON
	o := orm.NewOrm()
	goodsProductTable := new(base.ShopGoodsAttribute)
	_, err = o.QueryTable(goodsProductTable).Filter("goodsID", goodsID).All(&goodsAttr)
	if err != nil {
		return goodsAttrList, err
	}

	for _, v := range goodsAttr {
		goodsAttrRtnJSON.ID = v.ID
		goodsAttrRtnJSON.GoodsID = v.GoodsID
		goodsAttrRtnJSON.Attribute = v.Attribute
		goodsAttrRtnJSON.Value = v.Value
		goodsAttrRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		goodsAttrRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		goodsAttrList = append(goodsAttrList, goodsAttrRtnJSON)
	}
	return goodsAttrList, nil
}
