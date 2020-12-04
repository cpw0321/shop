// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"errors"
	"fmt"
	"math/rand"
	"shop/base"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
	"strconv"
	"strings"
	"time"

	"github.com/astaxie/beego/orm"
)

func init() {
	orm.RegisterModel(new(base.ShopOrder))
}

func OrderDetail(userID int, orderID int) (orderDetailInfo response.OrderDetailRtnJSON, err error) {
	o := orm.NewOrm()
	var order base.ShopOrder
	orderTable := new(base.ShopOrder)
	err = o.QueryTable(orderTable).Filter("ID", orderID).Filter("UserID", userID).One(&order)
	if err != nil {
		return orderDetailInfo, err
	}

	var orderInfo response.OrderInfo
	orderInfo.ID = order.ID
	orderInfo.UserID = order.UserID
	orderInfo.OrderSn = order.OrderSn
	orderInfo.OrderStatus = order.OrderStatus
	orderInfo.AftersaleStatus = order.AftersaleStatus
	orderInfo.Consignee = order.Consignee
	orderInfo.Mobile = order.Mobile
	orderInfo.Address = order.Address
	orderInfo.Message = order.Message
	orderInfo.GoodsPrice = order.GoodsPrice
	orderInfo.FreightPrice = order.FreightPrice
	orderInfo.CouponPrice = order.CouponPrice
	orderInfo.IntegralPrice = order.IntegralPrice
	orderInfo.GrouponPrice = order.GrouponPrice
	orderInfo.OrderPrice = order.OrderPrice
	orderInfo.ActualPrice = order.ActualPrice
	orderInfo.PayID = order.PayID
	orderInfo.PayTime = utils.FormatTimestampStr(order.PayTime)
	orderInfo.ShipSn = order.ShipSn
	orderInfo.ShipChannel = order.ShipChannel
	orderInfo.ShipTime = utils.FormatTimestampStr(order.ShipTime)
	orderInfo.RefundAmount = order.RefundAmount
	orderInfo.RefundType = order.RefundType
	orderInfo.RefundContent = order.RefundContent
	orderInfo.RefundTime = utils.FormatTimestampStr(order.RefundTime)
	orderInfo.ConfirmTime = utils.FormatTimestampStr(order.ConfirmTime)
	orderInfo.Comments = order.Comments
	orderInfo.EndTime = utils.FormatTimestampStr(order.EndTime)
	orderInfo.AddTime = utils.FormatTimestampStr(order.AddTime)
	orderInfo.UpdateTime = utils.FormatTimestampStr(order.UpdateTime)
	orderInfo.OrderStatusText = GetorderStatusText(order.OrderStatus)

	orderDetailInfo.OrderInfo = orderInfo
	handleOption, err := GetOrderHandleOption(orderID)
	if err != nil {
		logger.Logger.Errorf("GetOrderHandleOption is failed! err:[%v].", err)
	}
	orderDetailInfo.OrderInfo.HandleOption = handleOption

	orderGoodsTable := new(base.ShopOrderGoods)
	_, err = o.QueryTable(orderGoodsTable).Filter("OrderID", orderDetailInfo.OrderInfo.ID).All(&orderDetailInfo.OrderGoods)
	if err != nil {
		return orderDetailInfo, err
	}

	//var orderExpress ShopOrderExpress
	//orderExpressTable := new(ShopOrderExpress)
	//err = o.QueryTable(orderExpressTable).Filter("ShipSn", orderDetailInfo.OrderInfo.ShipSn).One(&orderExpress)
	//if err != nil {
	//	return orderDetailInfo, err
	//}

	return orderDetailInfo, nil
}

func GenerateOrderNumber() string {

	year := time.Now().Year()     //年
	month := time.Now().Month()   //月
	day := time.Now().Day()       //日
	hour := time.Now().Hour()     //小时
	minute := time.Now().Minute() //分钟
	second := time.Now().Second() //秒

	stryear := strconv.Itoa(year)        //年
	strmonth := strconv.Itoa(int(month)) //月
	strday := strconv.Itoa(day)          //日
	strhour := strconv.Itoa(hour)        //小时
	strminute := strconv.Itoa(minute)    //分钟
	strsecond := strconv.Itoa(second)    //秒

	strmonth2 := fmt.Sprintf("%02s", strmonth)
	strday2 := fmt.Sprintf("%02s", strday)
	strhour2 := fmt.Sprintf("%02s", strhour)
	strminute2 := fmt.Sprintf("%02s", strminute)
	strsecond2 := fmt.Sprintf("%02s", strsecond)

	randnum := rand.Intn(999999-100000) + 100000
	strrandnum := strconv.Itoa(randnum)

	return stryear + strmonth2 + strday2 + strhour2 + strminute2 + strsecond2 + strrandnum
}

