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
	orm.RegisterModel(new(base.ShopOrderGoods))
}

// GetOrderGoodsComment 获取订单商品待评论列表
func GetOrderGoodsComment(orderID int, goodsID int) (orderGoods base.ShopOrderGoods, err error) {
	o := orm.NewOrm()
	orderGoodsTable := new(base.ShopOrderGoods)
	err = o.QueryTable(orderGoodsTable).Filter("OrderID", orderID).Filter("GoodsID", goodsID).Filter("Comment", 0).One(&orderGoods)
	if err != nil {
		return orderGoods, err
	}
	return orderGoods, nil
}

// GetOrderGoods 获取订单商品信息
func GetOrderGoods(ID int) (orderGoods base.ShopOrderGoods, err error) {
	o := orm.NewOrm()
	orderGoodsTable := new(base.ShopOrderGoods)
	err = o.QueryTable(orderGoodsTable).Filter("ID", ID).One(&orderGoods)
	if err != nil {
		return orderGoods, err
	}
	return orderGoods, nil
}

// AddOrderGoodsComment 订单商品评论
func AddOrderGoodsComment(ID int, comment *base.ShopComment) (err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		return err
	}
	_, err = o.Insert(comment)
	if err != nil {
		o.Rollback()
		return err
	}
	orderGoodsTable := new(base.ShopOrderGoods)
	_, err = o.QueryTable(orderGoodsTable).Filter("ID", ID).Update(orm.Params{
		"comment": comment.ID})
	if err != nil {
		o.Rollback()
		return err
	}

	o.Commit()
	return nil
}

// GetOrderGoodsListByID 获取订单商品信息列表
func GetOrderGoodsListByOrderID(orderID int) (goodsList []response.OrderGoods, err error) {
	var orderGoodsList []base.ShopOrderGoods
	var goods response.OrderGoods
	o := orm.NewOrm()
	orderGoodsTable := new(base.ShopOrderGoods)
	_, err = o.QueryTable(orderGoodsTable).Filter("OrderID", orderID).All(&orderGoodsList)
	if err != nil {
		return goodsList, err
	}
	for _, v := range orderGoodsList {
		var specifications []string
		err = json.Unmarshal([]byte(v.Specifications), &specifications)
		if err != nil {
			logger.Logger.Errorf("get goodsProduct is failed! goodsProduct:[%v].", v)
		}
		goods.ID = v.ID
		goods.OrderID = v.OrderID
		goods.GoodsID = v.GoodsID
		goods.GoodsName = v.GoodsName
		goods.GoodsSn = v.GoodsSn
		goods.ProductID = v.ProductID
		goods.PicURL = v.PicURL
		goods.Price = v.Price
		goods.Number = v.Number
		goods.Specifications = specifications
		goods.Comment = v.Comment
		goods.AddTime = utils.FormatTimestampStr(v.AddTime)
		goodsList = append(goodsList, goods)
	}
	return goodsList, nil
}

// GetOrderGoods 获取订单商品信息通过orderID
func GetOrderGoodsByOrderID(orderID int) (goods response.OrderGoods, err error) {
	var orderGoods base.ShopOrderGoods
	o := orm.NewOrm()
	orderGoodsTable := new(base.ShopOrderGoods)
	err = o.QueryTable(orderGoodsTable).Filter("OrderID", orderID).One(&orderGoods)
	if err != nil {
		return goods, err
	}

	var specifications []string
	err = json.Unmarshal([]byte(orderGoods.Specifications), &specifications)
	if err != nil {
		logger.Logger.Errorf("get goodsProduct is failed! goodsProduct:[%v].", orderGoods)
	}
	goods.ID = orderGoods.ID
	goods.GoodsName = orderGoods.GoodsName
	goods.PicURL = orderGoods.PicURL
	goods.Price = orderGoods.Price
	goods.Number = orderGoods.Number
	goods.Specifications = specifications

	return goods, nil
}
