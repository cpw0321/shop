// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"shop/common"
	"shop/logger"
	"shop/utils"
	"strings"
)

// MainController ...
type MainController struct {
	UserID   int    `json:"userid"`
	Username string `json:"username"`
	BaseController
}

// Prepare 重写prepare方法
func (c *MainController) Prepare() {
	//jreq, err := json.Marshal(c.Input())
	//if err != nil {
	//	logger.Logger.Error(err.Error())
	//}
	//logger.Logger.Info("request path :", c.Controller)
	//logger.Logger.Info("method:", c.Ctx.Request.Method, "request header :", c.Ctx.Request.Header)
	//logger.Logger.Info("request param :", string(jreq))
	//logger.Logger.Info("request param :", string(c.Ctx.Input.RequestBody))

	token := strings.TrimSpace(c.Ctx.Request.Header.Get("X-Shop-Token"))
	tokenCustomClaims, err := utils.ParseToken(token)
	if tokenCustomClaims != nil && err == nil {
		c.UserID = tokenCustomClaims.UserID
		c.Username = tokenCustomClaims.Username
	}
	if err != nil {
		logger.Logger.Error(err.Error())
		c.RespJSON(common.ErrAuth.Error(), common.AUTH_INVALID_ACCOUNT, "")
		return
	}
	//logger.Logger.Infof("user info : userID:[%v] username:[%v].", c.UserID, c.Username)
}
