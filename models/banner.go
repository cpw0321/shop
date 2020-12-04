// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
	"strings"
)

func init() {
	orm.RegisterModel(new(base.ShopAd))
}

// AddBanner 创建banner
func AddBanner(m *base.ShopAd) (ID int64, err error) {
	o := orm.NewOrm()
	ID, err = o.Insert(m)
	if err != nil {
		logger.Logger.Errorf("add banner info failed! err:[%v].", err)
		return ID, err
	}
	return ID, nil
}

// GetBannerByID 获取banner信息通过ID
func GetBannerByID(ID int) (banner *base.ShopAd, err error) {
	o := orm.NewOrm()
	banner = &base.ShopAd{ID: ID}
	err = o.Read(banner)
	if err != nil {
		return nil, err
	}
	return banner, nil
}

// GetAllBanner 获取banner信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllBanner(query map[string]string, sortby string, order string,
	page int, limit int) (bannerRtnJSONList []response.BannerRtnJSON, err error) {
	bannerRtnJSONList = []response.BannerRtnJSON{}
	var bannerList []base.ShopAd
	var bannerRtnJSON response.BannerRtnJSON
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopAd))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
	}

	// order
	orderby := ""
	if order == "desc" {
		orderby = "-" + sortby
	} else {
		orderby = sortby
	}

	qs = qs.OrderBy(orderby)
	_, err = qs.Limit(limit, offset).All(&bannerList)
	if err != nil {
		return bannerRtnJSONList, err
	}
	for _, v := range bannerList {
		bannerRtnJSON.ID = v.ID
		bannerRtnJSON.Name = v.Name
		bannerRtnJSON.Link = v.Link
		bannerRtnJSON.URL = v.URL
		bannerRtnJSON.Position = v.Position
		bannerRtnJSON.Content = v.Content
		bannerRtnJSON.StartTime = utils.FormatTimestampStr(v.StartTime)
		bannerRtnJSON.EndTime = utils.FormatTimestampStr(v.EndTime)
		bannerRtnJSON.Enabled = v.Enabled
		bannerRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		bannerRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		bannerRtnJSONList = append(bannerRtnJSONList, bannerRtnJSON)
	}
	return bannerRtnJSONList, nil
}
