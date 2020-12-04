// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"encoding/json"
	"io/ioutil"
	"net/http"
	"shop/base"
	"shop/base/request"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"shop/utils"
	"time"

	"strings"

	"github.com/astaxie/beego"
)

// AuthController ...
type AuthController struct {
	BaseController
}

// AuthLoginByWeixin 微信登录
// @Tags 前台/微信登录
// @Summary 微信登录
// @Produce  application/JSON
// @Param request.AuthLoginBody body request.AuthLoginBody true "用户信息"
// @Success 200 {object} response.AuthUserRtnJSON ""
// @router /auth/loginbyweixin [post]
func (c *AuthController) AuthLoginByWeixin() {
	var reqUserBody request.AuthLoginBody
	var userInfo base.ShopUser // 用户信息
	var authUserRtnJSON response.AuthUserRtnJSON
	body := c.Ctx.Input.RequestBody

	err := json.Unmarshal(body, &reqUserBody)
	if err != nil {
		logger.Logger.Errorf("json unmarshal user info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	appid := beego.AppConfig.String("weixin::appID")
	secret := beego.AppConfig.String("weixin::secret")
	CodeToSessURL := beego.AppConfig.String("weixin::codeToSessURL")
	CodeToSessURL = strings.Replace(CodeToSessURL, "{appid}", appid, -1)
	CodeToSessURL = strings.Replace(CodeToSessURL, "{secret}", secret, -1)
	CodeToSessURL = strings.Replace(CodeToSessURL, "{code}", reqUserBody.Code, -1)
	// 获取用户OpenID等信息
	resp, err := http.Get(CodeToSessURL)
	if err != nil {
		logger.Logger.Errorf("get CodeToSessURL failed err:", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	defer resp.Body.Close()

	data, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		logger.Logger.Errorf("ioutil read failed err:", err)
		return
	}

	var value request.WXInfo
	err = json.Unmarshal(data, &value)
	if err != nil {
		logger.Logger.Errorf("JSON Unmarshal failed err:", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 判断用户的openid是否正常获取
	if value.OpenID == "" {
		logger.Logger.Errorf("get code and openid data:[%v]", string(data))
		c.RespJSON(common.ErrAuth.Error(), common.FCODE, "")
		return
	}

	clientIP := c.Ctx.Input.IP()

	// 根据openID查询用户是否存在
	isOK := models.GetUserIsExistByOpenID(value.OpenID)
	// 不存在创建用户
	if !isOK {
		userInfo = base.ShopUser{
			OpenID:      value.OpenID,
			SessionKey:  value.SessionKey,
			Username:    reqUserBody.UserInfo.NickName,
			Nickname:    reqUserBody.UserInfo.NickName,
			Gender:      reqUserBody.UserInfo.Gender,
			Avatar:      reqUserBody.UserInfo.AvatarURL,
			AddTime:     time.Now(),
			LastLoginIP: clientIP,
		}
		_, err := models.AddUser(&userInfo)
		if err != nil {
			logger.Logger.Errorf("add user info failed! err:[%v].", err)
			c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
			return
		}
	}
	userInfo, err = models.GetUserByOpenID(value.OpenID)
	if err != nil {
		logger.Logger.Errorf("get user info by openID failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 微信用户头像，名称改变更新用户信息
	if userInfo.Nickname != reqUserBody.UserInfo.NickName || userInfo.Avatar != reqUserBody.UserInfo.AvatarURL {
		userInfo.Username = reqUserBody.UserInfo.NickName
		userInfo.Nickname = reqUserBody.UserInfo.NickName
		userInfo.Avatar = reqUserBody.UserInfo.AvatarURL
		userInfo.UpdateTime = time.Now()
	}
	userInfo.LastLoginTime = time.Now()
	userInfo.LastLoginIP = clientIP
	err = models.UpdateUser(&userInfo)
	if err != nil {
		logger.Logger.Error("update user info failed! err:[%v]", err)
	}

	// 用户ID和nickName创建token
	token, err := utils.CreateToken(userInfo.ID, userInfo.Nickname)
	if err != nil {
		logger.Logger.Errorf("create token failed! err:[%v].", err)
		c.RespJSON(common.ErrCreateData.Error(), common.FCODE, "")
		return
	}

	authUserRtnJSON.UserInfo.NickName = userInfo.Nickname
	authUserRtnJSON.UserInfo.AvatarURL = userInfo.Avatar
	authUserRtnJSON.Token = token
	c.RespJSON(common.SUCCESEE, common.SCODE, authUserRtnJSON)
}

// AuthLogin ...
func (c *AuthController) AuthLogin() {

}

// AuthRegister ...
func (c *UserController) AuthRegister() {

}

// AuthSendCode ...
func (c *UserController) AuthSendCode() {

}
