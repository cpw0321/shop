// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
)

// IssueController ...
type IssueController struct {
	BaseController
}

// IssueList 获取问题列表
// @Tags 前台/获取问题列表
// @Summary 获取问题列表
// @Produce  application/JSON
// @Success 200 {object} response.IssueListRtnJSON ""
// @router /issue/list [get]
func (c *IssueController) IssueList() {
	var issueListRtnJSON response.IssueListRtnJSON

	issueRtnJSONList, err := models.GetIssue()
	if err != nil {
		logger.Logger.Error("GetIssue is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	issueListRtnJSON.List = issueRtnJSONList
	c.RespJSON(common.SUCCESEE, common.SCODE, issueListRtnJSON)
}
