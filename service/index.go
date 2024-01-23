package service

import (
	"hello_goland/config"
	"hello_goland/dao"
	"hello_goland/models"
	"html/template"
)

func GetAllIndexInfo(page, pageSize int) (*models.HomeResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPage(page, pageSize)
	var postMores []models.PostMore
	for _, post := range posts {
		categoryName := dao.GetCategoryNameByID(post.CategoryId)
		userName := dao.GetUserNameByID(post.UserId)
		content := []rune(post.Content)
		if len(content) > 100 {
			content = content[0:100]
		}
		postMore := models.PostMore{
			post.Pid,
			post.Title,
			post.Slug,
			template.HTML(content),
			post.CategoryId,
			categoryName,
			post.UserId,
			userName,
			post.ViewCount,
			post.Type,
			models.DateDay(post.CreateAt),
			models.DateDay(post.CreateAt),
		}
		postMores = append(postMores, postMore)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		1,
		1,
		[]int{1},
		true,
	}
	return hr, nil

}
