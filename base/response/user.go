// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// UserIndexRtnJSON 用户个人中心，订单数量详情返回体
type UserIndexRtnJSON struct {
	Order IndexOrderData `json:"order"`
}

// IndexOrderData 首页订单数量
type IndexOrderData struct {
	Uncomment   int64 `json:"uncomment"`   // 未评论
	Unpaid      int64 `json:"unpaid"`      // 未付款
	Unrecv      int64 `json:"unrecv"`      // 未收货
	Unship      int64 `json:"unship"`      // 未发货
	Unaftersale int64 `json:"unaftersale"` // 售后
}

type HomeChannelRtnJSON struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	IconURL string `json:"iconUrl"`
}

// HomeAboutRtnJSON 关于我们
type HomeAboutRtnJSON struct {
	List []AboutRtnJSON `json:"list"`
}

type AboutRtnJSON struct {
	ID         int    `json:"id"`
	Title      string `json:"title"`
	Context    string `json:"context"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
}
