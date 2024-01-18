package models

import (
	"hello_goland/config"
	"html/template"
	"time"
)

type Post struct {
	Pid        int       `json:"pid"`
	Title      string    `json:"title"`
	Slug       string    `json:"slug"`
	Content    string    `json:"content"`
	Markdown   string    `json:"markdown"`
	CategoryID int       `json:"categoryID"`
	UserID     int       `json:"userID"`
	ViewCount  int       `json:"viewCount"`
	Type       int       `json:"type"`
	CreateAt   time.Time `json:"createAt"`
	UpdateAt   time.Time `json:"updateAt"`
}
type Postmore struct {
	Pid        int           `json:"pid"`
	Title      string        `json:"title"`
	Slug       string        `json:"slug"`
	Content    template.HTML `json:"content"`
	Markdown   string        `json:"markdown"`
	CategoryID int           `json:"categoryID"`
	UserID     int           `json:"userID"`
	ViewCount  int           `json:"viewCount"`
	Type       int           `json:"type"`
	CreateAt   time.Time     `json:"createAt"`
	UpdateAt   time.Time     `json:"updateAt"`
}
type PostReq struct {
	Pid        int    `json:"pid"`
	Title      string `json:"title"`
	Slug       string `json:"slug"`
	Content    string `json:"content"`
	Markdown   string `json:"markdown"`
	CategoryID int    `json:"categoryID"`
	Type       int    `json:"type"`
}
type SearchResp struct {
	Pid   int    `orm:"pid" json:"pid"`
	Title string `orm:"title" json:"title"`
}
type PostRes struct {
	config.Viewer
	config.SystemConfig
	Article Postmore
}
