package main

import (
	"time"

	"github.com/quarkcms/quark-go/pkg/app/handler/admin"
	"github.com/quarkcms/quark-go/pkg/app/install"
	"github.com/quarkcms/quark-go/pkg/app/middleware"
	"github.com/quarkcms/quark-go/pkg/builder"
	"github.com/quarkcms/quark-go/pkg/dal/db"
	"github.com/quarkcms/quark-lite/dashboard"
	"github.com/quarkcms/quark-lite/login"
	"github.com/quarkcms/quark-lite/model"
	"github.com/quarkcms/quark-lite/resource"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

// 布局配置
func layout() *builder.AdminLayout {

	return &builder.AdminLayout{
		Title:     "QuarkLite",
		Copyright: time.Now().Format("2006") + " " + "QuarkLite",
		Links: []map[string]interface{}{
			{
				"key":   "Github",
				"title": "Github",
				"href":  "https://github.com/quarkcms",
			},
		},
	}
}

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
		Providers: append(admin.Providers, providers()...),
		DBConfig: &builder.DBConfig{
			Dialector: sqlite.Open(dsn),
			Opts:      &gorm.Config{},
		},
		AdminLayout: layout(),
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
