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
	"time"
)

// AfterSaleController ...
type AfterSaleController struct {
	MainController
}

// AftersaleSubmit 提交售后
// @Tags 前台/提交订单
// @Summary 提交订单
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.AftersaleBody body request.AftersaleBody true "售后内容"
// @Success 200 {object} controllers.ResponseData ""
// @router /aftersale/submit [post]
func (c *AfterSaleController) AftersaleSubmit() {
	var aftersale request.AftersaleBody
	var aftersaleInfo base.ShopAftersale
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &aftersale)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal add cart Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}
	orderID, _ := strconv.Atoi(aftersale.OrderID)
	orderInfo, err := models.GetOrderByID(orderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
	}
	isOK := models.GetAftersaleByOrderIDIsExist(orderID)
	if !isOK {
		aftersaleInfo.OrderID = orderID
		aftersaleInfo.AftersaleSn = orderInfo.OrderSn
		aftersaleInfo.UserID = c.UserID
		aftersaleInfo.Type = aftersale.Type
		aftersaleInfo.Reason = aftersale.Reason
		aftersaleInfo.Amount = aftersale.Amount
		if aftersale.Pictures == nil {
			aftersaleInfo.Pictures = "[]"
		} else {
			pictures, err := json.Marshal(aftersale.Pictures)
			if err != nil {
				logger.Logger.Errorf("JSON marshal aftersale Info is failed! err:[%v].", err)
			}
			aftersaleInfo.Pictures = string(pictures)
		}

		aftersaleInfo.Comment = aftersale.TypeDesc
		aftersaleInfo.Status = 1
		aftersaleInfo.AddTime = time.Now()
		aftersaleInfo.UpdateTime = time.Now()
		_, err = models.AddAftersale(&aftersaleInfo)
		if err != nil {
			logger.Logger.Errorf("AddAftersale is failed! err:[%v].", err)
			c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
			return
		}
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// AftersaleList 获取售后列表
// @Tags 前台/获取售后列表
// @Summary 获取售后列表
// @Description 获取售后列表
// @Produce  application/JSON
// @Param   token header   string  true  "token"
// @Param page query int true "分页"
// @Param limit query int true "页大小"
// @Param status query string false "排序字段"
// @Param status query int false "售后状态"
// @Success 200 {object} response.AftersaleRtnJSON ""
// @router /aftersale/list [get]
func (c *AfterSaleController) AftersaleList() {
	status, _ := c.GetInt("status")
	//sort := c.GetString("sort")
	//order := c.GetString("order")
	page, _ := c.GetInt("page")
	limit, _ := c.GetInt("limit")
	if page <= 1 {
		page = 1
	}
	if limit <= 10 {
		limit = 10
	}
	var query = make(map[string]string)
	if status != 0 {
		query["status"] = strconv.Itoa(status)
	}
	if query["deleted"] == "" {
		query["deleted"] = "0"
	}

	query["userID"] = strconv.Itoa(c.UserID)
	var aftersaleRtnJSON response.AftersaleRtnJSON
	aftersaleRtnJSONList := []response.AftersaleRtn{}
	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopAftersale))
	aftersaleRtnJSON.Limit = limit
	aftersaleRtnJSON.Page = page
	if total == 0 {
		aftersaleRtnJSON.List = aftersaleRtnJSONList
		aftersaleRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, aftersaleRtnJSON)
		return
	}
	limit = 100
	aftersaleRtnJSONList, err := models.GetAllAftersale(query, "id", "desc", page, limit)
	if err != nil {
		logger.Logger.Errorf("get topic list is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	aftersaleRtnJSON.Total = total
	aftersaleRtnJSON.List = aftersaleRtnJSONList
	c.RespJSON(common.SUCCESEE, common.SCODE, aftersaleRtnJSON)
}

// AftersaleDetail 获取售后详情
// @Tags 前台/获取售后详情
// @Summary 获取售后详情
// @Description 获取售后详情
// @Produce  application/JSON
// @Param   token header   string  true  "token"
// @Param orderId query int true "订单状态"
// @Success 200 {object} response.AftersaleRtnJSON ""
// @router /aftersale/list [get]
func (c *AfterSaleController) AftersaleDetail() {
	orderID, _ := c.GetInt("orderId")

	var aftersaleDetailRtnJSON response.AftersaleDetailRtnJSON
	var err error
	aftersaleDetailRtnJSON.Order, err = models.GetOrderByID(orderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	aftersaleDetailRtnJSON.OrderGoods, err = models.GetOrderGoodsListByOrderID(orderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderGoodsListByOrderID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	aftersaleDetailRtnJSON.Aftersale, err = models.GetAftersaleByOrderID(orderID)
	if err != nil {
		logger.Logger.Errorf("GetAftersaleByOrderID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, aftersaleDetailRtnJSON)
}
