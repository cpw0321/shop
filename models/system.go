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
	orm.RegisterModel(new(base.ShopSystem))
}

// GetSystem 获取小程序配置
func GetSystem() (systemListRespon []response.SystemRespon, err error) {
	var systemList []base.ShopSystem
	var systemRespon response.SystemRespon
	o := orm.NewOrm()
	systemTable := new(base.ShopSystem)
	_, err = o.QueryTable(systemTable).Filter("Deleted", 0).All(&systemList)
	if err != nil {
		logger.Logger.Errorf("QueryTable system table is failed! err;[%v].", err)
		return []response.SystemRespon{}, err
	}
	for _, v := range systemList {
		systemRespon.ID = v.ID
		systemRespon.KeyName = v.KeyName
		systemRespon.KeyValue = v.KeyValue
		systemRespon.AddTime = utils.FormatTimestampStr(v.AddTime)
		systemRespon.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		systemListRespon = append(systemListRespon, systemRespon)
	}
	return systemListRespon, nil
}

// GetSystemByName 获取单个设置的值
func GetSystemByName(keyName string) (keyValue string, err error) {
	var system base.ShopSystem
	o := orm.NewOrm()
	systemTable := new(base.ShopSystem)
	err = o.QueryTable(systemTable).Filter("Deleted", 0).Filter("KeyName", keyName).One(&system)
	if err != nil {
		logger.Logger.Errorf("QueryTable system table is failed! err;[%v].", err)
		return "", err
	}

	return system.KeyValue, nil
}
