// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// GoodsListRtnJSON 商品列表返回体
type GoodsListRtnJSON struct {
	Limit              int               `json:"limit"`
	Page               int               `json:"page"`
	Total              int               `json:"total"`
	List               []GoodsRtnJSON    `json:"list"`
	FilterCategoryList []CategoryRtnJSON `json:"filterCategoryList"`
}

// GoodsProductRtnJSON 商品货品表,即某种规格对应某种价格
type GoodsProductRtnJSON struct {
	ID             int      `json:"id"`
	GoodsID        int      `json:"goodsId"`
	Specifications []string `json:"specifications"`
	Price          float64  `json:"price"`
	Number         int      `json:"number"`
	URL            string   `json:"url"`
	AddTime        string   `json:"addTime"`
	UpdateTime     string   `json:"updateTime"`
}

// GoodsAttributeRtnJSON 商品属性
type GoodsAttributeRtnJSON struct {
	ID         int    `json:"id"`
	GoodsID    int    `json:"goodsId"`
	Attribute  string `json:"attribute"`
	Value      string `json:"value"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
}

// GoodsSpecificationRtnJSON 商品规格返回体
type GoodsSpecificationRtnJSON struct {
	ID            int    `json:"id"`
	GoodsID       int    `json:"goodsId"`
	Specification string `json:"specification"`
	Value         string `json:"value"`
	PicURL        string `json:"picUrl"`
	AddTime       string `json:"addTime"`
	UpdateTime    string `json:"updateTime"`
}

// GoodsDetailRtnJSON 商品详情返回体
type GoodsDetailRtnJSON struct {
	SpecificationList []SpecificationItem     `json:"specificationList"`
	Issues            []IssueRtnJSON          `json:"issue"`
	UserHasCollect    int                     `json:"userHasCollect"`
	ShareImage        string                  `json:"shareImage"`
	Comment           GoodsComment            `json:"comment"`
	Share             bool                    `json:"share"`
	Attribute         []GoodsAttributeRtnJSON `json:"attribute"`
	Goods             GoodsRtnJSON            `json:"info"`
	ProductList       []GoodsProductRtnJSON   `json:"productList"`
	Brand             BrandRespon             `json:"brand"`
}

// SpecificationItem 商品规格
type SpecificationItem struct {
	Name      string                      `json:"name"`
	ValueList []GoodsSpecificationRtnJSON `json:"valueList"`
}

// GoodsComment 商品评论
type GoodsComment struct {
	Count int64               `json:"count"`
	Data  []CommentListRespon `json:"data"`
}
