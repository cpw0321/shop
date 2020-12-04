// Copyright 2020 The shop Authors

// Package models implements models.
package models

import (
	"encoding/json"
	"errors"
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/response"
	"shop/logger"
	"shop/utils"
	"strconv"
	"strings"
	"time"
)

func init() {
	orm.RegisterModel(new(base.ShopCouponUser))
}

// GetCouponUserList 获取用户可以使用的优惠卷
func GetCouponUserList(userID int) (couponUserList []base.ShopCouponUser, err error) {
	o := orm.NewOrm()
	couponUserTable := new(base.ShopCouponUser)
	_, err = o.QueryTable(couponUserTable).Filter("UserID", userID).Filter("Status", 0).All(&couponUserList)
	if err != nil {
		return couponUserList, err
	}

	return couponUserList, nil
}

// GetAllCouponUser 获取分类信息，返回数组
// page: 请求页码,和通常计算机概念中数组下标从0开始不同，这里的page参数应该从1开始，1即代表第一页数据;
// limit: 每一页数量, 分页大小
// sortby: 排序字段, 例如"add_time"或者"ID";
// order: 升序降序, 只能是"desc"或者"asc"。
func GetAllCouponUser(query map[string]string, sortby string, order string,
	page int, limit int) (couponUserRtnJSONList []response.CouponUserRtnJSON, err error) {
	var couponUserList []base.ShopCouponUser
	if page <= 1 {
		page = 1
	}
	offset := (page - 1) * limit
	o := orm.NewOrm()
	qs := o.QueryTable(new(base.ShopCouponUser))
	// query
	for k, v := range query {
		k = strings.Replace(k, ".", "__", -1)
		if strings.Contains(k, "isnull") {
			qs = qs.Filter(k, (v == "true" || v == "1"))
		}
		qs = qs.Filter(k, v)
	}

	// order
	orderby := ""
	if order == "desc" {
		orderby = "-" + sortby
	} else {
		orderby = sortby
	}

	qs = qs.OrderBy(orderby)
	_, err = qs.Limit(limit, offset).All(&couponUserList)
	if err != nil {
		return couponUserRtnJSONList, err
	}

	var couponUserRtnJSON response.CouponUserRtnJSON
	for _, v := range couponUserList {
		couponUserRtnJSON.ID = v.ID
		couponUserRtnJSON.CID = v.CouponID

		couponInfo, err := GetOneCoupon(v.CouponID)
		if err != nil {
			logger.Logger.Errorf("GetOneCoupon is failed! err:[%v].", err)
			continue
		}
		couponUserRtnJSON.Name = couponInfo.Name
		couponUserRtnJSON.Desc = couponInfo.Desc
		couponUserRtnJSON.Tag = couponInfo.Tag
		couponUserRtnJSON.Discount = couponInfo.Discount
		couponUserRtnJSON.Min = couponInfo.Min
		couponUserRtnJSON.Type = couponInfo.Type
		couponUserRtnJSON.StartTime = utils.FormatTimestampStr(v.StartTime)
		couponUserRtnJSON.EndTime = utils.FormatTimestampStr(v.EndTime)
		couponUserRtnJSONList = append(couponUserRtnJSONList, couponUserRtnJSON)
	}
	return couponUserRtnJSONList, nil
}

// UpdateCouponUserStatus 更新优惠卷的使用状态， 0未使用；1已使用；2已过期；3已经下架
func UpdateCouponUserStatus(userID int, couponUserID int, status int) (err error) {
	o := orm.NewOrm()
	couponUserTable := new(base.ShopCouponUser)
	var couponUser base.ShopCouponUser
	err = o.QueryTable(couponUserTable).Filter("ID", couponUserID).One(&couponUser)
	if err != nil {
		return err
	}

	// 判断用户使用的优惠卷是否为自己名下
	if couponUser.UserID != userID {
		logger.Logger.Errorf("couponUser userID is no equal intput userid: %v, %v", couponUser.UserID, userID)
		return errors.New("internal service err")
	}

	couponUser.Status = status
	_, err = o.Update(&couponUser)
	if err != nil {
		return err
	}

	return nil
}

