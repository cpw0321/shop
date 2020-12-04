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
	orm.RegisterModel(new(base.ShopGoodsSpecification))
}

// GetGoodsSpecificationList 获取单个商品的规格列表
func GetGoodsSpecification(goodsID int) (goodsSpecificationList []response.GoodsSpecificationRtnJSON, err error) {
	var specificationList []base.ShopGoodsSpecification
	var goodsSpecificationRtnJSON response.GoodsSpecificationRtnJSON
	o := orm.NewOrm()

	goodsSpecificationTable := new(base.ShopGoodsSpecification)
	_, err = o.QueryTable(goodsSpecificationTable).Filter("GoodsID", goodsID).All(&specificationList)
	if err != nil {
		return goodsSpecificationList, err
	}

	for _, v := range specificationList {
		goodsSpecificationRtnJSON.ID = v.ID
		goodsSpecificationRtnJSON.GoodsID = v.GoodsID
		goodsSpecificationRtnJSON.Specification = v.Specification
		goodsSpecificationRtnJSON.Value = v.Value
		goodsSpecificationRtnJSON.PicURL = v.PicURL
		goodsSpecificationRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		goodsSpecificationRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		goodsSpecificationList = append(goodsSpecificationList, goodsSpecificationRtnJSON)
	}
	return goodsSpecificationList, nil
}

// GetInfoBySpecification 根据规格名称获取相同种类的规格列表
func GetInfoBySpecification(goodsID int, specification string) (goodsSpecificationList []response.GoodsSpecificationRtnJSON, err error) {
	var specificationList []base.ShopGoodsSpecification
	var goodsSpecificationRtnJSON response.GoodsSpecificationRtnJSON

	o := orm.NewOrm()

	goodsSpecificationTable := new(base.ShopGoodsSpecification)
	_, err = o.QueryTable(goodsSpecificationTable).Filter("GoodsID", goodsID).Filter("Specification", specification).All(&specificationList)
	if err != nil {
		return goodsSpecificationList, err
	}

	for _, v := range specificationList {
		goodsSpecificationRtnJSON.ID = v.ID
		goodsSpecificationRtnJSON.GoodsID = v.GoodsID
		goodsSpecificationRtnJSON.Specification = v.Specification
		goodsSpecificationRtnJSON.Value = v.Value
		goodsSpecificationRtnJSON.PicURL = v.PicURL
		goodsSpecificationRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		goodsSpecificationRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		goodsSpecificationList = append(goodsSpecificationList, goodsSpecificationRtnJSON)
	}

	return goodsSpecificationList, nil
}
