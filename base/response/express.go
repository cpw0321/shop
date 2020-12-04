// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// R 必填 O 非必填
// EBusinessID	String	用户ID	R
// OrderCode	String	订单编号	O
// ShipperCode	String	快递公司编码	R
// LogisticCode	String	物流运单号	O
// Success	Bool	成功与否	R
// Reason	String	失败原因	O
// State	String	物流状态：2-在途中,3-签收,4-问题件	R

//Traces
// AcceptTime	String	时间	R
// AcceptStation	String	描述	R
// Remark	String	备注

/*
{
	"EBusinessID": "1109259",
	"OrderCode": "",
	"ShipperCode": "SF",
	"LogisticCode": "118461988807",
	"Success": true,
	"State": 3,
	"Reason": null,
		"Traces": [{
		"AcceptTime": "2014/06/25 08:05:37",
		"AcceptStation": "正在派件..(派件人:邓裕富,电话:18718866310)[深圳 市]",
		"Remark": null
		},
		{
		"AcceptTime": "2014/06/25 04:01:28",
		"AcceptStation": "快件在 深圳集散中心 ,准备送往下一站 深圳 [深圳市]",
		"Remark": null
		},
		{
		"AcceptTime": "2014/06/25 01:41:06",
		"AcceptStation": "快件在 深圳集散中心 [深圳市]",
		"Remark": null
		},
		{
		"AcceptTime": "2014/06/24 20:18:58",
		"AcceptStation": "已收件[深圳市]",
		"Remark": null
		},
		{
		"AcceptTime": "2014/06/24 20:55:28",
		"AcceptStation": "快件在 深圳 ,准备送往下一站 深圳集散中心 [深圳市]",
		"Remark": null
		},
		{
		"AcceptTime": "2014/06/25 10:23:03",
		"AcceptStation": "派件已签收[深圳市]",
		"Remark": null
		},
		{
		"AcceptTime": "2014/06/25 10:23:03",
		"AcceptStation": "签收人是：已签收[深圳市]",
		"Remark": null
		}
	]
}
*/

type ExpressRtnInfo struct {
	ShipperCode  string   `json:"shipperCode"`  // 快递公司编码
	ShipperName  string   `json:"shipperName"`  // 快递公司名称
	LogisticCode string   `json:"logisticCode"` // 物流运单号
	IsFinish     int      `json:"isFinish"`
	Success      bool     `json:"success"`
	Traces       []Traces `json:"traces"`
	RequestTime  string   `json:"requestTime"`
}

type Traces struct {
	AcceptTime    string `json:"acceptTime"`
	AcceptStation string `json:"acceptStation"`
	Remark        string `json:"remark"`
}
