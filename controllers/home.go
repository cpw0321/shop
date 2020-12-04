// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"strconv"
)

// HomeController ...
type HomeController struct {
	BaseController
}

// HomeIndex 首页数据
// @Tags 前台/首页数据
// @Summary 首页数据
// @Produce  application/JSON
// @Success 200 {object} response.HomeRtnJSON ""
// @router /home/index [get]
func (c *HomeController) HomeIndex() {
	var homeRtnJSON response.HomeRtnJSON
	var bannerLimit int
	var couponLimit int
	var newLimit int
	var hotLimit int
	var topicLimit int

	systemListRespon, err := models.GetSystem()
	if err != nil {
		logger.Logger.Errorf("GetSystem is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	for _, v := range systemListRespon {
		switch v.KeyName {
		case common.SHOP_WX_INDEX_NEW:
			newLimit, _ = strconv.Atoi(v.KeyValue)
		case common.SHOP_WX_INDEX_HOT:
			hotLimit, _ = strconv.Atoi(v.KeyValue)
		case common.SHOP_WX_INDEX_TOPIC:
			topicLimit, _ = strconv.Atoi(v.KeyValue)
		case common.SHOP_WX_INDEX_COUPON:
			couponLimit, _ = strconv.Atoi(v.KeyValue)
		case common.SHOP_WX_INDEX_BRAND:
			bannerLimit, _ = strconv.Atoi(v.KeyValue)
		}
	}

	// 获取首页轮播图
	bannerRtnJSONList, err := models.GetAllBanner(map[string]string{"position": strconv.Itoa(1)}, "ID", "desc", 1, bannerLimit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	// 获取首页分类信息
	//channelList, err := models.GetAllCategory(map[string]string{"pID": strconv.Itoa(0)}, []string{}, []string{}, 1, 8)
	//if err != nil {
	//	logger.Logger.Error("get banners info failed! err:[%v].", err)
	//}

	// 获取首页优惠卷
	couponList, err := models.GetAllCoupon(map[string]string{"status": strconv.Itoa(0)}, "ID", "desc", 1, couponLimit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	// 获取首页热门商品
	newGoodsList, err := models.GetAllGoods(map[string]string{"is_new": strconv.Itoa(1)}, "ID", "desc", 1, hotLimit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	// 获取首页新品
	hostGoodsList, err := models.GetAllGoods(map[string]string{"is_hot": strconv.Itoa(1)}, "ID", "desc", 1, newLimit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	// 获取首页主题
	topicRtnJSONList, err := models.GetAllTopic(map[string]string{}, "ID", "desc", 1, topicLimit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	homeRtnJSON.Banners = bannerRtnJSONList
	homeRtnJSON.CouponList = couponList
	homeRtnJSON.NewGoods = newGoodsList
	homeRtnJSON.HotGoods = hostGoodsList
	homeRtnJSON.TopicList = topicRtnJSONList
	c.RespJSON(common.SUCCESEE, common.SCODE, homeRtnJSON)
}

// HomeAbout 关于我们
// @Tags 前台/关于我们
// @Summary 关于我们
// @Produce  application/JSON
// @Success 200 {object} response.HomeAboutRtnJSON ""
// @router wx/home/about [get]
func (c *HomeController) HomeAbout() {
	var homeAboutRtnJSON response.HomeAboutRtnJSON

	aboutList, err := models.GetAbout()
	if err != nil {
		logger.Logger.Error("GetAbout is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	homeAboutRtnJSON.List = aboutList
	c.RespJSON(common.SUCCESEE, common.SCODE, homeAboutRtnJSON)
}
