package service

import (
	"codecloud/models/dao"
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
	if err != nil {
		return err
	}
	return nil
}

func (u *UserService) CheckAccount(name string, password string) (string, bool) {
	id, pd := userDao.CheckAccount(name)
	//没有用户
	if pd == "" {
		return id, false
	} else if pd == password {
		return id, true
	}
	return id, false
}
