package views

import (
	"errors"
	"hello_goland/common"
	"hello_goland/service"
	"log"
	"net/http"
)

func (*HTMLApi) Index(w http.ResponseWriter, r *http.Request) {
	index := common.Template.Index
	//页面上涉及到的所有的数据，必须有定义
	hr, err := service.GetAllIndexInfo()
	if err != nil {
		log.Println("获取数据出错:", err)
		index.WriteError(w, errors.New("系统错误，联系管理员！"))
	}
	index.WriteData(w, hr)
}
