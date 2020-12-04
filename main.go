// Copyright 2020 The shop Authors

// Package main implements main.
package main

import (
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/plugins/cors"
	"shop/common"
	"shop/logger"
	"shop/models"
	_ "shop/routers"
)

// init ...
func init() {
	// 初始化日志
	logger.InitLogger()
	// 注册数据库
	models.RegisterDB()
	// 初始化微信支付
	common.InitWxPay()
}

// main ...
func main() {
	if beego.BConfig.RunMode == "dev" {
		beego.BConfig.WebConfig.DirectoryIndex = true
		beego.BConfig.WebConfig.StaticDir["/swagger"] = "swagger"
	}

	// 跨域
	// InsertFilter 是提供一个过滤函数
	beego.InsertFilter("*", beego.BeforeRouter, cors.Allow(&cors.Options{
		// 允许访问所有源
		AllowAllOrigins: true,
		// 可选参数"GET", "POST", "PUT", "DELETE", "OPTIONS" (*为所有)
		// 其中Options跨域复杂请求预检
		AllowMethods: []string{"*"},
		// 指的是允许的Header的种类
		AllowHeaders: []string{"*"},
		// 公开的HTTP标头列表
		ExposeHeaders: []string{"Content-Length"},
		// 如果设置，则允许共享身份验证凭据，例如cookie
		AllowCredentials: true,
	}))

	// 处理订单超时问题
	go common.HandleOrder()

	beego.Run()
}
