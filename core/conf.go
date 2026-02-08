package core

import (
	"fmt"
	"os"

	"gvb_server/config"
	"gvb_server/global"

	"gopkg.in/yaml.v3"
)

func Initconf() {
	const configFile = "settings.yaml"

	// 1. 读取 yaml 文件
	data, err := os.ReadFile(configFile)
	if err != nil {
		panic(fmt.Sprintf("配置文件读取失败: %v", err))
	}

	// 2. 初始化 Config
	conf := new(config.Config)

	// 3. 解析 yaml → 结构体
	err = yaml.Unmarshal(data, conf)
	if err != nil {
		panic(fmt.Sprintf("配置文件解析失败: %v", err))
	}

	// 4. 赋值给全局变量
	global.Config = conf

	fmt.Println("config yamlFile load Init success.")
}
