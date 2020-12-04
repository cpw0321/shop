// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"github.com/astaxie/beego"
	"os"
	"path"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"strconv"
	"time"
)

// StorageController ...
type StorageController struct {
	BaseController
}

// StorageUpload 上传文件
// @Tags 前台/上传文件
// @Summary 上传文件
// @Produce  application/JSON
// @Param token header string true "token"
// @Param file body string true "file"
// @Success 200 {object} response.UploadRtnJSON ""
// @router wx/storage/upload [post]
func (c *StorageController) StorageUpload() {
	prefixURL := beego.AppConfig.String("upload::prefixURL")
	f, h, err := c.GetFile("file") // 获取上传的文件
	if err != nil {
		logger.Logger.Error("get file faile,err:", err)
		c.RespJSON(common.FAILED, common.FCODE, "")
	}
	defer f.Close()
	//验证后缀名是否符合要求
	// var AllowExtMap map[string]bool = map[string]bool{
	//     ".jpg":true,
	//     ".jpeg":true,
	//     ".png":true,
	// }
	// if _,ok:=AllowExtMap[ext];!ok{
	//     c.Ctx.WriteString( "后缀名不符合上传要求" )
	//     return
	// }
	// 创建目录
	uploadDir := "static/upload/img/" + time.Now().Format("20060102/")
	err = os.MkdirAll(uploadDir, os.ModePerm)
	if err != nil {
		logger.Logger.Errorf("MkdirAll is failed! err:[%v]. v :%v", err, uploadDir)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}
	// 构造文件名称
	fileExt := path.Ext(h.Filename)
	NewFileName := strconv.FormatInt(time.Now().Unix(), 10) + fileExt

	filePath := uploadDir + NewFileName

	// 关闭上传的文件，不然的话会出现临时文件不能清除的情况
	defer f.Close()
	err = c.SaveToFile("file", filePath)
	if err != nil {
		logger.Logger.Errorf("SaveToFile is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}
	picURL := prefixURL + "/" + filePath
	var uploadRtnJSON response.UploadRtnJSON
	uploadRtnJSON.URL = picURL

	c.RespJSON(common.SUCCESEE, common.SCODE, uploadRtnJSON)
}
