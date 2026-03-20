package main

import (
	"fmt"
	"gvb_server/models"

	"github.com/fatih/structs"
)

type AdvertRequest struct {
	models.MODEL `structs:"-"`
	Title        string ` json:"title" binding:"required"  msg:"请输入标题" structs:"title"`      //标题
	Href         string ` json:"href" binding:"required,url" msg:"跳转连接非法" structs:"-"`       //跳转链接
	Image        string ` json:"image" binding:"required,url" msg:"图片地址非法" structs:"-"`      //图片
	IsShow       bool   ` json:"is_show" binding:"required" msg:"请选择是否展示" structs:"is_show"` //是否展示
}

func main() {
	u1 := AdvertRequest{
		Title:  "xxx",
		Href:   "xxx",
		Image:  "xxx",
		IsShow: true,
	}
	m3 := structs.Map(&u1)
	fmt.Println(m3)
}
