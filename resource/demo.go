package resource

import (
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/actions"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/service/searches"
	"github.com/quarkcms/quark-go/v2/pkg/app/admin/template/resource"
	"github.com/quarkcms/quark-go/v2/pkg/builder"
	"github.com/quarkcms/quark-lite/model"
)

type Demo struct {
	resource.Template
}

// 初始化
func (p *Demo) Init(ctx *builder.Context) interface{} {

	// 标题
	p.Title = "Demo"

	// 模型
	p.Model = &model.Demo{}

	// 分页
	p.PerPage = 10

	return p
}

func (p *Demo) Fields(ctx *builder.Context) []interface{} {
	field := &resource.Field{}

	return []interface{}{
		field.ID("id", "ID"),

		field.Text("name", "名称"),

		field.Datetime("created_at", "创建时间"),
	}
}

// 搜索
func (p *Demo) Searches(ctx *builder.Context) []interface{} {
	return []interface{}{
		searches.Input("name", "名称"),
	}
}

// 行为
func (p *Demo) Actions(ctx *builder.Context) []interface{} {
	return []interface{}{
		actions.CreateLink(),
		actions.BatchDelete(),
		actions.BatchDisable(),
		actions.BatchEnable(),
		actions.EditLink(),
		actions.Delete(),
		actions.FormSubmit(),
		actions.FormReset(),
		actions.FormBack(),
		actions.FormExtraBack(),
	}
}
