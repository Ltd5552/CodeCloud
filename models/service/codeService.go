package service

import (
	"codecloud/domain"
	"codecloud/models/dao"
	"fmt"
)

type CodeService struct {
}

var codeDao dao.CodeDao

func (c *CodeService) Save(cid string, uid string, code string, result string, time string, Type string) bool {
	err := codeDao.Save(cid, uid, code, result, time, Type)
	if err != nil {
		return false
	}
	return true
}

func (c *CodeService) GetList(uid string) []domain.List {
	all, err := codeDao.GetAll(uid)
	if err != nil {
		fmt.Println("get all err", err)
	}
	return all
}

func (c *CodeService) GetDetail(cid string) domain.Code {
	detail, err := codeDao.GetDetail(cid)
	if err != nil {
		fmt.Println("get detail failed,", err)
	}
	return detail
}

func (c *CodeService) DeleteOne(cid string) {
	err := codeDao.DeleteOne(cid)
	if err != nil {
		fmt.Println("删除失败，", err)
	}
	fmt.Println("删除成功！")
}

func (c *CodeService) GetCode(cid string) string {
	code, err := codeDao.GetCode(cid)
	if err != nil {
		return "null"
	}
	return code
}
