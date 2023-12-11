package serializer

import (
	"wild_goose_gin/models"
	"wild_goose_gin/utils"
)

type User struct {
	ID        uint   `json:"id"`
	UserName  string `json:"user_name"`
	NickName  string `json:"nick_name"`
	Email     string `json:"email"`
	Tel       string `json:"tel"`
	RoleTitle string `json:"role"`
	CreatedAt string `json:"created_at"`
	DeletedAt string `json:"deleted_at"`
}

func BuildUser(item models.User) User {
	return User{
		ID:        item.ID,
		UserName:  item.UserName,
		NickName:  item.NickName,
		Tel:       item.Tel,
		Email:     item.Email,
		RoleTitle: item.Role.Title,
		CreatedAt: utils.TimeFormat(item.CreatedAt),
		DeletedAt: utils.TimeFormat(item.CreatedAt),
	}
}

func BuildUsers(items []models.User) (users []User) {
	for _, item := range items {
		user := BuildUser(item)
		users = append(users, user)
	}
	return users
}
