package router

import (
	"gvb_server/api"
)

func (router RouterGroup) SettingRouter() {
	settingsApi := api.ApiGroupApp.SettingsApi
	router.GET("setting/:name", settingsApi.SettingsInfoView)
	router.PUT("setting/:name", settingsApi.SettingsInfoUpdateView)
}
