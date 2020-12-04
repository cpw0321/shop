// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"encoding/json"
	"shop/base"
	"shop/base/request"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"shop/utils"
	"strconv"
	"time"
)

// CommentController ...
type CommentController struct {
	MainController
}

// CommentPost 发表评论
// @Tags 前台/发表评论
// @Summary 发表评论
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CommentBody body request.CommentBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /comment/post [post]
func (c *CommentController) CommentPost() {
	var commentBody request.CommentBody
	var comment base.ShopComment
	var hasPicture int

	err := json.Unmarshal(c.Ctx.Input.RequestBody, &commentBody)
	if err != nil {
		logger.Logger.Errorf("input para JSON unmarshal is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	comment.ValueID = commentBody.ValueID
	comment.Type = commentBody.Type
	comment.Content = commentBody.Content
	comment.UserID = c.UserID
	if commentBody.HasPicture {
		hasPicture = 1
	} else {
		hasPicture = 0
	}
	comment.HasPicture = hasPicture
	picURLs, err := json.Marshal(commentBody.PicURLs)
	if err != nil {
		logger.Logger.Error("PicURLs JSON failed! err:[%v].", err)
	}
	comment.PicURLs = string(picURLs)
	comment.Star = commentBody.Star
	comment.AddTime = time.Now()

	err = models.AddComment(&comment)
	if err != nil {
		logger.Logger.Errorf("add comment is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// CommentCount 获取评论数量
// @Tags 前台/发表评论
// @Summary 发表评论
// @Produce  application/JSON
// @Param token header string true "token"
// @Param valueId query int true "valueId"
// @Param type query int true "type"
// @Success 200 {object} response.CommentCountRtnJSON ""
// @router /comment/count [get]
func (c *BaseController) CommentCount() {
	valueID, _ := c.GetInt("valueId")
	typeStatus, _ := c.GetInt("type")
	allCount, hasPicCount, err := models.GetCommentCount(typeStatus, valueID)
	if err != nil {
		logger.Logger.Errorf("GetCommentCount is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var commentCountRtnJSON response.CommentCountRtnJSON
	commentCountRtnJSON.AllCount = allCount
	commentCountRtnJSON.HasPicCount = hasPicCount
	c.RespJSON(common.SUCCESEE, common.SCODE, commentCountRtnJSON)
}

// CommentList 获取评论列表
// @Tags 前台/获取评论列表
// @Summary 获取评论列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param valueId query int true "valueId"
// @Param type query int true "type"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Param showType query int true "类型"
// @Success 200 {object} response.CommentCountRtnJSON ""
// @router /comment/list [get]
func (c *BaseController) CommentList() {
	typeID, _ := c.GetInt("type") // 显示是商品还是专题，0是商品，1是专题
	valueID, _ := c.GetInt("valueId")
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	showType, _ := c.GetInt("showType") // 显示 0 所有评论 1 图片评论（非0）  对应数据库has_picture字段
	if limit <= 0 {
		limit = 10
	}
	if page <= 1 {
		page = 1
	}

	var query = make(map[string]string)
	query["Type"] = strconv.Itoa(typeID)
	query["ValueID"] = strconv.Itoa(valueID)
	if showType == 1 {
		query["HasPicture"] = strconv.Itoa(showType)
	}

	var commentRtnJSON response.CommentRtnJSON
	var commentListArray []response.CommentListRtnJSON
	var commentListRtnJSON response.CommentListRtnJSON

	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopComment))
	rtnInfo := make(map[string]interface{})
	commentRtnJSON.Limit = limit
	commentRtnJSON.Page = page
	if total == 0 {
		commentRtnJSON.List = commentListArray
		commentRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, rtnInfo)
		return
	}

	limit = 100
	commentList, err := models.GetAllComment(query, "ID", "desc", page, limit)
	if err != nil {
		logger.Logger.Errorf("GetAllComment is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	for _, v := range commentList {
		user, err := models.GetUserByID(v.UserID)
		if err != nil {
			logger.Logger.Errorf("get user info is failed! err:[%v].", err)
		}
		var picList []string
		err = json.Unmarshal([]byte(v.PicURLs), &picList)
		if err != nil {
			logger.Logger.Errorf("get PicURLs is failed! goodsProduct:[%v].", v)
		}
		commentListRtnJSON.UserInfo.Nickname = user.Nickname
		commentListRtnJSON.UserInfo.AvatarURL = user.Avatar
		commentListRtnJSON.ID = v.ID
		commentListRtnJSON.Content = v.Content
		commentListRtnJSON.PicList = picList
		commentListRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		commentListArray = append(commentListArray, commentListRtnJSON)
	}

	commentRtnJSON.List = commentListArray
	commentRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, commentRtnJSON)
}
