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
	"time"

	"github.com/astaxie/beego/orm"
)

// CollectController ...
type CollectController struct {
	MainController
}

// CollectList 收藏夹显示列表
// @Tags 前台/收藏夹显示列表
// @Summary 收藏夹显示列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param type query int true "收藏类型id"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Success 200 {object} response.CollectListRtnJSON ""
// @router /collect/list [get]
func (c *CollectController) CollectList() {
	typeID, _ := c.GetInt("type")
	page, _ := c.GetInt("page")
	if page <= 1 {
		page = 1
	}
	limit, _ := c.GetInt("limit")
	if limit <= 10 {
		limit = 10
	}

	var collectListRtnJSON response.CollectListRtnJSON
	colletcRespon := response.CollectRespon{}
	colletcResponList := []response.CollectRespon{}

	limit = 100
	// 获取收藏中所有的东西，按分页显示
	collectList, err := models.GetCollect(c.UserID, typeID, page, limit)
	if err != nil {
		logger.Logger.Errorf("get my collect list by userID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	for _, v := range collectList {
		// 判断是商品还是专题，0是商品， 1是专题
		if v.Type == 0 {
			goods, err := models.GetOneGoods(v.ValueID)
			if err != nil {
				logger.Logger.Errorf("get my collect list goods is failed! err:[%v]  v:[%v].", err, v)
			}
			colletcRespon.ID = v.ID
			colletcRespon.Type = v.Type
			colletcRespon.ValueID = v.ValueID
			colletcRespon.Name = goods.Name
			colletcRespon.Brief = goods.Brief
			colletcRespon.PicURL = goods.PicURL
			colletcRespon.RetailPrice = goods.RetailPrice
			colletcResponList = append(colletcResponList, colletcRespon)
		}
	}

	collectListRtnJSON.List = colletcResponList

	c.RespJSON(common.SUCCESEE, common.SCODE, collectListRtnJSON)
}

// CollectAddOrDelete 添加或取消收藏
// @Tags 前台/添加或取消收藏
// @Summary 添加或取消收藏
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CollectBody body request.CollectBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /collect/addordelete [post]
func (c *CollectController) CollectAddOrDelete() {
	var collectBody request.CollectBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &collectBody)
	if err != nil {
		logger.Logger.Errorf("input para json unmarshal is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	collect, err := models.GetOneCollect(c.UserID, collectBody.Type, collectBody.ValueID)
	if err != nil {
		// 未收藏就新增
		if err.Error() == orm.ErrNoRows.Error() {
			err = models.AddCollect(&base.ShopCollect{
				Type:    collectBody.Type,
				ValueID: collectBody.ValueID,
				UserID:  c.UserID,
				AddTime: time.Now(),
			})
			if err != nil {
				logger.Logger.Error("AddCollect is failed! err:[%v].", err)
				c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
				return
			}
			c.RespJSON(common.SUCCESEE, common.SCODE, "")
		}
		logger.Logger.Errorf("get my collect info by userID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 已收藏就删除
	collect.UpdateTime = time.Now()
	collect.Deleted = 1
	err = models.DeleteCollect(&collect)
	if err != nil {
		logger.Logger.Error("DeleteCollect is failed! err:[%v].", err)
		c.RespJSON(common.ErrDeleteData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}
