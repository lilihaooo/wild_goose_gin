package user_api

import (
	"github.com/gin-gonic/gin"
	"strings"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/pkg/response"
	"wild_goose_gin/utils/jwts"
)

type userInfoRes struct {
	UserID     uint   `json:"user_id"`
	UserName   string `json:"user_name"`
	NickName   string `json:"nick_name"`
	Profile    string `json:"profile"`
	Tel        string `json:"tel"`
	Avatar     string `json:"avatar"`
	Email      string `json:"email"`
	GroupID    uint   `json:"group_id"`
	RoleIDs    []uint `json:"role_ids"`
	RoleTitle  string `json:"role_title"`
	GroupTitle string `json:"group_title"`
}

func (UserApi) UserInfo(c *gin.Context) {
	userID := c.MustGet("user").(*jwts.Payload).UserID
	var model models.User
	model.ID = userID
	user, err := model.GetUserInfoByID()
	if err != nil {
		response.FailWithMsg(c, response.FAIL_OPER, "")
		return
	}
	// 处理图片路径
	avatarPath := ""
	if user.AvatarImage != nil {
		if user.AvatarImage.Type == common_type.Local {
			// uploads/file/logo.jpeg
			// api/static/file/logo.jpeg
			avatarPath = strings.Replace(user.AvatarImage.Path, "uploads", "api/static", 1)
		}
	}

	var userRolesIDs []uint
	userRolesTitle := ""
	if len(user.Roles) > 0 {
		for _, one := range user.Roles {
			userRolesIDs = append(userRolesIDs, one.ID)
			userRolesTitle += one.Title + " "
		}
		userRolesTitle = userRolesTitle[:len(userRolesTitle)-1]
	}

	userGroupID := 0
	userGroupTitle := ""
	if user.Group != nil {
		userGroupID = int(*user.GroupID)
		userGroupTitle = user.Group.Name
	}

	res := userInfoRes{
		UserID:     user.ID,
		UserName:   user.UserName,
		NickName:   user.NickName,
		Profile:    user.Profile,
		Tel:        user.Tel,
		Email:      user.Email,
		Avatar:     avatarPath,
		RoleIDs:    userRolesIDs,
		RoleTitle:  userRolesTitle,
		GroupTitle: userGroupTitle,
		GroupID:    uint(userGroupID),
	}
	response.OkWithData(c, res)
}