func OrderDelete(orderID int) (err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	_, err = o.QueryTable(orderTable).Filter("ID", orderID).Update(orm.Params{
		"Deleted": 1,
	})
	return err
}

// OrderSubmit 提交订单
func OrderSubmit(userID, cartID, addressID, couponID int, message string) (orderSubmitRtnJSON response.OrderSubmitRtnJSON, err error) {
	o := orm.NewOrm()
	err = o.Begin()
	if err != nil {
		logger.Logger.Errorf("orm begin is failed! err:[%v].", err)
		return orderSubmitRtnJSON, err
	}
	addressTable := new(base.ShopAddress)
	var address base.ShopAddress
	err = o.QueryTable(addressTable).Filter("ID", addressID).Filter("Deleted", 0).One(&address)
	if err != nil {
		return orderSubmitRtnJSON, err
	}

	cartTable := new(base.ShopCart)
	var cart base.ShopCart
	var carts []base.ShopCart
	if cartID == 0 {
		_, err = o.QueryTable(cartTable).Filter("UserID", userID).Filter("checked", 1).Filter("Deleted", 0).All(&carts)
		if err != nil && err == orm.ErrNoRows {
			err = errors.New("请选择商品")
			return orderSubmitRtnJSON, err
		}
	} else {
		err = o.QueryTable(cartTable).Filter("ID", cartID).Filter("Deleted", 0).One(&cart)
		if err != nil {
			logger.Logger.Errorf("query cart by cartID is failed! err:[%v].", err)
			return orderSubmitRtnJSON, err
		}
		carts = append(carts, cart)
	}

	couponTable := new(base.ShopCoupon)
	var coupon base.ShopCoupon
	err = o.QueryTable(couponTable).Filter("ID", couponID).Filter("Deleted", 0).One(&coupon)
	if err != nil && err != orm.ErrNoRows {
		return orderSubmitRtnJSON, err
	}

	var goodsTotalPrice float64 = 0

	for _, val := range carts {
		goodsTotalPrice += float64(val.Number) * val.Price
	}

	var freightPrice float64 = 0
	var expressFreightMin float64
	var expressFreightValue float64
	systemListRespon, err := GetSystem()
	if err != nil {
		logger.Logger.Errorf("GetSystem is failed! err:[%v].", err)
	} else {
		for _, v := range systemListRespon {
			switch v.KeyName {
			case "Shop_express_freight_min":
				expressFreightMin, _ = strconv.ParseFloat(v.KeyValue, 64)
			case "Shop_express_freight_value":
				expressFreightValue, _ = strconv.ParseFloat(v.KeyValue, 64)
			}
		}
	}

	if goodsTotalPrice < expressFreightMin {
		freightPrice = expressFreightValue
	} else {
		freightPrice = 0
	}

	var couponPrice float64
	couponPrice = coupon.Discount
	orderTotalPrice := goodsTotalPrice + freightPrice - couponPrice
	actualPrice := orderTotalPrice - 0

	orderInfo := base.ShopOrder{
		UserID:       userID,
		OrderSn:      GenerateOrderNumber(),
		OrderStatus:  101,
		Consignee:    address.Name,
		Mobile:       address.Tel,
		Address:      address.Province + address.City + address.County + address.AddressDetail,
		Message:      message,
		GoodsPrice:   goodsTotalPrice,
		FreightPrice: freightPrice,
		CouponPrice:  couponPrice,
		OrderPrice:   orderTotalPrice,
		ActualPrice:  actualPrice,
		AddTime:      time.Now(),
	}

	orderID, err := o.Insert(&orderInfo)
	if err != nil {
		o.Rollback()
		err = errors.New("订单提交失败")
		return orderSubmitRtnJSON, err
	}

	for _, item := range carts {
		ordergood := base.ShopOrderGoods{
			OrderID:        int(orderID),
			GoodsID:        item.GoodsID,
			GoodsName:      item.GoodsName,
			GoodsSn:        item.GoodsSn,
			ProductID:      item.ProductID,
			Number:         item.Number,
			Price:          item.Price,
			Specifications: item.Specifications,
			PicURL:         item.PicURL,
			AddTime:        time.Now(),
		}
		_, err = o.Insert(&ordergood)
		if err != nil {
			o.Rollback()
			return orderSubmitRtnJSON, err
		}
	}

	o.Commit()

	// 删除购物车中信息
	for _, v1 := range carts {
		// 删除购物车里的信息
		err = DeleteCartByID(v1.ID)
		if err != nil {
			logger.Logger.Error("DeleteCartByID is failed! err:[%v].", err)
		}
	}

	orderSubmitRtnJSON.OrderID = int(orderID)

	return orderSubmitRtnJSON, nil
}

