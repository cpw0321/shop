// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
	"strings"
)

func init() {
	orm.RegisterModel(new(base.ShopComment))
}

// GetAllComment 获取评论信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllComment(query map[string]string, sortby string, order string,
	page int, limit int) (commentList []base.ShopComment, err error) {
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopComment))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
	}

	// order
	if sortby != "" {
		orderby := ""
		if order == "desc" {
			orderby = "-" + sortby
		} else {
			orderby = sortby
		}
		qs = qs.OrderBy(orderby)
	}

	_, err = qs.Limit(limit, offset).All(&commentList)
	if err != nil {
		return commentList, err
	}
	return commentList, nil
}

// GetCommentCount 获取评论总数和图片总数
func GetCommentCount(typeID int, valueID int) (allCount int64, hasPicCount int64, err error) {
	o := orm.NewOrm()
	commentTable := new(base.ShopComment)
	allCount, err = o.QueryTable(commentTable).Filter("Type", typeID).Filter("ValueID", valueID).Count()
	if err != nil {
		return allCount, hasPicCount, err
	}
	hasPicCount, err = o.QueryTable(commentTable).Filter("Type", typeID).Filter("ValueID", valueID).Filter("HasPicture", `1`).Count()
	if err != nil {
		return allCount, hasPicCount, err
	}
	return allCount, hasPicCount, nil
}

// AddComment 发表评论
func AddComment(comment *base.ShopComment) (err error) {
	o := orm.NewOrm()
	_, err = o.Insert(comment)
	if err != nil {
		return err
	}
	return nil
}

// GetCommentListByGoosID 获取商品相关的评论
func GetCommentListByGoosID(goodsID int) (commentList []base.ShopComment, err error) {
	o := orm.NewOrm()
	commentTable := new(base.ShopComment)
	_, err = o.QueryTable(commentTable).Filter("ValueID", goodsID).Filter("Deleted", 0).Limit(5).All(&commentList)
	if err != nil {
		return []base.ShopComment{}, err
	}

	return commentList, nil
}
