package flag

import (
	sys_flag "flag"
	"fmt"
	"github.com/fatih/structs"
	"time"
	"wild_goose_gin/flag/fake"
)

type Option struct {
	DB   bool
	User string
	Es   string
	Fake string
}

// Parse 解析命令行参数
func Parse() Option {
	db := sys_flag.Bool("db", false, "初始化数据库")    // false 默认值, 含有-db就为true
	fake := sys_flag.String("fake", "", "生成模拟数据") // false 默认值, 含有-db就为true
	user := sys_flag.String("u", "", "创建用户")
	es := sys_flag.String("es", "", "es操作")
	// 解析命令行参数写入注册的flag里
	sys_flag.Parse()
	return Option{
		DB:   *db,
		User: *user,
		Es:   *es,
		Fake: *fake,
	}
}

// IsWebStop 是否停止web项目
func IsWebStop(option Option) (f bool) {
	crMap := structs.Map(&option)
	for _, v := range crMap {
		switch val := v.(type) {
		case string:
			if val != "" {
				f = true
			}
		case bool:
			if val == true {
				f = true
			}
		}

	}
	return f
}

// SwitchOption 根据命令执行不同的函数
func SwitchOption(option Option) {
	if option.DB {
		Makemigrations()
		return
	}
	if option.User == "admin" {
		// 控制台创建admin
		CreateAdmin()
		return
	}

	// 伪造数据
	if option.Fake == "material" {
		startTime := time.Now() // 记录开始时间
		for i := 0; i < 20; i++ {
			fake.FakeMaterial()
		}
		elapsedTime := time.Since(startTime) // 计算经过的时间
		fmt.Printf("总共执行时间: %s\n", elapsedTime)
	}

}