// GetAllOrder 获取订单列表，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllOrder(query map[string]string, sortby string, order string,
	page int, limit int) (ordersRtnJSON []response.OrderRtnJSON, err error) {
	var orderInfo response.OrderRtnJSON
	var orderList []base.ShopOrder
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopOrder))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
	}

	// order
	if sortby != "" {
		orderby := ""
		if order == "desc" {
			orderby = "-" + sortby
		} else {
			orderby = sortby
		}
		qs = qs.OrderBy(orderby)
	}

	_, err = qs.Limit(limit, offset).All(&orderList)
	if err != nil {
		return ordersRtnJSON, err
	}

	for _, v := range orderList {
		goodsList, err := GetOrderGoodsListByOrderID(v.ID)
		if err != nil {
			logger.Logger.Errorf("GetOrderGoodsListByOrderID is failed! err:[%v].", err)
		}
		orderInfo.ID = v.ID
		orderInfo.GoodsList = goodsList
		orderInfo.OrderSn = v.OrderSn
		orderInfo.ActualPrice = v.ActualPrice
		orderInfo.OrderStatusText = GetorderStatusText(v.OrderStatus)
		ordersRtnJSON = append(ordersRtnJSON, orderInfo)
	}

	return ordersRtnJSON, nil
}

// UpdataOrderStatus 修改订单状态
func UpdataOrderStatus(ID int, status int) (err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	_, err = o.QueryTable(orderTable).Filter("ID", ID).Update(orm.Params{
		"OrderStatus": status,
	})
	if err != nil {
		return err
	}

	return nil
}

// UpdataOrderStatusBySn 修改订单状态
func UpdataOrderStatusBySn(orderSn string, status int) (err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	_, err = o.QueryTable(orderTable).Filter("OrderSn", orderSn).Update(orm.Params{
		"OrderStatus": status,
	})
	if err != nil {
		return err
	}

	return nil
}

// GetOrderStatusCount 获取订单状态的数量
func GetOrderStatusCount(UserID int) (orderRtnJSON response.UserIndexRtnJSON) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	aftersaleTable := new(base.ShopAftersale)
	unpaid, _ := o.QueryTable(orderTable).Filter("UserID", UserID).Filter("OrderStatus", 101).Count()
	unship, _ := o.QueryTable(orderTable).Filter("UserID", UserID).Filter("OrderStatus", 201).Count()
	unrecv, _ := o.QueryTable(orderTable).Filter("UserID", UserID).Filter("OrderStatus", 301).Count()
	uncomment, _ := o.QueryTable(orderTable).Filter("UserID", UserID).Filter("OrderStatus__in", 401, 402).Count()
	unaftersale, _ := o.QueryTable(aftersaleTable).Filter("UserID", UserID).Filter("Status__in", 1, 2).Count()
	orderRtnJSON.Order.Unpaid = unpaid
	orderRtnJSON.Order.Unship = unship
	orderRtnJSON.Order.Unrecv = unrecv
	orderRtnJSON.Order.Uncomment = uncomment
	orderRtnJSON.Order.Unaftersale = unaftersale

	return orderRtnJSON
}

func GetOrderHandleOption(orderID int) (handleOption response.OrderHandleOption, err error) {

	// 订单流程：下单成功－》支付订单－》发货－》收货
	// 订单状态：
	// 101 订单生成，未支付；102，下单未支付用户取消；103，下单未支付超期系统自动取消
	// 201 支付完成，商家未发货；202，订单生产，已付款未发货，用户申请退款；203，管理员执行退款操作，确认退款成功；
	// 301 商家发货，用户未确认；
	// 401 用户确认收货，订单结束； 402 用户没有确认收货，但是快递反馈已收货后，超过一定时间，系统自动确认收货，订单结束。

	handleOption = response.OrderHandleOption{false, false, false, false, false, false, false, false, false}

	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	var order base.ShopOrder
	err = o.QueryTable(orderTable).Filter("ID", orderID).Filter("Deleted", 0).One(&order)
	if err != nil {
		logger.Logger.Errorf("QueryTable orderTable is failed! err:[%v].", err)
		return handleOption, err
	}

	switch order.OrderStatus {
	case 101:
		// 如果订单没有被取消，且没有支付，则可支付，可取消
		handleOption.Cancel = true
		handleOption.Pay = true
	case 102, 103:
		// 如果订单已经取消或是已完成，则可删除
		handleOption.Delete = true
	case 201:
		// 如果订单已付款，没有发货，则可退款
		handleOption.Refund = true
	case 202, 204:
		// 如果订单申请退款中，没有相关操作
	case 203:
		// 如果订单已经退款，则可删除
		handleOption.Delete = true
	case 301:
		// 如果订单已经发货，没有收货，则可收货操作,
		// 此时不能取消订单
		handleOption.Confirm = true
		handleOption.Ship = true
	case 401, 402:
		// 如果订单已经支付，且已经收货，则可删除、去评论、申请售后和再次购买
		handleOption.Delete = true
		handleOption.Comment = true
		handleOption.Rebuy = true
		handleOption.Aftersale = true
		handleOption.Ship = true
	}

	return handleOption, nil
}

