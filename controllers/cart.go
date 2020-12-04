// Copyright 2020 The shop Authors

// Package controllers implements controllers.
package controllers

import (
	"encoding/json"
	"fmt"
	"github.com/astaxie/beego/orm"
	"shop/base"
	"shop/base/request"
	"shop/base/response"
	"shop/common"
	"shop/logger"
	"shop/models"
	"strconv"
	"time"
)

// CartController ...
type CartController struct {
	MainController
}

// CartIndex 获取购物车的数据
// @Tags 前台/获取购物车的数据
// @Summary 获取购物车的数据
// @Produce  application/JSON
// @Param token header string true "token"
// @Success 200 {object} response.IndexCartRtnJSON ""
// @router /cart/index [get]
func (c *CartController) CartIndex() {
	cartList, err := models.GetAllCart(c.UserID)
	if err != nil {
		logger.Logger.Errorf("get my cart list info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	cartTotal, err := models.GetCartTotal(c.UserID)
	if err != nil {
		logger.Logger.Errorf("get my cart list total info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var indexCartRtnJSON response.IndexCartRtnJSON
	var cart response.Cart
	var crtList []response.Cart
	crtList = []response.Cart{}
	for _, v := range cartList {
		cart.ID = v.ID
		cart.UserID = v.UserID
		cart.GoodsID = v.GoodsID
		cart.GoodsSn = v.GoodsSn
		cart.GoodsName = v.GoodsName
		cart.ProductID = v.ProductID
		cart.Price = v.Price
		cart.Number = v.Number
		cart.Specifications = v.Specifications
		if v.Checked == 0 {
			cart.Checked = false
		} else {
			cart.Checked = true
		}
		cart.PicURL = v.PicURL
		cart.AddTime = v.AddTime
		cart.UpdateTime = v.UpdateTime
		crtList = append(crtList, cart)
	}
	indexCartRtnJSON.CartList = crtList
	indexCartRtnJSON.CartTotal = cartTotal
	c.RespJSON(common.SUCCESEE, common.SCODE, indexCartRtnJSON)
}

// CartAdd 添加到购物车
// @Tags 前台/添加到购物车
// @Summary 添加到购物车
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CartAddBody body request.CartAddBody true "请求体"
// @Success 200 {object} controllers.ResponseData "数量"
// @router /cart/add [post]
func (c *CartController) CartAdd() {
	var reqCartBody request.CartAddBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqCartBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal add cart Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 获取商品信息
	goodsInfo, err := models.GetOneGoods(reqCartBody.GoodsID)
	if err != nil {
		logger.Logger.Errorf("get one goods info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 获取商品规格信息
	goodsProductInfo, err := models.GetGoodsOneProductByID(reqCartBody.ProductID)
	if err != nil {
		logger.Logger.Errorf("get one goods info product failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 查询数据购物车是否存在该商品信息，如果存在更新数量，如果不存在就新增
	cartInfo, err := models.GetCartIsExist(c.UserID, reqCartBody.GoodsID, reqCartBody.ProductID)
	if err != nil {
		if err.Error() == orm.ErrNoRows.Error() {
			specifications, err := json.Marshal(goodsProductInfo.Specifications)
			if err != nil {
				logger.Logger.Errorf("JSON marshal specifications is failed! goodsProductInfo:[%v] err:[%v]", goodsProductInfo, err)
			}
			cartData := base.ShopCart{
				UserID:         c.UserID,
				GoodsID:        goodsInfo.ID,
				GoodsSn:        goodsInfo.GoodsSn,
				GoodsName:      goodsInfo.Name,
				ProductID:      goodsProductInfo.ID,
				Price:          goodsProductInfo.Price,
				Number:         reqCartBody.Number,
				Specifications: string(specifications),
				PicURL:         goodsInfo.PicURL,
				AddTime:        time.Now(),
			}
			// 添加到购物车
			_, err = models.AddCart(&cartData)
			if err != nil {
				logger.Logger.Errorf("add goods into cart is failed! err:[%v].", err)
				c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
				return
			}

			// 获取购物车的数量
			number, err := models.GetCartGoodsCount(c.UserID)
			if err != nil {
				logger.Logger.Errorf("GetCartGoodsCount is failed! err:[%v].", err)
				c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
				return
			}
			c.RespJSON(common.SUCCESEE, common.SCODE, int(number))
			return
		}

		logger.Logger.Errorf("get cart info is exist is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 判断库存是否足够
	if goodsProductInfo.Number < reqCartBody.Number {
		logger.Logger.Error("the number is not enough! request: number:[%v]  goodsproduct: ID:[%v] number:[%v].", reqCartBody.Number, goodsProductInfo.ID, goodsProductInfo.Number)
		c.RespJSON(common.ErrNoEnough.Error(), common.FCODE, "")
		return
	}

	// 更新购物车商品数量
	_, err = models.UpdateCartNumberByID(cartInfo.ID, reqCartBody.Number)
	if err != nil {
		logger.Logger.Errorf("update cart info is failed! err:[%v].", err)
		c.RespJSON(common.ErrUpdateData.Error(), common.FCODE, "")
		return
	}

	// 获取购物车的数量
	number, err := models.GetCartGoodsCount(c.UserID)
	if err != nil {
		logger.Logger.Errorf("GetCartGoodsCount is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, int(number))
}

// CartUpdate 更新购物车信息
// @Tags 前台/更新购物车信息
// @Summary 更新购物车信息
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CartUpdateBody body request.CartUpdateBody true "请求体"
// @Success 200 {object} controllers.ResponseData ""
// @router /cart/update [post]
func (c *CartController) CartUpdate() {
	var reqCartBody request.CartUpdateBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqCartBody)
	if err != nil {
		logger.Logger.Errorf("input para JSON unmarshal is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 查询购物车的信息
	cartInfo, err := models.GetCartByID(reqCartBody.ID)
	if err != nil {
		logger.Logger.Error("get cart info by ID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 查询商品信息，商品不存在则报错
	_, err = models.GetOneGoods(reqCartBody.GoodsID)
	if err != nil {
		logger.Logger.Error("get goods info by ID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 查询商品规格信息
	goodsProductInfo, err := models.GetGoodsOneProductByID(reqCartBody.ProductID)
	if err != nil {
		logger.Logger.Error("get goods product info by ID is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	// 判断库存是否足够
	if goodsProductInfo.Number < reqCartBody.Number {
		logger.Logger.Error("the number is not enough! request: number:[%v]  goodsproduct: ID:[%v] number:[%v].", reqCartBody.Number, goodsProductInfo.ID, goodsProductInfo.Number)
		c.RespJSON(common.ErrNoEnough.Error(), common.FCODE, "")
		return
	}

	// 更新购物车商品数量
	_, err = models.UpdateCartNumberByID(cartInfo.ID, reqCartBody.Number)
	if err != nil {
		logger.Logger.Errorf("update cart into is failed! err:[%v].", err)
		c.RespJSON(common.ErrUpdateData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, "")
}

// CartDelete 删除购物车信息，支持批量删除（数组）
// @Tags 前台/删除购物车信息
// @Summary 删除购物车信息
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CartDeleteBody body request.CartDeleteBody true "请求体"
// @Success 200 {object} response.IndexCartRtnJSON ""
// @router /cart/delete [post]
func (c *CartController) CartDelete() {
	var reqCartBody request.CartDeleteBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqCartBody)
	if err != nil {
		logger.Logger.Errorf("input para JSON unmarshal is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 通过productID和userID批量删除购物车信息
	_, err = models.DeleteCartByProductID(c.UserID, reqCartBody.ProductIDs)
	if err != nil {
		logger.Logger.Errorf("delete cart info is failed! prodouctID:[%v] err:[%v].", reqCartBody.ProductIDs, err)
		c.RespJSON(common.ErrDeleteData.Error(), common.FCODE, "")
		return
	}

	cartList, err := models.GetAllCart(c.UserID)
	if err != nil {
		logger.Logger.Errorf("get my cart list info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	cartTotal, err := models.GetCartTotal(c.UserID)
	if err != nil {
		logger.Logger.Errorf("get my cart list total info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var indexCartRtnJSON response.IndexCartRtnJSON
	cartListRtn := []response.Cart{}
	var cart response.Cart
	for _, v := range cartList {
		cart.ID = v.ID
		cart.UserID = v.UserID
		cart.GoodsID = v.GoodsID
		cart.GoodsSn = v.GoodsSn
		cart.GoodsName = v.GoodsName
		cart.ProductID = v.ProductID
		cart.Price = v.Price
		cart.Number = v.Number
		cart.Specifications = v.Specifications
		if v.Checked == 0 {
			cart.Checked = false
		} else {
			cart.Checked = true
		}
		cart.PicURL = v.PicURL
		cart.AddTime = v.AddTime
		cart.UpdateTime = v.UpdateTime
		cartListRtn = append(cartListRtn, cart)
	}
	indexCartRtnJSON.CartList = cartListRtn
	indexCartRtnJSON.CartTotal = cartTotal

	c.RespJSON(common.SUCCESEE, common.SCODE, indexCartRtnJSON)
}

// CartChecked 选择或取消选择商品
// @Tags 前台/选择或取消选择商品
// @Summary 选择或取消选择商品
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CartCheckedBody body request.CartCheckedBody true "请求体"
// @Success 200 {object} response.IndexCartRtnJSON ""
// @router /cart/checked [post]
func (c *CartController) CartChecked() {
	var reqCartCheckedBody request.CartCheckedBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqCartCheckedBody)
	if err != nil {
		logger.Logger.Errorf("input para JSON unmarshal is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}
	if len(reqCartCheckedBody.ProductIDs) == 0 {
		logger.Logger.Errorf("iProductIDs is empty! requestCartBody:[%v].", reqCartCheckedBody)
		c.RespJSON(common.ErrSelect.Error(), common.FCODE, "")
		return
	}

	// 更新商品是否选择
	_, err = models.UpdateCartByProductID(c.UserID, reqCartCheckedBody.ProductIDs, reqCartCheckedBody.IsChecked)
	if err != nil {
		logger.Logger.Errorf("delete cart info is failed! prodouctID:[%v] err:[%v].", reqCartCheckedBody.ProductIDs, err)
		c.RespJSON(common.ErrDeleteData.Error(), common.FCODE, "")
		return
	}

	cartList, err := models.GetAllCart(c.UserID)
	if err != nil {
		logger.Logger.Errorf("get my cart list info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	cartTotal, err := models.GetCartTotal(c.UserID)
	if err != nil {
		logger.Logger.Errorf("get my cart list total info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	var indexCartRtnJSON response.IndexCartRtnJSON
	cartListRtn := []response.Cart{}
	var cart response.Cart
	for _, v := range cartList {
		cart.ID = v.ID
		cart.UserID = v.UserID
		cart.GoodsID = v.GoodsID
		cart.GoodsSn = v.GoodsSn
		cart.GoodsName = v.GoodsName
		cart.ProductID = v.ProductID
		cart.Price = v.Price
		cart.Number = v.Number
		cart.Specifications = v.Specifications
		if v.Checked == 0 {
			cart.Checked = false
		} else {
			cart.Checked = true
		}
		cart.PicURL = v.PicURL
		cart.AddTime = v.AddTime
		cart.UpdateTime = v.UpdateTime
		cartListRtn = append(cartListRtn, cart)
	}
	indexCartRtnJSON.CartList = cartListRtn
	indexCartRtnJSON.CartTotal = cartTotal

	c.RespJSON(common.SUCCESEE, common.SCODE, indexCartRtnJSON)
}

// CartGoodsCount 获取用户购物车商品数量
// @Tags 前台/获取用户购物车商品数量
// @Summary 获取用户购物车商品数量
// @Produce  application/JSON
// @Param token header string true "token"
// @Success 200 {object} controllers.ResponseData "数量"
// @router /cart/goodscount [get]
func (c *CartController) CartGoodsCount() {
	num, err := models.GetCartGoodsCount(c.UserID)
	if err != nil {
		logger.Logger.Errorf("delete cart info is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	c.RespJSON(common.SUCCESEE, common.SCODE, num)
}

// CartCheckout 下单前信息确认
// userID    用户ID
// cartID    购物车商品ID： 如果购物车商品ID是空，则下单当前用户所有购物车商品;如果购物车商品ID非空，则只下单当前购物车商品
// addressID 收货地址ID： 如果收货地址ID是空，则查询当前用户的默认地址
// couponID  优惠券ID： 如果优惠券ID是空，则自动选择合适的优惠券
// couponID与userCouponID同时获取到的，一一对应，用哪个都可以

// @Tags 前台/下单前信息确认
// @Summary 下单前信息确认
// @Produce  application/JSON
// @Param token header string true "token"
// @Param cartId query int true "cartId"
// @Param addressId query int true "addressId"
// @Param couponId query int true "couponId"
// @Param userCouponId query int true "userCouponId"
// @Success 200 {object} response.CheckoutRtnJSON ""
// @router /cart/checkout [get]
func (c *CartController) CartCheckout() {
	cartID, _ := c.GetInt("cartId")
	addressID, _ := c.GetInt("addressId")
	couponID, _ := c.GetInt("couponId")
	userCouponID, _ := c.GetInt("userCouponId")

	var address base.ShopAddress
	var err error
	// 获取地址
	if addressID == 0 {
		address, err = models.GetAddressDefault(c.UserID)
		if err != nil && err.Error() != common.ErrNoRow.Error() {
			logger.Logger.Errorf("get default address info is failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}
	} else {
		address, err = models.GetAddressByID(c.UserID, addressID)
		if err != nil {
			logger.Logger.Errorf("get address info by ID is failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}
	}

	var availableCouponLength int
	var freightPrice float64
	var couponPrice float64
	var goodstotalprice float64
	var ordertotalprice float64
	var actualPrice float64

	checkedCartList := []base.ShopCart{}

	if cartID > 0 {
		cartInfo, err := models.GetCartByID(cartID)
		if err != nil {
			logger.Logger.Errorf("gGetCartByID is failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}
		checkedCartList = append(checkedCartList, cartInfo)
		goodstotalprice = cartInfo.Price * float64(cartInfo.Number)

	} else {
		// 获取选定的购物车列表
		checkedCartList, err = models.GetCheckAllCart(c.UserID)
		if err != nil {
			logger.Logger.Errorf("get my cart list info is failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}

		cartTotal, err := models.GetCartTotal(c.UserID)
		if err != nil {
			logger.Logger.Errorf("get my cart list total info is failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}

		goodstotalprice = cartTotal.CheckedGoodsAmount
	}

	// 优惠卷使用情况
	// 1. 自动使用优惠券，则选择合适优惠券
	// 2. 用户选择优惠券，则测试优惠券是否合适
	couponIsOKList := []base.ShopCouponUser{}
	// 获取优惠卷的可用数量
	couponUserList, err := models.GetCouponUserList(c.UserID)
	if err != nil {
		logger.Logger.Errorf("GetCouponUserList is failed! err:[%v].", err)
	}
	if len(couponUserList) > 0 {
		for _, v := range couponUserList {
			isOK := models.CheckCoupon(v, goodstotalprice, checkedCartList)
			if isOK {
				availableCouponLength++
			}
			couponIsOKList = append(couponIsOKList, v)
		}
	}

	// 获取优惠卷的价格, ID大于0 获取正常优惠卷, ID小于0 获取可用优惠卷第一个
	if couponID > 0 {
		coupon, err := models.GetOneCoupon(couponID)
		if err != nil {
			logger.Logger.Errorf("GetOneCoupon is failed! err:[%v].", err)
			c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
			return
		}
		couponPrice = coupon.Discount
	} else {
		if len(couponIsOKList) > 0 {
			userCouponID = couponIsOKList[0].ID
			couponID = couponIsOKList[0].CouponID
			coupon, err := models.GetOneCoupon(couponID)
			if err != nil {
				logger.Logger.Errorf("GetOneCoupon is failed! err:[%v].", err)
				c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
				return
			}
			couponPrice = coupon.Discount
		}
	}

	// 获取邮费价格
	// freightPrice = 0
	var expressFreightMin float64
	var expressFreightValue float64
	systemListRespon, err := models.GetSystem()
	if err != nil {
		logger.Logger.Errorf("GetSystem is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}
	for _, v := range systemListRespon {
		switch v.KeyName {
		case common.SHOP_EXPRESS_FREIGHT_MIN:
			expressFreightMin, _ = strconv.ParseFloat(v.KeyValue, 64)
		case common.SHOP_EXPRESS_FREIGHT_VALUE:
			expressFreightValue, _ = strconv.ParseFloat(v.KeyValue, 64)
		}
	}
	if goodstotalprice < expressFreightMin {
		freightPrice = expressFreightValue
	} else {
		freightPrice = 0
	}

	ordertotalprice = goodstotalprice + freightPrice - couponPrice
	actualPrice = ordertotalprice
	// float保留小数位两位，比如0.99
	ordertotalprice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", ordertotalprice), 64)
	actualPrice, _ = strconv.ParseFloat(fmt.Sprintf("%.2f", actualPrice), 64)

	var checkoutRtnJSON response.CartCheckoutRtnJSON
	checkoutRtnJSON.AddressID = address.ID
	checkoutRtnJSON.Address = address
	checkoutRtnJSON.CartID = cartID
	checkoutRtnJSON.UserCouponID = userCouponID
	checkoutRtnJSON.CouponID = couponID
	checkoutRtnJSON.CheckedGoodsList = checkedCartList
	checkoutRtnJSON.ActualPrice = actualPrice
	checkoutRtnJSON.OrderTotalPrice = ordertotalprice
	checkoutRtnJSON.FreightPrice = freightPrice
	checkoutRtnJSON.GoodsTotalPrice = goodstotalprice
	checkoutRtnJSON.CouponPrice = couponPrice
	checkoutRtnJSON.AvailableCouponLength = availableCouponLength

	c.RespJSON(common.SUCCESEE, common.SCODE, checkoutRtnJSON)
}

// CartFastadd 立即购买
// @Tags 前台/立即购买
// @Summary 立即购买
// @Produce  application/JSON
// @Param token header string true "token"
// @Param request.CartAddBody query request.CartAddBody true "请求体"
// @Success 200 {object} controllers.ResponseData "id"
// @router /cart/fastadd [post]
func (c *CartController) CartFastadd() {
	var reqCartBody request.CartAddBody
	err := json.Unmarshal(c.Ctx.Input.RequestBody, &reqCartBody)
	if err != nil {
		logger.Logger.Errorf("JSON unmarshal add cart Info is failed! err:[%v].", err)
		c.RespJSON(common.ErrStructJSON.Error(), common.FCODE, "")
		return
	}

	// 获取商品信息
	goodsInfo, err := models.GetOneGoods(reqCartBody.GoodsID)
	if err != nil {
		logger.Logger.Errorf("get one goods info failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 获取商品规格信息
	goodsProductInfo, err := models.GetGoodsOneProductByID(reqCartBody.ProductID)
	if err != nil {
		logger.Logger.Errorf("get one goods info product failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 查询数据购物车是否存在该商品信息，如果存在更新数量，如果不存在就新增
	cartInfo, err := models.GetCartIsExist(c.UserID, reqCartBody.GoodsID, reqCartBody.ProductID)
	if err != nil {
		if err.Error() == common.ErrNoRow.Error() {
			specifications, err := json.Marshal(goodsProductInfo.Specifications)
			if err != nil {
				logger.Logger.Errorf("JSON marshal specifications is failed! goodsProductInfo:[%v] err:[%v]", goodsProductInfo, err)
			}
			cartData := base.ShopCart{
				UserID:         c.UserID,
				GoodsID:        goodsInfo.ID,
				GoodsSn:        goodsInfo.GoodsSn,
				GoodsName:      goodsInfo.Name,
				ProductID:      goodsProductInfo.ID,
				Price:          goodsProductInfo.Price,
				Number:         reqCartBody.Number,
				Specifications: string(specifications),
				PicURL:         goodsInfo.PicURL,
				AddTime:        time.Now(),
			}
			// 添加到购物车
			_, err = models.AddCart(&cartData)
			if err != nil {
				logger.Logger.Errorf("add goods into cart is failed! err:[%v].", err)
				c.RespJSON(common.ErrAddData.Error(), common.FCODE, "")
				return
			}
			c.RespJSON(common.SUCCESEE, common.SCODE, cartData.ID)
			return
		}

		logger.Logger.Errorf("get cart info is exist is failed! err:[%v].", err)
		c.RespJSON(common.ErrGetData.Error(), common.FCODE, "")
		return
	}

	// 判断库存是否足够
	if goodsProductInfo.Number < reqCartBody.Number {
		logger.Logger.Error("the number is not enough! request: number:[%v]  goodsproduct: ID:[%v] number:[%v].", reqCartBody.Number, goodsProductInfo.ID, goodsProductInfo.Number)
		c.RespJSON(common.ErrNoEnough.Error(), common.FCODE, "")
		return
	}

	// 更新购物车商品数量
	_, err = models.UpdateCartNumberByID(cartInfo.ID, reqCartBody.Number)
	if err != nil {
		logger.Logger.Errorf("update cart info is failed! err:[%v].", err)
		c.RespJSON(common.ErrUpdateData.Error(), common.FCODE, "")
		return
	}

	c.RespJSON(common.SUCCESEE, common.SCODE, cartInfo.ID)
}
