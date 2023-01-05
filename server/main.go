package main

import (
	"pt-auto/app"
)

//主程序
func main() {
	//初始化config
	app.Config()
	//初始化redis
	// app.Redis()
	//初始化log
	app.ZapLog()
	//定时任务
	app.Schedule()
	//初始话web
	app.Web()
}
