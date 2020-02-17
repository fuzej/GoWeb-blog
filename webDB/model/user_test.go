package model

import (
	"fmt"
	"testing"
)

func TestMain(m *testing.M) {
	fmt.Printf("测试开始")
	m.Run()
}
func TestUser(t *testing.T) {
	fmt.Printf("子方法")
	t.Run("测试添加用户", TestAddUser)
}

func TestAddUser(t *testing.T) {
	fmt.Println("测试")
	user := &User{}
	user.AddUser()
	user.AddUser2()
}
