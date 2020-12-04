// Copyright 2020 The shop Authors

// Package response implements response body.
package response

// CommentCountRtnJSON 评论数量返回体
type CommentCountRtnJSON struct {
	AllCount    int64 `json:"allCount"`
	HasPicCount int64 `json:"hasPicCount"`
}

// CommentRtnJSON 评论返回体
type CommentRtnJSON struct {
	Total int64                `json:"total"`
	Limit int                  `json:"limit"`
	Page  int                  `json:"page"`
	List  []CommentListRtnJSON `json:"list"`
}

type CommentListRtnJSON struct {
	ID       int             `json:"id"`
	UserInfo UserInfoRtnJSON `json:"userInfo"`
	Content  string          `json:"content"`
	PicList  []string        `json:"picList"`
	AddTime  string          `json:"addTime"`
}

type UserInfoRtnJSON struct {
	Nickname  string `json:"nickname"`
	AvatarURL string `json:"avatarUrl"`
}

type CommentListRespon struct {
	ID       int      `json:"id"`
	Nickname string   `json:"nickname"`
	Avatar   string   `json:"avatar"`
	Content  string   `json:"content"`
	PicList  []string `json:"picList"`
	AddTime  string   `json:"addTime"`
}
