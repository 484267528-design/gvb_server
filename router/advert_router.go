package router

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	ImageApi := api.ApiGroupApp.AdvertApi
	router.POST("adverts", ImageApi.AdvertCreateView)

}
