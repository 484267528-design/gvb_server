package menu_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

func (MenuApi) MenuRemoveView(c *gin.Context) {
	var cr models.RemoveResult
	err := c.ShouldBindJSON(&cr)
	if err != nil {
		res.FailWithCode(res.ArgumentError, c)
		return
	}
	var MenuList []models.MenuModel
	count := global.DB.Find(&MenuList, cr.IDList).RowsAffected
	if count == 0 {
		res.FailWithMessage("菜单不存在", c)
		return
	}
	err = global.DB.Transaction(func(tx *gorm.DB) error {
		err = global.DB.Model(&MenuList).Association("Banners").Clear()
		if err != nil {
			global.Log.Error(err)
			return err
		}

		err = global.DB.Delete(&MenuList).Error
		if err != nil {
			global.Log.Error(err)
			return err
		}
		return nil
	})
	if err != nil {
		global.Log.Error(err)
		res.FailWithMessage("删除菜单失败", c)
		return
	}
	res.OkWithMessage(fmt.Sprintf("共删除%d条菜单", count), c)

}
