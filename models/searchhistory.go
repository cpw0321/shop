// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
)

func init() {
	orm.RegisterModel(new(base.ShopSearchHistory))
}

// GetSearchHistory 获取历史搜索记录
func GetSearchHistory(userID int) (searchHistoryList []base.ShopSearchHistory, err error) {
	o := orm.NewOrm()
	searchHistoryTable := new(base.ShopSearchHistory)
	_, err = o.QueryTable(searchHistoryTable).Filter("UserID", userID).Filter("Deleted", 0).All(&searchHistoryList)
	if err != nil {
		return searchHistoryList, err
	}
	return searchHistoryList, nil
}

// SearchHistoryClear 用户搜索记录清理
func SearchHistoryClear(userID int) (err error) {
	o := orm.NewOrm()
	SearchHistory := new(base.ShopSearchHistory)
	_, err = o.QueryTable(SearchHistory).Filter("UserID", userID).Update(orm.Params{
		"Deleted": 1})

	return err
}

// AddSearchHistory 添加用户搜索记录
func AddSearchHistory(searchHistory base.ShopSearchHistory) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(&searchHistory)
	return err
}

// GetSearchHistoryByName 根据搜索名查看是否已存在搜索历史的表中
func GetSearchHistoryByName(name string) (isOK bool) {
	o := orm.NewOrm()
	searchHistoryTable := new(base.ShopSearchHistory)
	isOK = o.QueryTable(searchHistoryTable).Filter("Keyword", name).Exist()
	return isOK
}
