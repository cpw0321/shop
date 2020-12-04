// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"strings"
)

// GetTotal 算数据总数
func GetTotal(query map[string]string, tableName interface{}) (int64, error) {
	o := orm.NewOrm()
	qs := o.QueryTable(tableName)
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		switch k {
		case "name":
			qs = qs.Filter("name__contains", v)
		default:
			qs = qs.Filter(k, v)
		}
	}
	return qs.Count()
}
