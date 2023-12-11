package flag

import (
	"fmt"
	"wild_goose_gin/global"
	"wild_goose_gin/models"
	"wild_goose_gin/models/common_type"
	"wild_goose_gin/utils"
)

type AdminCreateRequest struct {
	NickName   string ` validate:"required,max=36"`
	UserName   string `validate:"required,max=36"`
	Password   string `validate:"required,max=6"`
	RePassword string `validate:"required,max=6"`
	Email      string
}

func CreateAdmin() {
	var cr AdminCreateRequest
	fmt.Print("请输入用户名: ")
	fmt.Scan(&cr.UserName)
	fmt.Print("请输入昵称: ")
	fmt.Scan(&cr.NickName)
	fmt.Print("请输入密码: ")
	fmt.Scan(&cr.Password)
	fmt.Print("请再次输入密码: ")
	fmt.Scan(&cr.RePassword)
	if cr.Password != cr.RePassword {
		fmt.Println("两次输入的密码不一致!!!")
		return
	}
	fmt.Print("请输入用邮箱: ")
	fmt.Scanln(&cr.Email) // Scanln 可以不输入

	// 验证参数
	vErr := utils.ZhValidate(&cr)
	if vErr != "" {
		fmt.Printf("err: %s!!! \n", vErr)
		return
	}
	// 判断用户是否存在
	var userModel models.User
	if err := global.DB.Take(userModel, "user_name = ?", cr.UserName).Error; err == nil {
		fmt.Println("用户已存在!!!")
		return
	}

	// 对密码进行加密处理
	hashPassword, err := utils.HashPassword(cr.Password)
	if err != nil {
		fmt.Println("密码加密失败!!!")
		return
	}
	userModel.UserName = cr.UserName
	userModel.NickName = cr.NickName
	userModel.Password = hashPassword
	userModel.Addr = "内网地址"
	userModel.IP = "127.0.0.1"
	userModel.Email = cr.Email
	userModel.Role.ID = common_type.Admin
	if err = global.DB.Create(&userModel).Error; err != nil {
		fmt.Println("添加失败!!!")
		return
	}
	fmt.Println("添加成功!!!")
}
