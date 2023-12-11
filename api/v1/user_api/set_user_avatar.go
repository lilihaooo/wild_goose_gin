package user_api

import (
	"github.com/gin-gonic/gin"
	"io"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/service"
	"wild_goose_gin/utils"
	"wild_goose_gin/utils/jwts"
)

func (UserApi) SetAvatar(c *gin.Context) {
	file, err := c.FormFile("image")
	if err != nil {
		response.FailWithMsg(c, response.INVALID_PARAMS, "无法解析上传的表单数据")
		return
	}
	// 先查询是否存在该图片, 如果存在直接修改图片id
	var user models.User
	payload := c.MustGet("user").(*jwts.Payload)
	user.ID = payload.UserID
	// 读取文件并将文件Md5加密(唯一, 同一文件的Md5值是相同的)
	fileObj, err := file.Open()
	if err != nil {
		global.Logrus.Error(err)
		return
	}
	byteData, err := io.ReadAll(fileObj)
	imageHash := utils.Md5(byteData)
	var image models.Image
	image.Hash = imageHash
	imageID := image.GetImageIDByHash()
	if imageID != 0 {
		user.AvatarImageID = &imageID
		if err := user.UpdateAvatarImageID(); err != nil {
			response.FailWithMsg(c, response.FAIL_OPER, "用户头像更改失败")
			return
		}
		response.OkWithMsg(c, "更新头像成功")
		return
	}

	// 不存在则上传
	fileUploadRes := service.AppService.ImageService.ImageUploadService(c, file)
	if fileUploadRes.IsSuccess {
		// 上传成功, 修改用户的image_id
		user.AvatarImageID = &fileUploadRes.FileID
		if err := user.UpdateAvatarImageID(); err != nil {
			response.FailWithMsg(c, response.FAIL_OPER, "用户头像更改失败")
			return
		}
		response.OkWithMsg(c, fileUploadRes.Msg)
		return
	}
	response.FailWithMsg(c, response.FAIL_OPER, fileUploadRes.Msg)

}
