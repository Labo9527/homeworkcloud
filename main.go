package main

import (
	"./service"
)

func main() {
	service.NewServer("8080") //开启服务器，默认端口8080
}