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
)

// AddressController ...
type AddressController struct {
	MainController
}

// AddressList 获取地址列表
// @Tags 前台/获取地址列表
// @Summary 获取地址列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Success 200 {object} response.AddressListRtnJSON ""
// @router /address/list [get]
func (c *AddressController) AddressList() {
	addressList, total, err := models.GetAddress(c.UserID)
	if err != nil {
		logger.Logger.Error("get address list is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var addressListRtnJSON response.AddressListRtnJSON
	addressListRtnJSON.List = addressList
	addressListRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, addressListRtnJSON)
}

// AddressDetail 获取地址详情
// @Tags 前台/获取地址详情
// @Summary 获取地址详情
// @Produce  application/JSON
// @Param token header string true "token"
// @Param token query int true "ID"
// @Success 200 {object} response.AddressListRtnJSON ""
// @router /address/list [get]
func (c *AddressController) AddressDetail() {
	id, _ := c.GetInt("id")

	address, err := models.GetAddressByID(c.UserID, id)
	if err != nil {
		logger.Logger.Error("get address list is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, address)
}

// AddressSave 保存地址
// @Tags 前台/保存地址
// @Summary 保存地址
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.AddressSaveBody body request.AddressSaveBody true "地址内容"
// @Success 200 {object} controllers.ResponseData ""
// @router /address/save [post]
func (c *AddressController) AddressSave() {
	var addressSaveBody request.AddressSaveBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &addressSaveBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal add address Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}
	var isDefault int
	if addressSaveBody.IsDefault == true {
		isDefault = 1
	} else {
		isDefault = 0
	}

	address := base.ShopAddress{
		Name:          addressSaveBody.Name,
		Tel:           addressSaveBody.Tel,
		UserID:        c.UserID,
		Province:      addressSaveBody.Province,
		City:          addressSaveBody.City,
		County:        addressSaveBody.County,
		AddressDetail: addressSaveBody.AddressDetail,
		AreaCode:      addressSaveBody.AreaCode,
		IsDefault:     isDefault,
		AddTime:       time.Now(),
	}

	err = models.AddAddress(address)
	if err != nil {
		logger.Logger.Error("AddAddress is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// AddressDelete 删除地址
// @Tags 前台/删除地址
// @Summary 删除地址
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.AddressDeleteBody body request.AddressDeleteBody true "要删除的地址id"
// @Success 200 {object} controllers.ResponseData ""
// @router /address/delete [post]
func (c *AddressController) AddressDelete() {
	var addressDeleteBody request.AddressDeleteBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &addressDeleteBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal add cart Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	err = models.DeleteAddress(c.UserID, addressDeleteBody.ID)
	if err != nil {
		logger.Logger.Error("DeleteAddress is failed! err:[%v].", err)
		c.RespJSON(common.ErrDeleteData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}
