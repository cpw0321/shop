// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
)

func init() {
	orm.RegisterModel(new(base.ShopRegion))
}

func GetAllRegion(PID int) (regionList []base.ShopRegion, err error) {
	o := orm.NewOrm()
	regiontable := new(base.ShopRegion)
	_, err = o.QueryTable(regiontable).Filter("PID", PID).All(&regionList)
	return regionList, err
}
