package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

// ImageRemoveView 删除图片
// @Summary 删除图片
// @Description 根据ID列表删除图片
// @Tags 图片管理
// @Accept json
// @Produce json
// @Param data body models.RemoveResult true "图片ID列表"
// @Success 200 {object} res.Response
// @Router /api/images [delete]
func (ImagesApi) ImageRemoveView(c *gin.Context) {
	var cr models.RemoveResult
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var imageList []models.BannerModel
	count := global.DB.Find(&imageList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("文件不存在", c)
		return
	}
	global.DB.Delete(&imageList)
	res.OkWithMessage(fmt.Sprintf("共删除%d张图片", count), c)

}
