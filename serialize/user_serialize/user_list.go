package user_serialize

import (
	"wild_goose_gin/models"
	"wild_goose_gin/utils"
)

type User struct {
	ID         uint   `json:"id"`
	UserName   string `json:"user_name"`
	NickName   string `json:"nick_name"`
	Email      string `json:"email"`
	Tel        string `json:"tel"`
	RoleID     uint   `json:"role_id"`
	RoleTitle  string `json:"role"`
	GroupID    uint   `json:"group_id"`
	GroupTitle string `json:"group"`
	CreatedAt  string `json:"created_at"`
	DeletedAt  string `json:"deleted_at"`
}

func BuildUser(item models.User) User {
	groupTitle := ""
	groupID := uint(0)
	if item.Group != nil {
		groupTitle = item.Group.Name
		groupID = item.Group.ID
	}
	return User{
		ID:         item.ID,
		UserName:   item.UserName,
		NickName:   item.NickName,
		Tel:        item.Tel,
		Email:      item.Email,
		RoleID:     item.Role.ID,
		RoleTitle:  item.Role.Title,
		GroupID:    groupID,
		GroupTitle: groupTitle,
		CreatedAt:  utils.TimeFormat_YMD(item.CreatedAt),
		DeletedAt:  utils.TimeFormat_YMD(item.CreatedAt),
	}
}
func BuildUsers(items []models.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
