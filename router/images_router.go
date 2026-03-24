package router

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	ImageApi := api.ApiGroupApp.ImagesApi
	router.GET("images", ImageApi.ImageListView)
	router.GET("images_names", ImageApi.ImageNameList)
	router.POST("images", ImageApi.ImageUploadView)
	router.DELETE("images", ImageApi.ImageRemoveView)
	router.PUT("images", ImageApi.ImageUpdateView)
}
