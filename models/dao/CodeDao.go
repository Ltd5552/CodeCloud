package dao

import (
	"codecloud/domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CodeDao struct {
}

func (c *CodeDao) Save(cid string, uid string, code string, result string, error string, time string, Type string) error {
	sqlStr := "INSERT INTO Code(cid, uid, code, result, time, type,errors) VALUE(?,?,?,?,?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("close err:", err)
		}
	}(stmt)
	_, err = stmt.Exec(cid, uid, code, result, time, Type, error)
	if err != nil {
		return err
	}
	return nil
}

// GetAll get /code/history
func (c *CodeDao) GetAll(uid string) ([]domain.List, error) {

	var Lists []domain.List

	sqlStr := "SELECT cid,time FROM Code WHERE uid = ?"
	query, err := db.Query(sqlStr, uid)
	if err != nil {
		return Lists, err
	}

	for query.Next() {
		var list domain.List
		err := query.Scan(&list.CodeID, &list.RunTime)
		if err != nil {
			return Lists, err
		}
		Lists = append(Lists, list)
	}

	return Lists, nil
}

// GetDetail get /code/detail
func (c *CodeDao) GetDetail(cid string) (domain.Code, error) {
	var code domain.Code

	sqlStr := "SELECT * FROM Code WHERE cid = ?"
	query, err := db.Query(sqlStr, cid)
	if err != nil {
		return code, err
	}

	for query.Next() {
		err := query.Scan(&code.CodeID, &code.UserID, &code.Code, &code.Result, &code.Time, &code.Type, &code.Errors)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return code, err
		}
	}
	return code, nil
}

// DeleteOne delete /code/detail
func (c *CodeDao) DeleteOne(cid string) error {
	sqlStr := "DELETE FROM Code WHERE cid = ?"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		return err
	}
	defer func(stmt *sql.Stmt) {
		err := stmt.Close()
		if err != nil {
			fmt.Println("Close err", err)
		}
	}(stmt)
	_, err = stmt.Exec(cid)
	if err != nil {
		return err
	}
	return nil
}

// GetCode get /code/runAgain
func (c *CodeDao) GetCode(cid string) (string, error) {
	var code string

	sqlStr := "SELECT code FROM Code WHERE cid = ?"
	query, err := db.Query(sqlStr, cid)
	if err != nil {
		return code, err
	}

	for query.Next() {
		err := query.Scan(&code)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return code, err
		}
	}
	return code, nil
}
