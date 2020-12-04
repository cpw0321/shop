// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/base"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"strconv"
)

// CouponController ...
type CouponController struct {
	BaseController
}

// CouponList 获取优惠劵列表
// @Tags 前台/获取评论列表
// @Summary 获取评论列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Success 200 {object} response.CouponRtnJSON ""
// @router /coupon/list [get]
func (c *CouponController) CouponList() {
	page, _ := c.GetInt("page")
	if page <= 1 {
		page = 1
	}
	limit, _ := c.GetInt("limit")
	if limit <= 10 {
		limit = 10
	}

	var couponList []response.CouponRtnJSON

	// 计算数据总数
	total, _ := models.GetTotal(map[string]string{"status": strconv.Itoa(0)}, new(base.ShopCoupon))
	var couponListRtnJSON response.CouponListRtnJSON
	couponListRtnJSON.Limit = limit
	couponListRtnJSON.Page = page
	if total == 0 {
		couponListRtnJSON.List = couponList
		couponListRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, couponListRtnJSON)
		return
	}

	// 获取可以使用的优惠卷
	couponList, err := models.GetAllCoupon(map[string]string{"status": strconv.Itoa(0)}, "ID", "asc", page, limit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	couponListRtnJSON.List = couponList
	couponListRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, couponListRtnJSON)
}
