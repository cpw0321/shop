// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// HomeRtnJSON 首页返回体
type HomeRtnJSON struct {
	Banners    []BannerRtnJSON `json:"banner"`       // 首页轮播图
	CouponList []CouponRtnJSON `json:"couponList"`   // 首页优惠卷
	NewGoods   []GoodsRtnJSON  `json:"newGoodsList"` // 好物推荐
	HotGoods   []GoodsRtnJSON  `json:"hotGoodsList"` // 热门商品
	TopicList  []TopicJSON     `json:"topicList"`    // 专题
	//Channel    []HomeChannelRtnJSON `json:"channel"`      // 首页分类
	//BrandList    []models.ShopBrand   `json:"brandList"` // 品牌制造商
	//CategoryList []newCategoryList      `json:"floorGoodsList"` // 分类推荐
}

// BannerRtnJSON 轮播图返回体
type BannerRtnJSON struct {
	ID         int    `json:"id"`
	Name       string `json:"name"`
	Link       string `json:"link"`
	URL        string `json:"url"`
	Position   int    `json:"position"`
	Content    string `json:"content"`
	StartTime  string `json:"startTime"`
	EndTime    string `json:"endTime"`
	Enabled    bool   `json:"enabled"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
}

// GoodsRtnJSON 商品信息返回体
type GoodsRtnJSON struct {
	ID           int      `json:"id"`
	GoodsSn      string   `json:"goodsSn"`
	Name         string   `json:"name"`
	CategoryID   int      `json:"categoryId"`
	BrandID      int      `json:"brandId"`
	Gallery      []string `json:"gallery"` // 商品宣传图片列表，采用JSON数组格式
	Keywords     string   `json:"keywords"`
	Brief        string   `json:"brief"`
	PicURL       string   `json:"picUrl"` // 商品页面商品图片
	ShareURL     string   `json:"shareUrl"`
	Unit         string   `json:"unit"`
	CounterPrice float64  `json:"counterPrice"`
	RetailPrice  float64  `json:"retailPrice"`
	Detail       string   `json:"detail"`
	AddTime      string   `json:"addTime"`
	UpdateTime   string   `json:"updateTime"`
}
