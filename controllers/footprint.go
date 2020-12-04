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
)

// FootprintController ...
type FootprintController struct {
	MainController
}

// FootprintDelete 删除足迹
// @Tags 前台/删除足迹
// @Summary 删除足迹
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.FootprintBody body request.FootprintBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /footprint/delete [post]
func (c *FootprintController) FootprintDelete() {
	var footprintBody request.FootprintBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &footprintBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal FootprintDelete Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	err = models.DeleteFootprint(footprintBody.ID)
	if err != nil {
		logger.Logger.Errorf("DeleteFootprint is failed! err:[%v].", err)
		c.RespJSON(common.ErrDeleteData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// FootprintDelete 足迹列表
// @Tags 前台/足迹列表
// @Summary 足迹列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param page query int true "页数"
// @Param limit query int true "页数大小"
// @Success 200 {object} response.FootprintListRtnJSON ""
// @router /footprint/list [get]
func (c *FootprintController) FootprintList() {
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	if page <= 1 {
		page = 1
	}
	if limit <= 10 {
		limit = 10
	}

	var footprintListRtnJSON response.FootprintListRtnJSON
	var footprintList []base.ShopFootprint
	var footprintGoodsList []response.FootprintGoods
	footprintGoodsList = []response.FootprintGoods{}
	var footprintGoods response.FootprintGoods
	query := make(map[string]string)
	query["UserID"] = strconv.Itoa(c.UserID)
	query["deleted"] = "0"

	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopFootprint))
	footprintListRtnJSON.Limit = limit
	footprintListRtnJSON.Page = page
	if total == 0 {
		footprintListRtnJSON.List = footprintGoodsList
		footprintListRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, footprintListRtnJSON)
		return
	}
	limit = 100
	footprintList, err := models.GetAllFootprint(query, "ID", "desc", page, limit)
	if err != nil {
		logger.Logger.Errorf("GetAllFootprint is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	for _, v := range footprintList {
		goods, err := models.GetOneGoods(v.GoodsID)
		if err != nil {
			logger.Logger.Errorf("GetOneGoods is failed! err:[%v].", err)
		}
		footprintGoods.ID = v.ID
		footprintGoods.GoodsID = goods.ID
		footprintGoods.Name = goods.Name
		footprintGoods.Brief = goods.Brief
		footprintGoods.PicURL = goods.PicURL
		footprintGoods.RetailPrice = goods.RetailPrice
		footprintGoods.AddTime = utils.FormatTimestampStr(v.AddTime)
		footprintGoodsList = append(footprintGoodsList, footprintGoods)
	}

	footprintListRtnJSON.List = footprintGoodsList
	footprintListRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, footprintListRtnJSON)
}
