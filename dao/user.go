package dao

import "log"

func GetUserNameByID(userID int) string {
	row := DB.QueryRow("select user_name from blog_user where uid=?", userID)
	if row.Err() != nil {
		log.Println(row.Err())
	}
	var userName string
	_ = row.Scan(&userName)
	return userName
}
