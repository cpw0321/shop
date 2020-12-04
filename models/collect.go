// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
)

func init() {
	orm.RegisterModel(new(base.ShopCollect))
}

// GetCollect 根据userID获取收藏的商品
func GetCollect(userID int, typeID int, page int, limit int) (collectList []base.ShopCollect, err error) {
	offset := (page - 1) * limit
	o := orm.NewOrm()
	collectTable := new(base.ShopCollect)
	_, err = o.QueryTable(collectTable).Filter("UserID", userID).Filter("Type", typeID).Filter("deleted", 0).Limit(limit, offset).All(&collectList)
	if err != nil {
		return collectList, err
	}
	return collectList, nil
}

// GetOneCollect 根据用户userID获取一种类型的收藏
func GetOneCollect(userID int, typeID int, valueID int) (collect base.ShopCollect, err error) {
	o := orm.NewOrm()
	collecttable := new(base.ShopCollect)
	qs := o.QueryTable(collecttable)
	err = qs.Filter("UserID", userID).Filter("Type", typeID).Filter("ValueID", valueID).Filter("Deleted", 0).One(&collect)
	if err != nil {
		return collect, err
	}
	return collect, nil
}

// AddCollect 新增收藏
func AddCollect(collect *base.ShopCollect) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(collect)
	if err != nil {
		return err
	}
	return nil
}

// DeleteCollect 删除收藏
func DeleteCollect(collect *base.ShopCollect) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(collect)
	if err != nil {
		return err
	}
	return nil
}

// GetCollectExist 根据用户userID获取一种类型的收藏是否收藏
func GetCollectExist(userID int, typeID int, valueID int) (isOK bool) {
	o := orm.NewOrm()
	collecttable := new(base.ShopCollect)
	qs := o.QueryTable(collecttable)
	isOK = qs.Filter("UserID", userID).Filter("Type", typeID).Filter("ValueID", valueID).Filter("deleted", 0).Exist()
	return isOK
}
