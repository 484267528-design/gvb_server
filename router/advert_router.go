package router

import "gvb_server/api"

func (router RouterGroup) AdvertRouter() {
	ImageApi := api.ApiGroupApp.AdvertApi
	router.POST("adverts", ImageApi.AdvertCreateView)
	router.GET("adverts", ImageApi.AdvertListView)
	router.PUT("adverts/:id", ImageApi.AdvertUpdateView)
	router.DELETE("adverts", ImageApi.AdvertRemoveView)

}
