// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/base/response"
	"shop/common"
	"shop/models"
)

// UserController ...
type UserController struct {
	MainController
}

// UserIndex 个人中心订单状态数量
// @Tags 前台/个人中心订单状态数量
// @Summary 个人中心订单状态数量
// @Produce  application/JSON
// @Param token header string true "token"
// @Success 200 {object} response.UserIndexRtnJSON ""
// @router wx/user/index [get]
func (c *UserController) UserIndex() {
	var userIndexRtnJSON response.UserIndexRtnJSON
	userIndexRtnJSON = models.GetOrderStatusCount(c.UserID)
	c.RespJSON(common.SUCCESEE, common.SCODE, userIndexRtnJSON)
}
