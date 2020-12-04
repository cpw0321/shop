// Copyright 2020 The shop Authors

// Package response implements response body.
package response

type IssueListRtnJSON struct {
	List []IssueRtnJSON `json:"list"`
}

type IssueRtnJSON struct {
	ID         int    `json:"id"`
	Question   string `json:"question"`
	Answer     string `json:"answer"`
	AddTime    string `json:"addTime"`
	UpdateTime string `json:"updateTime"`
	Deleted    bool   `json:"deleted"`
}
