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
	orm.RegisterModel(new(base.ShopIssue))
}

func GetIssue() (issueRtnJSONList []response.IssueRtnJSON, err error) {
	var issueList []base.ShopIssue
	var issueRtnJSON response.IssueRtnJSON
	o := orm.NewOrm()
	issueTable := new(base.ShopIssue)
	_, err = o.QueryTable(issueTable).All(&issueList)
	if err != nil {
		return issueRtnJSONList, err
	}
	for _, v := range issueList {
		issueRtnJSON.ID = v.ID
		issueRtnJSON.Question = v.Question
		issueRtnJSON.Answer = v.Answer
		issueRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		issueRtnJSON.UpdateTime = utils.FormatTimestampStr(v.UpdateTime)
		issueRtnJSONList = append(issueRtnJSONList, issueRtnJSON)
	}

	return issueRtnJSONList, nil
}
