package dao

import (
	domain2 "codecloud/domain"
	"database/sql"
	"fmt"
	_ "github.com/go-sql-driver/mysql"
)

type CodeDao struct {
}

func (c *CodeDao) Save(cid string, uid string, code string, result string, time string, Type string) error {
	sqlStr := "INSERT INTO Code(cid, uid, code, result, time, type) VALUE(?,?,?,?,?,?)"
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
	_, err = stmt.Exec(cid, uid, code, result, time, Type)
	if err != nil {
		return err
	}
	return nil
}

// GetAll get /code/history
func (c *CodeDao) GetAll() ([]domain2.List, error) {

	var Lists []domain2.List

	sqlStr := "SELECT cid,time FROM Code"
	query, err := db.Query(sqlStr)
	if err != nil {
		return Lists, err
	}

	for query.Next() {
		var list domain2.List
		err := query.Scan(&list.Cid, &list.Time)
		if err != nil {
			return Lists, err
		}
		Lists = append(Lists, list)
	}
	return Lists, nil
}

// GetDetail get /code/detail
func (c *CodeDao) GetDetail(cid string) (domain2.Code, error) {
	var code domain2.Code

	sqlStr := "SELECT * FROM Code WHERE cid = ?"
	query, err := db.Query(sqlStr, cid)
	if err != nil {
		return code, err
	}

	for query.Next() {
		err := query.Scan(&code.CodeID, &code.UserID, &code.Code, &code.Time, &code.Type)
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
