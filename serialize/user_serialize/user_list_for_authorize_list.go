package user_serialize

import (
	"wild_goose_gin/models"
)

type UserForAuthorize struct {
	ID       uint   `json:"id"`
	UserName string `json:"user_name"`
	Group    string `json:"group"`
}

func BuildUserForAuthorize(item models.User) UserForAuthorize {
	groupName := ""
	if item.Group != nil {
		groupName = item.Group.Name
	}
	return UserForAuthorize{
		ID:       item.ID,
		UserName: item.UserName,
		Group:    groupName,
		//Role:     item.Role.Title,
	}
}

func BuildUsersListForAuthorize(items []models.User) (users []UserForAuthorize) {
	for _, item := range items {
		user := BuildUserForAuthorize(item)
		users = append(users, user)
	}
	return users
}
