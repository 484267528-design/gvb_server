package images_api

import (
	"gvb_server/models"
	"gvb_server/models/res"
	"gvb_server/service/common"

	"github.com/gin-gonic/gin"
)

// ImageListView 图片列表
// @Summary 图片列表
// @Description 获取图片列表
// @Tags 图片管理
// @Accept json
// @Produce json
// @Param page query models.PageInfo false "查询参数"
// @Success 200 {object} res.Response{data=res.ListResponse[models.BannerModel]}
// @Router /api/images [get]
func (ImagesApi) ImageListView(c *gin.Context) {
	var cr models.PageInfo
	err := c.ShouldBindQuery(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	list, count, err := common.ComList(models.BannerModel{}, common.Option{
		cr,
		true,
	})
	res.OkWithList(list, count, c)

	return

}
