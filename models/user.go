package models

import (
	"wild_goose_gin/global"
	"wild_goose_gin/pkg/request"
)

type User struct {
	Model
	NickName      string `gorm:"comment:昵称" json:"nick_name"`
	UserName      string `gorm:"comment:用户名" json:"user_name"`
	Password      string `gorm:"comment:密码" json:"-"`
	AvatarImageID *uint  `gorm:"comment:头像ID" json:"avatar_image_id"`
	AvatarImage   *Image `gorm:"foreignKey:AvatarImageID"`
	Tel           string `gorm:"comment:电话" json:"tel"`
	Email         string `gorm:"comment:邮箱" json:"email"`
	Addr          string `gorm:"comment:地址" json:"addr"`
	IP            string `gorm:"comment:ip地址" json:"ip"`
	Profile       string `gorm:"comment:简介" json:"profile"`
	RoleID        uint
	Role          *Role
	GroupID       *uint `gorm:"comment:小组  1 燃油  2 液压  3 电源  4 电气" json:"group_id"`
	Group         *Group
}

func (u *User) UpdateAvatarImageID() error {
	return global.DB.Model(u).Update("avatar_image_id", u.AvatarImageID).Error
}

func (u *User) GetUserInfoByID() (user *User, err error) {
	err = global.DB.Preload("AvatarImage").Preload("Role").Take(u, u.ID).Error
	return u, err
}

func (u *User) GetUserInfoByUserName() (user *User, err error) {
	err = global.DB.Where("user_name = ?", u.UserName).Preload("Role").Take(u).Error
	return u, err
}

func (u *User) GetUserRoleID() (roleID uint, err error) {
	err = global.DB.Take(u).Error
	return u.RoleID, nil
}

func (u *User) GetUserList(pReq *request.PaginationReq) (users []User, count int64, err error) {
	limit := pReq.PageSize
	offset := pReq.GetOffset()
	db := global.DB.Model(u)
	db.Count(&count)
	err = global.DB.Preload("Role").Limit(limit).Offset(offset).Find(&users).Error
	return
}

func (u *User) AddOneRecord() error {
	return global.DB.Create(u).Error
}

func (u *User) DeleteByID() error {
	return global.DB.Delete(u).Error
}
