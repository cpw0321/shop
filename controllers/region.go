// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/common"
	"shop/logger"
	"shop/models"
)

// RegionController ...
type RegionController struct {
	MainController
}

// RegionList 区域列表
// @Tags 前台/区域列表
// @Summary 区域列表
// @Produce  application/JSON
// @Success 200 {object} []base.ShopRegion "区域列表"
// @router wx/region/list [get]
func (c *RegionController) RegionList() {
	parentID, _ := c.GetInt("parentId")

	regionList, err := models.GetAllRegion(parentID)
	if err != nil {
		logger.Logger.Error("GetIssue is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, regionList)
}
