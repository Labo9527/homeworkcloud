package service

import (
	"fmt"
	"github.com/go-martini/martini" //使用martini框架
)

func second1() string {	//分别设定几个处理函数 second1, second2, second3
	return "second1"
}

func second2() string{
	return "second2"
}

func second3() string{
	return "second3"
}

func NewServer(port string) {
	m := martini.Classic()	//创建实例

	m.Group("/first", func(r martini.Router) {	//	采用了路由分组处理的方法
		r.Get("/second1", second1)
		r.Get("/second2", second2)
		r.Get("/second3", second3)
	})

	m.Use(func(c martini.Context) {
		fmt.Println("i am middle man")
	})

	m.Get("/", func(params martini.Params) string {	//普通路由处理
		return "hello world"
	})

	m.Get("/tom", func(params martini.Params) string {	//普通路由处理
		return "hello tom"
	})

	m.RunOnAddr(":"+port)	//监听指定端口
}