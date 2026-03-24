package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageResponse struct {
	ID   int    `json:"id"`
	Path string `json:"path"`
	Name string `json:"name"`
}

// ImageNameList 图片名称列表
// @Summary 图片名称列表
// @Description 图片名称列表
// @Tags 图片管理
// @Accept json
// @Produce json
// @Success 200 {object} res.Response{data=[]ImageResponse}
// @Router /api/images_names [get]
func (ImagesApi) ImageNameList(c *gin.Context) {
	var imageList []ImageResponse
	global.DB.Model(models.BannerModel{}).Select("id", "path", "name").Scan(&imageList)
	res.OkWithData(imageList, c)
}
