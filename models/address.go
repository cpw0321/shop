// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
)

func init() {
	orm.RegisterModel(new(base.ShopAddress))
}

// GetAddressIsDefault 获取默认地址
func GetAddressDefault(userID int) (address base.ShopAddress, err error) {
	o := orm.NewOrm()
	addressTable := new(base.ShopAddress)
	err = o.QueryTable(addressTable).Filter("UserID", userID).Filter("IsDefault", 1).Filter("Deleted", 0).One(&address)
	if err != nil {
		return address, err
	}
	return address, nil
}

// GetAddressByID 获取地址通过ID
func GetAddressByID(userID, ID int) (address base.ShopAddress, err error) {
	o := orm.NewOrm()
	addressTable := new(base.ShopAddress)
	err = o.QueryTable(addressTable).Filter("UserID", userID).Filter("ID", ID).Filter("Deleted", 0).One(&address)
	if err != nil {
		return address, err
	}
	return address, nil
}

// GetAddress 获取所有的地址和数量
func GetAddress(userID int) (addressList []base.ShopAddress, total int64, err error) {
	o := orm.NewOrm()
	addressTable := new(base.ShopAddress)
	_, err = o.QueryTable(addressTable).Filter("UserID", userID).Filter("Deleted", 0).All(&addressList)
	if err != nil {
		return addressList, total, err
	}
	total, err = o.QueryTable(addressTable).Filter("UserID", userID).Filter("Deleted", 0).Count()
	if err != nil {
		return addressList, total, err
	}
	return addressList, total, nil
}

// DeleteAddress 删除地址
func DeleteAddress(userID int, ID int) (err error) {
	o := orm.NewOrm()
	addressTable := new(base.ShopAddress)
	_, err = o.QueryTable(addressTable).Filter("UserID", userID).Filter("ID", ID).Update(orm.Params{
		"Deleted": 1})
	if err != nil {
		return err
	}
	return nil
}

// AddAddress 新增地址
func AddAddress(address base.ShopAddress) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&address)
	if err != nil {
		return err
	}
	return nil
}
