// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
)

func init() {
	orm.RegisterModel(new(base.ShopKeyword))
}

// GetKeywordDefault 获取默认关键字
func GetKeywordDefault() (keyword base.ShopKeyword, err error) {
	o := orm.NewOrm()
	keywordsTable := new(base.ShopKeyword)
	err = o.QueryTable(keywordsTable).Filter("IsDefault", 1).One(&keyword)
	if err != nil {
		return keyword, err
	}
	return keyword, nil
}

// GetKeywordHot 获取默认关键字
func GetKeywordHot() (keywordList []base.ShopKeyword, err error) {
	o := orm.NewOrm()
	keywordsTable := new(base.ShopKeyword)
	_, err = o.QueryTable(keywordsTable).Filter("IsHot", 1).All(&keywordList)
	if err != nil {
		return keywordList, err
	}
	return keywordList, nil
}

// SearchHelper 搜索帮助
func SearchHelper(keyword string) (keywordList []string, err error) {
	o := orm.NewOrm()
	keywordstable := new(base.ShopKeyword)

	var keywords []base.ShopKeyword
	_, err = o.QueryTable(keywordstable).Filter("Keyword__contains", keyword).Distinct().Limit(10).All(&keywords)
	if err != nil {
		return keywordList, err
	}

	for _, v := range keywords {
		keywordList = append(keywordList, v.Keyword)
	}
	return keywordList, nil
}
