package dao

import (
	"CodeCloud/models/domain"
)

type UserDao struct {
}

func (u *UserDao) findName(user domain.User) domain.User {

	return user
}
