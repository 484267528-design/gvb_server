package router

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	ImageApi := api.ApiGroupApp.ImagesApi
	router.GET("images", ImageApi.ImageListView)
	router.POST("images", ImageApi.ImageUploadView)
}
