package dao

import (
	"fmt"
)

type UserDao struct {
}

func (u *UserDao) Save(uid string, name string, password string) error {
	sqlStr := "INSERT INTO User(uid,name,password) VALUE(?,?,?)"
	stmt, err := db.Prepare(sqlStr)
	if err != nil {
		fmt.Println(err)
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec(uid, name, password)
	if err != nil {
		fmt.Println(err)
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
		err := query.Scan(&name)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return ""
		}
	}
	return name
}

func (u *UserDao) CheckAccount(name string) (string, string) {
	var password string
	var id string
	sqlStr := "SELECT uid,password FROM User WHERE name = ?"
	query, err := db.Query(sqlStr, name)
	if err != nil {
		fmt.Println("query failed,err:", err)
		return "", ""
	}
	for query.Next() {
		err := query.Scan(&id, &password)
		if err != nil {
			fmt.Println("Scan failed,err:", err)
			return "", ""
		}
	}
	return id, password

}
