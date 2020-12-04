// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// AddressSaveBody 添加地址请求体
type AddressSaveBody struct {
	ID            int    `json:"id"`            // 地址ID
	Name          string `json:"name"`          // 收货人名称
	Tel           string `json:"tel"`           // 手机号码
	Province      string `json:"province"`      // 省
	City          string `json:"city"`          // 市
	County        string `json:"county"`        // 区县
	AddressDetail string `json:"addressDetail"` // 详细收货地址
	AreaCode      string `json:"areaCode"`      // 区域码
	IsDefault     bool   `json:"isDefault"`     // 是否默认地址
}

// AddressDeleteBody 地址删除请求体
type AddressDeleteBody struct {
	ID int `json:"id"` // 地址id
}
