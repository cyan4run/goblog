package service

import (
	"hello_goland/config"
	"hello_goland/dao"
	"hello_goland/models"
	"html/template"
)

func GetPostsByCategoryID(cID, page, pageSize int) (*models.CategoryResponse, error) {
	categorys, err := dao.GetAllCategory()
	if err != nil {
		return nil, err
	}
	posts, err := dao.GetPostPageByCategoryID(cID, page, pageSize)
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
	total := dao.CountGetAllPostByCategoryID(cID)
	pagesCount := (total-1)/10 + 1
	var pages []int
	for i := 0; i < pagesCount; i++ {
		pages = append(pages, i+1)
	}
	var hr = &models.HomeResponse{
		config.Cfg.Viewer,
		categorys,
		postMores,
		total,
		page,
		pages,
		page != pagesCount,
	}
	categoryName := dao.GetCategoryNameByID(cID)
	categoryResponse := &models.CategoryResponse{
		hr,
		categoryName,
	}
	return categoryResponse, nil
}
