package api

import (
	"hello_goland/common"
	"hello_goland/service"
	"net/http"
)

func (*Api) Login(w http.ResponseWriter, r *http.Request) {
	param := common.GetRequestJsonParam(r)
	userName := param["username"].(string)
	passwd := param["passwd"].(string)
	loginRes, err := service.Login(userName, passwd)
	if err != nil {
		common.Error(w, err)
		return
	}
	common.Success(w, loginRes)

}
