package service

import (
	"errors"
	"fmt"
	"hello_goland/dao"
	"hello_goland/models"
	"hello_goland/utils"
)

func Login(userName, passwd string) (*models.LoginRes, error) {
	fmt.Println(passwd)
	passwd = utils.Md5Crypt(passwd, "lincoxi")
	//加盐两次，得到的密码进行对照，前端一次，后端一次
	fmt.Println(passwd)
	user := dao.GetUser(userName, passwd)
	if user == nil {
		return nil, errors.New("账号密码不正确")
	}
	uid := user.Uid
	//token jwt 令牌
	token, err := utils.Award(&uid)
	if err != nil {
		return nil, errors.New("token生成失败")
	}
	var userInfo models.UserInfo
	userInfo.Uid = user.Uid
	userInfo.UserName = user.UserName
	userInfo.Avatar = user.Avatar
	var lr = &models.LoginRes{
		token,
		userInfo,
	}
	return lr, nil
}
