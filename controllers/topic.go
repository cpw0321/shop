// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/base"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
)

// TopicController ...
type TopicController struct {
	BaseController
}

// TopicList 获取主题列表
// @Tags 前台/获取主题列表
// @Summary 获取主题列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param page query int true "page"
// @Param limit query int true "limit"
// @Success 200 {object} response.TopicRtnJSON ""
// @router wx/topic/list [get]
func (c *TopicController) TopicList() {
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	if page <= 1 {
		page = 1
	}
	if limit <= 10 {
		limit = 10
	}

	var topicRtnJSON response.TopicRtnJSON
	var topicList []response.TopicJSON
	// 计算数据总数
	total, _ := models.GetTotal(map[string]string{}, new(base.ShopTopic))
	topicRtnJSON.Limit = limit
	topicRtnJSON.Page = page
	if total == 0 {
		topicRtnJSON.List = topicList
		topicRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, topicRtnJSON)
		return
	}

	topicList, err := models.GetAllTopic(map[string]string{}, "ID", "desc", page, limit)
	if err != nil {
		logger.Logger.Errorf("get topic list is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	topicRtnJSON.Total = total
	topicRtnJSON.List = topicList
	c.RespJSON(common.SUCCESEE, common.SCODE, topicRtnJSON)
}

// TopicDetail 获取专题详情
// @Tags 前台/获取专题详情
// @Summary 获取专题详情
// @Produce  application/JSON
// @Success 200 {object} response.TopicDetailRtnJSON ""
// @router wx/topic/detail [get]
func (c *TopicController) TopicDetail() {
	id, _ := c.GetInt("id")

	topic, err := models.GetOneTopic(id)
	if err != nil {
		logger.Logger.Errorf("get topic info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var topicDetailRtnJSON response.TopicDetailRtnJSON
	topicDetailRtnJSON.Topic = topic
	c.RespJSON(common.SUCCESEE, common.SCODE, topicDetailRtnJSON)
}

// Topic_Related 相关专题，即所有的专题
// @Tags 前台/相关专题
// @Summary 相关专题
// @Produce  application/JSON
// @Success 200 {object} response.TopicRelatedRtnJSON ""
// @router wx/topic/related [get]
func (c *TopicController) TopicRelated() {
	topicList, err := models.GetTopic()
	if err != nil {
		logger.Logger.Errorf("get topic list is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	var topicRelatedRtnJSON response.TopicRelatedRtnJSON
	topicRelatedRtnJSON.List = topicList
	c.RespJSON(common.SUCCESEE, common.SCODE, topicRelatedRtnJSON)
}
