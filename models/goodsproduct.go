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
)

func init() {
	orm.RegisterModel(new(base.ShopGoodsProduct))
}

// GetGoodsProduct 根据商品ID获取商品规格信息，比如：标准规格对应price 100
func GetGoodsProduct(goodsID int) (goodsProductList []response.GoodsProductRtnJSON, err error) {
	var goodsProduct []base.ShopGoodsProduct
	var goodsProductRtnJSON response.GoodsProductRtnJSON
	o := orm.NewOrm()
	goodsProductTable := new(base.ShopGoodsProduct)
	_, err = o.QueryTable(goodsProductTable).Filter("goodsID", goodsID).All(&goodsProduct)
	if err != nil {
		return goodsProductList, err
	}

	for _, v := range goodsProduct {
		var specifications []string
		err = json.Unmarshal([]byte(v.Specifications), &specifications)
		if err != nil {
			logger.Logger.Errorf("get goodsProduct is failed! goodsProduct:[%v].", v)
		}
		goodsProductRtnJSON.ID = v.ID
		goodsProductRtnJSON.GoodsID = v.GoodsID
		goodsProductRtnJSON.Specifications = specifications
		goodsProductRtnJSON.Price = v.Price
		goodsProductRtnJSON.Number = v.Number
		goodsProductRtnJSON.URL = v.URL
		goodsProductRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		goodsProductRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		goodsProductList = append(goodsProductList, goodsProductRtnJSON)
	}
	return goodsProductList, nil
}

// GetGoodsOneProductByID 根据规格ID获取商品一个规格信息
func GetGoodsOneProductByID(ID int) (goodsProductRtnJSON response.GoodsProductRtnJSON, err error) {
	var goodsProduct base.ShopGoodsProduct
	o := orm.NewOrm()
	goodsProductTable := new(base.ShopGoodsProduct)
	err = o.QueryTable(goodsProductTable).Filter("ID", ID).One(&goodsProduct)
	if err != nil {
		return goodsProductRtnJSON, err
	}

	var specifications []string
	err = json.Unmarshal([]byte(goodsProduct.Specifications), &specifications)
	if err != nil {
		logger.Logger.Errorf("get goodsProduct is failed! goodsProduct:[%v].", goodsProduct)
	}
	goodsProductRtnJSON.ID = goodsProduct.ID
	goodsProductRtnJSON.GoodsID = goodsProduct.GoodsID
	goodsProductRtnJSON.Specifications = specifications
	goodsProductRtnJSON.Price = goodsProduct.Price
	goodsProductRtnJSON.Number = goodsProduct.Number
	goodsProductRtnJSON.URL = goodsProduct.URL
	goodsProductRtnJSON.AddTime = utils.FormatTimestampStr(goodsProduct.AddTime)
	goodsProductRtnJSON.UpdateTime = utils.FormatTimestampStr(goodsProduct.UpdateTime)

	return goodsProductRtnJSON, nil
}
