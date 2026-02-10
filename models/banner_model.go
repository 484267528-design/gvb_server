package models

type BannerModel struct {
	MODEL
	Path string `json:"path"` //图片路径
	Hash string `json:"hash"` //图片Hash值，判断重复图片
	Name string `json:"name"` //图片名称
}
