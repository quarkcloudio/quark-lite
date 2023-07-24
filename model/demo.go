package model

import (
	"time"

	appmodel "github.com/quarkcms/quark-go/v2/pkg/app/admin/model"
	"github.com/quarkcms/quark-go/v2/pkg/dal/db"
	"gorm.io/gorm"
)

// 模型
type Demo struct {
	Id        int            `json:"id" gorm:"autoIncrement"`
	Name      string         `json:"name" gorm:"size:200;not null"`
	Status    int            `json:"status" gorm:"size:4;not null;default:1"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `json:"deleted_at"`
}

// Seeder
func (m *Demo) Seeder() {

	// 如果菜单已存在，不执行Seeder操作
	if (&appmodel.Menu{}).IsExist(18) {
		return
	}

	// 创建菜单
	menuSeeders := []*appmodel.Menu{
		{Id: 18, Name: "演示管理", GuardName: "admin", Icon: "icon-book", Type: 1, Pid: 0, Sort: 0, Path: "/demo", Show: 1, Status: 1},
		{Id: 19, Name: "演示列表", GuardName: "admin", Icon: "", Type: 2, Pid: 18, Sort: 0, Path: "/api/admin/demo/index", Show: 1, Status: 1},
	}
	db.Client.Create(&menuSeeders)
}
