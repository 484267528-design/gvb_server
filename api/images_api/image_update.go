package images_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type ImageUpdateRequest struct {
	ID   uint   `json:"id" binding:"required" mag:"请选择文件id"`
	Name string `json:"name" binding:"required" mag:"请输入文件名称"`
}

// ImageUpdateView 图片修改
// @Summary 修改图片名称
// @Description 根据ID修改图片名称
// @Tags 图片管理
// @Accept json
// @Produce json
// @Param data body images_api.ImageUpdateRequest true "图片ID和名称"
// @Success 200 {object} res.Response
// @Router /api/images [put]
func (ImagesApi) ImageUpdateView(c *gin.Context) {
	var cr ImageUpdateRequest
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithError(err, &cr, c)
		return
	}
	var imageModel models.BannerModel
	err = global.DB.Take(&imageModel, cr.ID).Error
	if err != nil {
		res.FailWithMessage("文件不存在", c)
		return
	}
	err = global.DB.Model(&imageModel).Update("name", cr.Name).Error
	if err != nil {
		res.FailWithMessage(err.Error(), c)
		return
	}
	res.OkWithMessage("图片名称修改成功", c)
	return

}
