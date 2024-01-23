package views

import (
	"errors"
	"hello_goland/common"
	"hello_goland/service"
	"log"
	"net/http"
	"strconv"
	"strings"
)

func (*HTMLApi) Category(w http.ResponseWriter, r *http.Request) {
	categoryTemplate := common.Template.Category
	path := r.URL.Path
	cIDStr := strings.TrimPrefix(path, "/c/")
	cID, err := strconv.Atoi(cIDStr)
	if err != nil {
		categoryTemplate.WriteError(w, errors.New("不识别此请求路径！"))
		return
	}
	if err := r.ParseForm(); err != nil {
		log.Println("表单获取失败：", err)
		categoryTemplate.WriteError(w, errors.New("系统错误，请联系管理员!"))
		return
	}
	pageStr := r.Form.Get("page")
	if pageStr == "" {
		pageStr = "1"
	}
	page, _ := strconv.Atoi(pageStr)
	pageSize := 10
	categoryResponse, err := service.GetPostsByCategoryID(cID, page, pageSize)
	if err != nil {
		categoryTemplate.WriteError(w, err)
		return
	}
	categoryTemplate.WriteData(w, categoryResponse)
}
