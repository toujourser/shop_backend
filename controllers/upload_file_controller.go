package controllers

import (
	"github.com/kataras/iris/v12"
	"vue_shop/common"
	"vue_shop/services"
)

type UploadFileController struct {
	Ctx     iris.Context
	Service *services.UploadFileService
	common.Common
}

func NewUploadFileController() *UploadFileController {
	return &UploadFileController{Service: services.NewUploadFileService()}
}

func (c *UploadFileController) Post() {
	file, handle, err := c.Ctx.FormFile("file")
	if err != nil {
		c.ReturnJson(400, err.Error())
		return
	}

	fileBuffer := make([]byte, handle.Size)
	_, err = file.Read(fileBuffer)
	url := "http://192.168.237.130:8888/"

	if filePath, err := c.Service.UploadImg(fileBuffer, handle.Size, handle.Filename); err != nil {
		c.ReturnJson(500, err.Error())
		return
	} else {
		data := map[string]string{
			"tmp_path": filePath,
			"url": url + filePath,
		}
		c.ReturnSuccess(data)
	}

}
