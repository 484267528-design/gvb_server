package images_api

import (
	"fmt"
	"gvb_server/global"
	"gvb_server/models"
	"gvb_server/models/ctype"
	"gvb_server/models/res"
	"gvb_server/plugins/qiniu"
	"gvb_server/utils"
	"io"
	"os"
	"path"
	"strings"

	"github.com/gin-gonic/gin"
)

var (
	//图片白名单
	WhiteImageList = []string{
		"jpg",
		"jpeg",
		"png",
		"gif",
		"ico",
		"tiff",
		"svg",
		"webp",
	}
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
		fileName := file.Filename
		nameList := strings.Split(fileName, ".")
		suffix := strings.ToLower(nameList[len(nameList)-1]) //图片文件后缀
		if !utils.InList(suffix, WhiteImageList) {
			resList = append(resList, FileUploadResponse{
				Filename:  file.Filename,
				IsSuccess: false,
				Msg:       "非法文件",
			})
			continue

		}
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

		fileObj, err := file.Open()
		if err != nil {
			global.Log.Error(err)
		}
		byteData, err := io.ReadAll(fileObj)
		imageHash := utils.Md5(byteData)
		//去数据库中查这个图片是否存在
		var bannerModel models.BannerModel
		err = global.DB.Take(&bannerModel, "hash = ?", imageHash).Error
		if err == nil {
			//找到了
			resList = append(resList, FileUploadResponse{
				Filename:  bannerModel.Path,
				IsSuccess: false,
				Msg:       "图片已存在",
			})
			continue
		}
		if global.Config.QiNiu.Enable {
			filePath, err = qiniu.UploadImage(byteData, fileName, "gvb")
			if err != nil {
				global.Log.Error(err)
				continue
			}
			resList = append(resList, FileUploadResponse{
				Filename:  filePath,
				IsSuccess: true,
				Msg:       "上传七牛成功",
			})
			global.DB.Create(&models.BannerModel{
				Path:      filePath,
				Hash:      imageHash,
				Name:      fileName,
				ImageType: ctype.QiNiu,
			})
			continue
		}

		err = c.SaveUploadedFile(file, filePath)
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
		//图片入库
		global.DB.Create(&models.BannerModel{
			Path:      filePath,
			Hash:      imageHash,
			Name:      fileName,
			ImageType: ctype.Local,
		})

	}
	res.OkWithData(resList, c)
}
