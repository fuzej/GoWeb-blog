package model

import (
	"fmt"
	"github.com/fuzej/WebStudy-尚硅谷/webDB/util"
)

type User struct {
	User     string
	Password string
}

func (user *User) AddUser() error {
	sqlStr := "insert into userinfo(user,password) values (?,?)"
	//预编译
	inStmt, err := util.DB.Prepare(sqlStr)
	if err != nil {
		fmt.Println("错误", err.Error())
		return err
	}
	_, err2 := inStmt.Exec("root", "123456")
	if err2 != nil {
		fmt.Println("执行异常", err2.Error())
		return err2
	}
	return nil
}
func (user *User) AddUser2() error {
	sqlStr := "insert into userinfo(user,password) values (?,?)"

	_, err := util.DB.Exec(sqlStr, "as", "123")
	if err != nil {
		fmt.Println("2执行异常", err.Error())
		return err
	}
	return nil
}
