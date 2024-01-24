package router

import (
	"hello_goland/api"
	"hello_goland/views"
	"net/http"
)

func Router() {
	http.HandleFunc("/", views.HTML.Index)
	//http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/c/", views.HTML.Category)
	http.HandleFunc("/login", views.HTML.Login)
	http.HandleFunc("/api/v1/login", api.API.Login)
	http.HandleFunc("/p/", views.HTML.Detail)
	http.HandleFunc("/writing", views.HTML.Writing)
	http.HandleFunc("/api/v1/post", api.API.SaveAndUpdatePost)
	http.HandleFunc("/api/v1/post/", api.API.GetPost)
	http.HandleFunc("/api/v1/qiniu/token", api.API.QiniuToken)
	http.HandleFunc("/pigeonhole", views.HTML.Pigeonhole)
	http.HandleFunc("/api/v1/post/search", api.API.SearchPost)
	http.Handle("/resource/", http.StripPrefix("/resource/", http.FileServer(http.Dir("public/resource/"))))
}
