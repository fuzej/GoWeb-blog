package main

import (
	"github.com/fuzej/WebStudy-尚硅谷/bookstore/controller"
	"net/http"
)

func main() {
	//设置处理静态资源，如css和js文件
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("views/static"))))
	//直接去html页面
	http.Handle("/pages/", http.StripPrefix("/pages/", http.FileServer(http.Dir("views/pages"))))
	//去首页
	http.HandleFunc("/main", controller.GetPageBooksByPrice)
	//去登录
	http.HandleFunc("/login", controller.Login)
	//去注销
	http.HandleFunc("/logout", controller.Logout)
	//去注册
	http.HandleFunc("/regist", controller.Regist)
	//通过Ajax请求验证用户名是否可用
	http.HandleFunc("/checkUserName", controller.CheckUserName)
	//获取所有图书
	// http.HandleFunc("/getBooks", controller.GetBooks)
	//获取带分页的图书信息
	http.HandleFunc("/getPageBooks", controller.GetPageBooks)
	http.HandleFunc("/getPageBooksByPrice", controller.GetPageBooksByPrice)
	//添加图书
	// http.HandleFunc("/addBook", controller.AddBook)
	//删除图书
	http.HandleFunc("/deleteBook", controller.DeleteBook)
	//去更新图书的页面
	http.HandleFunc("/toUpdateBookPage", controller.ToUpdateBookPage)
	//更新或添加图书
	http.HandleFunc("/updateOraddBook", controller.UpdateOrAddBook)
	//添加图书到购物车中
	http.HandleFunc("/addBook2Cart", controller.AddBook2Cart)
	//获取购物车信息
	http.HandleFunc("/getCartInfo", controller.GetCartInfo)
	//清空购物车
	http.HandleFunc("/deleteCart", controller.DeleteCart)

	http.ListenAndServe(":8080", nil)
}
