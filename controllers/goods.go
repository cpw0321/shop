// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"encoding/json"
	"shop/base"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"shop/utils"
	"strconv"
	"strings"
	"time"
)

// GoodsController ...
type GoodsController struct {
	BaseController
}

// GoodsDetail 获得商品的详情
// @Tags 前台/获得商品的详情
// @Summary 获得商品的详情
// @Produce  application/JSON
// @Param token header string true "token"
// @Param id query int true "商品id"
// @Success 200 {object} response.GoodsDetailRtnJSON ""
// @router /goods/detail [get]
func (c *GoodsController) GoodsDetail() {
	var detailRtnJSON response.GoodsDetailRtnJSON
	goodsID, _ := c.GetInt("id")
	//userID, _ := c.GetInt("userID")
	var userID int
	token := strings.TrimSpace(c.Ctx.Request.Header.Get("X-Shop-Token"))
	if token != "" {
		tokenCustomClaims, err := utils.ParseToken(token)
		if tokenCustomClaims != nil && err == nil {
			userID = tokenCustomClaims.UserID
		}
	}

	// 获取单个商品信息
	goodsInfo, err := models.GetOneGoods(goodsID)
	if err != nil {
		logger.Logger.Error("get one goods info failed! err:[%v]", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
	}

	// 获取商品规格列表
	goodsSpecificationList, err := models.GetGoodsSpecification(goodsInfo.ID)
	if err != nil {
		logger.Logger.Error("get one goods specificationinfo list failed! err:[%v]", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
	}

	var specificationItem response.SpecificationItem
	var specificationItemList []response.SpecificationItem
	m := make(map[string]bool) // map的值不重要，用来去重
	// 拼装specificationList中name和vauleList字段
	for _, v := range goodsSpecificationList {
		specificationItem.Name = v.Specification
		specificationItem.ValueList, err = models.GetInfoBySpecification(goodsInfo.ID, v.Specification)
		if err != nil {
			logger.Logger.Error("GetInfoBySpecification is failed! err:[%v]", err)
		}
		// 去重，相同name的规格只显示一次不重复显示
		if _, ok := m[v.Specification]; !ok {
			specificationItemList = append(specificationItemList, specificationItem)
			m[v.Specification] = true
		}
	}

	// 获取问题列表
	issueList, err := models.GetIssue()
	if err != nil {
		logger.Logger.Error("get issue list info failed! err:[%v]", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
	}

	// 获取商品规格详细信息
	goodsProductList, err := models.GetGoodsProduct(goodsInfo.ID)
	if err != nil {
		logger.Logger.Error("get one goods product list info failed! err:[%v]", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
	}

	// 获取评论
	var comment response.GoodsComment
	allCount, _, err := models.GetCommentCount(0, goodsID)
	if err != nil {
		logger.Logger.Errorf("GetCommentCount is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	comment.Count = allCount

	var commentListArray []response.CommentListRespon
	var commentListRtnJSON response.CommentListRespon
	commentList, err := models.GetCommentListByGoosID(goodsID)
	if err != nil {
		logger.Logger.Errorf("GetCommentListByGoosID is failed! err:[%v].", err)
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
		commentListRtnJSON.Nickname = user.Nickname
		commentListRtnJSON.Avatar = user.Avatar
		commentListRtnJSON.ID = v.ID
		commentListRtnJSON.Content = v.Content
		commentListRtnJSON.PicList = picList
		commentListRtnJSON.AddTime = utils.FormatTimestampStr(v.AddTime)
		commentListArray = append(commentListArray, commentListRtnJSON)
	}
	comment.Data = commentListArray

	var share bool
	wxShareStr, err := models.GetSystemByName(common.SHOP_WX_SHARE)
	if err != nil {
		logger.Logger.Errorf("models.GetSystemByName is failed! err:[%v].", err)
	}
	if wxShareStr == "0" {
		share = false
	} else {
		share = true
	}

	var hasCollect int
	// 获取是否收藏
	isOK := models.GetCollectExist(userID, 0, goodsID)
	if isOK {
		hasCollect = 1
	} else {
		hasCollect = 0
	}

	// 获取品牌信息
	//brand, err := models.GetBrandByID(goodsInfo.BrandID)
	//if err != nil {
	//	logger.Logger.Errorf("GetBrandByID is failed! err:[%v].", err)
	//}

	// 添加用户足迹
	if userID != 0 {
		footprint := base.ShopFootprint{
			UserID:  userID,
			GoodsID: goodsInfo.ID,
			AddTime: time.Now(),
		}
		// 添加足迹
		err = models.AddFootprint(footprint)
		if err != nil {
			logger.Logger.Errorf("AddFootprint is failed! err:[%v].", err)
		}
	}

	// 获取商品属性
	goodsAttrList, err := models.GetGoodsAttrList(goodsInfo.ID)
	if err != nil {
		logger.Logger.Errorf("GetGoodsAttrList is failed! err:[%v].", err)
	}

	detailRtnJSON.SpecificationList = specificationItemList
	detailRtnJSON.Issues = issueList
	detailRtnJSON.UserHasCollect = hasCollect
	detailRtnJSON.ShareImage = goodsInfo.ShareURL
	detailRtnJSON.Comment = comment
	detailRtnJSON.Share = share
	detailRtnJSON.Goods = goodsInfo
	detailRtnJSON.ProductList = goodsProductList
	//detailRtnJSON.Brand = brand
	detailRtnJSON.Attribute = goodsAttrList

	c.RespJSON(common.SUCCESEE, common.SCODE, detailRtnJSON)
}

// GoodsCategory 获得商品分类数据
// @Tags 前台/获得商品分类数据
// @Summary 获得商品分类数据
// @Produce  application/JSON
// @Param id query int true "分类id"
// @Param page query int true "分页"
// @Param limit query int true "页大小"
// @Success 200 {object} response.GoodsCategoryRtnJSON ""
// @router /goods/detail [get]
func (c *GoodsController) GoodsCategory() {
	var goodsCategoryRtnJSON response.GoodsCategoryRtnJSON
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	limit, _ := c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}

	categoryID, _ := c.GetInt("id")

	// 获取当前商品的分类
	currentCategory, err := models.GetOneCategory(categoryID)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	// 根据分类ID获取当前一级分类菜单信息
	parentcategory, err := models.GetOneCategory(currentCategory.PID)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	// 获取一级分类菜单下的子菜单
	brotherCategory, err := models.GetAllCategory(map[string]string{"pID": strconv.Itoa(parentcategory.ID)}, "ID", "asc", page, limit)
	if err != nil {
		logger.Logger.Error("get banners info failed! err:[%v].", err)
	}

	goodsCategoryRtnJSON.CurrentCategory = currentCategory
	goodsCategoryRtnJSON.ParentCategory = parentcategory
	goodsCategoryRtnJSON.BrotherCategory = brotherCategory

	c.RespJSON(common.SUCCESEE, common.SCODE, goodsCategoryRtnJSON)
}

// GoodsList 商品列表
// @Tags 前台/获得商品分类数据
// @Summary 获得商品分类数据
// @Produce  application/JSON
// @Param isNew query bool true "是否新品，true或者false"
// @Param isHot query bool true "是否热卖商品，true或者false"
// @Param keyword query string true "关键字，如果设置则查询是否匹配关键字"
// @Param brandId query int true "品牌商ID，如果设置则查询品牌商所属商品"
// @Param categoryId query int true "商品分类ID，如果设置则查询分类所属商品"
// @Param page query int true "分页"
// @Param limit query int true "页大小"
// @Param sort query string true "排序字段"
// @Param order query string true "升序降序"
// @Success 200 {object} response.GoodsListRtnJSON ""
// @router /goods/list [get]
func (c *GoodsController) GoodsList() {
	var goodsListRtnJSON response.GoodsListRtnJSON
	goodsRtnJSONList := []response.GoodsRtnJSON{}
	filterCategoryList := []response.CategoryRtnJSON{}

	var isNewStr string
	var isHotStr string
	var query = make(map[string]string)

	sort := c.GetString("sort")
	order := c.GetString("order")
	isNew, _ := c.GetBool("isNew")
	if isNew {
		isNewStr = "1"
		query["isNew"] = isNewStr
	}
	isHot, _ := c.GetBool("isHot")
	if isHot {
		isHotStr = "1"
		query["isHot"] = isHotStr
	}
	keyword := c.GetString("keyword")
	categoryID, _ := c.GetInt("categoryId")
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	limit, _ := c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}

	if keyword != "" {
		query["name"] = keyword
	}
	// 商品分类ID不能为0
	if categoryID != 0 {
		query["categoryID"] = strconv.Itoa(categoryID)
	}
	if query["deleted"] == "" {
		query["deleted"] = "0"
	}

	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopGoods))
	goodsListRtnJSON.Limit = limit
	goodsListRtnJSON.Page = page
	if total == 0 {
		goodsListRtnJSON.List = goodsRtnJSONList
		goodsListRtnJSON.FilterCategoryList = filterCategoryList
		goodsListRtnJSON.Total = int(total)
		c.RespJSON(common.SUCCESEE, common.SCODE, goodsListRtnJSON)
		return
	}
	limit = 100
	goodsRtnJSONList, err := models.GetAllGoods(query, sort, order, page, limit)
	if err != nil {
		logger.Logger.Errorf("GetAllGoods is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 查找商品的CategoryID分类ID并利用map去重
	categoryIDList := []int{}
	m := make(map[int]bool) //map的值不重要
	for _, v := range goodsRtnJSONList {
		if _, ok := m[v.CategoryID]; !ok {
			categoryIDList = append(categoryIDList, v.CategoryID)
			m[v.CategoryID] = true
		}
	}

	// 获取分类即二级菜单信息
	for _, ID := range categoryIDList {
		category, err := models.GetOneCategory(ID)
		if err != nil {
			logger.Logger.Error("GetCategoryByIDs is failed! err:[%v].", err)
		}
		filterCategoryList = append(filterCategoryList, category)
	}

	goodsListRtnJSON.List = goodsRtnJSONList
	goodsListRtnJSON.FilterCategoryList = filterCategoryList
	goodsListRtnJSON.Total = int(total)
	c.RespJSON(common.SUCCESEE, common.SCODE, goodsListRtnJSON)
}

// GoodsCount 商品数量
// @Tags 前台/商品数量
// @Summary 商品数量
// @Produce  application/JSON
// @Success 200 {object} response.FootprintListRtnJSON ""
// @router /footprint/list [get]
func (c *GoodsController) GoodsCount() {
	count, err := models.GetGoodsCount()
	if err != nil {
		logger.Logger.Errorf("get goods count is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, count)
}