// GetOrderByID 获取订单信息通过订单ID
func GetOrderByID(orderID int) (orderInfo response.OrderInfo, err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	var order base.ShopOrder
	err = o.QueryTable(orderTable).Filter("ID", orderID).Filter("Deleted", 0).One(&order)
	if err != nil {
		logger.Logger.Errorf("query order by orderID is failed! err:[%v].", err)
		return orderInfo, err
	}
	orderInfo.ID = order.ID
	orderInfo.UserID = order.UserID
	orderInfo.OrderSn = order.OrderSn
	orderInfo.OrderStatus = order.OrderStatus
	orderInfo.Consignee = order.Consignee
	orderInfo.Mobile = order.Mobile
	orderInfo.Address = order.Address
	orderInfo.Message = order.Message
	orderInfo.GoodsPrice = order.GoodsPrice
	orderInfo.FreightPrice = order.FreightPrice
	orderInfo.CouponPrice = order.CouponPrice
	orderInfo.IntegralPrice = order.IntegralPrice
	orderInfo.GrouponPrice = order.GrouponPrice
	orderInfo.OrderPrice = order.OrderPrice
	orderInfo.ActualPrice = order.ActualPrice
	orderInfo.PayID = order.PayID
	orderInfo.PayTime = utils.FormatTimestampStr(order.PayTime)
	orderInfo.ShipSn = order.ShipSn
	orderInfo.ShipChannel = order.ShipChannel
	orderInfo.ShipTime = utils.FormatTimestampStr(order.ShipTime)
	orderInfo.RefundAmount = order.RefundAmount
	orderInfo.RefundType = order.RefundType
	orderInfo.RefundContent = order.RefundContent
	orderInfo.RefundTime = utils.FormatTimestampStr(order.RefundTime)
	orderInfo.ConfirmTime = utils.FormatTimestampStr(order.ConfirmTime)
	orderInfo.Comments = order.Comments
	orderInfo.EndTime = utils.FormatTimestampStr(order.EndTime)
	orderInfo.AddTime = utils.FormatTimestampStr(order.AddTime)
	orderInfo.UpdateTime = utils.FormatTimestampStr(order.UpdateTime)
	orderInfo.OrderStatusText = GetorderStatusText(order.OrderStatus)
	handleOption, err := GetOrderHandleOption(order.ID)
	if err != nil {
		logger.Logger.Errorf("GetOrderHandleOption is failed! err:[%order].", err)
	}
	orderInfo.HandleOption = handleOption

	return orderInfo, nil
}

// GetorderStatusText 获取订单状态
func GetorderStatusText(status int) (orderStatusText string) {

	//当101用户未付款时，此时用户可以进行的操作是取消或者付款
	//当201支付完成而商家未发货时，此时用户可以退款
	//当301商家已发货时，此时用户可以有确认收货
	//当401用户确认收货以后，此时用户可以进行的操作是退货、删除、去评价或者再次购买
	//当402系统自动确认收货以后，此时用户可以删除、去评价、或者再次购买

	switch status {
	case 101:
		return "未付款"
	case 102:
		return "已取消"
	case 103:
		return "已取消(系统)"
	case 201:
		return "已付款"
	case 202:
		return "订单取消，退款中"
	case 203:
		return "已退款"
	case 301:
		return "已发货"
	case 401:
		return "已收货"
	case 402:
		return "已收货(系统)"
	case 501:
		return "已完成"
	}
	return "orderStatus不支持"
}

// GetOrderListByStatus 根据订单状态获取订单信息
func GetOrderListByStatus(orderStatus int) (orderList []base.ShopOrder, err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	_, err = o.QueryTable(orderTable).Filter("OrderStatus", orderStatus).Filter("Deleted", 0).All(&orderList)
	if err != nil {
		logger.Logger.Errorf("query order by orderID is failed! err:[%v].", err)
		return orderList, err
	}

	return orderList, nil
}

// UpdateOrder 更新订单信息
func UpdateOrder(order base.ShopOrder) (err error) {
	o := orm.NewOrm()
	_, err = o.Update(&order)
	if err != nil {
		return err
	}
	return nil
}

// GetOrderBySn 根据订单号获取订单信息
func GetOrderBySn(orderSn string) (order base.ShopOrder, err error) {
	o := orm.NewOrm()
	orderTable := new(base.ShopOrder)
	err = o.QueryTable(orderTable).Filter("OrderSn", orderSn).One(&order)
	if err != nil {
		return order, err
	}

	return order, nil
}
