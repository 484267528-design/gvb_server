package advert_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type AdvertRequest struct {
	Title  string ` json:"title" binding:"required" msg:"请输入标题"`       //标题
	Href   string ` json:"href" binding:"required,url" msg:"跳转连接非法"`  //跳转链接
	Image  string ` json:"image" binding:"required,url" msg:"图片地址非法"` //图片
	IsShow bool   ` json:"is_show" binding:"required" msg:"请选择是否展示"` //是否展示
}

func (AdvertApi) AdvertCreateView(c *gin.Context) {
	var cr AdvertRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	//重复的判断
	var advert models.AdvertModel
	err = global.DB.Take(&advert, "title=?", cr.Title).Error
	if err == nil {
		res.FailWithMessage("该广告已存在", c)
		return
	}

	err = global.DB.Create(&models.AdvertModel{
		Title:  cr.Title,
		Href:   cr.Href,
		Image:  cr.Image,
		IsShow: cr.IsShow,
	}).Error
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("添加广告失败", c)
		return
	}
	res.OkWithMessage("添加广告成功", c)

}
