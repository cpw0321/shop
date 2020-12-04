// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"encoding/json"
	"shop/base"
	"shop/base/request"
	"shop/common"
	"shop/logger"
	"shop/models"
	"time"
)

// FeedbackController ...
type FeedbackController struct {
	MainController
}

// FeedbackSubmit 反馈
// @Tags 前台/反馈
// @Summary 反馈
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.FeedbackBody body request.FeedbackBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /feedback/submit [post]
func (c *FeedbackController) FeedbackSubmit() {
	var feedbackBody request.FeedbackBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &feedbackBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal feedbackBody Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	var picURLList string
	if feedbackBody.PicURLs != nil {
		picURLs, err := json.Marshal(feedbackBody.PicURLs)
		if err != nil {
			logger.Logger.Errorf("JSON marshal feedbackBody picURLs is failed! err:[%v].", err)
		}
		picURLList = string(picURLs)
	} else {
		picURLList = "[]"
	}
	var hasPicture int
	if feedbackBody.HasPicture == true {
		hasPicture = 1
	} else {
		hasPicture = 0
	}

	feedback := base.ShopFeedback{
		UserID:     c.UserID,
		Username:   c.Username,
		Mobile:     feedbackBody.Mobile,
		FeedType:   feedbackBody.FeedType,
		Content:    feedbackBody.Content,
		HasPicture: hasPicture,
		PicURLs:    picURLList,
		AddTime:    time.Now(),
	}

	err = models.FeedbackSubmit(&feedback)
	if err != nil {
		logger.Logger.Error("FeedbackSubmit is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}
