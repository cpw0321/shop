// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"encoding/json"
	"shop/base"
	"shop/base/request"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"strconv"
	"time"
)

// OrderController ...
type OrderController struct {
	MainController
}

// OrderList 订单列表
// @Tags 前台/订单列表
// @Summary 订单列表
// @Produce  application/JSON
// @Param token header string true "token"
// @Param showType query int true "showType"
// @Param sort query string true "排序字段"
// @Param order query string true "排序顺序"
// @Param page query int true "页数"
// @Param limit query int true "页大小"
// @Success 200 {object} response.OrderListRtnJSON ""
// @router /order/list [get]
func (c *OrderController) OrderList() {
	var orderListRtnJSON response.OrderListRtnJSON
	query := make(map[string]string)

	showType, _ := c.GetInt("showType")
	sort := c.GetString("sort")
	if sort == "" {
		sort = "ID"
	}
	order := c.GetString("order")
	if order == "" {
		order = "desc"
	}
	page, _ := c.GetInt("page")
	if page <= 0 {
		page = 1
	}
	limit, _ := c.GetInt("limit")
	if limit <= 0 {
		limit = 10
	}

	var orderStatus string
	var status int
	if showType != 0 {
		switch showType {
		case 1:
			// 待付款订单
			status = 101
		case 2:
			// 待发货订单
			status = 201
		case 3:
			// 待收货订单
			status = 301
		case 4:
			// 待评价订单
			status = 401
		}
		orderStatus = strconv.Itoa(status)
		query["OrderStatus"] = orderStatus
	}

	if query["deleted"] == "" {
		query["deleted"] = "0"
	}
	query["userID"] = strconv.Itoa(c.UserID)

	var orderList []response.OrderRtnJSON
	orderList = []response.OrderRtnJSON{}
	// 计算数据总数
	total, _ := models.GetTotal(query, new(base.ShopOrder))
	orderListRtnJSON.Limit = limit
	orderListRtnJSON.Page = page
	if total == 0 {
		orderListRtnJSON.List = orderList
		orderListRtnJSON.Total = total
		c.RespJSON(common.SUCCESEE, common.SCODE, orderListRtnJSON)
		return
	}

	limit = 100
	orderList, err := models.GetAllOrder(query, sort, order, page, limit)
	if err != nil {
		logger.Logger.Errorf("GetAllOrder is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	orderListRtnJSON.List = orderList
	orderListRtnJSON.Total = total
	c.RespJSON(common.SUCCESEE, common.SCODE, orderListRtnJSON)
}

// OrderDetail 订单详情
// @Tags 前台/订单详情
// @Summary 订单详情
// @Produce  application/JSON
// @Param token header string true "token"
// @Param orderId query int true "订单id"
// @Success 200 {object} response.OrderDetailRtnJSON ""
// @router /order/detail [get]
func (c *OrderController) OrderDetail() {
	orderID, _ := c.GetInt("orderId")

	OrderDetailInfo, err := models.OrderDetail(c.UserID, orderID)
	if err != nil {
		logger.Logger.Error("get OrderDetail is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, OrderDetailInfo)
}

// OrderSubmit 提交订单
// @Tags 前台/提交订单
// @Summary 提交订单
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderSubmitBody body request.OrderSubmitBody true "订单请求体"
// @Success 200 {object} response.OrderSubmitRtnJSON ""
// @router /order/submit [post]
func (c *OrderController) OrderSubmit() {
	var orderSubmitBody request.OrderSubmitBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderSubmitBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Submit Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	var orderSubmitRtnJSON response.OrderSubmitRtnJSON
	// 提交订单，写入订单表ShopOrder和订单商品表ShopOrderGoods
	orderSubmitRtnJSON, err = models.OrderSubmit(c.UserID, orderSubmitBody.CartID, orderSubmitBody.AddressID, orderSubmitBody.CouponID, orderSubmitBody.Message)
	if err != nil {
		logger.Logger.Error("OrderSubmit is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	// 更新用户优惠卷表中的优惠卷的状态
	if orderSubmitBody.UserCouponID > 0 {
		status := 1 // 1 代表优惠卷已经使用
		err = models.UpdateCouponUserStatus(c.UserID, orderSubmitBody.UserCouponID, status)
		if err != nil {
			logger.Logger.Error("UpdateCouponUserStatus is failed! err:[%v].", err)
			c.RespJSON(common.ErrUpdateData.Error(), common.FCODE, "")
			return
		}
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, orderSubmitRtnJSON)
}

// OrderCancel 取消订单
// @Tags 前台/取消订单
// @Summary 取消订单
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderCancelBody body request.OrderCancelBody true "订单请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /order/cancel [post]
func (c *OrderController) OrderCancel() {
	var orderCancelBody request.OrderCancelBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderCancelBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Submit Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}
	// 获取订单信息
	orderInfo, err := models.GetOrderByID(orderCancelBody.OrderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
		c.RespJSON("订单信息查询失败", common.FCODE, "")
		return
	}
	// 检查订单能否取消
	handleOption, err := models.GetOrderHandleOption(orderInfo.ID)
	if err != nil {
		logger.Logger.Errorf("GetOrderHandleOption is failed! err:[%v].", err)
		c.RespJSON("获取订单信息状态失败", common.FCODE, "")
		return
	}
	if !handleOption.Cancel {
		logger.Logger.Errorf("handleOption:[%v]  handleOption.Cancel:[%v].", handleOption, handleOption.Cancel)
		c.RespJSON("订单不能取消", common.FCODE, "")
		return
	}

	err = models.UpdataOrderStatus(orderCancelBody.OrderID, common.STATUS_CANCEL)
	if err != nil {
		logger.Logger.Error("OrderCancel is failed! err:[%v].", err)
		c.RespJSON("订单取消失败", common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// OrderRefund 订单退款
// @Tags 前台/订单退款
// @Summary 订单退款
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderRefundBody body request.OrderRefundBody true "订单请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /order/refund [post]
func (c *OrderController) OrderRefund() {
	var orderRefundBody request.OrderRefundBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderRefundBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Submit Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 获取订单信息
	orderInfo, err := models.GetOrderByID(orderRefundBody.OrderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
		c.RespJSON("订单信息查询失败", common.FCODE, "")
		return
	}
	// 检查订单能否取消
	handleOption, err := models.GetOrderHandleOption(orderInfo.ID)
	if err != nil {
		logger.Logger.Errorf("GetOrderHandleOption is failed! err:[%v].", err)
		c.RespJSON("获取订单信息状态失败", common.FCODE, "")
		return
	}
	if !handleOption.Refund {
		logger.Logger.Errorf("handleOption:[%v]  handleOption.Cancel:[%v].", handleOption, handleOption.Cancel)
		c.RespJSON("订单不能退款", common.FCODE, "")
		return
	}

	err = models.UpdataOrderStatus(orderRefundBody.OrderID, common.STATUS_REFUND)
	if err != nil {
		logger.Logger.Error("OrderCancel is failed! err:[%v].", err)
		c.RespJSON("订单退款失败", common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// OrderConfirm 确认收货
// @Tags 前台/确认收货
// @Summary 确认收货
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderConfirmBody body request.OrderConfirmBody true "订单请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /order/confirm [post]
func (c *OrderController) OrderConfirm() {
	var orderConfirmBody request.OrderConfirmBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderConfirmBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Submit Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 获取订单信息
	orderInfo, err := models.GetOrderByID(orderConfirmBody.OrderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
		c.RespJSON("订单信息查询失败", common.FCODE, "")
		return
	}
	// 检查订单能否取消
	handleOption, err := models.GetOrderHandleOption(orderInfo.ID)
	if err != nil {
		logger.Logger.Errorf("GetOrderHandleOption is failed! err:[%v].", err)
		c.RespJSON("获取订单信息状态失败", common.FCODE, "")
		return
	}
	if !handleOption.Confirm {
		logger.Logger.Errorf("handleOption:[%v]  handleOption.Cancel:[%v].", handleOption, handleOption.Cancel)
		c.RespJSON("订单不能确认收货", common.FCODE, "")
		return
	}

	err = models.UpdataOrderStatus(orderConfirmBody.OrderID, common.STATUS_CONFIRM)
	if err != nil {
		logger.Logger.Error("OrderCancel is failed! err:[%v].", err)
		c.RespJSON("订单确认收货失败", common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// OrderGoods 代评价商品信息
// @Tags 前台/代评价商品信息
// @Summary 代评价商品信息
// @Produce  application/JSON
// @Param token header string true "token"
// @Param orderId header int true "订单id"
// @Param goodsId header int true "商品id"
// @Success 200 {object} base.ShopOrderGoods "代评价商品信息"
// @router /order/goods [get]
func (c *OrderController) OrderGoods() {
	orderID, _ := c.GetInt("orderId")
	goodsID, _ := c.GetInt("goodsId")

	orderGoodsInfo, err := models.GetOrderGoodsComment(orderID, goodsID)
	if err != nil {
		logger.Logger.Errorf("GetOrderGoodsList is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, orderGoodsInfo)
}

// OrderComment 评价订单商品信息
// @Tags 前台/评价订单商品信息
// @Summary 评价订单商品信息
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderCommentBody body request.OrderCommentBody true "订单请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /order/comment [post]
func (c *OrderController) OrderComment() {
	var orderCommentBody request.OrderCommentBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderCommentBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Comment Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	orderGoods, err := models.GetOrderGoods(orderCommentBody.OrderGoodsID)
	if err != nil {
		logger.Logger.Errorf("GetOrderGoods is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var comment base.ShopComment
	comment.ValueID = orderGoods.GoodsID // 商品ID
	comment.Type = 0                     // 0 商品评论， 1 专题评论
	comment.Content = orderCommentBody.Content
	comment.UserID = c.UserID
	var hasPicture int
	if orderCommentBody.HasPicture {
		hasPicture = 1
	} else {
		hasPicture = 0
	}
	comment.HasPicture = hasPicture
	picURLs, err := json.Marshal(orderCommentBody.PicURLs)
	if err != nil {
		logger.Logger.Errorf("marshal picURLs into string is failed! err:[%v].", err)
	}
	comment.PicURLs = string(picURLs)
	comment.Star = orderCommentBody.Star
	comment.AddTime = time.Now()

	err = models.AddOrderGoodsComment(orderGoods.ID, &comment)
	if err != nil {
		logger.Logger.Errorf("AddOrderGoodsComment is failed! err:[%v].", err)
		c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
		return
	}

	// 根据orderID的获取订单商品是否评论，更改订单状态
	orderGoodsInfo, err := models.GetOrderGoods(orderCommentBody.OrderGoodsID)
	if err != nil {
		logger.Logger.Errorf("GetOrderGoods is failed! err:[%v].", err)
	}
	goodsList, err := models.GetOrderGoodsListByOrderID(orderGoodsInfo.OrderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderGoodsListByOrderID is failed! err:[%v]", err)
	}
	var isOK int
	isOK = 1
	for _, v := range goodsList {
		isOK = isOK & v.Comment
	}
	if isOK == 1 {
		// 更改订单为已经评论
		err := models.UpdataOrderStatus(orderGoods.OrderID, common.STATUS_COMMENT)
		if err != nil {
			logger.Logger.Errorf("UpdataOrderStatus is failed! orderID:[%v] status:[%v] err:[%v].", orderGoods.OrderID, common.STATUS_COMMENT, err)
		}
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// OrderPrepay 订单付款
// @Tags 前台/订单付款
// @Summary 订单付款
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderPrepayBody body request.OrderPrepayBody true "订单请求体"
// @Success 200 {object} response.UnifiedOrderRtnJSON ""
// @router /order/prepay [post]
func (c *OrderController) OrderPrepay() {
	var orderPrepayBody request.OrderPrepayBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderPrepayBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Prepay Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	orderInfo, err := models.GetOrderByID(orderPrepayBody.OrderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
		if err.Error() == common.ErrNoRow.Error() {
			c.RespJSON("订单已取消", common.FCODE, "")
			return
		}
		c.RespJSON("获取订单信息失败", common.FCODE, "")
		return
	}

	if orderInfo.PayID != "" {
		c.RespJSON("订单已支付，请不要重复操作", common.FCODE, "")
		return
	}
	// 根据订单ID获取商品名
	goods, err := models.GetOrderGoodsByOrderID(orderInfo.ID)
	if err != nil {
		logger.Logger.Errorf("GetOrderGoodsByOrderID is failed! err:[%v].", err)
	}

	userInfo, err := models.GetUserByID(c.UserID)
	if err != nil {
		logger.Logger.Errorf("GetUserByID is failed! err:[err].", err)
		c.RespJSON("微信支付失败", common.FCODE, "")
		return
	}
	if userInfo.OpenID == "" {
		logger.Logger.Errorf("userInfo.OpenID is null")
		c.RespJSON("微信支付失败", common.FCODE, "")
		return
	}

	payinfo := common.PayInfo{
		OpenID:         userInfo.OpenID,
		Body:           goods.GoodsName,
		OutTradeNo:     orderInfo.OrderSn,
		TotalFee:       orderInfo.ActualPrice,
		SpbillCreateIp: c.Ctx.Input.IP(),
	}
	// 统一下单
	unifiedOrderRtnJSON, err := common.UnifiedOrder(payinfo)
	if err != nil {
		logger.Logger.Errorf("UnifiedOrder is failed! err:[err].", err)
		c.RespJSON("微信支付失败", common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, unifiedOrderRtnJSON)
}

// OrderPrepay 删除订单
// @Tags 前台/删除订单
// @Summary 删除订单
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.OrderDeleteBody body request.OrderDeleteBody true "订单请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /order/delete [post]
func (c *OrderController) OrderDelete() {
	var orderDeleteBody request.OrderDeleteBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &orderDeleteBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal Order_Prepay Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}
	// 获取订单信息
	orderInfo, err := models.GetOrderByID(orderDeleteBody.OrderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderByID is failed! err:[%v].", err)
		c.RespJSON("订单信息查询失败", common.FCODE, "")
		return
	}
	// 检查订单能否取消
	handleOption, err := models.GetOrderHandleOption(orderInfo.ID)
	if err != nil {
		logger.Logger.Errorf("GetOrderHandleOption is failed! err:[%v].", err)
		c.RespJSON("获取订单信息状态失败", common.FCODE, "")
		return
	}
	if !handleOption.Delete {
		logger.Logger.Errorf("handleOption:[%v]  handleOption.Cancel:[%v].", handleOption, handleOption.Cancel)
		c.RespJSON("订单不能删除", common.FCODE, "")
		return
	}

	err = models.OrderDelete(orderInfo.ID)
	if err != nil {
		logger.Logger.Error("OrderDelete is failed! err:[%v].", err)
		c.RespJSON("删除订单失败", common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// OrderPrepay 物流信息查询
// @Tags 前台/物流信息查询
// @Summary 物流信息查询
// @Produce  application/JSON
// @Param token header string true "token"
// @Param orderId query int true "订单id"
// @Success 200 {object} response.ExpressRtnInfo ""
// @router /express/query [get]
func (c *OrderController) ExpressQuery() {
	orderID, _ := c.GetInt("orderId")

	// 查询订单信息
	orderInfo, err := models.GetOrderByID(orderID)
	if err != nil {
		logger.Logger.Error("GetOrderByID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var expressInfo base.ShopOrderExpress

	expressInfo.OrderID = orderInfo.ID
	expressInfo.ShipperCode = orderInfo.ShipSn
	expressInfo.ShipperName = orderInfo.ShipChannel

	// 查询物流进度
	expressserviceres := common.QueryExpress(orderInfo.ShipChannel, orderInfo.ShipSn, "")
	traces, _ := json.Marshal(expressserviceres.Traces)
	if expressserviceres.Success {
		expressInfo.Traces = string(traces)
		expressInfo.IsFinish = expressserviceres.IsFinish
		expressInfo.RequestTime = time.Now()
	}

	expressProvIDer, err := common.GetOrderChannel()
	if err != nil {
		logger.Logger.Errorf("GetOrderChannel is failed! err:[%v].", err)
	} else {
		for _, v := range expressProvIDer {
			if v.Code == orderInfo.ShipChannel {
				expressserviceres.ShipperName = v.Name
			}
		}
	}

	// 更新数据库订单查询物流信息
	expressInfo, err = models.GetExpressByOrderID(orderID)
	if err != nil {
		if err.Error() == common.ErrNoRow.Error() {
			// 新建物流和订单关系
			expressInfo.AddTime = time.Now()
			expressInfo.RequestTime = time.Now()
			expressInfo.UpdateTime = time.Now()
			err := models.AddOrderExpress(expressInfo)
			if err != nil {
				logger.Logger.Error("AddOrderExpress is failed! err:[%v].", err)
			}
		} else {
			logger.Logger.Error("GetExpressByOrderID is failed! err:[%v].", err)
		}
	} else {
		expressInfo.RequestCount += 1
		expressInfo.RequestTime = time.Now()
		expressInfo.UpdateTime = time.Now()
		// 更新物流进度到订单表
		err = models.UpdataOrderExpress(expressInfo)
		if err != nil {
			logger.Logger.Error("UpdataOrderExpress is failed! err:[%v].", err)
		}
	}

	// 实现物流信息倒序
	length := len(expressserviceres.Traces)
	if length > 1 {
		for i := 0; i < length/2; i++ {
			temp := expressserviceres.Traces[length-1-i]
			expressserviceres.Traces[length-1-i] = expressserviceres.Traces[i]
			expressserviceres.Traces[i] = temp
		}
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, expressserviceres)
}
