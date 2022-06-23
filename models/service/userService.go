package service

import (
	"codecloud/models/dao"
	"errors"
)

type UserService struct {
}

var userDao dao.UserDao

func (u *UserService) FindName(userId string) string {
	name := userDao.FindName(userId)
	if name != "" {
		return name
	}
	return "null"
}

func (u *UserService) Save(name string, password string) error {
	err := userDao.Save(name, password)
	if err == nil {
		return errors.New("保存失败")
	}
	return nil
}
