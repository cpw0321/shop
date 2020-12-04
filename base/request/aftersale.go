// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// AftersaleBody 售后请求体
type AftersaleBody struct {
	Amount   float64  `json:"amount"`
	OrderID  string   `json:"orderid"`
	Pictures []string `json:"pictures"`
	Reason   string   `json:"reason"`
	Type     int      `json:"type"`
	TypeDesc string   `json:"typeDesc"`
}
