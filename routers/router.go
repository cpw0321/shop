// Copyright 2020 The shop Authors

// Package routers implements routers.
// @APIVersion 1.0.0
// @Title shop API
// @Description swagger documents for your API
// @Contact shop
// @TermsOfServiceUrl
// @License Apache 2.0
// @LicenseUrl http://www.apache.org/licenses/LICENSE-2.0.html
package routers

import (
	"github.com/astaxie/beego"
	"shop/controllers"
)

func init() {
	beego.Router("wx/auth/loginbyweixin", &controllers.AuthController{}, "post:AuthLoginByWeixin") // 微信登录
	// beego.Router("wx/auth/login", &controllers.AuthController{}, "post:AuthLogin")                   // 账号登录
	// beego.Router("wx/auth/register", &controllers.AuthController{}, "post:AuthRegister")             // 账号注册
	// beego.Router("wx/auth/reset", &controllers.AuthController{}, "post:AuthReset")                   // 账号密码重置
	// beego.Router("wx/auth/regCaptcha", &controllers.AuthController{}, "post:AuthSendCode")         // 验证码

	beego.Router("wx/home/index", &controllers.HomeController{}, "get:HomeIndex") // 首页数据接口
	beego.Router("wx/home/about", &controllers.HomeController{}, "get:HomeAbout") // 关于我们

	beego.Router("wx/catalog/index", &controllers.CatalogController{}, "get:CatalogIndex")     // 获取分类目录全部分类数据
	beego.Router("wx/catalog/current", &controllers.CatalogController{}, "get:CatalogCurrent") // 获取分类目录当前分类数据

	beego.Router("wx/goods/count", &controllers.GoodsController{}, "get:GoodsCount")       // 统计商品总数
	beego.Router("wx/goods/list", &controllers.GoodsController{}, "get:GoodsList")         // 获得商品列表
	beego.Router("wx/goods/category", &controllers.GoodsController{}, "get:GoodsCategory") // 获得商品分类数据
	beego.Router("wx/goods/detail", &controllers.GoodsController{}, "get:GoodsDetail")     // 获得商品的详情
	// beego.Router("wx/goods/related", &controllers.GoodsController{}, "get:Goods_Related")   //大家都在看

	// 购物车
	beego.Router("wx/cart/index", &controllers.CartController{}, "get:CartIndex")           // 获取购物车的数据
	beego.Router("wx/cart/add", &controllers.CartController{}, "post:CartAdd")              // 添加商品到购物车
	beego.Router("wx/cart/fastadd", &controllers.CartController{}, "post:CartFastadd")      // 立即购买商品
	beego.Router("wx/cart/update", &controllers.CartController{}, "post:CartUpdate")        // 更新购物车的商品
	beego.Router("wx/cart/delete", &controllers.CartController{}, "post:CartDelete")        // 删除购物车的商品
	beego.Router("wx/cart/checked", &controllers.CartController{}, "post:CartChecked")      // 选择或取消选择商品
	beego.Router("wx/cart/goodscount", &controllers.CartController{}, "get:CartGoodsCount") // 获取购物车商品件数
	beego.Router("wx/cart/checkout", &controllers.CartController{}, "get:CartCheckout")     // 下单前信息确认

	// 收藏
	beego.Router("wx/collect/list", &controllers.CollectController{}, "get:CollectList")                // 收藏列表
	beego.Router("wx/collect/addordelete", &controllers.CollectController{}, "post:CollectAddOrDelete") // 添加或取消收藏

	// 评论
	beego.Router("wx/comment/list", &controllers.BaseController{}, "get:CommentList")     // 评论列表
	beego.Router("wx/comment/count", &controllers.BaseController{}, "get:CommentCount")   // 评论总数
	beego.Router("wx/comment/post", &controllers.CommentController{}, "post:CommentPost") // 发表评论

	// 专题
	beego.Router("wx/topic/list", &controllers.TopicController{}, "get:TopicList")       // 专题列表
	beego.Router("wx/topic/detail", &controllers.TopicController{}, "get:TopicDetail")   // 专题详情
	beego.Router("wx/topic/related", &controllers.TopicController{}, "get:TopicRelated") // 相关专题

	// 搜索
	beego.Router("wx/search/index", &controllers.SearchController{}, "get:SearchIndex")                // 搜索关键字
	beego.Router("wx/search/result", &controllers.SearchController{}, "get:SearchResult")              // 搜索结果
	beego.Router("wx/search/helper", &controllers.SearchController{}, "get:SearchHelper")              // 搜索帮助
	beego.Router("wx/search/clearhistory", &controllers.SearchController{}, "post:SearchClearHistory") // 搜索历史清理

	// 收获地址
	beego.Router("wx/address/list", &controllers.AddressController{}, "get:AddressList")
	beego.Router("wx/address/detail", &controllers.AddressController{}, "get:AddressDetail")
	beego.Router("wx/address/save", &controllers.AddressController{}, "post:AddressSave")
	beego.Router("wx/address/delete", &controllers.AddressController{}, "post:AddressDelete")

	// 物流查询
	beego.Router("wx/express/query", &controllers.OrderController{}, "get:ExpressQuery")

	// 区域列表
	beego.Router("wx/region/list", &controllers.RegionController{}, "get:RegionList")

	// 订单
	beego.Router("wx/order/submit", &controllers.OrderController{}, "post:OrderSubmit")   // 提交订单
	beego.Router("wx/order/prepay", &controllers.OrderController{}, "post:OrderPrepay")   // 订单的预支付会话
	beego.Router("wx/order/list", &controllers.OrderController{}, "get:OrderList")        // 订单列表
	beego.Router("wx/order/detail", &controllers.OrderController{}, "get:OrderDetail")    // 订单详情
	beego.Router("wx/order/cancel", &controllers.OrderController{}, "post:OrderCancel")   // 取消订单
	beego.Router("wx/order/refund", &controllers.OrderController{}, "post:OrderRefund")   // 退款取消订单
	beego.Router("wx/order/delete", &controllers.OrderController{}, "post:OrderDelete")   // 删除订单
	beego.Router("wx/order/confirm", &controllers.OrderController{}, "post:OrderConfirm") // 确认收货
	beego.Router("wx/order/goods", &controllers.OrderController{}, "get:OrderGoods")      // 代评价商品信息
	beego.Router("wx/order/comment", &controllers.OrderController{}, "post:OrderComment") // 评价订单商品信息

	// 足迹
	beego.Router("wx/footprint/list", &controllers.FootprintController{}, "get:FootprintList")
	beego.Router("wx/footprint/delete", &controllers.FootprintController{}, "post:FootprintDelete")

	// 反馈
	beego.Router("wx/feedback/submit", &controllers.FeedbackController{}, "post:FeedbackSubmit")
	// 图片上传
	beego.Router("wx/storage/upload", &controllers.StorageController{}, "post:StorageUpload")

	// 优惠卷
	beego.Router("wx/coupon/list", &controllers.CouponController{}, "get:CouponList")                 // 优惠券列表
	beego.Router("wx/coupon/mylist", &controllers.CouponUserController{}, "get:CouponMylist")         // 我的优惠券列表
	beego.Router("wx/coupon/selectlist", &controllers.CouponUserController{}, "get:CouponSelectList") // 当前订单可用优惠券列表
	beego.Router("wx/coupon/receive", &controllers.CouponUserController{}, "post:CouponReceive")      // 优惠券领取
	beego.Router("wx/coupon/exchange", &controllers.CouponUserController{}, "post:CouponExchange")    // 优惠券兑换

	// 个人中心订单状态数量
	beego.Router("wx/user/index", &controllers.UserController{}, "get:UserIndex")

	// 帮助信息
	beego.Router("wx/issue/list", &controllers.IssueController{}, "get:IssueList")

	// 售后
	beego.Router("wx/aftersale/submit", &controllers.AfterSaleController{}, "post:AftersaleSubmit") // 提交售后申请
	beego.Router("wx/aftersale/list", &controllers.AfterSaleController{}, "get:AftersaleList")      // 售后列表
	beego.Router("wx/aftersale/detail", &controllers.AfterSaleController{}, "get:AftersaleDetail")  // 售后详情

	beego.Router("wx/paynotify", &controllers.BaseController{}, "post:PayNotify") // 支付回调函数

	beego.Router("wx/test", &controllers.BaseController{}, "get:Test") // 物流测试

	// 生成swagger
	//ns := beego.NewNamespace("/v1",
	//	beego.NSInclude(
	//		&controllers.AuthController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.HomeController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.CatalogController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.GoodsController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.CartController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.CollectController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.CommentController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.TopicController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.SearchController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.AddressController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.RegionController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.OrderController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.FootprintController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.FeedbackController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.StorageController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.CouponUserController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.UserController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.IssueController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.AfterSaleController{},
	//	),
	//	beego.NSInclude(
	//		&controllers.BaseController{},
	//	),
	//)
	//beego.AddNamespace(ns)
}
