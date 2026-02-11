package router

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("setting", settingsApi.SettingsInfoView)
	router.PUT("setting", settingsApi.SettingsInfoUpdateView)
}
