// Copyright 2020 The shop Authors

// Package response implements response body.
package response

import "shop/base"

// AddressListRtnJSON 地址列表
type AddressListRtnJSON struct {
	List  []base.ShopAddress `json:"list"`  // 列表
	Total int64              `json:"total"` // 总数
}
