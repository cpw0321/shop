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
	"strconv"
)

// CouponUserController ...
type CouponUserController struct {
	MainController
}

// CouponMylist 我的优惠劵列表
// @Tags 前台/我的优惠券列表
// @Summary 我的优惠券列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Param status query int true "优惠劵状态"
// @Success 200 {object} response.CouponUserListRtnJSON ""
// @router /coupon/mylist [get]
func (c *CouponUserController) CouponMylist() {
	// 优惠劵状态
	status, _ := c.GetInt("status")
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	if page <= 1 {
		page = 1
	}
	if limit <= 10 {
		limit = 10
	}

	query := make(map[string]string)
	query["UserID"] = strconv.Itoa(c.UserID)
	query["status"] = strconv.Itoa(status)
	var couponUserList []response.CouponUserRtnJSON
	couponUserList = []response.CouponUserRtnJSON{}
	var couponUserListRtnJSON response.CouponUserListRtnJSON

	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopCouponUser))
	couponUserListRtnJSON.Limit = limit
	couponUserListRtnJSON.Page = page
	if total == 0 {
		couponUserListRtnJSON.List = couponUserList
		couponUserListRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, couponUserListRtnJSON)
		return
	}

	// 获取我的优惠卷
	couponUserList, err := models.GetAllCouponUser(query, "ID", "asc", page, limit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	couponUserListRtnJSON.List = couponUserList
	couponUserListRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, couponUserListRtnJSON)
}

// CouponSelectList 当前订单可用优惠券列表
// @Tags 前台/当前订单可用优惠券列表
// @Summary 当前订单可用优惠券列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Success 200 {object} response.CouponUserListRtnJSON ""
// @router /coupon/selectlist [get]
func (c *CouponUserController) CouponSelectList() {
	var couponUserListRtnJSON response.CouponUserListRtnJSON
	couponUserRtnJSONList, err := models.GetSelectCouponUserList(c.UserID)
	if err != nil {
		logger.Logger.Errorf("GetSelectCouponUserList is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	couponUserListRtnJSON.List = couponUserRtnJSONList
	c.RespJSON(common.SUCCESEE, common.SCODE, couponUserListRtnJSON)
}

// CouponReceive 优惠券领取
// @Tags 前台/优惠券领取
// @Summary 优惠券领取
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CouponReceiveBody body request.CouponReceiveBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /coupon/receive [post]
func (c *CouponUserController) CouponReceive() {
	var couponBody request.CouponReceiveBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &couponBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 先查询是否已经领取过
	isOK := models.GetCouponUserIsExist(c.UserID, couponBody.CouponID)
	if isOK {
		logger.Logger.Error("Coupon is Received! UserID:[%v] CouponID:[%v].", c.UserID, couponBody.CouponID)
		c.RespJSON("不可重复领取", common.FCODE, "")
		return
	}
	err = models.CouponReceive(c.UserID, couponBody.CouponID)
	if err != nil {
		logger.Logger.Error("CouponReceive is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// CouponExchange 优惠券兑换
// @Tags 前台/优惠券兑换
// @Summary 优惠券兑换
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CouponExchangeBody body request.CouponExchangeBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /coupon/exchange [post]
func (c *CouponUserController) CouponExchange() {
	var couponExchangeBody request.CouponExchangeBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &couponExchangeBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	err = models.CouponExchange(c.UserID, couponExchangeBody.Code)
	if err != nil {
		logger.Logger.Error("CouponExchange is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}
