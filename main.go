package main

import (
	"fmt"
	"gvb_server/core"
	_ "gvb_server/docs"
	"gvb_server/flag"
	"gvb_server/global"
	"gvb_server/router"
)

// @title gvb_server API文档
// @version 1.0
// @description gvb_server API文档
// @host localhost:8080
// @BasePath /
func main() {
	core.Initconf() //读取配置文件
	fmt.Println(global.Config)

	global.Log = core.InitLogger() //初始化日志

	global.DB = core.InitGorm() //连接数据库
	fmt.Println(global.DB)

	option := flag.Parse()
	if flag.IsWebStop(option) {
		flag.SwitchOption(option)
		return
	} //命令行参数绑定

	Router := router.InitRouter()
	addr := global.Config.System.Addr()
	global.Log.Info(fmt.Sprintf("gvb_server at %s", addr))
	err := Router.Run(addr)
	if err != nil {
		global.Log.Fatalf(err.Error())
	}

}
