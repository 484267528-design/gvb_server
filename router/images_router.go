package router

import "gvb_server/api"

func (router RouterGroup) ImagesRouter() {
	ImageApi := api.ApiGroupApp.ImagesApi
	router.POST("images", ImageApi.ImageUploadView)
}
