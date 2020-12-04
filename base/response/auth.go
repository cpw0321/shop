// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// AuthUserRtnJSON 用户登录返回体
type AuthUserRtnJSON struct {
	UserInfo UserData `json:"userInfo"`
	Token    string   `json:"token"`
}

// UserData 用户信息
type UserData struct {
	NickName  string `json:"nickName"`
	AvatarURL string `json:"avatarUrl"`
}
