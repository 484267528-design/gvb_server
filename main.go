package main

import (
	"fmt"
	"gvb_server/core"
	"gvb_server/global"
	"gvb_server/router"
)

func main() {
	core.Initconf() //读取配置文件
	fmt.Println(global.Config)

	global.Log = core.InitLogger() //初始化日志

	global.DB = core.InitGorm() //连接数据库
	fmt.Println(global.DB)

	Router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Info(fmt.Sprintf("gvb_server at %s", addr))
	Router.Run(addr)
}
