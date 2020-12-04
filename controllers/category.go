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

// CatalogController ...
type CatalogController struct {
	BaseController
}

// CatalogIndex 获取分类目录全部分类数据
// @Tags 前台/获取分类目录
// @Summary 获取分类目录
// @Produce  application/JSON
// @Param token header string true "token"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Success 200 {object} response.CatelogIndexRtnJSON ""
// @router /catalog/index [get]
func (c *CatalogController) CatalogIndex() {
	var limit int
	var page int
	if page <= 0 {
		page = 1
	}
	limit, _ = c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}

	var catelogIndexRtnJSON response.CatelogIndexRtnJSON

	// 获取一级分类菜单
	categoryList, err := models.GetAllCategory(map[string]string{"pID": strconv.Itoa(0)}, "ID", "asc", page, limit)
	if err != nil {
		logger.Logger.Error("GetAllCategory info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 获取当前分类菜单信息
	currentCategory, err := models.GetOneCategory(categoryList[0].ID)
	if err != nil {
		logger.Logger.Error("GetOneCategory info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var currentSubCategory []response.CategoryRtnJSON
	// 获取当前分类二级子分类信息
	if currentCategory.ID > 0 {
		currentSubCategory, err = models.GetAllCategory(map[string]string{"pID": strconv.Itoa(currentCategory.ID)}, "ID", "asc", page, 100)
		if err != nil {
			logger.Logger.Error("GetAllCategory info failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}
	}

	catelogIndexRtnJSON.CategoryList = categoryList
	catelogIndexRtnJSON.CurrentCategory = currentCategory
	catelogIndexRtnJSON.CurrentSubCategory = currentSubCategory

	c.RespJSON(common.SUCCESEE, common.SCODE, catelogIndexRtnJSON)
}

// CatalogCurrent 获取分类目录当前分类数据
// @Tags 前台/获取分类目录当前分类数据
// @Summary 获取分类目录当前分类数据
// @Produce  application/JSON
// @Param token header string true "token"
// @Param id query int true "分类id"
// @Param limit query int true "页大小"
// @Success 200 {object} response.CatalogCurrentRtnJSON ""
// @router /catalog/current [get]
func (c *CatalogController) CatalogCurrent() {
	var catalogCurrentRtnJSON response.CatalogCurrentRtnJSON
	var currentSubCategory []response.CategoryRtnJSON

	categoryID, _ := c.GetInt("id")

	// 获取当前分类菜单信息
	currentCategory, err := models.GetOneCategory(categoryID)
	if err != nil {
		logger.Logger.Error("GetOneCategory info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 获取当前分类二级子分类信息
	if currentCategory.ID > 0 {
		currentSubCategory, err = models.GetAllCategory(map[string]string{"pID": strconv.Itoa(currentCategory.ID)}, "ID", "asc", 1, 100) // 1 第1页  100 100个数据
		if err != nil {
			logger.Logger.Error("GetAllCategory info failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}
	}

	catalogCurrentRtnJSON.CurrentCategory = currentCategory
	catalogCurrentRtnJSON.CurrentSubCategory = currentSubCategory

	c.RespJSON(common.SUCCESEE, common.SCODE, catalogCurrentRtnJSON)
}
