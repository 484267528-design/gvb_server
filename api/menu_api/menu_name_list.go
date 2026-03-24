package menu_api

import (
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/res"

	"github.com/gin-gonic/gin"
)

type MenuNameResponse struct {
	ID    uint   `json:"id"`
	Title string `json:"title"`
	Path  string `json:"path"`
}

func (MenuApi) MenuNameList(c *gin.Context) {
	var MenuNameList []MenuNameResponse
	global.DB.Model(&models.MenuModel{}).Select("id", "title", "path").Scan(&MenuNameList)
	res.OkWithData(MenuNameList, c)
}
