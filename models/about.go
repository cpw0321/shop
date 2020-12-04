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
	orm.RegisterModel(new(base.ShopAbout))
}

func GetAbout() (aboutRtnJSONList []response.AboutRtnJSON, err error) {
	var aboutList []base.ShopAbout
	var aboutRtnJSON response.AboutRtnJSON
	o := orm.NewOrm()
	aboutTable := new(base.ShopAbout)
	_, err = o.QueryTable(aboutTable).All(&aboutList)
	if err != nil {
		return aboutRtnJSONList, err
	}
	for _, v := range aboutList {
		aboutRtnJSON.ID = v.ID
		aboutRtnJSON.Title = v.Title
		aboutRtnJSON.Context = v.Context
		aboutRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		aboutRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		aboutRtnJSONList = append(aboutRtnJSONList, aboutRtnJSON)
	}

	return aboutRtnJSONList, nil
}
