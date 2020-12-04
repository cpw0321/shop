// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"github.com/astaxie/beego/orm"
	"shop/base"
)

func init() {
	orm.RegisterModel(new(base.ShopFeedback))
}

// FeedbackSubmit 提交反馈
func FeedbackSubmit(feedback *base.ShopFeedback) error {
	o := orm.NewOrm()
	_, err := o.Insert(feedback)
	if err != nil {
		return err
	}
	return nil
}
