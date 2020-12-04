package response

import (
	"shop/base"
)

// SearchIndexRtnJSON 搜索关键字返回体
type SearchIndexRtnJSON struct {
	DefaultKeyword     base.ShopKeyword         `json:"defaultKeyword"`
	HistoryKeyworkList []base.ShopSearchHistory `json:"historyKeywordList"`
	HotKeywordList     []base.ShopKeyword       `json:"hotKeywordList"`
}

// SearchResultRtnJSON 搜索结果返回体
type SearchResultRtnJSON struct {
	Total              int64             `json:"total"`
	Limit              int               `json:"limit"`
	Page               int               `json:"page"`
	List               []GoodsRtnJSON    `json:"list"`
	FilterCategoryList []CategoryRtnJSON `json:"filterCategoryList"`
}
