package dao

import (
	"fmt"
	"github.com/fuzej/WebStudy-尚硅谷/bookstore/model"
	"testing"
)

func TestUser(t *testing.T) {
	fmt.Println("测试userdao中的函数")
	//t.Run("验证用户名或密码：", testLogin)
	//t.Run("验证用户名：", testRegist)
	// t.Run("保存用户：", testSave)
	//book
	//t.Run("测试数据添加",TestGetBooks)
	//t.Run("测试数据添加",TestAddBook)
	t.Run("测试数据添加", TestDeleteBook)
}

func testLogin(t *testing.T) {
	user, _ := CheckUserNameAndPassword("admin", "123456")
	fmt.Println("获取用户信息是：", user)
}
func testRegist(t *testing.T) {
	user, _ := CheckUserName("admin")
	fmt.Println("获取用户信息是：", user)
}
func testSave(t *testing.T) {
	SaveUser("admin3", "123456", "admin@atguigu.com")
}
func TestGetBooks(t *testing.T) {
	fmt.Println("测试book-get")
	books, _ := GetBooks()
	for k, v := range books {
		fmt.Printf("第%v本信息：%v\n", k+1, v)
	}
}
func TestAddBook(t *testing.T) {
	fmt.Printf("测试添加book")
	book := &model.Book{
		ID:      0,
		Title:   "三国",
		Author:  "罗gg",
		Price:   12.22,
		Sales:   51,
		Stock:   15,
		ImgPath: "/static/img/default.jpg",
	}
	//
	AddBook(book)
}
func TestDeleteBook(t *testing.T) {
	fmt.Printf("测试添加book")

	DeleteBook("32")
}
