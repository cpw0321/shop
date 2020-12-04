// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/base"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"time"
)

// SearchController ...
type SearchController struct {
	MainController
}

// SearchIndex 搜索关键字
// @Tags 前台/搜索关键字
// @Summary 搜索关键字
// @Produce  application/JSON
// @Success 200 {object} response.SearchIndexRtnJSON ""
// @router wx/search/index [get]
func (c *SearchController) SearchIndex() {
	var searchIndexRtnJSON response.SearchIndexRtnJSON
	defaultKeyword, err := models.GetKeywordDefault()
	if err != nil {
		logger.Logger.Error("GetKeywordDefault is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	hotKeywordsList, err := models.GetKeywordHot()
	if err != nil {
		logger.Logger.Error("GetKeywordHot is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	historyKeyworkList, err := models.GetSearchHistory(c.UserID)
	if err != nil {
		logger.Logger.Error("searchHistoryList is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	searchIndexRtnJSON.DefaultKeyword = defaultKeyword
	searchIndexRtnJSON.HistoryKeyworkList = historyKeyworkList
	searchIndexRtnJSON.HotKeywordList = hotKeywordsList

	c.RespJSON(common.SUCCESEE, common.SCODE, searchIndexRtnJSON)
}

// SearchIndex 搜索帮助
// @Tags 前台/搜索帮助
// @Summary 搜索帮助
// @Produce  application/JSON
// @Param keyword query string true "关键字"
// @Success 200 {object} string "字符串数组"
// @router wx/search/helper [get]
func (c *SearchController) SearchHelper() {
	keyword := c.GetString("keyword")

	keywordList, err := models.SearchHelper(keyword)
	if err != nil {
		logger.Logger.Error("SearchHelper is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, keywordList)
}

// SearchClearHistory 搜索历史清理
// @Tags 前台/搜索历史清理
// @Summary 搜索历史清理
// @Produce  application/JSON
// @Param token header string true "token"
// @Success 200 {object} controllers.ResponseData ""
// @router wx/search/clearhistory [get]
func (c *SearchController) SearchClearHistory() {
	err := models.SearchHistoryClear(c.UserID)
	if err != nil {
		logger.Logger.Error("SearchHistoryClear is failed! err:[%v].", err)
		c.RespJSON(common.ErrDeleteData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// SearchResult 搜索结果
// @Tags 前台/搜索结果
// @Summary 搜索结果
// @Produce  application/JSON
// @Param categoryId query int true "categoryId"
// @Param name query string true "name"
// @Param sort query string true "排序字段"
// @Param order query string true "排序顺序"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Success 200 {object} response.SearchResultRtnJSON ""
// @router /search/result [get]
func (c *SearchController) SearchResult() {
	var query = make(map[string]string)
	var searchResultRtnJSON response.SearchResultRtnJSON
	goodsArray := []response.GoodsRtnJSON{}

	categoryID := c.GetString("categoryId")
	sort := c.GetString("sort")
	order := c.GetString("order")
	name := c.GetString("name")
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	limit, _ := c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}

	query["name"] = name
	if query["deleted"] == "" {
		query["deleted"] = "0"
	}

	// 搜索词不存在搜索历史中时插入
	isOK := models.GetSearchHistoryByName(name)
	if !isOK {
		// 写入搜索记录
		var searchHistory base.ShopSearchHistory
		searchHistory.UserID = c.UserID
		searchHistory.Keyword = name
		searchHistory.AddTime = time.Now()
		err := models.AddSearchHistory(searchHistory)
		if err != nil {
			logger.Logger.Errorf("AddSearchHistory is failed! err:[%v].", err)
		}
	}

	// 获取搜索结果的分类
	goodsArray, err := models.GetAllGoods(query, sort, order, page, limit)
	if err != nil {
		logger.Logger.Errorf("GetAllGoods is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 查找商品的CategoryID分类ID并利用map去重
	categoryIDList := []int{}
	m := make(map[int]bool) //map的值不重要
	for _, v := range goodsArray {
		if _, ok := m[v.CategoryID]; !ok {
			categoryIDList = append(categoryIDList, v.CategoryID)
			m[v.CategoryID] = true
		}
	}

	// 获取分类即二级菜单信息
	filterCategoryList := []response.CategoryRtnJSON{}
	for _, ID := range categoryIDList {
		category, err := models.GetOneCategory(ID)
		if err != nil {
			logger.Logger.Errorf("GetCategoryByIDs is failed! err:[%v].", err)
		}
		filterCategoryList = append(filterCategoryList, category)
	}

	if categoryID != "0" {
		query["categoryID"] = categoryID
	}
	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopGoods))

	searchResultRtnJSON.Limit = limit
	searchResultRtnJSON.Page = page
	if total == 0 {
		searchResultRtnJSON.List = goodsArray
		searchResultRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, searchResultRtnJSON)
		return
	}

	goodsList, err := models.GetAllGoods(query, sort, order, page, limit)
	if err != nil {
		logger.Logger.Errorf("GetAllGoods is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	searchResultRtnJSON.List = goodsList
	searchResultRtnJSON.FilterCategoryList = filterCategoryList
	searchResultRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, searchResultRtnJSON)
}
