package main

import (
	"github.com/glebarez/sqlite"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/install"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/middleware"
	"github.com/quarkcloudio/quark-go/v2/pkg/app/admin/service"
	"github.com/quarkcloudio/quark-go/v2/pkg/builder"
	"github.com/quarkcloudio/quark-go/v2/pkg/dal/db"
	"github.com/quarkcloudio/quark-lite/dashboard"
	"github.com/quarkcloudio/quark-lite/layout"
	"github.com/quarkcloudio/quark-lite/login"
	"github.com/quarkcloudio/quark-lite/model"
	"github.com/quarkcloudio/quark-lite/resource"
	"gorm.io/gorm"
)

// 自动构建数据库
func migrate() {

	// 迁移数据
	db.Client.AutoMigrate(
		&model.Demo{},
	)

	// 数据填充
	(&model.Demo{}).Seeder()
}

// 注册服务
func providers() []interface{} {

	return []interface{}{
		&login.Index{},
		&layout.Index{},
		&dashboard.Index{},
		&resource.Demo{},
	}
}

func main() {

	// 数据库配置信息
	dsn := "./data.db"

	// 配置资源
	config := &builder.Config{
		AppKey:    "abcdefg",
		Providers: append(service.Providers, providers()...),
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open(dsn),
			Opts:      &gorm.Config{},
		},
	}

	// 实例化对象
	b := builder.New(config)

	// 静态文件
	b.Static("/", "./web/app")

	// 自动构建数据库、拉取静态文件
	install.Handle()

	// 自动构建本地数据库
	migrate()

	// 后台中间件
	b.Use(middleware.Handle)

	// 重定向到后台管理
	b.GET("/", func(ctx *builder.Context) error {
		return ctx.Redirect(301, "/admin/")
	})

	// 启动服务
	b.Run(":3000")
}
