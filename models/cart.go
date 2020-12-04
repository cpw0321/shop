// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/logger"
)

func init() {
	orm.RegisterModel(new(base.ShopCart))
}

// GetAllCart 获取个人购物车所有数据
func GetAllCart(userID int) (cartList []base.ShopCart, err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	_, err = o.QueryTable(cartTable).Filter("UserID", userID).Filter("Deleted", 0).All(&cartList)
	if err != nil {
		return cartList, err
	}
	return cartList, nil
}

// AddCart 添加购物车
func AddCart(m *base.ShopCart) (ID int64, err error) {
	o := orm.NewOrm()
	ID, err = o.Insert(m)
	return
}

// GetCartIsExist 查询用户添加的商品是否已经存在购物车里
func GetCartIsExist(userID, goodsID, productID int) (cart base.ShopCart, err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	err = o.QueryTable(cartTable).Filter("UserID", userID).Filter("GoodsID", goodsID).Filter("ProductID", productID).Filter("Deleted", 0).One(&cart)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

// UpdateCartNumberByID 更新数量通过ID
func UpdateCartNumberByID(ID, num int) (n int64, err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	n, err = o.QueryTable(cartTable).Filter("ID", ID).Update(orm.Params{
		"number": num})
	if err != nil {
		return n, err
	}
	return n, nil
}

// GetCartByID 查询购物车信息通过ID
func GetCartByID(ID int) (cart base.ShopCart, err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	err = o.QueryTable(cartTable).Filter("ID", ID).One(&cart)
	if err != nil {
		return cart, err
	}
	return cart, nil
}

// DeleteCartByID 批量删除购物车通过productID
func DeleteCartByProductID(userID int, productIDs []int) (n int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		logger.Logger.Errorf("orm begin is failed! err:[%v].", err)
		return n, err
	}
	cartTable := new(base.ShopCart)
	for _, v := range productIDs {
		n, err = o.QueryTable(cartTable).Filter("userID", userID).Filter("productID", v).Update(orm.Params{
			"deleted": 1})
		if err != nil {
			o.Rollback()
			return n, err
		}
	}

	o.Commit()
	return n, nil
}

// DeleteCartByID 批量选定购物车通过productID
func UpdateCartByProductID(userID int, productIDs []int, checked int) (n int64, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		logger.Logger.Errorf("orm begin is failed! err:[%v].", err)
		return n, err
	}
	cartTable := new(base.ShopCart)
	for _, v := range productIDs {
		n, err = o.QueryTable(cartTable).Filter("userID", userID).Filter("productID", v).Update(orm.Params{
			"checked": checked})
		if err != nil {
			o.Rollback()
			return n, err
		}
	}

	o.Commit()
	return n, nil
}

// GetCartGoodsCount 获取用户购物车商品数量
func GetCartGoodsCount(userID int) (num int64, err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	num, err = o.QueryTable(cartTable).Filter("userID", userID).Filter("Deleted", 0).Count()
	if err != nil {
		return num, err
	}
	return num, nil
}

// GetCartTotal 获取用户购物车中数量和价格
func GetCartTotal(userID int) (cartTotal response.CartTotal, err error) {
	var goodsCount int
	var goodsAmount float64
	var checkedGoodsCount int
	var checkedGoodsAmount float64

	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	var carts []base.ShopCart
	_, err = o.QueryTable(cartTable).Filter("UserID", userID).Filter("Deleted", 0).All(&carts)
	if err != nil {
		return cartTotal, err
	}

	for _, val := range carts {
		goodsCount += val.Number
		goodsAmount += float64(val.Number) * val.Price
		if val.Checked == 1 {
			checkedGoodsCount += val.Number
			checkedGoodsAmount += float64(val.Number) * val.Price
		}
	}
	cartTotal.GoodsCount = goodsCount
	cartTotal.GoodsAmount = goodsAmount
	cartTotal.CheckedGoodsCount = checkedGoodsCount
	cartTotal.CheckedGoodsAmount = checkedGoodsAmount

	return cartTotal, nil
}

// GetCheckAllCart 获取个人购物车选定的数据
func GetCheckAllCart(userID int) (cartList []base.ShopCart, err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	_, err = o.QueryTable(cartTable).Filter("UserID", userID).Filter("Checked", 1).Filter("Deleted", 0).All(&cartList)
	if err != nil {
		return cartList, err
	}
	return cartList, nil
}

// DeleteCartByID 批量删除购物车通过ID
func DeleteCartByID(ID int) (err error) {
	o := orm.NewOrm()
	cartTable := new(base.ShopCart)
	_, err = o.QueryTable(cartTable).Filter("ID", ID).Filter("Deleted", 0).Update(orm.Params{
		"Deleted": 1})
	if err != nil {
		return err
	}
	return nil
}
