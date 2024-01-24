package service

import (
	"hello_goland/config"
	"hello_goland/dao"
	"hello_goland/models"
	"html/template"
	"log"
)

func GetPostDetail(pid int) (*models.PostRes, error) {
	post, err := dao.GetPostByID(pid)

	if err != nil {
		return nil, err
	}
	categoryName := dao.GetCategoryNameByID(post.CategoryId)
	userName := dao.GetUserNameByID(post.UserId)
	postMore := models.PostMore{
		post.Pid,
		post.Title,
		post.Slug,
		template.HTML(post.Content),
		post.CategoryId,
		categoryName,
		post.UserId,
		userName,
		post.ViewCount,
		post.Type,
		models.DateDay(post.CreateAt),
		models.DateDay(post.CreateAt),
	}
	var postRes = &models.PostRes{
		config.Cfg.Viewer,
		config.Cfg.System,
		postMore,
	}
	return postRes, nil

}
func Writing() (wr models.WritingRes) {
	wr.Title = config.Cfg.Viewer.Title
	wr.CdnURL = config.Cfg.System.CdnURL
	category, err := dao.GetAllCategory()
	if err != nil {
		log.Println(err)
		return
	}

	wr.Categorys = category
	return
}
func SavePost(post *models.Post) {
	dao.SavePost(post)
}
func UpdatePost(post *models.Post) {
	dao.UpdatePost(post)
}
func SearchPost(condition string)[]models.SearchResp{
	posts,_:=dao.GetPostSearch(condition)
	var searchResps []models.SearchResp
	for _,post:=range posts{
		searchResps=append(searchResps,models.SearchResp{
			post.Pid,
			post.Title,

		})
	}
	return searchResps
}