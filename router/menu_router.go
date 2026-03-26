package router

import (
	"gvb_server/api"
)

func (router RouterGroup) MenuRouter() {
	app := api.ApiGroupApp.MenuApi
	router.POST("menus", app.MenuCreatView)
	router.GET("menus", app.MenuListView)
	router.GET("menus_names", app.MenuNameList)
	router.PUT("menus/:id", app.MenuUpdateView)
}
