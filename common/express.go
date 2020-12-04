// Copyright 2020 The shop Authors

// Package common implements common.
package common

import (
	"encoding/json"
	"github.com/astaxie/beego"
	"github.com/astaxie/beego/httplib"
	"gopkg.in/yaml.v2"
	"io/ioutil"
	"net/url"
	"runtime"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
)

type RequestData struct {
	OrderCode    string `json:"OrderCode"`
	ShipperCode  string `json:"ShipperCode"`
	LogisticCode string `json:"LogisticCode"`
}

type ExpressResult struct {
	Success bool              `json:"success"`
	State   int               `json:"state"`
	Traces  []response.Traces `json:"traces"`
}

func QueryExpress(shippercode, logisticcode string, ordercode string) (expressinfo response.ExpressRtnInfo) {
	expressinfo = response.ExpressRtnInfo{
		Success:      false,
		ShipperCode:  shippercode,
		ShipperName:  "",
		LogisticCode: logisticcode,
		IsFinish:     0,
		Traces:       make([]response.Traces, 0),
	}
	fromData := GenerateFromData(shippercode, logisticcode, ordercode)
	postURL := beego.AppConfig.String("express::requestUrl")

	req := httplib.Post(postURL)
	req.Header("Content-Type", "application/x-www-form-urlencoded")
	req.Param("EBusinessID", fromData["EBusinessID"])
	req.Param("RequestData", fromData["RequestData"])
	req.Param("RequestType", fromData["RequestType"])
	req.Param("DataType", fromData["DataType"])
	req.Param("DataSign", fromData["DataSign"])

	var res ExpressResult
	req.ToJSON(&res)

	expressinfo.Success = res.Success
	if res.State == 3 {
		expressinfo.IsFinish = 1
	}
	expressinfo.Traces = append(expressinfo.Traces, res.Traces...)

	return expressinfo

}

func GenerateFromData(shippercode, logisticcode, ordercode string) map[string]string {
	requestData := RequestData{
		OrderCode:    ordercode,
		ShipperCode:  shippercode,
		LogisticCode: logisticcode,
	}
	requestDataByte, err := json.Marshal(requestData)
	if err != nil {
		panic(err)
	}
	requestDataString := string(requestDataByte)
	requestDataStringEscape := url.QueryEscape(requestDataString)

	var fromData map[string]string
	fromData = make(map[string]string)
	fromData["EBusinessID"] = beego.AppConfig.String("express::appID")
	fromData["RequestData"] = requestDataStringEscape
	fromData["RequestType"] = "1002"
	fromData["DataType"] = "2"
	fromData["DataSign"] = GenerateDataSign(requestDataString)
	return fromData
}

func GenerateDataSign(requestdata string) string {
	appkey := beego.AppConfig.String("express::appkey")
	md5str := utils.Md5(requestdata + appkey)
	base64str := utils.Base64Encode(md5str)
	rv := url.QueryEscape(base64str)
	return rv
}

type ExpressProvider struct {
	Code string `json:"code"`
	Name string `json:"name"`
}

// 获取物流公司供应商
func GetOrderChannel() ([]ExpressProvider, error) {
	var content []byte
	var err error
	// 根据不同的操作系统，做不同的路径转换
	if runtime.GOOS == "windows" {
		// 1.从conf/express配置文件中读取供应商信息
		content, err = ioutil.ReadFile(".//conf//express.yaml")
	} else {
		content, err = ioutil.ReadFile("./conf/express.yaml")
	}

	if err != nil {
		logger.Logger.Error("read express conf failed.err:", err)
		return nil, err
	}
	orderChan := []ExpressProvider{}
	err = yaml.Unmarshal(content, &orderChan)
	if err != nil {
		logger.Logger.Error("yaml unmarshal content failed. err:", err)
		return nil, err
	}
	return orderChan, err
}
