package services

import (
	"path"
	"vue_shop/repositories"
	"vue_shop/utils"
)

type UploadFileService struct {
	repo *repositories.UploadFileRepositories
}

func NewUploadFileService() *UploadFileService {
	return &UploadFileService{repo: repositories.NewUploadFileRepositories()}
}

func (s *UploadFileService) UploadImg(buffer []byte, size int64, filename string) (string, error) {
	fileext := path.Ext(filename)
	filePath, err := utils.UploadByBuffer(buffer, fileext[1:])
	return filePath, err
}
