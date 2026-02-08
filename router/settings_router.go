package router

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingRouter() {
	settings := api.ApiGroupApp.SettingsApi
	router.GET("setting", settings.SettingsInfoView)
}
