package image_ser

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"io"
	"mime/multipart"
	"os"
	"path/filepath"
	"strings"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"

	"wild_goose_gin/plugins/qi_niu"
	"wild_goose_gin/utils"
)

type FileUploadResponse struct {
	FileID    uint   `json:"file_id"`
	Name      string `json:"name"`
	IsSuccess bool   `json:"is_success"`
	Msg       string `json:"msg"`
}

func (ImageService) ImageUploadService(c *gin.Context, file *multipart.FileHeader) (fileUploadRes FileUploadResponse) {
	fileUploadRes.Name = file.Filename
	ext := strings.ToLower(filepath.Ext(file.Filename))
	// 检查文件类型是否在白名单中
	ok := utils.InListStr(ext, global.Config.WhiteList)
	if !ok {
		fileUploadRes.IsSuccess = false
		fileUploadRes.Msg = "非法文件"
		return
	}
	// 判断是否超出指定大小
	fileSizeInMB := float64(file.Size) / (1024 * 1024) // file.Size单位为字节将其转为MB
	formattedSize := fmt.Sprintf("%.2f", fileSizeInMB) //保留2位小数
	if fileSizeInMB > global.Config.Upload.Size {
		fileUploadRes.IsSuccess = false
		fileUploadRes.Msg = fmt.Sprintf("图片不能大于%.2fMB, 当前文件大小为%sMB", global.Config.Upload.Size, formattedSize)
		return
	}

	// 读取文件并将文件Md5加密(唯一, 同一文件的Md5值是相同的)
	fileObj, err := file.Open()
	if err != nil {
		global.Logrus.Error(err)
		return
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	// 去数据库中查询这个文件是否存在
	var banner models.Image
	err = global.DB.Take(&banner, "hash = ?", imageHash).Error
	if err == nil {
		fileUploadRes.IsSuccess = false
		fileUploadRes.Msg = fmt.Sprintf("图片已存在")
		return
	}
	var image models.Image

	// 上传至七牛云
	if global.Config.QiNiu.Enable {
		// 上传七牛
		filename := fmt.Sprintf("%s%s", "wild_goose_gin", file.Filename)
		path, err := qi_niu.UploadImage(byteData, filename, global.Config.QiNiu.Prefix)
		if err != nil {
			fileUploadRes.IsSuccess = false
			fileUploadRes.Msg = err.Error()
		}
		//写入数据库
		image.Path = path
		image.Hash = imageHash
		image.Name = filename
		image.Type = common_type.QiNiu
		if err := global.DB.Create(&image).Error; err != nil {
			fileUploadRes.IsSuccess = false
			fileUploadRes.Msg = fmt.Sprintf("图片入库失败")
			return
		}
		fileUploadRes.IsSuccess = true
		fileUploadRes.FileID = image.ID
		fileUploadRes.Msg = fmt.Sprintf("图片上传成功")
		return
	}

	// 上传本地
	uploadDir := global.Config.Upload.Path
	// 检查并创建存储上传文件的目录
	if _, err = os.Stat(uploadDir); os.IsNotExist(err) {
		if err = os.MkdirAll(uploadDir, os.ModePerm); err != nil {
			response.FailWithMsg(c, response.FAIL_OPER, "无法创建上传目录")
			return
		}
	}
	uploadPath := filepath.Join(global.Config.Upload.Path, file.Filename)
	if err = c.SaveUploadedFile(file, uploadPath); err != nil {
		fileUploadRes.IsSuccess = false
		fileUploadRes.Msg = fmt.Sprintf("图片保存失败")
		return
	}

	image.Path = uploadPath
	image.Hash = imageHash
	image.Name = file.Filename
	image.Type = common_type.Local
	//入数据库
	if err := global.DB.Create(&image).Error; err != nil {
		os.Remove(uploadPath)
		fileUploadRes.IsSuccess = false
		fileUploadRes.Msg = fmt.Sprintf("图片入库失败")
		return
	}
	fileUploadRes.IsSuccess = true
	fileUploadRes.Msg = fmt.Sprintf("图片上传成功")
	fileUploadRes.FileID = image.ID
	return
}
