// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/utils"
	"strings"
)

func init() {
	orm.RegisterModel(new(base.ShopTopic))
}

// GetAllTopic 获取topic信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllTopic(query map[string]string, sortby string, order string,
	page int, limit int) (topicJSONList []response.TopicJSON, err error) {
	topicJSONList = []response.TopicJSON{}
	var topicList []base.ShopTopic
	var topicJSON response.TopicJSON
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopTopic))
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
	_, err = qs.Limit(limit, offset).All(&topicList)
	if err != nil {
		return topicJSONList, err
	}

	for _, v := range topicList {
		topicJSON.ID = v.ID
		topicJSON.Title = v.Title
		topicJSON.Subtitle = v.Subtitle
		topicJSON.PicURL = v.PicURL
		topicJSON.SortOrder = v.SortOrder
		topicJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		topicJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		topicJSONList = append(topicJSONList, topicJSON)
	}

	return topicJSONList, nil
}

// GetOneTopic 获取专题详情
func GetOneTopic(ID int) (topicDetailJSON response.TopicDetailJSON, err error) {
	var topic base.ShopTopic
	o := orm.NewOrm()
	topicTable := new(base.ShopTopic)
	err = o.QueryTable(topicTable).Filter("ID", ID).One(&topic)
	if err != nil {
		return topicDetailJSON, err
	}

	topicDetailJSON.ID = topic.ID
	topicDetailJSON.Title = topic.Title
	topicDetailJSON.Subtitle = topic.Subtitle
	topicDetailJSON.PicURL = topic.PicURL
	topicDetailJSON.SortOrder = topic.SortOrder
	topicDetailJSON.Content = topic.Content
	topicDetailJSON.ReadCount = topic.ReadCount
	topicDetailJSON.AddTime = utils.FormatTimestampStr(topic.AddTime)
	topicDetailJSON.UpdateTime = utils.FormatTimestampStr(topic.UpdateTime)

	return topicDetailJSON, nil
}

// GetTopic 获取所有topic
func GetTopic() (topicList []base.ShopTopic, err error) {
	o := orm.NewOrm()
	topictable := new(base.ShopTopic)
	_, err = o.QueryTable(topictable).All(&topicList)
	if err != nil {
		return topicList, err
	}
	return topicList, nil
}
