// Copyright 2020 The shop Authors

// Package common implements common.
package common

import (
	"fmt"
	"github.com/astaxie/beego"
	"github.com/iGoogle-ink/gopay"
	"github.com/iGoogle-ink/gopay/wechat"
	"github.com/iGoogle-ink/gotil"
	"shop/base/response"
	"shop/logger"
	"strconv"
	"time"
)

var client *wechat.Client

func InitWxPay() {
	appId := beego.AppConfig.String("weixin::appID")
	mchId := beego.AppConfig.String("weixin::mchID")
	apiKey := beego.AppConfig.String("weixin::apikey")
	//初始化微信客户端
	//    appId：应用ID
	//    mchId：商户ID
	//    apiKey：API秘钥值
	//    isProd：是否是正式环境
	client = wechat.NewClient(appId, mchId, apiKey, true)

	//设置国家
	client.SetCountry(wechat.China)
}

type PayInfo struct {
	OpenID         string
	Body           string
	OutTradeNo     string
	TotalFee       float64
	SpbillCreateIp string
}

// UnifiedOrder 统一下单接口
func UnifiedOrder(payInfo PayInfo) (unifiedOrderRtnJson response.UnifiedOrderRtnJSON, err error) {
	appId := beego.AppConfig.String("weixin::appID")
	apiKey := beego.AppConfig.String("weixin::apikey")
	notifyUrl := beego.AppConfig.String("weixin::notifyUrl")

	//初始化参数Map
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", payInfo.OutTradeNo)
	bm.Set("total_fee", payInfo.TotalFee*100)
	bm.Set("spbill_create_ip", payInfo.SpbillCreateIp)
	bm.Set("notify_url", notifyUrl)
	bm.Set("trade_type", wechat.TradeType_Mini)
	bm.Set("body", payInfo.Body)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", wechat.SignType_MD5)
	bm.Set("openid", payInfo.OpenID)

	//请求支付下单，成功后得到结果
	wxRsp, err := client.UnifiedOrder(bm)
	if err != nil {
		logger.Logger.Errorf("UnifiedOrder is failed! err:[%v].", err)
		return unifiedOrderRtnJson, err
	}

	// 获取微信小程序支付的 paySign
	//    appId：AppID
	//    nonceStr：随机字符串
	//    prepayId：统一下单成功后得到的值
	//    signType：签名方式，务必与统一下单时用的签名方式一致
	//    timeStamp：时间
	//    apiKey：API秘钥值
	timeStamp := strconv.FormatInt(time.Now().Unix(), 10)
	prepayId := "prepay_id=" + wxRsp.PrepayId // 此处的 wxRsp.PrepayId ,统一下单成功后得到
	paySign := wechat.GetMiniPaySign(appId, wxRsp.NonceStr, prepayId, wechat.SignType_MD5, timeStamp, apiKey)

	unifiedOrderRtnJson.TimeStamp = timeStamp
	unifiedOrderRtnJson.NonceStr = wxRsp.NonceStr
	unifiedOrderRtnJson.Package = prepayId
	unifiedOrderRtnJson.SignType = wechat.SignType_MD5
	unifiedOrderRtnJson.PaySign = paySign

	return unifiedOrderRtnJson, nil
}

// QueryOrder 查询订单
func QueryOrder(outTradeNo string) (orderResult response.OrderResult, err error) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", wechat.SignType_MD5)

	// 请求订单查询，成功后得到结果
	wxRsp, _, err := client.QueryOrder(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return orderResult, err
	}

	orderResult.TradeState = wxRsp.TradeState
	orderResult.TransactionID = wxRsp.TransactionId
	orderResult.TradeState = wxRsp.TradeState
	cashFee, _ := strconv.Atoi(wxRsp.CashFee)
	orderResult.CashFee = cashFee

	return orderResult, nil
}

// Refund 发起退款
func Refund(outTradeNo string, totalFee float64, refundFee float64) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", wechat.SignType_MD5)
	s := gotil.GetRandomString(64)
	fmt.Println("out_refund_no:", s)
	bm.Set("out_refund_no", s)
	bm.Set("total_fee", totalFee)
	bm.Set("refund_fee", refundFee)
	bm.Set("notify_url", "https://www.gopay.ink")

	//请求申请退款（沙箱环境下，证书路径参数可传空）
	//    body：参数Body
	//    certFilePath：cert证书路径
	//    keyFilePath：Key证书路径
	//    pkcs12FilePath：p12证书路径
	wxRsp, resBm, err := client.Refund(bm, "iguiyu_cert/apiclient_cert.pem", "iguiyu_cert/apiclient_key.pem", "iguiyu_cert/apiclient_cert.p12")
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", *wxRsp)
	fmt.Println("resBm:", resBm)
}

// QueryRefund 查询退款
func QueryRefund(outTradeNo string) {
	// 初始化参数结构体
	bm := make(gopay.BodyMap)
	bm.Set("out_trade_no", outTradeNo)
	bm.Set("nonce_str", gotil.GetRandomString(32))
	bm.Set("sign_type", wechat.SignType_MD5)

	//请求申请退款
	wxRsp, resBm, err := client.QueryRefund(bm)
	if err != nil {
		fmt.Println("Error:", err)
		return
	}
	fmt.Println("wxRsp：", *wxRsp)
	fmt.Println("resBm：", resBm)

}
