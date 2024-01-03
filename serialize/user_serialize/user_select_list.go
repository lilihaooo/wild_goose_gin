package user_serialize

import (
	"strconv"
	"wild_goose_gin/models"
)

type UserSelect struct {
	ID       string `json:"id"`
	UserName string `json:"user_name"`
}

func BuildUserSelect(item models.User) UserSelect {
	return UserSelect{
		ID:       strconv.Itoa(int(item.ID)),
		UserName: item.UserName,
	}
}

func BuildUsersSelect(items []models.User) (users []UserSelect) {
	for _, item := range items {
		user := BuildUserSelect(item)
		users = append(users, user)
	}
	return users
}
