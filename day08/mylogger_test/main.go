package main

import (
	"Go/day08/mylogger"
)

var log mylogger.Logger // 声明一个全局的接口变量

// 测试我们自己的日志库
func main() {
	// log := mylogger.Newlog("error")
	log = mylogger.NewConsoleLogger("Info")                                     // 终端日志
	log = mylogger.NewFileLogger("Info", "./", "zhoulinwang.log", 10*1024*1024) // 文件日志

	for {
		id := 100
		name := "理想"
		log.Debug("这是一条Debug日志")
		log.Info("这是一条info日志")
		log.Warning("这是一条Warning日志")
		log.Error("这是一条Error日志,id:%d,name:%s", id, name)
		log.Fatal("这是一条Fatal日志")
		// time.Sleep(time.Second)
	}
}