// CheckCoupon 检查优惠卷是否可用
func CheckCoupon(couponUser base.ShopCouponUser, checkedGoodsPrice float64, checkedCartList []base.ShopCart) (isOk bool) {
	couponInfo, err := GetOneCoupon(couponUser.CouponID)
	if err != nil {
		logger.Logger.Errorf("GetOneCoupon is failed! err:[%v].", err)
		return false
	}

	// 判断是否超期
	if time.Now().Unix() > couponUser.EndTime.Unix() {
		return false
	}

	// 检测是否满足最低消费
	if couponInfo.Min > checkedGoodsPrice {
		return false
	}

	// 检测商品是否符合可使用优惠券的商品或分类
	// GoodsType 商品限制类型，如果0则全商品，如果是1则是类目限制，如果是2则是商品限制
	// GoodsValue 商品限制值，goods_type如果是0则空集合，如果是1则是类目集合，如果是2则是商品集合
	var goodsValue []string
	if couponInfo.GoodsType != 0 {
		// 获取商品类目和商品
		for _, v := range checkedCartList {
			err = json.Unmarshal([]byte(couponInfo.GoodsValue), &goodsValue)
			if err != nil {
				logger.Logger.Errorf("get goodsProduct is failed! couponInfo.GoodsValue:[%v].", couponInfo.GoodsValue)
				continue
			}

			// 获取商品信息
			goodsInfo, err := GetOneGoods(v.GoodsID)
			if err != nil {
				logger.Logger.Errorf("get goodsProduct is failed! couponInfo.GoodsValue:[%v].", couponInfo.GoodsValue)
				continue
			}

			// 判断是否符合分类
			if couponInfo.GoodsType == 1 {
				for _, v := range goodsValue {
					categoryID, _ := strconv.Atoi(v)
					if categoryID == goodsInfo.CategoryID {
						return true
					}
				}
				return false
			}

			// 判断是否符合商品
			if couponInfo.GoodsType == 2 {
				for _, v := range goodsValue {
					categoryID, _ := strconv.Atoi(v)
					if categoryID == goodsInfo.ID {
						return true
					}
				}
				return false
			}
		}

	}

	return true
}

// GetSelectCouponUserList 获取可选的优惠卷的列表
func GetSelectCouponUserList(userID int) (couponUserRtnJSONList []response.CouponUserRtnJSON, err error) {
	var couponUserList []base.ShopCouponUser
	couponUserRtnJSONList = []response.CouponUserRtnJSON{}
	o := orm.NewOrm()
	couponUserTable := new(base.ShopCouponUser)
	_, err = o.QueryTable(couponUserTable).Filter("UserID", userID).Filter("Status", 0).Filter("Deleted", 0).All(&couponUserList)
	if err != nil {
		logger.Logger.Errorf("get user coupon is failed!, err:[%v].", err)
		return []response.CouponUserRtnJSON{}, err
	}

	var couponUserRtnJSON response.CouponUserRtnJSON
	for _, v := range couponUserList {
		couponUserRtnJSON.ID = v.ID
		couponUserRtnJSON.CID = v.CouponID

		couponInfo, err := GetOneCoupon(v.CouponID)
		if err != nil {
			logger.Logger.Errorf("GetOneCoupon is failed! err:[%v].", err)
			continue
		}
		couponUserRtnJSON.Name = couponInfo.Name
		couponUserRtnJSON.Desc = couponInfo.Desc
		couponUserRtnJSON.Tag = couponInfo.Tag
		couponUserRtnJSON.Discount = couponInfo.Discount
		couponUserRtnJSON.Min = couponInfo.Min
		couponUserRtnJSON.Type = couponInfo.Type
		couponUserRtnJSON.StartTime = utils.FormatTimestampStr(v.StartTime)
		couponUserRtnJSON.EndTime = utils.FormatTimestampStr(v.EndTime)

		cartTotal, err := GetCartTotal(userID)
		if err != nil {
			logger.Logger.Errorf("GetCartTotal is failed! err:[%v].", err)
			continue
		}
		checkedCartList, err := GetCheckAllCart(userID)
		if err != nil {
			logger.Logger.Errorf("GetCheckAllCart is failed! err:[%v].", err)
			continue
		}
		isOK := CheckCoupon(v, cartTotal.CheckedGoodsAmount, checkedCartList)
		if isOK {
			couponUserRtnJSON.Available = true
		} else {
			couponUserRtnJSON.Available = false
		}

		couponUserRtnJSONList = append(couponUserRtnJSONList, couponUserRtnJSON)
	}
	return couponUserRtnJSONList, nil
}

// GetCouponUserIsExist 获取用户优惠卷是否已经领取过
func GetCouponUserIsExist(userID int, couponID int) (isOK bool) {
	o := orm.NewOrm()
	couponUserTable := new(base.ShopCouponUser)
	isOK = o.QueryTable(couponUserTable).Filter("UserID", userID).Filter("CouponID", couponID).Exist()
	return isOK
}
