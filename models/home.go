package models

import "hello_goland/config"

type HomeResponse struct {
	config.Viewer
	Categorys []Category
	Posts     []Postmore
	Total     int
	Page      int
	Pages     []int
	PageEnd   bool
}
