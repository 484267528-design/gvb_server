package core

import (
	"gvb_server/config"
	"gvb_server/global"
	"io/fs"
	"io/ioutil"
	"log"

	"gopkg.in/yaml.v3"
)

const configFile = "settings.yaml"

func Initconf() {
	c := &config.Config{}
	yamlConf, err := ioutil.ReadFile(configFile)
	if err != nil {
		panic(err)
	}
	err = yaml.Unmarshal(yamlConf, c)
	if err != nil {
		panic(err)
	}
	log.Println("config yamlFile Init success")
	global.Config = c
}
func SetYaml() error {
	bytedata, err := yaml.Marshal(global.Config)
	if err != nil {
		return err
	}
	err = ioutil.WriteFile(configFile, bytedata, fs.ModePerm)
	if err != nil {
		global.Log.Error(err)
		return err
	}
	global.Log.Info("配置文件成功")
	return nil
}
