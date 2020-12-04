// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"github.com/astaxie/beego"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"net/http"
	"shop/common"
	"shop/logger"
	"shop/models"
	"time"
)

// BaseController 基础controller不带个人信息验证
type BaseController struct {
	beego.Controller
}

// ResponseData 返回体
type ResponseData struct {
	Errno  int         `json:"errno"`
	Errmsg string      `json:"errmsg"`
	Data   interface{} `json:"data"`
}

// Prepare 覆盖Controller的方法
func (c *BaseController) Prepare() {
	//jreq, err := json.Marshal(c.Input())
	//if err != nil {
	//	logger.Logger.Error(err.Error())
	//}
	//logger.Logger.Debug("request path :", c.Controller)
	//logger.Logger.Debug("method:", c.Ctx.Request.Method, " request header :", c.Ctx.Request.Header)
	//logger.Logger.Info("request param :", string(jreq))
	//logger.Logger.Debug("request param :", string(c.Ctx.Input.RequestBody))
}

// RespJSON 返回json
func (c *BaseController) RespJSON(errMsg string, statusCode int, data interface{}) {
	var res ResponseData
	res.Errmsg = errMsg
	res.Errno = statusCode
	res.Data = data
	//msgResp, err := json.Marshal(msg)
	//if err != nil {
	//	logger.Logger.Error(err.Error())
	//}
	//logger.Logger.Info("response data :", string(msgResp))

	c.Data["json"] = res
	c.ServeJSON()
}

// PayNotify 支付异步通知
func (c *BaseController) PayNotify() {
	// ====支付异步通知参数解析和验签Sign====
	// 解析支付异步通知的参数
	//    req：*http.Request
	//    返回参数 notifyReq：通知的参数
	//    返回参数 err：错误信息
	notifyReq, err := wechat.ParseNotify(c.Ctx.Request)
	// 查询订单支付结果
	orderResult, err := common.QueryOrder(notifyReq.OutTradeNo)
	if err != nil {
		logger.Logger.Errorf("QueryOrder is failed! err:[err].", err)
	}
	if orderResult.TradeState == "SUCCESS" {
		// 更新订单状态
		//err := models.UpdataOrderStatusBySn(notifyReq.OutTradeNo, common.STATUS_PAY)
		//if err != nil {
		//	logger.Logger.Errorf("UpdataOrderStatus is failed! err:[err].", err)
		//}
		// 更新订单信息
		orderInfo, err := models.GetOrderBySn(notifyReq.OutTradeNo)
		if err != nil {
			logger.Logger.Errorf("GetOrderBySn is failed! err:[err].", err)
		} else {
			payTime, err := time.ParseInLocation("20060102150405", notifyReq.TimeEnd, time.Local)
			if err != nil {
				logger.Logger.Errorf("TimeEnd string to time.Time is failed! err:[err].", err)
			}
			orderInfo.PayTime = payTime
			orderInfo.PayID = notifyReq.TransactionId
			orderInfo.OrderStatus = common.STATUS_PAY
			err = models.UpdateOrder(orderInfo)
			if err != nil {
				logger.Logger.Errorf("UpdateOrder is failed! err:[err].", err)
			}
		}
	}

	// ==异步通知，返回给微信平台的信息==
	rsp := new(wechat.NotifyResponse) // 回复微信的数据
	rsp.ReturnCode = gopay.SUCCESS
	rsp.ReturnMsg = gopay.OK
	c.Ctx.ResponseWriter.WriteHeader(http.StatusOK)
	c.Ctx.WriteString(rsp.ToXmlString())
}

// Test 测试物流信息
func (c *BaseController) Test() {
	resp := common.QueryExpress("ZTO", "订单号", "")
	c.RespJSON("success", 0, resp)

}
