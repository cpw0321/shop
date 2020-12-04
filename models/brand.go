// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
)

func init() {
	orm.RegisterModel(new(base.ShopBrand))
}

// GetBrandByID 根据ID获取品牌信息
func GetBrandByID(ID int) (brandRespon response.BrandRespon, err error) {
	var brand base.ShopBrand
	o := orm.NewOrm()
	brandTable := new(base.ShopBrand)
	err = o.QueryTable(brandTable).Filter("ID", ID).Filter("Deleted", 0).One(&brand)
	if err != nil {
		logger.Logger.Errorf("get brand by ID is failed! err:[%v].", err)
		return response.BrandRespon{}, err
	}
	brandRespon.ID = brand.ID
	brandRespon.Name = brand.Name
	brandRespon.PicURL = brand.PicURL
	brandRespon.Desc = brand.Desc
	brandRespon.FloorPrice = brand.FloorPrice
	brandRespon.SortOrder = brand.SortOrder
	brandRespon.AddTime = utils.FormatTimestampStr(brand.AddTime)
	brandRespon.UpdateTime = utils.FormatTimestampStr(brand.UpdateTime)

	return brandRespon, nil
}
