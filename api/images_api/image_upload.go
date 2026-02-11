package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models/res"
	"os"
	"path"

	"github.com/gin-gonic/gin"
)

type FileUploadResponse struct {
	Filename  string `json:"file_name"`
	IsSuccess bool   `json:"is_success"` //是否上传成功
	Msg       string `json:"msg"`
}

// 上传单给图片，返回图片的url
func (ImagesApi) ImageUploadView(c *gin.Context) {
	form, err := c.MultipartForm()
	if err != nil {
		res.FailWithMessage(err.Error(), c)
	}
	fileList, ok := form.File["images"]
	if !ok {
		res.FailWithMessage("不存在的文件", c)
		return
	}
	//判断路径是否存在
	basePath := global.Config.Upload.Path
	_, err = os.ReadDir(basePath)
	if err != nil {
		err = os.Mkdir(global.Config.Upload.Path, os.ModePerm)
		if err != nil {
			global.Log.Error(err)
		}
	}
	//不存在就创建
	var resList []FileUploadResponse
	for _, file := range fileList {
		filePath := path.Join(basePath, file.Filename)
		//判断大小
		size := float64(file.Size) / float64(1024*1024)
		if size >= float64(global.Config.Upload.Size) {
			resList = append(resList, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Msg:       fmt.Sprintf("图片大小超过设定大小，当前大小为：%.2fMB,设定大小为%dMB", size, global.Config.Upload.Size),
			})
			continue
		}
		err := c.SaveUploadedFile(file, filePath)
		if err != nil {
			global.Log.Error(err)
			resList = append(resList, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Msg:       err.Error(),
			})

			continue
		}
		resList = append(resList, FileUploadResponse{
			Filename:  filePath,
			IsSuccess: true,
			Msg:       "上传成功",
		})

	}
	res.OkWithData(resList, c)
}
