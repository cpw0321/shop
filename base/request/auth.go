// Copyright 2020 The shop Authors

// Package request implements request body.
package request

// AuthLoginBody 请求体
type AuthLoginBody struct {
	Code     string     `json:"code"`
	UserInfo WXUserBody `json:"userInfo"`
}

// WXUserBody 请求体中用户信息
type WXUserBody struct {
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
	Gender    int    `json:"gender"`
	Country   string `json:"country"`
	Province  string `json:"province"`
	City      string `json:"city"`
	Language  string `json:"language"`
}

// WXInfo 用户openID信息
type WXInfo struct {
	SessionKey string `json:"sessionKey"`
	OpenID     string `json:"openid"`
}
