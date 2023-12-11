package flag

import (
	"wild_goose_gin/global"
	"wild_goose_gin/models"
)

func Makemigrations() {
	var err error
	// 生成四张表的表结构
	err = global.DB.Set("gorm:table_options", "ENGINE=InnoDB").
		AutoMigrate(
			&models.User{},
			&models.Modify{},
			&models.Manual{},
			&models.Certificate{},
			&models.Component{},
			&models.Custom{},
			&models.Group{},
			&models.Task{},
			&models.Image{},
			&models.Menu{},
			&models.Route{},
			&models.Role{},
		)
	if err != nil {
		global.Logrus.Errorf("生成数据库表结构失败:%s", err)
		return
	}
	global.Logrus.Info("生成数据库表结构成功")
}
