package dao

import (
	"codecloud/utils"
	"database/sql"
	"fmt"
)

type UserDao struct {
}

func (u *UserDao) Save(name string, password string) error {
	uid := utils.RandStr(8)
	sqlStr := "INSERT INTO User(uid,name,password) VALUE(?,?,?)"
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
	_, err = stmt.Exec(uid, name, password)
	if err != nil {
		return err
	}
	return nil
}

// FindName get /user/findName
func (u *UserDao) FindName(userID string) string {
	sqlStr := "SELECT name FROM User WHERE uid = ?"
	query, err := db.Query(sqlStr, userID)
	if err != nil {
		fmt.Println("query failed,err:", err)
		return ""
	}
	var name string
	for query.Next() {
		err := query.Scan(name)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return ""
		}
	}
	return name
}
