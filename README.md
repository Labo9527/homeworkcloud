# 服务计算作业——开发 web 服务程序

## 17343140 杨泽涛

## 任务目标

1. 熟悉 go 服务器工作原理
2. 基于现有 web 库，编写一个简单 web 应用类似 cloudgo。
3. 使用 curl 工具访问 web 程序
4. 对 web 执行压力测试

## 任务要求

1. 编程 web 服务程序 类似 cloudgo 应用。
   - 要求有详细的注释
   - 是否使用框架、选哪个框架自己决定 请在 README.md 说明你决策的依据
2. 使用 curl 测试，将测试结果写入 README.md
3. 使用 ab 测试，将测试结果写入 README.md。并解释重要参数。

## 实验过程

本次作业是仿照课程主页来完成简单web服务程序cloudgo的开发工作，看了老师的示例代码之后感觉改动空间不大。。。问了一下老师，老师说可以按照课程主页来写就行了，但我还是做了一些改动，以体现自己的工作量（虽然不多）。。。

### 编写Cloudgo

#### 项目目录：

![image-20191112110921244](/Users/yang/Library/Application Support/typora-user-images/image-20191112110921244.png)

目录结构很简单，只有一个main.go和一个service包

main.go:

```go
package main

import (
	"./service"
)

func main() {
	service.NewServer("8080") //开启服务器，默认端口8080
}
```

server.go

```go
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
```

 我的cloudgo程序使用了martini框架， 在我选择的时候有两条路摆在我的面前：1. 选择用web库手撸cloudgo 2.使用martini框架 促使我选择使用框架的原因有1. 老师的课程示例是使用Martini框架的 2.我上次作业用了web库，所以这次想试试框架



我做出的改动是利用了路由分组处理的方法，采用了Group方法来处理集中为不同的路由分配不同的处理函数。

其次是我尝试使用了一个简单的中间件，它会对每次请求都输出"i am middle man"，这是最简单的一种中间件。



![image-20191112113710033](/Users/yang/Library/Application Support/typora-user-images/image-20191112113710033.png)

![image-20191112113739685](/Users/yang/Library/Application Support/typora-user-images/image-20191112113739685.png)



#### curl测试

![image-20191112113814664](/Users/yang/Library/Application Support/typora-user-images/image-20191112113814664.png)



#### ab压力测试

![image-20191112113859371](/Users/yang/Library/Application Support/typora-user-images/image-20191112113859371.png)

## 实验总结

本次实验比较简单，仿照了课程主页做下来基本没有太大问题，唯一比较在意的就是改动空间不大，工作量不多，自己最后加了一个中间件和路由分组体现了一下工作量，对golang进行一些web开发上的应用让我体会到golang的强大，简简单单几行代码就可以实现js复杂的功能，希望golang继续发展壮大，早日造福业界。

