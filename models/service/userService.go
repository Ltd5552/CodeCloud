package service

import "codecloud/models/dao"

type UserService struct {
}

var userDao dao.UserDao

func (u *UserService) FindName(userId string) string {
	userDao.FindName(userId)

	return ""
}
